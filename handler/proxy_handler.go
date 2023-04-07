package handler

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func InitProxyHandler(mux *http.ServeMux) {
	mux.HandleFunc("/", handleProxy)
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	host := getHost(r)
	path := r.URL.Path
	fmt.Printf("host: [%s], path: [%s]\n", host, path)
	var target *url.URL
	if path == "/" {
		target, _ = url.Parse("https://naver.com")
	} else {
		u := fmt.Sprintf("https://search.naver.com/search.naver?query=%s", path[1:])
		target, _ = url.Parse(u)
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(target)
	reverseProxy.Transport = &http.Transport{
		DialTLSContext:  dialTLS,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	reverseProxy.ServeHTTP(w, r)
}

func getHost(r *http.Request) string {
	i := strings.Index(r.Host, ":")
	if i > 0 {
		return r.Host[:i]
	}
	return r.Host
}

func dialTLS(ctx context.Context, network, addr string) (net.Conn, error) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	cfg := &tls.Config{ServerName: host}

	tlsConn := tls.Client(conn, cfg)
	if err := tlsConn.Handshake(); err != nil {
		conn.Close()
		return nil, err
	}

	cs := tlsConn.ConnectionState()
	cert := cs.PeerCertificates[0]

	cert.VerifyHostname(host)
	log.Println(cert.Subject)

	return tlsConn, nil
}
