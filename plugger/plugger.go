package plugger

import (
    "fmt"
    "plugin"
    "os/user"
    "log"
    "regexp"
    "path/filepath"
    "io/ioutil"
)

// CallbackFunc is callback func
type CallbackFunc func(request interface{}) (response interface{}, err error)

// Caller is caller
type Caller struct {
	Name string
	callbacks map[string]CallbackFunc
}

// SetCallback is set callback
func (c *Caller)setCallback(methodName string, callback CallbackFunc) {
	c.callbacks[methodName] = callback
}

// Callback is  callback
func (c *Caller)Callback(methodName string, request interface{}) (response interface{}, err error) {
	callback, ok := c.callbacks[methodName]
	if !ok {
		return nil, fmt.Errorf("no callback (%v)", methodName)
	}
	return callback(request)
}

// Plugin is plugin
type Plugin interface {
	GetName() (pluginName string)
	Initialize() (err error)
	Finalize()
	Call(methodName string, request interface{}) (response interface{}, err error)
}

// PluginContext is plugin context
type PluginContext struct {
	caller *Caller
	plugin Plugin
}

// SetCallback is set callback
func (p *PluginContext)SetCallback(methodName string, callback CallbackFunc) {
	p.caller.setCallback(methodName, callback)
}

// Initialize is Initilize
func (p *PluginContext)Initialize() (error) {
	return p.plugin.Initialize()
}

// Finalize is Finalize
func (p *PluginContext)Finalize() {
	p.plugin.Finalize()
}

// Call is Call
func (p *PluginContext)Call(methodName string, request interface{}) (interface{}, error) {
	return p.plugin.Call(methodName, request)
}

const (
	// GetPluginInfo is GetPluginInfo symoble
	GetPluginInfo string = "GetPluginInfo"
)

// PluginNewFunc is PluginNewFunc
type PluginNewFunc func(caller *Caller, configFile string) (Plugin, error)

// GetPluginInfoFunc is GetPluginInfoFunc
type GetPluginInfoFunc func() (string, PluginNewFunc)

type pluginInfo struct {
     pluginFilePath string
     pluginNewFunc PluginNewFunc
}

var registeredPlugins = make(map[string]*pluginInfo)

func registerPlugin(pluginFilePath string,  getPluginInfoFunc GetPluginInfoFunc) {
    name, pluginNewFunc := getPluginInfoFunc()
    registeredPlugins[name] = &pluginInfo {
        pluginFilePath: pluginFilePath,
        pluginNewFunc: pluginNewFunc,
    }
    log.Printf("regiter plugin %v", name)
}

func getPluginSymbole(openedPlugin *plugin.Plugin) (GetPluginInfoFunc, error) {
    s, err := openedPlugin.Lookup(GetPluginInfo)
    if err != nil {
	return nil, fmt.Errorf("not found %v symbole: %w", GetPluginInfo, err)
    }
    return s.(func() (string, PluginNewFunc)), nil
}

func loadPlugin(pluginFilePath string) (error) {
    openedPlugin, err := plugin.Open(pluginFilePath)
    if err != nil {
	return fmt.Errorf("can not open plugin file (file = %v): %w", pluginFilePath, err)
    }
    f, err := getPluginSymbole(openedPlugin)
    if err != nil {
	return fmt.Errorf("not plugin file (file = %v): %w", pluginFilePath, err)
    }
    registerPlugin(pluginFilePath, f)
    return nil
}

func fixupPluginPath(pluginPath string) (string) {
    u, err := user.Current()
    if err != nil {
        return pluginPath
    }
    re := regexp.MustCompile("^~/")
    return re.ReplaceAllString(pluginPath, u.HomeDir+"/")
}

func loadPluginFiles(pluginPath string) (error) {
    if pluginPath == "" {
        return fmt.Errorf("invalid plugin path")
    }
    pluginPath = fixupPluginPath(pluginPath)
    fileList, err := ioutil.ReadDir(pluginPath)
    if err != nil {
	    return fmt.Errorf("can not read directory (path = %v): %w", pluginPath, err)
    }
    for _, file := range fileList {
        newPath := filepath.Join(pluginPath, file.Name())
        if file.IsDir() {
            err := loadPluginFiles(newPath)
            if err != nil {
                log.Printf("can not load plugin files (%v): %v", newPath, err)
            }
            continue
        }
	ext := filepath.Ext(file.Name())
	if ext != ".so" && ext != ".dylib" {
	    continue
	}
	err := loadPlugin(newPath)
	if err != nil {
	    log.Printf("can not load plugin file (%v): %v", newPath, err)
	    continue
	}
    }
    return nil
}

// LoadPlugins is load actor Plugins
func LoadPlugins(pluginPath string) (error) {
	return loadPluginFiles(pluginPath)
}

// GetPluginContext is get plugin context
func GetPluginContext(pluginName string, callerName string, configFilePath string) (*PluginContext, error) {
        info, ok := registeredPlugins[pluginName]
	if !ok {
		return nil, fmt.Errorf("not found plugin (%v)", pluginName)
	}
	caller := &Caller{
		Name : callerName,
	}
        pluginDir := filepath.Dir(info.pluginFilePath)
        pluginConfigPath := filepath.Join(pluginDir, configFilePath)
        newPlugin, err := info.pluginNewFunc(caller, pluginConfigPath)
        if err != nil {
		return nil, fmt.Errorf("can not create plugin instance (%v): %w", pluginName, err)
        }
	pluginCtx := &PluginContext{
		caller : caller,
		plugin : newPlugin,
	}
        return pluginCtx, err
}

