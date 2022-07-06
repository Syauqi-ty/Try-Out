package entity

import (
	"time"
)

type Test struct {
	ID          uint64     `json:"id" gorm:"primary_key;auto_increment"`
	Slug        string     `json:"slug"`
	Name        string     `json:"name"`
	Image       string     `json:"image"`
	Price       int        `json:"price"`
	Types       int        `json:"types"`
	Ma          int        `json:"ma"`
	Fi          int        `json:"fi"`
	Ki          int        `json:"ki"`
	Bi          int        `json:"bi"`
	Pu          int        `json:"pu"`
	Ppu         int        `json:"ppu"`
	Pmm         int        `json:"pmm"`
	Pk          int        `json:"pk"`
	Eng         int        `json:"eng"`
	Sos         int        `json:"sos"`
	Sej         int        `json:"sej"`
	Eko         int        `json:"eko"`
	Geo         int        `json:"geo"`
	Questions   []Question `json:"questions" gorm:"many2many:test_question"`
	ScheduledAt time.Time  `json:"scheduledAt" gorm:"column:scheduledAt"`
	EndsAt      time.Time  `json:"endsAt" gorm:"column:endsAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt" gorm:"column:updatedAt"`
}

type TestQuestionMin struct {
	ID          uint64        `json:"id" gorm:"primary_key;auto_increment"`
	Slug        string        `json:"slug"`
	Name        string        `json:"name"`
	Image       string        `json:"image"`
	Price       int           `json:"price"`
	Types       int        `json:"types"`
	Ma          int           `json:"ma"`
	Fi          int           `json:"fi"`
	Ki          int           `json:"ki"`
	Bi          int           `json:"bi"`
	Pu          int           `json:"pu"`
	Ppu         int           `json:"ppu"`
	Pmm         int           `json:"pmm"`
	Pk          int           `json:"pk"`
	Eng         int           `json:"eng"`
	Sos         int           `json:"sos"`
	Sej         int           `json:"sej"`
	Eko         int           `json:"eko"`
	Geo         int           `json:"geo"`
	Questions   []QuestionMin `json:"questions" gorm:"many2many:test_question;foreignKey:ID;joinForeignKey:test_id;References:ID;JoinReferences:question_id"`
	ScheduledAt time.Time     `json:"scheduledAt" gorm:"column:scheduledAt"`
	EndsAt      time.Time     `json:"endsAt" gorm:"column:endsAt"`
	CreatedAt   time.Time     `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt" gorm:"column:updatedAt"`
}
type AvaibleBattle struct {
	Name        string `json:"name"`
	BattleID 	uint64 `json:"battle_id"`
	Image       string `json:"image"`
	ScheduledAt string `json:"scheduledAt"`
	Durasi float64 `json:"durasi"`
	Jumlahsoal  int    `json:"jumlah"`
	Waktu       string `json:"waktu"`
	Types       int    `json:"types"`
	Dibeli      bool   `json:"dibeli"`
}

