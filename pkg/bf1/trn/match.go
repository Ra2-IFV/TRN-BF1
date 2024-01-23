package trn

import (
	"encoding/json"
	"log/slog"
	"time"
)

type matchListStruct struct {
	Data struct {
		Matches []struct {
			Attributes struct {
				ID string `json:"id"`
			} `json:"attributes"`
		} `json:"matches"`
	} `json:"data"`
}

type matchDetailStruct struct {
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

func GetMatchList(displayName string) (matchListStruct, error) {
	matchListData := matchListStruct{}
	url := "https://api.tracker.gg/api/v2/bf1/standard/matches/origin/"
	data, err := requestGetTRN(url, displayName)
	if err != nil {
		slog.Warn("Request TRN match list failed", "error", err)
		return matchListStruct{}, err
	}

	json.Unmarshal(data, &matchListData)
	return matchListData, nil
}

func GetMatchDetail(matchId string) (matchDetailStruct, error) {
	matchDetailData := matchDetailStruct{}
	url := "https://api.tracker.gg/api/v2/bf1/standard/matches/" + matchId
	data, err := requestGetTRN(url, matchId)
	if err != nil {
		slog.Warn("Request TRN match detail failed", "error", err)
		return matchDetailStruct{}, err
	}

	json.Unmarshal(data, &matchDetailData)
	return matchDetailData, nil
}
