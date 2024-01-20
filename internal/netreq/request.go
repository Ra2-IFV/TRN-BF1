package netreq

// https://github.com/KomeiDiSanXian/BFHelper/blob/master/bfhelper/pkg/netreq/netreq.go

import (
	"io"
	"log/slog"
	"net/http"
	//"compress/gzip"
	//"compress/flate"
)

// Init request struct
type Request struct {
	Method string
	Header map[string]string
	Body   io.ReadCloser
	//Transport &http.Transport
	Transport http.RoundTripper
	URL       string
}

func (r Request) client() *http.Client {
	return &http.Client{Transport: r.Transport}
}

// Init request
func (r Request) do() (*http.Response, error) {
	// Send GET request by default
	if r.Method == "" {
		r.Method = http.MethodGet
	}
	req, err := http.NewRequest(r.Method, r.URL, r.Body)
	if err != nil {
		slog.Error("Failed to initialize request", "error", err)
	}
	// https://www.runoob.com/go/go-range.html
	for k, v := range r.Header {
		// Maybe it's better? https://stackoverflow.com/a/68018927
		req.Header.Set(k, v)
	}
	return r.client().Do(req)
}

func (r Request) respBody() (io.ReadCloser, error) {
	resp, err := r.do()
	if err != nil {
		slog.Warn("Failed to receive response", "error", err)
		return nil, err
	}
	slog.Info("Response", "Content-Encoding", resp.Header.Get("Content-Encoding"))
	return resp.Body, nil
}

// Return a decoded response body
// PS: default http transport handles encoded response properly
// https://pkg.go.dev/net/http#Transport
func (r Request) ReadRespBodyByte() ([]byte, error) {
	body, err := r.respBody()
	if err != nil {
		slog.Error("Failed to read body", "error", err)
		return nil, err
	}
	defer body.Close()
	return io.ReadAll(body)
}
