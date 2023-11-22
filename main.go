package main

import (
	"compress/gzip"
	//"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type matches struct {
	Data struct {
		Matches []struct {
			Attributes struct {
				ID string `json:"id"`
			} `json:"attributes"`
		} `json:"matches"`
	} `json:"data"`
}

func main() {
	display_name := "SHlSAN13"
	fmt.Println("Proxy is set:", use_proxy())
	fmt.Println(get_match_id(display_name))
	//data := matches{}
	//json.Unmarshal([]byte(body), &data)
	//fmt.Printf("The most recent match id: %s", data.Data.Matches[0].Attributes.ID)
}

func use_proxy() bool {
	proxy := flag.String("proxy", "", "Set http proxy host. Format: host:port")
	flag.Parse()
	if *proxy != "" {
		os.Setenv("HTTP_PROXY", *proxy)
		os.Setenv("HTTPS_PROXY", *proxy)
		return true
	} else {
		return false
	}
}

func get_match_id(display_name string) []byte {
	method := "GET"
	url := "https://api.tracker.gg/api/v2/bf1/standard/matches/origin/" + display_name
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err) // handle error
	}
	req.Header = http.Header{
		"Host":            {"api.tracker.gg"},
		"Content-Type":    {"application/json"},
		"Accept-Encoding": {"gzip"},
		"User-Agent":      {"Tracker Network App/3.22.9"},
		"x-app-version":   {"3.22.9"},
	}
	fmt.Println("Sending request.")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error occoured when sending request.")
		panic(err) // handle error
	}
	defer resp.Body.Close()
	fmt.Println("Status code:", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Println("Is uncompressed:", resp.Uncompressed)
	fmt.Println("Content-Encoding:", resp.Header.Get("Content-Encoding"))
	tmpBody, _ := gzip.NewReader(resp.Body)
	body, err := io.ReadAll(tmpBody)
	if err != nil {
		fmt.Println("Error occoured when reading gzipped response.")
		panic(err) // handle error
	}
	return body
}
