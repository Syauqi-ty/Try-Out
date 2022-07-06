package entity

type Score struct {
	ID        uint64  `json:"id" gorm:"primary_key;auto_increment"`
	StudentID uint64  `json:"student_id"`
	TestID    uint64  `json:"test_id"`
	Student   Student `json:"student"`
	Test      Test    `json:"test"`
	Type      string  `json:"type"`
	X1        bool    `json:"x1"`
	X2        bool    `json:"x2"`
	X3        bool    `json:"x3"`
	X4        bool    `json:"x4"`
	X5        bool    `json:"x5"`
	X6        bool    `json:"x6"`
	X7        bool    `json:"x7"`
	X8        bool    `json:"x8"`
	X9        bool    `json:"x9"`
	X10       bool    `json:"x10"`
	X11       bool    `json:"x11"`
	X12       bool    `json:"x12"`
	X13       bool    `json:"x13"`
	X14       bool    `json:"x14"`
	X15       bool    `json:"x15"`
	X16       bool    `json:"x16"`
	X17       bool    `json:"x17"`
	X18       bool    `json:"x18"`
	X19       bool    `json:"x19"`
	X20       bool    `json:"x20"`
	Score     float64 `json:"score"`
}

type ScoreMin struct {
	ID        uint64  `json:"id" gorm:"primary_key;auto_increment"`
	StudentID uint64  `json:"student_id"`
	TestID    uint64  `json:"test_id"`
	Type      string  `json:"type"`
	X1        bool    `json:"x1"`
	X2        bool    `json:"x2"`
	X3        bool    `json:"x3"`
	X4        bool    `json:"x4"`
	X5        bool    `json:"x5"`
	X6        bool    `json:"x6"`
	X7        bool    `json:"x7"`
	X8        bool    `json:"x8"`
	X9        bool    `json:"x9"`
	X10       bool    `json:"x10"`
	X11       bool    `json:"x11"`
	X12       bool    `json:"x12"`
	X13       bool    `json:"x13"`
	X14       bool    `json:"x14"`
	X15       bool    `json:"x15"`
	X16       bool    `json:"x16"`
	X17       bool    `json:"x17"`
	X18       bool    `json:"x18"`
	X19       bool    `json:"x19"`
	X20       bool    `json:"x20"`
	Score     float64 `json:"score"`
}

type ScoreOnly struct {
	ID        uint64  `json:"id" gorm:"primary_key;auto_increment"`
	StudentID uint64  `json:"student_id"`
	TestID    uint64  `json:"test_id"`
	Type      string  `json:"type"`
	Score     float64 `json:"score"`
}
type ScoreReal struct {
	ID        uint64  `json:"id" gorm:"primary_key;auto_increment"`
	StudentID uint64  `json:"student_id"`
	TestID    uint64  `json:"test_id"`
	Score     float64 `json:"score"`
}

type Subtests struct {
	Pu  float64 `json:"pu"`
	Ppu float64 `json:"ppu"`
	Pk  float64 `json:"pk"`
	Pmm float64 `json:"pmm"`
	Eng float64 `json:"eng"`
}
type SubtestsTKA struct {
	Pu  float64 `json:"pu"`
	Ppu float64 `json:"ppu"`
	Pk  float64 `json:"pk"`
	Pmm float64 `json:"pmm"`
	Eng float64 `json:"eng"`
	Ma  float64 `json:"ma"`
	Fi  float64 `json:"fi"`
	Ki  float64 `json:"ki"`
	Bi  float64 `json:"bi"`
	Sos float64 `json:"sos"`
	Sej float64 `json:"sej"`
	Eko float64 `json:"eko"`
	Geo float64 `json:"geo"`
}
type SubtestsSaintek struct {
	Ma float64 `json:"ma"`
	Fi float64 `json:"fi"`
	Ki float64 `json:"ki"`
	Bi float64 `json:"bi"`
}
type SubtestsSoshum struct {
	Sos float64 `json:"sos"`
	Sej float64 `json:"sej"`
	Eko float64 `json:"eko"`
	Geo float64 `json:"geo"`
}
type QSubtests struct {
	Pu  []float64 `json:"pu"`
	Ppu []float64 `json:"ppu"`
	Pk  []float64 `json:"pk"`
	Pmm []float64 `json:"pmm"`
	Eng []float64 `json:"eng"`
}

