package main

import (
	"fmt"
	//"encoding/json"
	"net/http"
	"io"
	"os"
)

type BF1_GetMatchID struct {
	data struct {
		matches []struct {
			attributes struct {
				id 			string	 `json:"id"`
			}  			`json:"attributes"`
		    metadata 	struct {
                serverName  string 	 `json:"serverName"`
			}  			`json:"metadata"`
			}       `json:"matches"`
	}       `json:"data"`
}

func main() {
	method := "GET"
	url := "https://api.tracker.gg/api/v2/bf1/standard/profile/origin/wuhutakeoffyoo"
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:10809")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:10809")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Host", "api.tracker.gg")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("", "")
	req.Header.Add("User-Agent", "Tracker Network App/3.22.9")
	req.Header.Add("x-app-version", "3.22.9")
	fmt.Printf("Sending request.\n")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error occoured when sending request.\n")
		panic(err) // handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error occoured when reading request.\n")
		panic(err) // handle error
	}
	//var Message BF1_GetMatchID
	fmt.Printf(string(body))
	//fmt.Printf("%s", body)
	//json.Unmarshal([]byte(body.ToJsonString()), &Message)
	//fmt.Printf("The most recent match id: %s", Message.data.matches[0].attributes.id)
}