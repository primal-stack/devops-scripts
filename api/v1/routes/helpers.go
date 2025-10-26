package helpers

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// NewUUID generates a random UUID.
func NewUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// GetUUID generates a random UUID.
func GetUUID() (string, error) {
	return NewUUID()
}

// GetHostname gets the current hostname.
func GetHostname() (string, error) {
	return os.Hostname()
}

// GetIPAddress gets the current IP address.
func GetIPAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", errors.New("no non-loopback IPv4 address found")
}

// GetMACAddress gets the current MAC address.
func GetMACAddress(iface string) (string, error) {
	addrs, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		if strings.Contains(a.Name, iface) {
			// Handle error if there is no hardware address
			hw, err := a.HardwareAddr()
			if err != nil {
				return "", err
			}
			return hw.String(), nil
		}
	}
	return "", errors.New("interface not found")
}

// GetOSFamily gets the current OS family.
func GetOSFamily() string {
	if runtime.GOOS == "windows" {
		return "windows"
	}
	return "linux"
}