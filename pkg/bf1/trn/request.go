package trn

import (
	"crypto/tls"
	"log/slog"
	"net/http"

	"github.com/Ra2-IFV/TRN-BF1/pkg/netreq"
)

// Initialize a HTTP request for TRN API
func requestGetTRN(url string, displayName string) ([]byte, error) {
	method := "GET"
	header := map[string]string{
		"Content-Type": "application/json",
		// Won't parse gzipped body transparently if uncommented
		//"Accept-Encoding": "gzip",
		"User-Agent":    "Tracker Network App/3.22.9",
		"x-app-version": "3.22.9",
	}
	transport := &http.Transport{
		//Proxy:           http.ProxyURL(proxyUrl),
		// Disable http2 to bypass Cloudflare, choose one from below
		TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
		//TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	slog.Info(
		"Request",
		"url", url+displayName,
	)
	data, err := netreq.Request{
		Method:    method,
		Header:    header,
		URL:       url + displayName,
		Transport: transport,
	}.ReadRespBodyByte()
	if err != nil {
		slog.Warn("Request to TRN failed", "error", err)
		return nil, err
	}
	return data, nil
}
