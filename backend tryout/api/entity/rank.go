package entity

type RankFormat struct {
	Rank        int `json:"rank"`
	Participant int `json:"participant"`
}

type Rank struct {
	Public RankFormat `json:"public"`
	Pref1  RankFormat `json:"pref1"`
	Pref2  RankFormat `json:"pref2"`
}
