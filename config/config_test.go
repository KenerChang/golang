package config

import (
	"testing"
)

func TestParseEnvs(t *testing.T) {
	err := Parse("test.env")

	if err != nil {
		t.Errorf("test parseEnvs failed, %s", err.Error())
		return
	}

	target := map[string]string{"a": "b", "c": "d=d", "e": "fg"}
	for k, v := range store {
		if targetV, ok := target[k]; !ok || v != targetV {
			t.Errorf("parse envs failed, expect %v, got %v", target, store)
			return
		}
	}
}
