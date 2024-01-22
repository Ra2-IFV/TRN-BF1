package main

import (
	"crypto/tls"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/Ra2-IFV/TRN-BF1/pkg/netreq"
)

type Config struct {
	Proxy string `yaml:"proxy"`
}

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
					Value string `json:"value"`
				} `json:"kdRatio"`
				ScorePerMinute struct {
					Value string `json:"value"`
				} `json:"scorePerMinute"`
				KillsPerMinute struct {
					Value string `json:"value"`
				} `json:"killsPerMinute"`
				WinPercentage struct {
					Value string `json:"value"`
				} `json:"winPercentage"`
				ShotsAccuracy struct {
					Value string `json:"value"`
				} `json:"shotsAccuracy"`
				HeadshotsPercentage struct {
					Value string `json:"value"`
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

type match_detail struct {
	Data struct {
		Attributes struct {
			MapKey      string `json:"mapKey"`
			GamemodeKey string `json:"gamemodeKey"`
		} `json:"attributes"`
		Metadata struct {
			Timestamp  time.Time `json:"timestamp"`
			MapName    string    `json:"mapName"`
			ServerName string    `json:"serverName"`
			Teams      []struct {
				ID   uint8  `json:"id"`
				Name string `json:"name"`
			} `json:"teams"`
			Winner uint8 `json:"winner"`
		} `json:"metadata"`
		Segments []struct {
			Attributes struct {
				PlayerId string
			}
			Metadata struct {
				playerName string
			}
			Stats struct {
				KdRatio struct {
					Value string `json:"value"`
				} `json:"kdRatio"`
				KillsPerMinute struct {
					Value string `json:"value"`
				} `json:"killsPerMinute"`
				HeadshotsPercentage struct {
					Value string `json:"value"`
				} `json:"headshotsPercentage"`
				ShotAccuracy struct {
					Value string `json:"value"`
				} `json:"shotAccuracy"`
				Time struct {
					Value float32 `json:"value"`
				} `json:"time"`
				ScorePerMinute struct {
					// Original value is float64, like 12.1234567890
					Value uint16 `json:"value"`
				} `json:"scorePerMinute"`
			}
		} `json:"segments"`
	} `json:"data"`
}

func init() {}

func main() {
	display_name := "404-HK416-MOD3"
	get_player_info(display_name)
	//player_info, err := get_player_info(display_name)
	//match_list, _ := get_match_list(display_name)
	//fmt.Println("Player name", player_info)
	//fmt.Println("Rank", player_info.Data.Segments[0].Stats.Rank.Value)
	//fmt.Println("KD", player_info.Data.Segments[0].Stats.KdRatio.Value)
	//fmt.Println("KPM", player_info.Data.Segments[0].Stats.KillsPerMinute.Value)
	//fmt.Println("Heals", player_info.Data.Segments[0].Stats.Heals.Value)
	//fmt.Println("Revives", player_info.Data.Segments[0].Stats.Revive.Value)
	//fmt.Println("The last 3 matches ID")
	//fmt.Println(match_list.Data.Matches[0].Attributes.ID)
	//fmt.Println(match_list.Data.Matches[1].Attributes.ID)
	//fmt.Println(match_list.Data.Matches[2].Attributes.ID)
}

func requestGetTRN(url string, display_name string) ([]byte, error) {
	method := "GET"
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//Proxy:           http.ProxyURL(proxyUrl),
	}
	header := map[string]string{
		"Content-Type":    "application/json",
		"Accept-Encoding": "gzip",
		"User-Agent":      "Tracker Network App/3.22.9",
		"x-app-version":   "3.22.9",
	}
	data, err := netreq.Request{
		Method:    method,
		Header:    header,
		URL:       url + display_name,
		Transport: transport}.ReadRespBodyByte()
	if err != nil {
		return nil, err
	}
	slog.Info(
		"Request",
		"url", url+display_name,
	)
	return data, nil
}

func get_player_info(display_name string) (*player_info, error) {
	player_info := player_info{}
	url := "https://api.tracker.gg/api/v2/bf1/standard/profile/origin/"
	data, err := requestGetTRN(url, display_name)
	if err != nil {
		slog.Error("Failed to parse ")
		return nil, err
	}

	json.Unmarshal(data, &player_info)
	return &player_info, err
}

func get_match_list(display_name string) (*match_list, error) {
	match_list := match_list{}
	url := "http://api.tracker.gg/api/v2/bf1/standard/matches/origin/"
	data, err := requestGetTRN(url, display_name)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &match_list)
	return &match_list, nil
}

func get_match_detail(match_id string) (*match_detail, error) {
	match_detail := match_detail{}
	url := "http://api.tracker.gg/api/v2/bf1/standard/matches/" + match_id
	data, err := requestGetTRN(url, match_id)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &match_detail)
	return &match_detail, nil
}
