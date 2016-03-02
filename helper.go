package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

// genToken will create a unique token
// for a web form.  The token will ensure
// a double submit will not be sent.
func genToken() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

// OptionItem will help with comboboxes.
// A a selected option needs to be known.
// Use this to create a slice of ListOptions
// for the combobox to display.  The selected flag
// will be set for the selected item.  Only 1 selected
// flag should be set in a slice.
type OptionItem struct {
	Value    string // The value to equal when the option is selected
	Text     string // The text to display for the option
	Selected bool   // Flag if the option is selected
}

// GetExternalIP returns the non loopback local IP of the host to the external web.
func GetExternalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// GetLocalIP will get the local IP for the computer.
func GetLocalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
