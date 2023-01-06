package socks

import (
	"http2socks5/config"
	"net"
	"time"

	"golang.org/x/net/proxy"
)

func NewClient(target string, cfg config.Socks5) (net.Conn, error) {
	auth := proxy.Auth{User: cfg.User, Password: cfg.Password}

	dailer, err := proxy.SOCKS5("tcp", cfg.Host, &auth, &net.Dialer{
		Timeout:   60 * time.Second,
		KeepAlive: 30 * time.Second,
	})

	if err != nil {
		return nil, err
	}

	return dailer.Dial("tcp", target)
}
