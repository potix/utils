package configurator

import (
	"os"
	"testing"
)

type Config struct {
	A string
	B int
	C bool
}

func TestLoadToml(t *testing.T) {
	cf, err := NewConfigurator("./example/example.toml")
	if err != nil {
		t.Fatalf("can not create configurator: %v", err)
	}
	var conf Config
	if err = cf.Load(&conf); err != nil {
		t.Fatalf("can not load config: %v", err)
	}
	if conf.A != "aaa" {
		t.Fatalf("A is wrong: %v", conf.A)
	}
	if conf.B != 100 {
		t.Fatalf("B is wrong: %v", conf.B)
	}
	if conf.C != true {
		t.Fatalf("C is wrong: %v", conf.C)
	}
	if err = cf.Save(FormatTypeAuto, &conf); err != nil {
		t.Fatalf("can not save config: %v", err)
	}
	if err = cf.Load(&conf); err != nil {
		t.Fatalf("can not load config: %v", err)
	}
	if conf.A != "aaa" {
		t.Fatalf("A is wrong: %v", conf.A)
	}
	if conf.B != 100 {
		t.Fatalf("B is wrong: %v", conf.B)
	}
	if conf.C != true {
		t.Fatalf("C is wrong: %v", conf.C)
	}
}

func TestLoadJson(t *testing.T) {
	cf, err := NewConfigurator("./example/example.json")
	if err != nil {
		t.Fatalf("can not create configurator: %v", err)
	}
	var conf Config
	if err = cf.Load(&conf); err != nil {
		t.Fatalf("can not load config: %v", err)
	}
	if conf.A != "aaa" {
		t.Fatalf("A is wrong: %v", conf.A)
	}
	if conf.B != 100 {
		t.Fatalf("B is wrong: %v", conf.B)
	}
	if conf.C != true {
		t.Fatalf("C is wrong: %v", conf.C)
	}
	if err = cf.Save(FormatTypeAuto, &conf); err != nil {
		t.Fatalf("can not save config: %v", err)
	}
	if err = cf.Load(&conf); err != nil {
		t.Fatalf("can not load config: %v", err)
	}
	if conf.A != "aaa" {
		t.Fatalf("A is wrong: %v", conf.A)
	}
	if conf.B != 100 {
		t.Fatalf("B is wrong: %v", conf.B)
	}
	if conf.C != true {
		t.Fatalf("C is wrong: %v", conf.C)
	}
}

func TestLoadYaml(t *testing.T) {
	cf, err := NewConfigurator("./example/example.yaml")
	if err != nil {
		t.Fatalf("can not create configurator: %v", err)
	}
	var conf Config
	if err = cf.Load(&conf); err != nil {
		t.Fatalf("can not load config: %v", err)
	}
	if conf.A != "aaa" {
		t.Fatalf("A is wrong: %v", conf.A)
	}
	if conf.B != 100 {
		t.Fatalf("B is wrong: %v", conf.B)
	}
	if conf.C != true {
		t.Fatalf("C is wrong: %v", conf.C)
	}
	if err = cf.Save(FormatTypeAuto, &conf); err != nil {
		t.Fatalf("can not save config: %v", err)
	}
	if err = cf.Load(&conf); err != nil {
		t.Fatalf("can not load config: %v", err)
	}
	if conf.A != "aaa" {
		t.Fatalf("A is wrong: %v", conf.A)
	}
	if conf.B != 100 {
		t.Fatalf("B is wrong: %v", conf.B)
	}
	if conf.C != true {
		t.Fatalf("C is wrong: %v", conf.C)
	}
}

func TestSecret(t *testing.T) {
	if err := os.Chmod("./example/secret", 0644); err != nil {
		t.Fatalf("can not chmod: %v", err)
	}
	keys, err :=  LoadSecretFile("./example/secret")
	if err == nil {
		t.Fatalf("permission check error: %v", err)
	}
	if keys != nil {
		t.Fatalf("permission check error: %v", err)
	}
	if err := os.Chmod("./example/secret", 0600); err != nil {
		t.Fatalf("can not chmod: %v", err)
	}
	keys, err =  LoadSecretFile("./example/secret")
	if err != nil {
		t.Fatalf("load error: %v", err)
	}
	if keys == nil {
		t.Fatalf("load error: %v", err)
	}
	if len(keys) != 2 {
		t.Fatalf("wrong keys: %+v", keys)
	}
	if keys[0] != "ssss" {
		t.Fatalf("wrong key 0: %v", keys[0])
	}
	if keys[1] != "vvvv" {
		t.Fatalf("wrong key 1: %v", keys[1])
	}
}