type RaraSaintek struct {
	Subject   string  `json:"subject"`
	Score     float64 `json:"score"`
	Avg       float64 `json:"avg"`
	AvgSchool float64 `json:"avg_school"`
	Max       float64 `json:"max"`
	MaxSchool float64 `json:"max_school"`
}
type RaraSoshum struct {
	Subject   string  `json:"subject"`
	Score     float64 `json:"score"`
	Avg       float64 `json:"avg"`
	AvgSchool float64 `json:"avg_school"`
	Max       float64 `json:"max"`
	MaxSchool float64 `json:"max_school"`
}
type RaraTPS struct {
	Subject   string  `json:"subject"`
	Score     float64 `json:"score"`
	Avg       float64 `json:"avg"`
	AvgSchool float64 `json:"avg_school"`
	Max       float64 `json:"max"`
	MaxSchool float64 `json:"max_school"`
}
type Radar struct {
	Name    string        `json:"name"`
	Tps     []RaraTPS     `json:"tps"`
	Saintek []RaraSaintek `json:"saintek"`
	Soshum  []RaraSoshum  `json:"soshum"`
}
type RangeTps struct {
	Range string `json:"range"`
	Avg   int64  `json:"avg"`
	Pu    int64  `json:"pu"`
	Ppu   int64  `json:"ppu"`
	Pmm   int64  `json:"pk"`
	Pk    int64  `json:"pmm"`
	Eng   int64  `json:"eng"`
}
type RangeSaintek struct {
	Range string `json:"range"`
	Avg   int64  `json:"avg"`
	Ma    int64  `json:"ma"`
	Fi    int64  `json:"fi"`
	Ki    int64  `json:"ki"`
	Bi    int64  `json:"bi"`
}
type RangeSoshum struct {
	Range string `json:"range"`
	Avg   int64  `json:"avg"`
	Sos   int64  `json:"sos"`
	Sej   int64  `json:"sej"`
	Eko   int64  `json:"eko"`
	Geo   int64  `json:"geo"`
}
type Distribution struct {
	Tps     []RangeTps     `json:"tps"`
	Saintek []RangeSaintek `json:"saintek"`
	Soshum  []RangeSoshum  `json:"soshum"`
}
type NilaiTPS struct {
	Pu  []float64 `json:"pu"`
	Ppu []float64 `json:"ppu"`
	Pmm []float64 `json:"pmm"`
	Pk  []float64 `json:"pk"`
	Eng []float64 `json:"eng"`
}
type QSubtestsSaintek struct {
	Subject string  `json:"subject"`
	Q0      float64 `json:"q0"`
	Q1      float64 `json:"q1"`
	Q2      float64 `json:"q2"`
	Q3      float64 `json:"q3"`
	Q4      float64 `json:"q4"`
}
type QSubtestsSoshum struct {
	Subject string  `json:"subject"`
	Q0      float64 `json:"q0"`
	Q1      float64 `json:"q1"`
	Q2      float64 `json:"q2"`
	Q3      float64 `json:"q3"`
	Q4      float64 `json:"q4"`
}
type QSubtestsTps struct {
	Subject string  `json:"subject"`
	Q0      float64 `json:"q0"`
	Q1      float64 `json:"q1"`
	Q2      float64 `json:"q2"`
	Q3      float64 `json:"q3"`
	Q4      float64 `json:"q4"`
}

type Quartile struct {
	Tps     []QSubtestsTps     `json:"tps"`
	Saintek []QSubtestsSaintek `json:"saintek"`
	Soshum  []QSubtestsSoshum  `json:"soshum"`
}

type FScore struct {
	Test   TestBare `json:"test" gorm:"foreignKey:TestID;references:ID;"`
	Accum  float64  `json:"accum"`
	Avg    float64  `json:"avg"`
	Scores Subtests `json:"scores"`
}
type Leaderboard struct {
	Name      string  `json:"name"`
	Peringkat int     `json:"peringkat"`
	Pu        float64 `json:"pu"`
	Ppu       float64 `json:"ppu"`
	Pk        float64 `json:"pk"`
	Pmm       float64 `json:"pmm"`
	Eng       float64 `json:"eng"`
	Ma        float64 `json:"ma"`
	Fi        float64 `json:"fi"`
	Ki        float64 `json:"ki"`
	Bi        float64 `json:"bi"`
	Sos       float64 `json:"sos"`
	Sej       float64 `json:"sej"`
	Geo       float64 `json:"geo"`
	Eko       float64 `json:"eko"`
	Tps       float64 `json:"tps"`
	Saintek   float64 `json:"saintek"`
	Soshum    float64 `json:"soshum"`
}

