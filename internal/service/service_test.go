package service

import (
	"testing"

	"github.com/azhao1981/systemctl-add/pkg/cfg"
	"github.com/creasty/defaults"
	"github.com/magiconair/properties/assert"
)

func TestExc(t *testing.T) {
	t.Log("TestExc")
	body, err := Execute(map[string]interface{}{
		"Description": "freeswitch",
		"ExecStart":   "/usr/local/freeswitch/bin/freeswitch -nc -nonat",
	})
	assert.Equal(t, err, nil)

	t.Log(body)
}

func TestExcService(t *testing.T) {
	t.Log("TestExc")
	service := cfg.Service{
		Description: "freeswitch",
		ExecStart:   "/usr/local/freeswitch/bin/freeswitch -nc -nonat",
		User:        "ubuntu",
	}
	defaults.Set(&service)
	body, err := Execute(service)
	assert.Equal(t, err, nil)

	t.Log(body)
}
