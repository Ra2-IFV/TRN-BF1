package main

import (
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	//"compress/flate"
)

type player_info struct {
	Data struct {
		PlatformInfo struct {
			PlatformUserIdentifier string `json:"platformUserIdentifier"`
			AvatarURL              string `json:"avatarUrl"`
		} `json:"platformInfo"`
		Segments []struct {
			Stats struct {
				Rank struct {
					Metadata struct {
						ImageURL string `json:"imageUrl"`
					} `json:"metadata"`
					Value uint8 `json:"value"`
				} `json:"rank"`
				TimePlayed struct {
					Value uint16 `json:"value"`
				} `json:"timePlayed"`
				Kills struct {
					Value uint32 `json:"value"`
				} `json:"kills"`
				Deaths struct {
					Value uint32 `json:"value"`
				} `json:"deaths"`
				LongestHeadshot struct {
					Value uint16 `json:"value"`
				} `json:"longestHeadshot"`
				Repairs struct {
					Value uint32 `json:"value"`
				} `json:"repairs"`
				Heals struct {
					Value uint32 `json:"value"`
				} `json:"heals"`
				Revive struct {
					Value uint32 `json:"value"`
				} `json:"revive"`
				KdRatio struct {
					Value float32 `json:"value"`
				} `json:"kdRatio"`
				ScorePerMinute struct {
					Value float32 `json:"value"`
				} `json:"scorePerMinute"`
				KillsPerMinute struct {
					Value float32 `json:"value"`
				} `json:"killsPerMinute"`
				WinPercentage struct {
					Value float32 `json:"value"`
				} `json:"winPercentage"`
				ShotsAccuracy struct {
					Value float32 `json:"value"`
				} `json:"shotsAccuracy"`
				HeadshotsPercentage struct {
					Value float32 `json:"value"`
				} `json:"headshotsPercentage"`
			} `json:"stats"`
		} `json:"segments"`
		ExpiryDate string `json:"expiryDate"`
	}
}

type match_list struct {
	Data struct {
		Matches []struct {
			Attributes struct {
				ID string `json:"id"`
			} `json:"attributes"`
		} `json:"matches"`
	} `json:"data"`
}

func main() {
	display_name := "wuhutakeoffyoo"
	fmt.Println("Proxy is set:", use_proxy())
	match_list := get_match_list(display_name)
	player_info := get_player_info(display_name)
	fmt.Println("Player name", player_info.Data.PlatformInfo.PlatformUserIdentifier)
	fmt.Println("Rank", player_info.Data.Segments[0].Stats.Rank.Value)
	fmt.Println("Kills", player_info.Data.Segments[0].Stats.Kills.Value)
	fmt.Println("Deaths", player_info.Data.Segments[0].Stats.Deaths.Value)
	fmt.Println("KD", player_info.Data.Segments[0].Stats.KdRatio.Value)
	fmt.Println("KPM", player_info.Data.Segments[0].Stats.KillsPerMinute.Value)
	fmt.Println("Heals", player_info.Data.Segments[0].Stats.Heals.Value)
	fmt.Println("Revives", player_info.Data.Segments[0].Stats.Revive.Value)
	fmt.Println("The last 3 matches ID")
	fmt.Println(match_list.Data.Matches[0].Attributes.ID)
	fmt.Println(match_list.Data.Matches[1].Attributes.ID)
	fmt.Println(match_list.Data.Matches[2].Attributes.ID)
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

func send_get_request(url string, display_name string) []byte {
	method := "GET"
	client := &http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, err := http.NewRequest(method, url+display_name, nil)
	if err != nil {
		panic(err) // handle error
	}
	req.Header = http.Header{
		"Content-Type":    {"application/json"},
		"Accept-Encoding": {"gzip"},
		"User-Agent":      {"Tracker Network App/3.22.9"},
		"x-app-version":   {"3.22.9"},
	}
	slog.Info(
		"Request",
		"full_url", url+display_name,
	)
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to send request", "error", err)
	}
	defer resp.Body.Close()
	slog.Info(
		"Response",
		"status_code", resp.StatusCode,
		"description", http.StatusText(resp.StatusCode),
		"encoding", resp.Header.Get("Content-Encoding"),
	)
	return read_resp_body(resp.Header, resp.Body)
}

func read_resp_body(resp_header http.Header, resp_body io.ReadCloser) []byte {
	var reader io.ReadCloser
	switch resp_header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp_body)
		defer reader.Close()
	//case "deflate":
	//	reader = flate.NewReader(resp_body)
	//defer reader.Close()
	case "":
		reader = resp_body
		defer reader.Close()
	default:
		slog.Warn("Unsupported Content-Encoding.", "content-encoding", resp_header)
		reader = resp_body
	}
	body, err := io.ReadAll(reader)
	if err != nil {
		slog.Error("Failed to read body", "error", err) // handle error
	}
	return body
}

func get_player_info(display_name string) player_info {
	url := "https://api.tracker.gg/api/v2/bf1/standard/profile/origin/"
	data := send_get_request(url, display_name)
	player_info := player_info{}
	json.Unmarshal(data, &player_info)
	return player_info
}

func get_match_list(display_name string) match_list {
	url := "http://api.tracker.gg/api/v2/bf1/standard/matches/origin/"
	data := send_get_request(url, display_name)
	match_list := match_list{}
	json.Unmarshal(data, &match_list)
	return match_list
}
