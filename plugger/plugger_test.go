package plugger

import (
        "testing"
)

TestPlugger() {
	err = LoadPlugins("./plugin")
	if err != nil {
		t.Errorf("can not load plugins: %v", err)
	}
}

