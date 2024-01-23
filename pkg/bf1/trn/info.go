package trn

import (
	"encoding/json"
	"log/slog"
)

type playerInfoStruct struct {
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
					Value uint32 `json:"value"`
				} `json:"timePlayed"`
				Kills struct {
					Value uint32 `json:"value"`
				} `json:"kills"`
				Deaths struct {
					Value uint32 `json:"value"`
				} `json:"deaths"`
				LongestHeadshot struct {
					DisplayValue string `json:"displayValue"`
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
					DisplayValue string `json:"displayValue"`
				} `json:"kdRatio"`
				ScorePerMinute struct {
					DisplayValue string `json:"displayValue"`
				} `json:"scorePerMinute"`
				KillsPerMinute struct {
					DisplayValue string `json:"displayValue"`
				} `json:"killsPerMinute"`
				WinPercentage struct {
					DisplayValue string `json:"displayValue"`
				} `json:"winPercentage"`
				ShotsAccuracy struct {
					DisplayValuedisplay string `json:"displayvalue"`
				} `json:"shotsAccuracy"`
				HeadshotsPercentage struct {
					DisplayValue string `json:"displayvalue"`
				} `json:"headshotsPercentage"`
			} `json:"stats"`
		} `json:"segments"`
		ExpiryDate string `json:"expiryDate"`
	}
}

func GetPlayerInfo(displayName string) (playerInfoStruct, error) {
	playerInfoData := playerInfoStruct{}
	url := "https://api.tracker.gg/api/v2/bf1/standard/profile/origin/"
	data, err := requestGetTRN(url, displayName)
	if err != nil {
		slog.Warn("Request TRN player info failed", "error", err)
		return playerInfoStruct{}, err
	}
	//fmt.Println(string(data))
	json.Unmarshal(data, &playerInfoData)
	return playerInfoData, err
}
