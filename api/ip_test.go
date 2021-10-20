package api

import (
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIp(t *testing.T) {
	t.Run("external", func(t *testing.T) {
		external := externalIP()
		log.Println("External IP is", external)
		assert.NotNil(t, external)
		assert.NotEmpty(t, external)
		ip := net.ParseIP(external)
		assert.NotNil(t, ip)
		assert.NotEmpty(t, ip)
	})
	t.Run("outbound", func(t *testing.T) {
		external := OutboundIP()
		log.Println("Outbound IP is", external)
		assert.NotNil(t, external)
		assert.NotEmpty(t, external)
		ip := net.ParseIP(external)
		assert.NotNil(t, ip)
		assert.NotEmpty(t, ip)
	})
}
