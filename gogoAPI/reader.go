package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"strconv"
)

import loggly "github.com/jamespearly/loggly"

type Standings struct {
	Standings []Standing `json:"standings"`
}

type Standing struct {
	Table []Team `json:"table"`
}

type Team struct {
	Position int `json:"position"`
	Info Info `json:"team"`
	PlayedGames int`json:"playedGames"`
	Wins int `json:"won"`
	Draws int `json:"draw"`
	Losses int `json:"lost"`
	Points int `json:"points"`
	GoalsFor int `json:"goalsFor"`
	GoalsAgainst int `json:"goalsAgainst"`
	GoalDifference int `json:"goalDifference"`
}

type Info struct {
	Name string `json:"name"`
}

func main() {
	var keyword = "Passer"
	sender := loggly.New(keyword)
	url := "http://api.football-data.org/v2/competitions/PL/standings"
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	if(err != nil) {
		sender.EchoSend("error", "Problem with the API Get, " + err.Error())
	}

	request.Header.Set("X-Auth-Token", "78b11f075e39478ca4d95729a8da7271")
	answer, _ := client.Do(request)
	
	byteValue, _ := ioutil.ReadAll(answer.Body)
	var standings Standings
	var errer = json.Unmarshal(byteValue, &standings)

	if(errer != nil) {
		sender.Send("error", "Theres an error")
	}

	sender.EchoSend("info", "Everything worked")

	for i := 0; i < len(standings.Standings); i++ {
		for j := 0; j < len(standings.Standings[i].Table); j++ {	
			fmt.Println("Team: " + standings.Standings[i].Table[j].Info.Name)
			fmt.Println("Position: " + strconv.Itoa(standings.Standings[i].Table[j].Position))
			fmt.Println("Points: " + strconv.Itoa(standings.Standings[i].Table[j].Points))
		}
	}
}