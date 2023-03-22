package memberlist

import (
	"testing"
)

func Test_IsValidAddressDefaults(t *testing.T) {
	tests := []string{
		"127.0.0.1",
		"127.0.0.5",
		"10.0.0.9",
		"172.16.0.7",
		"192.168.2.1",
		"fe80::aede:48ff:fe00:1122",
		"::1",
	}
	config := DefaultLANConfig()
	for _, ip := range tests {
		if err := config.HostAllowed(ip); err != nil {
			t.Fatalf("IP %s Localhost Should be accepted for LAN", ip)
		}
	}
	config = DefaultWANConfig()
	for _, ip := range tests {
		if err := config.HostAllowed(ip); err != nil {
			t.Fatalf("IP %s Localhost Should be accepted for WAN", ip)
		}
	}
}
