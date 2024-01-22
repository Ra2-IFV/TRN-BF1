package trn

import (
	"crypto/tls"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/Ra2-IFV/TRN-BF1/pkg/netreq"
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

func requestGetTRN(url string, display_name string) ([]byte, error) {
	method := "GET"
	header := map[string]string{
		"Content-Type":    "application/json",
		"Accept-Encoding": "gzip",
		"User-Agent":      "Tracker Network App/3.22.9",
		"x-app-version":   "3.22.9",
	}
	transport := &http.Transport{
		// Disable http2 to bypass Cloudflare
		//TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{},
		//Proxy:           http.ProxyURL(proxyUrl),
	}
	slog.Info(
		"Request",
		"url", url+display_name,
	)
	data, err := netreq.Request{
		Method:    method,
		Header:    header,
		URL:       url + display_name,
		Transport: transport,
	}.ReadRespBodyByte()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetPlayerInfo(display_name string) (player_info, error) {
	playerInfo := player_info{}
	url := "https://api.tracker.gg/api/v2/bf1/standard/profile/origin/"
	data, err := requestGetTRN(url, display_name)
	if err != nil {
		return player_info{}, err
	}
	//fmt.Println(string(data))
	json.Unmarshal(data, &playerInfo)
	return playerInfo, err
}

func GetMatchList(display_name string) (match_list, error) {
	matchList := match_list{}
	url := "https://api.tracker.gg/api/v2/bf1/standard/matches/origin/"
	data, err := requestGetTRN(url, display_name)
	if err != nil {
		return match_list{}, err
	}

	json.Unmarshal(data, &matchList)
	return matchList, nil
}

func GetMatchDetail(match_id string) (match_detail, error) {
	matchDetail := match_detail{}
	url := "https://api.tracker.gg/api/v2/bf1/standard/matches/" + match_id
	data, err := requestGetTRN(url, match_id)
	if err != nil {
		return match_detail{}, err
	}

	json.Unmarshal(data, &matchDetail)
	return matchDetail, nil
}
