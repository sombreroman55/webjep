package main

type Clue struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Value    int    `json:"value"`
	Answered bool   `json:"answered"`
}

type Category struct {
	Title string  `json:"title"`
	Clues [5]Clue `json:"clues"`
}

type Round struct {
	Categories [6]Category `json:"categories"`
}

type FinalRound struct {
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Game struct {
	Rounds     [2]Round   `json:"rounds"`
	Finalround FinalRound `json:"finalRound"`
}

type Player struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
