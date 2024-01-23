package gateway

import (
	"log/slog"
	"net/http"

	"github.com/Ra2-IFV/TRN-BF1/pkg/netreq"
)

func main() {}

func requestGateway(platform string) []byte {
	method := "GET"
	header := map[string]string{
		"Content-Type": "application/json",
		"Connection":   "keep-alive",
	}
	transport := &http.Transport{}
	url := "https://sparta-gw-bf1.battlelog.com/jsonrpc/" + platform + "/api"
	slog.Info(
		"Request",
		"url", url,
	)
	data, err := netreq.Request{
		Method:    method,
		Header:    header,
		URL:       url,
		Transport: transport}.ReadRespBodyByte()
	if err != nil {
		slog.Error("Failed to send request", "error", err)
	}
	return data
}