type FScoreMin struct {
	Accum  float64  `json:"accum"`
	Avg    float64  `json:"avg"`
	Scores Subtests `json:"scores"`
}
type FScoreMinSchool struct {
	Accum  float64  `json:"accum"`
	Avg    float64  `json:"avg"`
	Scores Subtests `json:"scores"`
}
type ScoreNa struct {
	Accum  float64         `json:"accum"`
	Avg    float64         `json:"avg"`
	Scores SubtestsSaintek `json:"scores"`
}
type ScoreNaSchool struct {
	Accum  float64         `json:"accum"`
	Avg    float64         `json:"avg"`
	Scores SubtestsSaintek `json:"scores"`
}
type ScoreSos struct {
	Accum  float64        `json:"accum"`
	Avg    float64        `json:"avg"`
	Scores SubtestsSoshum `json:"scores"`
}
type ScoreSosSchool struct {
	Accum  float64        `json:"accum"`
	Avg    float64        `json:"avg"`
	Scores SubtestsSoshum `json:"scores"`
}
type FQuartile struct {
	Quartile []float64 `json:"quartile"`
	Scores   QSubtests `json:"scores"`
}

type UserScoreStat struct {
	Avg    FScoreMin `json:"avg"`
	Max    FScoreMin `json:"max"`
	Latest FScoreMin `json:"latest"`
}

type TestStats struct {
	Max      FScoreMin `json:"max"`
	Avg      FScoreMin `json:"avg"`
	Quartile FQuartile `json:"quartile"`
}
type Rerata struct {
	Tps     float64 `json:"tps"`
	Saintek float64 `json:"saintek"`
	Soshum  float64 `json:"soshum"`
}
type High struct {
	Tps        float64 `json:"tps"`
	Saintek    float64 `json:"saintek"`
	Soshum     float64 `json:"soshum"`
	Battlename string  `json:"battle_name"`
}
type IkutBattle struct {
	Name  string `json:"name" gorm:"foreignKey:TestID;references:ID;"`
	Score Rerata `json:"score"`
	Types int    `json:"types"`
}

type HighScore struct {
	Name   string `json:"name"`
	Target int    `json:"target"`
	Score  Rerata `json:"score"`
}

type HighScoreBattle struct {
	Name   string `json:"name"`
	Target int    `json:"target"`
	Score  High   `json:"score"`
}
type DataTps struct {
	Name         string  `json:"name"`
	Score        float64 `json:"score"`
	AvgSchool    float64 `json:"avg_school"`
	AvgNational  float64 `json:"avg_national"`
	MaxAvg       float64 `json:"max_avg"`
	MaxAvgSchool float64 `json:"max_avgschool"`
}
type DataSaintek struct {
	Name         string  `json:"name"`
	Score        float64 `json:"score"`
	AvgSchool    float64 `json:"avg_school"`
	AvgNational  float64 `json:"avg_national"`
	MaxAvg       float64 `json:"max_avg"`
	MaxAvgSchool float64 `json:"max_avgschool"`
}
type DataSoshum struct {
	Name         string  `json:"name"`
	Score        float64 `json:"score"`
	AvgSchool    float64 `json:"avg_school"`
	AvgNational  float64 `json:"avg_national"`
	MaxAvg       float64 `json:"max_avg"`
	MaxAvgSchool float64 `json:"max_avgschool"`
}
type ScoreGraph struct {
	Tps     []DataTps     `json:"tps"`
	Saintek []DataSaintek `json:"saintek"`
	Soshum  []DataSoshum  `json:"soshum"`
}
type BenarSalah struct {
	X1  int `json:"x1"`
	X2  int `json:"x2"`
	X3  int `json:"x3"`
	X4  int `json:"x4"`
	X5  int `json:"x5"`
	X6  int `json:"x6"`
	X7  int `json:"x7"`
	X8  int `json:"x8"`
	X9  int `json:"x9"`
	X10 int `json:"x10"`
	X11 int `json:"x11"`
	X12 int `json:"x12"`
	X13 int `json:"x13"`
	X14 int `json:"x14"`
	X15 int `json:"x15"`
	X16 int `json:"x16"`
	X17 int `json:"x17"`
	X18 int `json:"x18"`
	X19 int `json:"x19"`
	X20 int `json:"x20"`
}
type Nilai struct {
	Pu  float64 `json:"pu"`
	Ppu float64 `json:"ppu"`
	Pk  float64 `json:"pk"`
	Pmm float64 `json:"pmm"`
	Eng float64 `json:"eng"`
	Ma  float64 `json:"ma"`
	Fi  float64 `json:"fi"`
	Ki  float64 `json:"ki"`
	Bi  float64 `json:"bi"`
	Sos float64 `json:"sos"`
	Sej float64 `json:"sej"`
	Geo float64 `json:"geo"`
	Eko float64 `json:"eko"`
}

type CekPesan struct {
	Subject string `json:"subject"`
	Pesan   string `json:"pesan"`
}