type TestQuestionAndSolution struct {
	ID          uint64                `json:"id" gorm:"primary_key;auto_increment"`
	Slug        string                `json:"slug"`
	Name        string                `json:"name"`
	Image       string                `json:"image"`
	Price       int                   `json:"price"`
	Types       int        			  `json:"types"`
	Ma          int                   `json:"ma"`
	Fi          int                   `json:"fi"`
	Ki          int                   `json:"ki"`
	Bi          int                   `json:"bi"`
	Pu          int                   `json:"pu"`
	Ppu         int                   `json:"ppu"`
	Pmm         int                   `json:"pmm"`
	Pk          int                   `json:"pk"`
	Eng         int                   `json:"eng"`
	Sos         int                   `json:"sos"`
	Sej         int                   `json:"sej"`
	Eko         int                   `json:"eko"`
	Geo         int                   `json:"geo"`
	Questions   []QuestionAndSolution `json:"questions" gorm:"many2many:test_question;foreignKey:ID;joinForeignKey:test_id;References:ID;JoinReferences:question_id"`
	ScheduledAt time.Time             `json:"scheduledAt" gorm:"column:scheduledAt"`
	EndsAt      time.Time             `json:"endsAt" gorm:"column:endsAt"`
	CreatedAt   time.Time             `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time             `json:"updatedAt" gorm:"column:updatedAt"`
}

type TestMin struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Price       int       `json:"price"`
	Types       int       `json:"types"`
	Ma          int       `json:"ma"`
	Fi          int       `json:"fi"`
	Ki          int       `json:"ki"`
	Bi          int       `json:"bi"`
	Pu          int       `json:"pu"`
	Ppu         int       `json:"ppu"`
	Pmm         int       `json:"pmm"`
	Pk          int       `json:"pk"`
	Eng         int       `json:"eng"`
	Sos         int       `json:"sos"`
	Sej         int       `json:"sej"`
	Eko         int       `json:"eko"`
	Geo         int       `json:"geo"`
	ScheduledAt time.Time `json:"scheduledAt" gorm:"column:scheduledAt"`
	EndsAt      time.Time `json:"endsAt" gorm:"column:endsAt"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
type AvaibleTest struct {
	ID 			uint64    `json:"battle_id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Types       int       `json:"types"`
	Ma          int       `json:"ma"`
	Fi          int       `json:"fi"`
	Ki          int       `json:"ki"`
	Bi          int       `json:"bi"`
	Pu          int       `json:"pu"`
	Ppu         int       `json:"ppu"`
	Pmm         int       `json:"pmm"`
	Pk          int       `json:"pk"`
	Eng         int       `json:"eng"`
	Sos         int       `json:"sos"`
	Sej         int       `json:"sej"`
	Eko         int       `json:"eko"`
	Geo         int       `json:"geo"`
	ScheduledAt time.Time `json:"scheduledAt" gorm:"column:scheduledAt"`
	EndsAt      time.Time `json:"endsAt" gorm:"column:endsAt"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}
type NameDibeli struct {
	Name   string `json:"name"`
	Dibeli bool   `json:"dibeli"`
}
type NameType struct {
	Name  string `json:"name"`
	Types int    `json:"types"`
}
type BattleIkut struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

type TestBare struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	ScheduledAt time.Time `json:"scheduledAt" gorm:"column:scheduledAt"`
	EndsAt      time.Time `json:"endsAt" gorm:"column:endsAt"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	Types 		int 	  `json:"types"`
}

type TestAuth struct {
	ID         uint64 `json:"id" gorm:"primary_key;auto_increment"`
	TestID     uint64 `json:"test_id"`
	StudentID  uint64 `json:"student_id"`
	UsernameTO string `json:"username_to"`
	PasswordTO string `json:"password_to"`
	Pref1Uni   uint64 `json:"pref_1_uni"`
	Pref2Uni   uint64 `json:"pref_2_uni"`
	Pref1Prodi uint64 `json:"pref_1_prodi"`
	Pref2Prodi uint64 `json:"pref_2_prodi"`
}

type Last struct {
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	ScheduledAt time.Time `json:"scheduledAt"`
}
type IkutTest struct {
	Name     string `json:"name"`
	BattleID uint64 `json:"battle_id"`
}

type LoginSpace struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DataBattle struct {
	Name string `json:"name"`
	Image string `json:"image"`
	StartedAt time.Time `json:"started_at"`
	EndsAt time.Time `json:"ends_at"`
	Subtest []string `json:"subtest"`
}

type ResponLogin struct{
	UserID int `json:"user_id"`
	TestID int `json:"test_id"`
	Name string `json:"name"`
	DataTest DataBattle `json:"test"`
}
type Pertanyaan struct{
	Urutan int `json:"no"`
	Tests string `json:"soal"`
	Dijawab bool `json:"dijawab"`
	Jawab string `json:"jawabanbenar"`
	Jawaban string `json:"jawaban"`
}
type QuestionSpace struct {
	Type string `json:"type"`
	Waktu int `json:"waktu"`
	Tests []Pertanyaan `json:"soal"`
}

type Jawaban struct {
	UserID int `json:"user_id"`
	TestID int `json:"test_id"`
	Type string `json:"type"`
	Urutan string `json:"urutan"`
	Jawab int8 `json:"jawaban"`
}