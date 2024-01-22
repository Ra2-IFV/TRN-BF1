package main

import (
	"fmt"

	"github.com/Ra2-IFV/TRN-BF1/pkg/bf1/trn"
)

type Config struct {
	Proxy string `yaml:"proxy"`
}

func init() {}

func main() {
	display_name := "404-HK416-MOD3"
	playerInfo, _ := trn.GetPlayerInfo(display_name)
	//match_list, _ := get_match_list(display_name)
	fmt.Println("Player name", playerInfo)
	fmt.Println("Rank", playerInfo.Data.Segments[0].Stats.Rank.Value)
	//fmt.Println("KD", playerInfo.Data.Segments[0].Stats.KdRatio.Value)
	//fmt.Println("KPM", playerInfo.Data.Segments[0].Stats.KillsPerMinute.Value)
	//fmt.Println("Heals", playerInfo.Data.Segments[0].Stats.Heals.Value)
	//fmt.Println("Revives", playerInfo.Data.Segments[0].Stats.Revive.Value)
	//fmt.Println("The last 3 matches ID")
	//fmt.Println(match_list.Data.Matches[0].Attributes.ID)
	//fmt.Println(match_list.Data.Matches[1].Attributes.ID)
	//fmt.Println(match_list.Data.Matches[2].Attributes.ID)
}
