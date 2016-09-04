// Package reqip provides functions for extracting IP address from a
// request and associating it with a Context.
// package based on https://blog.golang.org/context/userip/userip.go
package reqip

import (
	"net/http"
	"net"
	"fmt"
)

// FromRequest extracts the IP address from req, if present.
func FromRequest(req *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("reqip: %q is not IP:port", req.RemoteAddr)
	}
	fmt.Println(ip)
	return ip
}