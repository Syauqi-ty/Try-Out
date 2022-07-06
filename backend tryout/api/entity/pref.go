package entity

type Uni struct {
	ID    uint64  `json:"id" gorm:"primary_key;auto_increment"`
	UniID uint64  `json:"uni_id"`
	Name  string  `json:"uni_name"`
	Sci   []Prodi `json:"sci" gorm:"foreignKey:UniID;references:UniID"`
	Soc   []Prodi `json:"soc" gorm:"foreignKey:UniID;references:UniID"`
}

type FUni struct {
	ID      uint64 `json:"id"`
	UniID   uint64 `json:"uni_id"`
	Name    string `json:"uni_name"`
	Studies SciSoc `json:"studies"`
}

type UniMin struct {
	ID    uint64 `json:"id"`
	UniID uint64 `json:"uni_id"`
	Name  string `json:"uni_name"`
}

type Prodi struct {
	ID      uint64 `json:"-" gorm:"primary_key;auto_increment"`
	UniID   uint64 `json:"-"`
	ProdiID uint64 `json:"prodi_id"`
	Name    string `json:"prodi_name"`
	Type    string `json:"type"`
}

type SciSoc struct {
	Sci []Prodi `json:"sci"`
	Soc []Prodi `json:"soc"`
}

type SetPrefBody struct {
	StudentID  int `json:"-"`
	TestID     int `json:"test_id"`
	Pref1Uni   int `json:"pref1_uni"`
	Pref2Uni   int `json:"pref2_uni"`
	Pref1Prodi int `json:"pref1_prodi"`
	Pref2Prodi int `json:"pref2_prodi"`
}
