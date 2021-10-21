package api

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

// OutboundIP returns external IP of current device
//
// NOTE:  Unfortunately, this will only work on networks
// that don't employ the use of NAT.
func OutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

// externalIP returns external IP of current device
func externalIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	type IP struct {
		Query string
	}
	var ip IP
	_ = json.Unmarshal(body, &ip)
	return ip.Query
}
