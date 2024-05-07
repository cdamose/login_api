package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	config := InitConfig()
	assert.Equal(t, config.Port, "8090")
}
