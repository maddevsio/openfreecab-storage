package conf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeCanGetConfiguration(t *testing.T) {
	cr := NewConfigurator()
	os.Clearenv()
	os.Setenv("HTTP_BIND_ADDR", ":9090")
	cr.Run()
	conf := cr.Get()
	assert.Equal(t, conf.HTTPBindAddr, ":9090")
}
