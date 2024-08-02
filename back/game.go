package main

import ()

type Clue struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Value    int
	Answered bool
}

type Category struct {
	Title string `json:"title"`
	Clues []Clue `json:"clues"`
}

type Round struct {
	Categories []Category `json:"categories"`
}

type FinalRound struct {
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Game struct {
	Rounds     []Round     `json:"rounds"`
	FinalRound *FinalRound `json:"finalRound"`
}

type GameSettings struct {
	PlayFinalJep         bool    `json:"playFinalJep"`
	BaseClueValue        int     `json:"baseClueValue"`
	InterClueInterval    float64 `json:"interClueInterval"`
	InterRoundMultiplier float64 `json:"InterRoundMultiplier"`
	NumRounds            int     `json:"numberOfRounds"`
	CategoriesPerRound   int     `json:"categoriesPerRound"`
	CluesPerCategory     int     `json:"cluesPerCategory"`
}

type Session struct {
	Game     Game         `json:"game"`
	Settings GameSettings `json:"settings"`
}

func computeClueValues(session *Session) {
	for r := range session.Game.Rounds {
		for c := range session.Game.Rounds[r].Categories {
			for q := range session.Game.Rounds[r].Categories[c].Clues {
				bcv := session.Settings.BaseClueValue
				irm := session.Settings.InterRoundMultiplier
				ici := session.Settings.InterClueInterval
				crbv := int(float64(bcv) + float64(r)*float64(bcv)*irm)
				ccv := int(float64(crbv) + float64(q)*float64(crbv)*ici)
				session.Game.Rounds[r].Categories[c].Clues[q].Value = ccv
				session.Game.Rounds[r].Categories[c].Clues[q].Answered = false
			}
		}
	}
}
