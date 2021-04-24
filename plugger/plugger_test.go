package plugger

import (
        "testing"
)

func TestPlugger(t *testing.T) {
	err := LoadPlugins("./plugin")
	if err != nil {
		t.Errorf("can not load plugins: %v", err)
	}
}

