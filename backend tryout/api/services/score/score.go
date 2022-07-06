package service

import (
	"math"
	entity "studybuddy-backend-fast/api/entity"
	questionRepository "studybuddy-backend-fast/api/repository/question"
	repository "studybuddy-backend-fast/api/repository/score"
	studentRepository "studybuddy-backend-fast/api/repository/student"
	testRepository "studybuddy-backend-fast/api/repository/test"
	testAuthRepository "studybuddy-backend-fast/api/repository/testauth"
)

type ScoreService interface {
	// Basic CRUD
	CreateScore(score entity.Score) entity.Score
	FindAllScore() []entity.ScoreMin
	FindScoreByID(id int) entity.Score
	UpdateScore(score entity.Score) entity.Score
	DeleteScore(score entity.Score)
	BenarSalah(testID int,StudentID int) entity.Nilai

	// PANTAT
	AnswerTestQuestion(TestID int, QuestionID int, StudentID int, Answer string)
	CheckAnswers(TestID int, StudentID int) ScoreCheck
	GetRank(TestID int, StudentID int) entity.Rank
	LeaderBoard(TestID int) []entity.ScoreOnly
	TypeLeaderBoard(TestID int, subtest string) []entity.ScoreOnly
	Jawab(jawab entity.Jawaban) entity.Jawaban

	// FORMATTED
	StudentFScore(studentsID int) []entity.FScore
	StudentFScoreOfTest(studentsID int, testID int) entity.FScore
	StudentFScoreMin(studentsID int) []entity.FScoreMin
	StudentFScoreMinOfTest(studentsID int, testID int) entity.FScoreMin

	// CALCULATIONS
	CalculateStudentStats(id int) entity.UserScoreStat
	CalculateTestStats(studentID int) entity.TestStats
	Graphic(test entity.Test, score entity.Score) []entity.IkutBattle
	OutputAllAverage(studentsID int) []entity.IkutBattle
	Battled(id int) int64
	HighScore(id int) entity.HighScore
	HighScoreBattle(id int) entity.HighScoreBattle
	GraphicData(id int) entity.ScoreGraph
	AllQuartil(TestID int) entity.Quartile
	Distribution(TestID int) entity.Distribution
	Rara(studentsID int) entity.Radar
	NewLeaderBoard(TestID int, studentID int) []entity.Leaderboard
	Lider(studentsID int, testID int) entity.Leaderboard
}

type scoreService struct {
	repo         repository.ScoreRepo
	testRepo     testRepository.TestRepo
	studentRepo  studentRepository.StudentRepo
	questionRepo questionRepository.QuestionRepo
	testAuthRepo testAuthRepository.TestAuthRepo
}

func Round(input float64) float64 {
	if input == 0 {
		return 0.0
	} else {
		return math.Round(input*100)/100
	}
}

type ScoreCheck struct {
	Pu  []bool `json:"pu"`
	Ppu []bool `json:"ppu"`
	Pmm []bool `json:"pmm"`
	Pk  []bool `json:"pk"`
	Eng []bool `json:"eng"`
}

func NewScoreService(rep repository.ScoreRepo, trep testRepository.TestRepo, srep studentRepository.StudentRepo, qrep questionRepository.QuestionRepo, tarep testAuthRepository.TestAuthRepo) ScoreService {
	return &scoreService{rep, trep, srep, qrep, tarep}
}

///////////////////////
// PRACTICAL USECASE //
///////////////////////

func (s *scoreService) LeaderBoard(TestID int) []entity.ScoreOnly {
	return s.repo.GroupScoreByTestID(TestID)
}

func (s *scoreService) TypeLeaderBoard(TestID int, subtest string) []entity.ScoreOnly {
	return s.repo.GroupScoreByTestIDWithType(TestID, subtest)
}
func reverse(numbers []entity.ScoreReal) []entity.ScoreReal {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
func (s *scoreService) removeDuplicatesTest(elements []entity.ScoreReal) []entity.ScoreReal {
	encountered := map[entity.ScoreReal]bool{}
	result := []entity.ScoreReal{}
	for key, _ := range elements {
		if encountered[elements[key]] != true {
			encountered[elements[key]] = true
			result = append(result, elements[key])
		}

	}
	return result
}

func (s *scoreService) NewLeaderBoard(TestID int, studentID int) []entity.Leaderboard {
	var c []entity.ScoreReal
	var ld entity.Leaderboard
	var arrayld []entity.Leaderboard
	a := s.repo.Leaderboard(TestID)
	b := s.repo.TigaBawah(TestID)
	d := s.repo.ScoreAnak(TestID, studentID)
	for i := range d{
		if d[i].Score == 0{
			c = append(c, a...)
			c = append(c, reverse(b)...)
			c = append(c, d...)
			c = s.removeDuplicatesTest(c)
		}
	}
	c = append(c, a...)
	c = append(c, d...)
	c = append(c, reverse(b)...)
	c = s.removeDuplicatesTest(c)
	for i := range c {
		student := s.studentRepo.FindOneById(int(c[i].StudentID))
		rank := s.GetRank(TestID, int(c[i].StudentID))
		nilai := s.Lider(int(c[i].StudentID), TestID)
		ld.Name = student.Name
		ld.Peringkat = rank.Public.Rank
		ld.Pu = nilai.Pu
		ld.Ppu = nilai.Ppu
		ld.Pmm = nilai.Pmm
		ld.Pk = nilai.Pk
		ld.Eng = nilai.Eng
		ld.Ma = nilai.Ma
		ld.Fi = nilai.Fi
		ld.Ki = nilai.Ki
		ld.Bi = nilai.Bi
		ld.Sos = nilai.Sos
		ld.Sej = nilai.Sej
		ld.Geo = nilai.Geo
		ld.Eko = nilai.Eko
		ld.Tps = nilai.Tps
		ld.Saintek = nilai.Saintek
		ld.Soshum = nilai.Soshum
		arrayld = append(arrayld, ld)
	}
	return arrayld
}
func (s *scoreService) Lider(studentsID int, testID int) entity.Leaderboard {
	var subtests entity.SubtestsTKA
	var i entity.Leaderboard
	scores := s.repo.FindScoreMinByStudentsIDAndTestID(studentsID, testID)
	for _, score := range scores {
		switch subtest := score.Type; subtest {
		case "pu":
			subtests.Pu = score.Score
			i.Pu = score.Score
		case "ppu":
			subtests.Ppu = score.Score
			i.Ppu = score.Score
		case "pmm":
			subtests.Pmm = score.Score
			i.Pmm = score.Score
		case "pk":
			subtests.Pk = score.Score
			i.Pk = score.Score
		case "eng":
			subtests.Eng = score.Score
			i.Eng = score.Score
		case "ma":
			subtests.Ma = score.Score
			i.Ma = score.Score
		case "fi":
			subtests.Fi = score.Score
			i.Fi = score.Score
		case "ki":
			subtests.Ki = score.Score
			i.Ki = score.Score
		case "bi":
			subtests.Bi = score.Score
			i.Bi = score.Score
		case "sos":
			subtests.Sos = score.Score
			i.Sos = score.Score
		case "sej":
			subtests.Sej = score.Score
			i.Sej = score.Score
		case "geo":
			subtests.Geo = score.Score
			i.Geo = score.Score
		case "eko":
			subtests.Eko = score.Score
			i.Eko = score.Score
		}
	}
	i.Tps = (i.Pu + i.Ppu + i.Pmm + i.Pk + i.Eng) / float64(5)
	i.Saintek = (i.Ma + i.Fi + i.Ki + i.Bi) / float64(4)
	i.Soshum = (i.Sej + i.Eko + i.Geo + i.Sos) / float64(4)
	return i
}

func (s *scoreService) CheckAnswers(TestID int, StudentID int) ScoreCheck {
	var scoreCheck ScoreCheck
	scores := s.repo.FindScoreMinByStudentsIDAndTestID(StudentID, TestID)

	for _, score := range scores {
		switch s := score.Type; s {
		case "pu":
			scoreCheck.Pu = []bool{score.X1, score.X2, score.X3, score.X4, score.X5, score.X6, score.X7, score.X8, score.X9, score.X10, score.X11, score.X12, score.X13, score.X14, score.X15, score.X16, score.X17, score.X18, score.X19, score.X20}
		case "ppu":
			scoreCheck.Ppu = []bool{score.X1, score.X2, score.X3, score.X4, score.X5, score.X6, score.X7, score.X8, score.X9, score.X10, score.X11, score.X12, score.X13, score.X14, score.X15, score.X16, score.X17, score.X18, score.X19, score.X20}
		case "pmm":
			scoreCheck.Pmm = []bool{score.X1, score.X2, score.X3, score.X4, score.X5, score.X6, score.X7, score.X8, score.X9, score.X10, score.X11, score.X12, score.X13, score.X14, score.X15, score.X16, score.X17, score.X18, score.X19, score.X20}
		case "pk":
			scoreCheck.Pk = []bool{score.X1, score.X2, score.X3, score.X4, score.X5, score.X6, score.X7, score.X8, score.X9, score.X10, score.X11, score.X12, score.X13, score.X14, score.X15, score.X16, score.X17, score.X18, score.X19, score.X20}
		case "eng":
			scoreCheck.Eng = []bool{score.X1, score.X2, score.X3, score.X4, score.X5, score.X6, score.X7, score.X8, score.X9, score.X10, score.X11, score.X12, score.X13, score.X14, score.X15, score.X16, score.X17, score.X18, score.X19, score.X20}
		}
	}

	return scoreCheck
}

func (s *scoreService) AnswerTestQuestion(TestID int, QuestionID int, StudentID int, Answer string) {
	question := s.questionRepo.FindQuestionMinByID(QuestionID)
	tquestions := s.questionRepo.FindQuestionsOfTest(TestID)

	var (
		qtype   string = ""
		mark    string
		correct bool = false
		puc     int  = 0
		qid     int  = 0
		ppuc    int  = 0
		pmmc    int  = 0
		pkc     int  = 0
		engc    int  = 0
	)

	if Answer == question.Answer {
		correct = true
	}

	for _, one := range tquestions {
		switch subtest := one.Type; subtest {
		case "pu":
			puc++
		case "ppu":
			ppuc++
		case "pmm":
			pmmc++
		case "pk":
			pkc++
		case "eng":
			engc++
		}
		if int(one.ID) == int(question.ID) {
			qtype = one.Type
			break
		}
	}

	switch s := qtype; s {
	case "pu":
		qid = puc
	case "ppu":
		qid = ppuc
	case "pmm":
		qid = pmmc
	case "pk":
		qid = pkc
	case "eng":
		qid = engc
	}

	markmap := []string{"x1", "x2", "x3", "x4", "x5", "x6", "x7", "x8", "x9", "x10", "x11", "x12", "x13", "x14", "x15", "x16", "x17", "x18", "x19", "x20"}
	mark = markmap[qid-1]

	s.repo.AnswerQuestion(TestID, StudentID, qtype, mark, correct)
}

func (s *scoreService) GetStudentRankOfTest(testID int, studentID int) int {
	scores := s.repo.GroupScoreByTestIDAndStudentID(testID, studentID)
	for i, score := range scores {
		if score.StudentID == uint64(studentID) {
			return i + 1
		}
	}
	return 0
}

func (s *scoreService) GetStudentRankOfTestWithFilter(testID int, studentID int, school string) int {
	return 1
}

///////////////////////////
// FORMAT SCORE USECASES //
///////////////////////////

func (s *scoreService) StudentFScore(studentsID int) []entity.FScore {
	var fscores []entity.FScore
	finishedTests := s.repo.GroupScoreByStudentID(studentsID)

	for _, t := range finishedTests {
		fscore := s.StudentFScoreOfTest(studentsID, int(t.TestID))
		fscores = append(fscores, fscore)
	}

	return fscores
}

func (s *scoreService) StudentFScoreOfTest(studentsID int, testID int) entity.FScore {
	var fscore entity.FScore
	var subtests entity.Subtests

	scores := s.repo.FindScoreMinByStudentsIDAndTestID(studentsID, testID)
	test := s.testRepo.FindTestByIDBare(testID)

	for _, score := range scores {
		switch subtest := score.Type; subtest {
		case "pu":
			subtests.Pu = score.Score
		case "ppu":
			subtests.Ppu = score.Score
		case "pmm":
			subtests.Pmm = score.Score
		case "pk":
			subtests.Pk = score.Score
		case "eng":
			subtests.Eng = score.Score
		}
	}
	fscore.Test = test
	fscore.Accum = subtests.Pu + subtests.Ppu + subtests.Pmm + subtests.Pk + subtests.Eng
	fscore.Avg = fscore.Accum / float64(5)
	fscore.Scores = subtests

	return fscore
}
func (s *scoreService) Average(studentsID int, testID int) entity.IkutBattle {
	var subtests entity.SubtestsTKA
	var i entity.IkutBattle
	scores := s.repo.FindScoreMinByStudentsIDAndTestID(studentsID, testID)
	test := s.testRepo.NameandType(testID)
	for _, score := range scores {
		switch subtest := score.Type; subtest {
		case "pu":
			subtests.Pu = score.Score
		case "ppu":
			subtests.Ppu = score.Score
		case "pmm":
			subtests.Pmm = score.Score
		case "pk":
			subtests.Pk = score.Score
		case "eng":
			subtests.Eng = score.Score
		case "ma":
			subtests.Ma = score.Score
		case "fi":
			subtests.Fi = score.Score
		case "ki":
			subtests.Ki = score.Score
		case "bi":
			subtests.Bi = score.Score
		case "sos":
			subtests.Sos = score.Score
		case "sej":
			subtests.Sej = score.Score
		case "geo":
			subtests.Geo = score.Score
		case "eko":
			subtests.Eko = score.Score
		}
	}
	i.Name = test.Name
	i.Types = test.Types
	i.Score.Tps = (subtests.Pu + subtests.Ppu + subtests.Pmm + subtests.Pk + subtests.Eng) / float64(5)
	i.Score.Saintek = (subtests.Ma + subtests.Fi + subtests.Ki + subtests.Bi) / float64(4)
	i.Score.Soshum = (subtests.Sej + subtests.Eko + subtests.Geo + subtests.Sos) / float64(4)
	return i
}
func (s *scoreService) DataTps(studentsID int, testID int) entity.DataTps {
	var subtests entity.SubtestsTKA
	var i entity.DataTps
	var avg entity.FScoreMin
	var avgs entity.FScoreMinSchool
	scores := s.repo.FindScoreMinByStudentsIDAndTestID(studentsID, testID)
	test := s.testRepo.NameandType(testID)
	var school string
	siswa := s.studentRepo.FindOneById(studentsID)
	school = siswa.School

	for _, score := range scores {
		switch subtest := score.Type; subtest {
		case "pu":
			subtests.Pu = score.Score
		case "ppu":
			subtests.Ppu = score.Score
		case "pmm":
			subtests.Pmm = score.Score
		case "pk":
			subtests.Pk = score.Score
		case "eng":
			subtests.Eng = score.Score
		}
		i.Score = Round((subtests.Pu + subtests.Ppu + subtests.Pmm + subtests.Pk + subtests.Eng) / float64(5))
	}
	for _, score := range s.repo.GroupScoreByTestIDAndType(testID) {
		avg.Accum = avg.Accum + score.Score
		switch s := score.Type; s {
		case "pu":
			avg.Scores.Pu = score.Score
		case "ppu":
			avg.Scores.Ppu = score.Score
		case "pmm":
			avg.Scores.Pmm = score.Score
		case "pk":
			avg.Scores.Pk = score.Score
		case "eng":
			avg.Scores.Eng = score.Score
		}
		i.AvgNational = Round(avg.Accum / float64(5))
	}
	for _, score := range s.repo.RataSekolah(testID, school) {
		avgs.Accum = avgs.Accum + score.Score
		switch s := score.Type; s {
		case "pu":
			avgs.Scores.Pu = score.Score
		case "ppu":
			avgs.Scores.Ppu = score.Score
		case "pmm":
			avgs.Scores.Pmm = score.Score
		case "pk":
			avgs.Scores.Pk = score.Score
		case "eng":
			avgs.Scores.Eng = score.Score
		}
		i.AvgSchool = Round(avgs.Accum / float64(5))
	}
	for _,score := range s.repo.MaxScore(testID){
		i.MaxAvg = Round(score.Score)
	}
	for _,score := range s.repo.MaxScoreSchool(testID,school){
		i.MaxAvgSchool = Round(score.Score)
	}
	i.Name = test.Name
	return i
}
func (s *scoreService) DataSaintek(studentsID int, testID int) entity.DataSaintek {
	var subtests entity.SubtestsTKA
	var avg entity.ScoreNa
	var i entity.DataSaintek
	var avgs entity.ScoreNaSchool
	scores := s.repo.FindScoreMinByStudentsIDAndTestID(studentsID, testID)
	test := s.testRepo.NameandType(testID)
	var school string
	siswa := s.studentRepo.FindOneById(studentsID)
	school = siswa.School
	for _, score := range scores {
		switch subtest := score.Type; subtest {
		case "ma":
			subtests.Ma = score.Score
		case "fi":
			subtests.Fi = score.Score
		case "ki":
			subtests.Ki = score.Score
		case "bi":
			subtests.Bi = score.Score
		}
		i.Score = Round((subtests.Ma + subtests.Fi + subtests.Ki + subtests.Bi) / float64(4))
	}
	for _, score := range s.repo.GroupScoreByTestIDAndType(testID) {

		switch subtests := score.Type; subtests {
		case "ma":
			avg.Scores.Ma = score.Score
		case "fi":
			avg.Scores.Fi = score.Score
		case "ki":
			avg.Scores.Ki = score.Score
		case "bi":
			avg.Scores.Bi = score.Score
		}
		avg.Accum = avg.Accum + score.Score
		i.AvgNational = Round(avg.Accum / float64(4))
	}
	for _, score := range s.repo.RataSekolah(testID, school) {
		avgs.Accum = avgs.Accum + score.Score
		switch subtests := score.Type; subtests {
		case "ma":
			avgs.Scores.Ma = score.Score
		case "fi":
			avgs.Scores.Fi = score.Score
		case "ki":
			avgs.Scores.Ki = score.Score
		case "bi":
			avgs.Scores.Bi = score.Score
		}
		i.AvgSchool = Round(avgs.Accum / float64(4))
	}
	for _,score := range s.repo.MaxAllAverageSaintek(uint64(testID)){
		i.MaxAvg = Round(score.Score / float64(4))
	}
	for _,score := range s.repo.MaxAllAverageSchoolSaintek(uint64(testID),school){
		i.MaxAvgSchool = Round(score.Score /float64(4))
	}
	i.Name = test.Name
	return i
}
func (s *scoreService) DataSoshum(studentsID int, testID int) entity.DataSoshum {
	var subtests entity.SubtestsTKA
	var i entity.DataSoshum
	var avg entity.ScoreSos
	var avgs entity.ScoreSosSchool
	rata := s.repo.GroupScoreByTestIDAndType(testID)
	scores := s.repo.FindScoreMinByStudentsIDAndTestID(studentsID, testID)
	test := s.testRepo.NameandType(testID)
	var school string
	siswa := s.studentRepo.FindOneById(studentsID)
	school = siswa.School
	ratas := s.repo.RataSekolah(testID, school)
	for _, score := range scores {
		switch subtest := score.Type; subtest {
		case "sos":
			subtests.Sos = score.Score
		case "sej":
			subtests.Sej = score.Score
		case "geo":
			subtests.Geo = score.Score
		case "eko":
			subtests.Eko = score.Score
		}
		i.Score = (subtests.Sej + subtests.Eko + subtests.Geo + subtests.Sos) / float64(4)
	}
	for _, score := range  rata {
		switch subtests := score.Type; subtests {
		case "sos":
			avg.Scores.Sos = score.Score
		case "sej":
			avg.Scores.Sej = score.Score
		case "geo":
			avg.Scores.Geo = score.Score
		case "eko":
			avg.Scores.Eko = score.Score
		}
		avg.Accum = avg.Accum + score.Score
		i.AvgNational = Round(avg.Accum / float64(4))
	}
	for _, score := range  ratas{
		avgs.Accum = avgs.Accum + score.Score
		switch subtests := score.Type; subtests {
		case "sos":
			avgs.Scores.Sos = score.Score
		case "sej":
			avgs.Scores.Sej = score.Score
		case "geo":
			avgs.Scores.Geo = score.Score
		case "eko":
			avgs.Scores.Eko = score.Score
		}
		i.AvgSchool = avgs.Accum / float64(4)
	}
	for _,score := range s.repo.MaxAllAverageSoshum(uint64(testID)){
		i.MaxAvg = Round(score.Score / float64(4))
	}
	for _,score := range s.repo.MaxAllAverageSchoolSoshum(uint64(testID),school){
		i.MaxAvgSchool = Round(score.Score / float64(4))
	}
	i.Name = test.Name
	return i
}
func (s *scoreService) GraphicData(id int) entity.ScoreGraph {
	var i entity.ScoreGraph
	var tps []entity.DataTps
	var saintek []entity.DataSaintek
	var soshum []entity.DataSoshum
	finishedTests := s.repo.GroupScoreByStudentID(id)
	for _, t := range finishedTests {
		wibu := s.testRepo.NameandType2(t.TestID)
		if wibu.Types == 0 {
			fscore := s.DataTps(id, int(t.TestID))
			tps = append(tps, fscore)
			i.Tps = tps
		} else if wibu.Types == 1 {
			fscore := s.DataTps(id, int(t.TestID))
			fscoresaintek := s.DataSaintek(id, int(t.TestID))
			tps = append(tps, fscore)
			saintek = append(saintek, fscoresaintek)
			i.Saintek = saintek
			i.Tps = tps
		} else if wibu.Types == 2 {
			fscore := s.DataTps(id, int(t.TestID))
			fscoresoshum := s.DataSoshum(id, int(t.TestID))
			tps = append(tps, fscore)
			soshum = append(soshum, fscoresoshum)
			i.Tps = tps
			i.Soshum = soshum
		} else if wibu.Types == 3 {
			fscore := s.DataTps(id, int(t.TestID))
			fscoresaintek := s.DataSaintek(id, int(t.TestID))
			fscoresoshum := s.DataSoshum(id, int(t.TestID))
			tps = append(tps, fscore)
			saintek = append(saintek, fscoresaintek)
			soshum = append(soshum, fscoresoshum)
			i.Tps = tps
			i.Saintek = saintek
			i.Soshum = soshum
		}
	}

	return i
}
func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
func contain(s []bool, str bool) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
func (s *scoreService) Battled(id int) int64 {
	var nilai int64
	var you []int
	test := s.OutputAllAverage(id)
	for t := range test {
		if len(test) == 0 {
			nilai = 0
		} else if len(test) == 1 {
			if test[t].Types == 1 {
				nilai = 1
			} else if test[t].Types == 2 {
				nilai = 2
			} else if test[t].Types == 3 {
				nilai = 3
			} else {
				nilai = 0
			}
		} else {
			you = append(you, test[t].Types)
			if contains(you, 3) == true {
				nilai = 3
			} else if contains(you, 1) == true && contains(you, 2) == true {
				nilai = 3
			} else if contains(you, 1) == true && contains(you, 2) == false {
				nilai = 1
			} else if contains(you, 1) == false && contains(you, 2) == true {
				nilai = 2
			} else {
				nilai = 0
			}
		}
	}
	return nilai
}
func (s *scoreService) OutputAllAverage(studentsID int) []entity.IkutBattle {
	var fscores []entity.IkutBattle
	finishedTests := s.repo.GroupScoreByStudentID(studentsID)

	for _, t := range finishedTests {
		fscore := s.Average(studentsID, int(t.TestID))
		fscores = append(fscores, fscore)
	}

	return fscores
}
func (s *scoreService) CariHighscore(id int) entity.IkutBattle {
	finishedTests := s.repo.GroupScoreByStudentID(id)
	var wawaw []float64
	var i entity.IkutBattle
	var avg float64
	var nilai float64
	for _, t := range finishedTests {
		fscore := s.Average(id, int(t.TestID))
		if fscore.Types == 3 {
			avg = (fscore.Score.Tps + fscore.Score.Saintek + fscore.Score.Soshum) / float64(3)
		} else {
			avg = (fscore.Score.Tps + fscore.Score.Saintek + fscore.Score.Soshum) / float64(2)
		}
		wawaw = append(wawaw, avg)
	}
	for _, t := range wawaw {
		if t > nilai {
			nilai = t
		}
	}
	for _, t := range finishedTests {
		fscore := s.Average(id, int(t.TestID))
		if (fscore.Score.Tps+fscore.Score.Saintek+fscore.Score.Soshum)/float64(3) == nilai || (fscore.Score.Tps+fscore.Score.Saintek+fscore.Score.Soshum)/float64(2) == nilai {
			i = fscore
		}
	}
	return i
}
func (s *scoreService) HighScoreBattle(id int) entity.HighScoreBattle {
	var h entity.HighScoreBattle
	student := s.studentRepo.FindOneById(id)
	all := s.CariHighscore(id)
	h.Name = student.Name
	h.Target = student.Target
	h.Score.Tps = all.Score.Tps
	h.Score.Saintek = all.Score.Saintek
	h.Score.Soshum = all.Score.Soshum
	h.Score.Battlename = all.Name
	return h
}
func (s *scoreService) StudentFScoreMin(studentsID int) []entity.FScoreMin {
	var fscores []entity.FScoreMin
	finishedTests := s.repo.GroupScoreByStudentID(studentsID)

	for _, t := range finishedTests {
		fscore := s.StudentFScoreMinOfTest(studentsID, int(t.TestID))
		fscores = append(fscores, fscore)
	}

	return fscores
}

func (s *scoreService) StudentFScoreMinOfTest(studentsID int, testID int) entity.FScoreMin {
	var fscore entity.FScoreMin
	var subtests entity.Subtests

	scores := s.repo.FindScoreMinByStudentsIDAndTestID(studentsID, testID)

	for _, score := range scores {
		switch subtest := score.Type; subtest {
		case "pu":
			subtests.Pu = score.Score
		case "ppu":
			subtests.Ppu = score.Score
		case "pmm":
			subtests.Pmm = score.Score
		case "pk":
			subtests.Pk = score.Score
		case "eng":
			subtests.Eng = score.Score
		}
	}
	fscore.Accum = subtests.Pu + subtests.Ppu + subtests.Pmm + subtests.Pk + subtests.Eng
	fscore.Avg = fscore.Accum / float64(5)
	fscore.Scores = subtests

	return fscore
}

func (s *scoreService) HighScore(id int) entity.HighScore {
	var h entity.HighScore
	var how []float64
	var how1 []float64
	var how2 []float64
	student := s.studentRepo.FindOneById(id)
	score := s.OutputAllAverage(id)
	var tps float64
	var saintek float64
	var soshum float64
	for t := range score {
		how = append(how, score[t].Score.Tps)
		how1 = append(how1, score[t].Score.Saintek)
		how2 = append(how2, score[t].Score.Soshum)
	}
	for _, v := range how {
		if v > tps {
			tps = v
		}
	}
	for _, v := range how1 {
		if v > saintek {
			saintek = v
		}
	}
	for _, v := range how2 {
		if v > soshum {
			soshum = v
		}
	}
	h.Score.Tps = tps
	h.Score.Saintek = saintek
	h.Score.Soshum = soshum
	h.Name = student.Name
	h.Target = student.Target
	return h
}

// Calculates one user stats (avg, latest, max)
func (s *scoreService) CalculateStudentStats(id int) entity.UserScoreStat {
	scores := s.repo.GroupScoreByStudentID(id)

	var max entity.FScoreMin = entity.FScoreMin{Avg: 0}
	var avg entity.FScoreMin = entity.FScoreMin{Avg: 0}
	var latest entity.FScoreMin

	var counter int = 0
	var maxTestID uint64

	for _, score := range scores {
		if score.Score >= max.Avg {
			max.Avg = score.Score
			maxTestID = score.TestID
		}
		counter++
	}

	for _, score := range s.repo.GroupScoreByStudentIDAndType(id) {
		avg.Accum = avg.Accum + score.Score

		switch s := score.Type; s {
		case "pu":
			avg.Scores.Pu = score.Score
		case "ppu":
			avg.Scores.Ppu = score.Score
		case "pmm":
			avg.Scores.Pmm = score.Score
		case "pk":
			avg.Scores.Pk = score.Score
		case "eng":
			avg.Scores.Eng = score.Score

		}
	}
	avg.Avg = avg.Accum / float64(5)

	latest = s.StudentFScoreMinOfTest(id, int(scores[counter-1].TestID))
	max = s.StudentFScoreMinOfTest(id, int(maxTestID))

	return entity.UserScoreStat{
		Avg:    avg,
		Max:    max,
		Latest: latest,
	}
}

// Calculates Test Stats
func (s *scoreService) CalculateTestStats(testID int) entity.TestStats {
	scores := s.repo.GroupScoreByTestID(testID)

	var max entity.FScoreMin = entity.FScoreMin{Avg: 0}
	var avg entity.FScoreMin

	maxStudentID := scores[0].StudentID

	for _, score := range s.repo.GroupScoreByTestIDAndType(testID) {

		avg.Accum = avg.Accum + score.Score

		switch s := score.Type; s {
		case "pu":
			avg.Scores.Pu = score.Score
		case "ppu":
			avg.Scores.Ppu = score.Score
		case "pmm":
			avg.Scores.Pmm = score.Score
		case "pk":
			avg.Scores.Pk = score.Score
		case "eng":
			avg.Scores.Eng = score.Score
		}
	}
	avg.Avg = avg.Accum / float64(5)

	max = s.StudentFScoreMinOfTest(int(maxStudentID), testID)
	la := s.GetQuartile(testID)

	return entity.TestStats{
		Max:      max,
		Avg:      avg,
		Quartile: la,
	}
}

func (s *scoreService) Rara(studentsID int) entity.Radar {
	var outTps []entity.RaraTPS
	var outSa []entity.RaraSaintek
	var outSos []entity.RaraSoshum
	var inTps [6]entity.RaraTPS
	var inSaintek [5]entity.RaraSaintek
	var inSoshum [5]entity.RaraSoshum
	var rara entity.Radar
	
	lastID := s.repo.GroupScoreByStudentIDLast(studentsID)
	for _, t := range lastID {
		testID := t.TestID
		siswa := s.studentRepo.FindOneById(studentsID)
		school := siswa.School
		pantek := s.testRepo.NameandType(int(testID))
		if pantek.Types == 0 {
			for _, subtest := range []string{"pu", "ppu", "pmm", "pk", "eng", "rata"} {
				skor := s.Lider(studentsID,int(testID))
				rataAll := s.repo.RataAll(testID, subtest)
				rataSekolah := s.repo.RataSekolahType(testID, school, subtest)
				maxAll := s.repo.MaxAll(testID, subtest)
				maxSekolah := s.repo.MaxSekolah(testID, school, subtest)
				maxAvgAll := s.repo.MaxAllAverage(testID)
				maxAvgSekolah := s.repo.MaxAllAverageSchool(testID, school)
				switch subtest {
				case "pu":
					inTps[0].Subject = "pu"
						inTps[0].Score = Round(skor.Pu)
					for _, rataall := range rataAll {
						inTps[0].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[0].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[0].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[0].MaxSchool = Round(maxSekolah.Score)
					}
	
				case "ppu":
					inTps[1].Subject = "ppu"
					inTps[1].Score = Round(skor.Ppu)
					for _, rataall := range rataAll {
						inTps[1].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[1].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[1].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[1].MaxSchool = Round(maxSekolah.Score)
					}
				case "pmm":
					inTps[2].Subject = "pmm"
						inTps[2].Score = Round(skor.Pmm)
					for _, rataall := range rataAll {
						inTps[2].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[2].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[2].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[2].MaxSchool = Round(maxSekolah.Score)
					}
				case "pk":
					inTps[3].Subject = "pk"
						inTps[3].Score = Round(skor.Pk)
					for _, rataall := range rataAll {
						inTps[3].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[3].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[3].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[3].MaxSchool = Round(maxSekolah.Score)
					}
				case "eng":
					inTps[4].Subject = "eng"
					inTps[4].Score = Round(skor.Eng)
					for _, rataall := range rataAll {
						inTps[4].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[4].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[4].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[4].MaxSchool = Round(maxSekolah.Score)
					}
				case "rata":
					inTps[5].Subject = "rata"
					inTps[5].Score = Round((inTps[0].Score + inTps[1].Score + inTps[2].Score + inTps[3].Score) / 4)
					inTps[5].Avg = Round((inTps[0].Avg + inTps[1].Avg + inTps[2].Avg + inTps[3].Avg) / 4)
					inTps[5].AvgSchool = Round((inTps[0].AvgSchool + inTps[1].AvgSchool + inTps[2].AvgSchool + inTps[3].AvgSchool) / 4)
					for _, maxAvgAll := range maxAvgAll {
						inTps[5].Max = Round(maxAvgAll.Score / 5)
					}
					for _, maxAvgSekolah := range maxAvgSekolah {
						inTps[5].MaxSchool = Round(maxAvgSekolah.Score / 5)
					}
				}
			}
		} else if pantek.Types == 1 {
			for _, subtest := range []string{"pu", "ppu", "pmm", "pk", "eng", "rata"} {
				skor := s.Lider(studentsID,int(testID))
				rataAll := s.repo.RataAll(testID, subtest)
				rataSekolah := s.repo.RataSekolahType(testID, school, subtest)
				maxAll := s.repo.MaxAll(testID, subtest)
				maxSekolah := s.repo.MaxSekolah(testID, school, subtest)
				maxAvgAll := s.repo.MaxAllAverage(testID)
				maxAvgSekolah := s.repo.MaxAllAverageSchool(testID, school)
				switch subtest {
				case "pu":
					inTps[0].Subject = "pu"
						inTps[0].Score = Round(skor.Pu)
					for _, rataall := range rataAll {
						inTps[0].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[0].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[0].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[0].MaxSchool = Round(maxSekolah.Score)
					}
	
				case "ppu":
					inTps[1].Subject = "ppu"
					inTps[1].Score = Round(skor.Ppu)
					for _, rataall := range rataAll {
						inTps[1].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[1].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[1].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[1].MaxSchool = Round(maxSekolah.Score)
					}
				case "pmm":
					inTps[2].Subject = "pmm"
						inTps[2].Score = Round(skor.Pmm)
					for _, rataall := range rataAll {
						inTps[2].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[2].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[2].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[2].MaxSchool = Round(maxSekolah.Score)
					}
				case "pk":
					inTps[3].Subject = "pk"
						inTps[3].Score = Round(skor.Pk)
					for _, rataall := range rataAll {
						inTps[3].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[3].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[3].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[3].MaxSchool = Round(maxSekolah.Score)
					}
				case "eng":
					inTps[4].Subject = "eng"
					inTps[4].Score = Round(skor.Eng)
					for _, rataall := range rataAll {
						inTps[4].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[4].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[4].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[4].MaxSchool = Round(maxSekolah.Score)
					}
				case "rata":
					inTps[5].Subject = "rata"
					inTps[5].Score = Round((inTps[0].Score + inTps[1].Score + inTps[2].Score + inTps[3].Score) / 4)
					inTps[5].Avg = Round((inTps[0].Avg + inTps[1].Avg + inTps[2].Avg + inTps[3].Avg) / 4)
					inTps[5].AvgSchool = Round((inTps[0].AvgSchool + inTps[1].AvgSchool + inTps[2].AvgSchool + inTps[3].AvgSchool) / 4)
					for _, maxAvgAll := range maxAvgAll {
						inTps[5].Max = Round(maxAvgAll.Score / 5)
					}
					for _, maxAvgSekolah := range maxAvgSekolah {
						inTps[5].MaxSchool = Round(maxAvgSekolah.Score / 5)
					}
				}
			}
			for _, subtest := range []string{"ma", "fi", "ki", "bi", "rata"} {
				skor := s.Lider(studentsID,int(testID))
				rataAll := s.repo.RataAll(testID, subtest)
				rataSekolah := s.repo.RataSekolahType(testID, school, subtest)
				maxAll := s.repo.MaxAll(testID, subtest)
				maxSekolah := s.repo.MaxSekolah(testID, school, subtest)
				maxAvgAll := s.repo.MaxAllAverageSaintek(testID)
				maxAvgSekolah := s.repo.MaxAllAverageSchoolSaintek(testID, school)
				switch subtest {
				case "ma":
					inSaintek[0].Subject = "ma"
					inSaintek[0].Score = Round(skor.Ma)
					for _, rataall := range rataAll {
						inSaintek[0].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSaintek[0].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSaintek[0].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSaintek[0].MaxSchool = Round(maxSekolah.Score)
					}
	
				case "fi":
					inSaintek[1].Subject = "fi"
						inSaintek[1].Score = Round(skor.Fi)
					for _, rataall := range rataAll {
						inSaintek[1].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSaintek[1].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSaintek[1].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSaintek[1].MaxSchool = Round(maxSekolah.Score)
					}
				case "ki":
					inSaintek[2].Subject = "ki"
					inSaintek[2].Score = Round(skor.Ki)
					for _, rataall := range rataAll {
						inSaintek[2].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSaintek[2].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSaintek[2].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSaintek[2].MaxSchool = Round(maxSekolah.Score)
					}
				case "bi":
					inSaintek[3].Subject = "bi"
					inSaintek[3].Score = Round(skor.Bi)
					for _, rataall := range rataAll {
						inSaintek[3].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSaintek[3].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSaintek[3].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSaintek[3].MaxSchool = Round(maxSekolah.Score)
					}
				case "rata":
					inSaintek[4].Subject = "rata"
					inSaintek[4].Score = (Round(inSaintek[0].Score + inSaintek[1].Score + inSaintek[2].Score + inSaintek[3].Score) / 4)
					inSaintek[4].Avg = Round((inSaintek[0].Avg + inSaintek[1].Avg + inSaintek[2].Avg + inSaintek[3].Avg) / 4)
					inSaintek[4].AvgSchool = Round((inSaintek[0].AvgSchool + inSaintek[1].AvgSchool + inSaintek[2].AvgSchool + inSaintek[3].AvgSchool) / 4)
					for _, maxAvgAll := range maxAvgAll {
						inSaintek[4].Max = Round(maxAvgAll.Score / 4)
					}
					for _, maxAvgSekolah := range maxAvgSekolah {
						inSaintek[4].MaxSchool = Round(maxAvgSekolah.Score / 4)
					}
				}
	
			}	
		} else if pantek.Types == 2 {
			for _, subtest := range []string{"pu", "ppu", "pmm", "pk", "eng", "rata"} {
				skor := s.Lider(studentsID,int(testID))
				rataAll := s.repo.RataAll(testID, subtest)
				rataSekolah := s.repo.RataSekolahType(testID, school, subtest)
				maxAll := s.repo.MaxAll(testID, subtest)
				maxSekolah := s.repo.MaxSekolah(testID, school, subtest)
				maxAvgAll := s.repo.MaxAllAverage(testID)
				maxAvgSekolah := s.repo.MaxAllAverageSchool(testID, school)
				switch subtest {
				case "pu":
					inTps[0].Subject = "pu"
						inTps[0].Score = Round(skor.Pu)
					for _, rataall := range rataAll {
						inTps[0].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[0].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[0].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[0].MaxSchool = Round(maxSekolah.Score)
					}
	
				case "ppu":
					inTps[1].Subject = "ppu"
					inTps[1].Score = Round(skor.Ppu)
					for _, rataall := range rataAll {
						inTps[1].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[1].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[1].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[1].MaxSchool = Round(maxSekolah.Score)
					}
				case "pmm":
					inTps[2].Subject = "pmm"
						inTps[2].Score = Round(skor.Pmm)
					for _, rataall := range rataAll {
						inTps[2].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[2].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[2].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[2].MaxSchool = Round(maxSekolah.Score)
					}
				case "pk":
					inTps[3].Subject = "pk"
						inTps[3].Score = Round(skor.Pk)
					for _, rataall := range rataAll {
						inTps[3].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[3].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[3].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[3].MaxSchool = Round(maxSekolah.Score)
					}
				case "eng":
					inTps[4].Subject = "eng"
					inTps[4].Score = Round(skor.Eng)
					for _, rataall := range rataAll {
						inTps[4].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[4].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[4].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[4].MaxSchool = Round(maxSekolah.Score)
					}
				case "rata":
					inTps[5].Subject = "rata"
					inTps[5].Score = Round((inTps[0].Score + inTps[1].Score + inTps[2].Score + inTps[3].Score) / 4)
					inTps[5].Avg = Round((inTps[0].Avg + inTps[1].Avg + inTps[2].Avg + inTps[3].Avg) / 4)
					inTps[5].AvgSchool = Round((inTps[0].AvgSchool + inTps[1].AvgSchool + inTps[2].AvgSchool + inTps[3].AvgSchool) / 4)
					for _, maxAvgAll := range maxAvgAll {
						inTps[5].Max = Round(maxAvgAll.Score / 5)
					}
					for _, maxAvgSekolah := range maxAvgSekolah {
						inTps[5].MaxSchool = Round(maxAvgSekolah.Score / 5)
					}
				}
			}
			for _, subtest := range []string{"sos", "sej", "geo", "eko", "rata"} {
				skor := s.Lider(studentsID,int(testID))
				rataAll := s.repo.RataAll(testID, subtest)
				rataSekolah := s.repo.RataSekolahType(testID, school, subtest)
				maxAll := s.repo.MaxAll(testID, subtest)
				maxSekolah := s.repo.MaxSekolah(testID, school, subtest)
				maxAvgAll := s.repo.MaxAllAverageSoshum(testID)
				maxAvgSekolah := s.repo.MaxAllAverageSchoolSoshum(testID, school)
				switch subtest {
				case "sos":
					inSoshum[0].Subject = "sos"
					inSoshum[0].Score = Round(skor.Sos)
					for _, rataall := range rataAll {
						inSoshum[0].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSoshum[0].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSoshum[0].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSoshum[0].MaxSchool = Round(maxSekolah.Score)
					}
	
				case "sej":
					inSoshum[1].Subject = "sej"
					inSoshum[1].Score = Round(skor.Sej)
					for _, rataall := range rataAll {
						inSoshum[1].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSoshum[1].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSoshum[1].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSoshum[1].MaxSchool = Round(maxSekolah.Score)
					}
				case "geo":
					inSoshum[2].Subject = "geo"
					inSoshum[2].Score = Round(skor.Geo)
					for _, rataall := range rataAll {
						inSoshum[2].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSoshum[2].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSoshum[2].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSoshum[2].MaxSchool = Round(maxSekolah.Score)
					}
				case "eko":
					inSoshum[3].Subject = "eko"
					inSoshum[3].Score = Round(skor.Eko)
					for _, rataall := range rataAll {
						inSoshum[3].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSoshum[3].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSoshum[3].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSoshum[3].MaxSchool = Round(maxSekolah.Score)
					}
				case "rata":
					inSoshum[4].Subject = "rata"
					inSoshum[4].Score = Round((inSoshum[0].Score + inSoshum[1].Score + inSoshum[2].Score + inSoshum[3].Score) / 4)
					inSoshum[4].Avg = Round((inSoshum[0].Avg + inSoshum[1].Avg + inSoshum[2].Avg + inSoshum[3].Avg) / 4)
					inSoshum[4].AvgSchool = Round((inSoshum[0].AvgSchool + inSoshum[1].AvgSchool + inSoshum[2].AvgSchool + inSoshum[3].AvgSchool) / 4)
					for _, maxAvgAll := range maxAvgAll {
						inSoshum[4].Max = Round(maxAvgAll.Score / 4)
					}
					for _, maxAvgSekolah := range maxAvgSekolah {
						inSoshum[4].MaxSchool = Round(maxAvgSekolah.Score / 4)
					}
				}
			}
		} else if pantek.Types == 3 {
			for _, subtest := range []string{"pu", "ppu", "pmm", "pk", "eng", "rata"} {
				skor := s.Lider(studentsID,int(testID))
				rataAll := s.repo.RataAll(testID, subtest)
				rataSekolah := s.repo.RataSekolahType(testID, school, subtest)
				maxAll := s.repo.MaxAll(testID, subtest)
				maxSekolah := s.repo.MaxSekolah(testID, school, subtest)
				maxAvgAll := s.repo.MaxAllAverage(testID)
				maxAvgSekolah := s.repo.MaxAllAverageSchool(testID, school)
				switch subtest {
				case "pu":
					inTps[0].Subject = "pu"
						inTps[0].Score = Round(skor.Pu)
					for _, rataall := range rataAll {
						inTps[0].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[0].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[0].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[0].MaxSchool = Round(maxSekolah.Score)
					}
	
				case "ppu":
					inTps[1].Subject = "ppu"
					inTps[1].Score = Round(skor.Ppu)
					for _, rataall := range rataAll {
						inTps[1].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[1].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[1].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[1].MaxSchool = Round(maxSekolah.Score)
					}
				case "pmm":
					inTps[2].Subject = "pmm"
						inTps[2].Score = Round(skor.Pmm)
					for _, rataall := range rataAll {
						inTps[2].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[2].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[2].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[2].MaxSchool = Round(maxSekolah.Score)
					}
				case "pk":
					inTps[3].Subject = "pk"
						inTps[3].Score = Round(skor.Pk)
					for _, rataall := range rataAll {
						inTps[3].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[3].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[3].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[3].MaxSchool = Round(maxSekolah.Score)
					}
				case "eng":
					inTps[4].Subject = "eng"
					inTps[4].Score = Round(skor.Eng)
					for _, rataall := range rataAll {
						inTps[4].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inTps[4].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inTps[4].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inTps[4].MaxSchool = Round(maxSekolah.Score)
					}
				case "rata":
					inTps[5].Subject = "rata"
					inTps[5].Score = Round((inTps[0].Score + inTps[1].Score + inTps[2].Score + inTps[3].Score) / 4)
					inTps[5].Avg = Round((inTps[0].Avg + inTps[1].Avg + inTps[2].Avg + inTps[3].Avg) / 4)
					inTps[5].AvgSchool = Round((inTps[0].AvgSchool + inTps[1].AvgSchool + inTps[2].AvgSchool + inTps[3].AvgSchool) / 4)
					for _, maxAvgAll := range maxAvgAll {
						inTps[5].Max = Round(maxAvgAll.Score / 5)
					}
					for _, maxAvgSekolah := range maxAvgSekolah {
						inTps[5].MaxSchool = Round(maxAvgSekolah.Score / 5)
					}
				}
			}
			for _, subtest := range []string{"ma", "fi", "ki", "bi", "rata"} {
				skor := s.Lider(studentsID,int(testID))
				rataAll := s.repo.RataAll(testID, subtest)
				rataSekolah := s.repo.RataSekolahType(testID, school, subtest)
				maxAll := s.repo.MaxAll(testID, subtest)
				maxSekolah := s.repo.MaxSekolah(testID, school, subtest)
				maxAvgAll := s.repo.MaxAllAverageSaintek(testID)
				maxAvgSekolah := s.repo.MaxAllAverageSchoolSaintek(testID, school)
				switch subtest {
				case "ma":
					inSaintek[0].Subject = "ma"
					inSaintek[0].Score = Round(skor.Ma)
					for _, rataall := range rataAll {
						inSaintek[0].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSaintek[0].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSaintek[0].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSaintek[0].MaxSchool = Round(maxSekolah.Score)
					}
	
				case "fi":
					inSaintek[1].Subject = "fi"
						inSaintek[1].Score = Round(skor.Fi)
					for _, rataall := range rataAll {
						inSaintek[1].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSaintek[1].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSaintek[1].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSaintek[1].MaxSchool = Round(maxSekolah.Score)
					}
				case "ki":
					inSaintek[2].Subject = "ki"
					inSaintek[2].Score = Round(skor.Ki)
					for _, rataall := range rataAll {
						inSaintek[2].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSaintek[2].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSaintek[2].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSaintek[2].MaxSchool = Round(maxSekolah.Score)
					}
				case "bi":
					inSaintek[3].Subject = "bi"
					inSaintek[3].Score = Round(skor.Bi)
					for _, rataall := range rataAll {
						inSaintek[3].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSaintek[3].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSaintek[3].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSaintek[3].MaxSchool = Round(maxSekolah.Score)
					}
				case "rata":
					inSaintek[4].Subject = "rata"
					inSaintek[4].Score = (Round(inSaintek[0].Score + inSaintek[1].Score + inSaintek[2].Score + inSaintek[3].Score) / 4)
					inSaintek[4].Avg = Round((inSaintek[0].Avg + inSaintek[1].Avg + inSaintek[2].Avg + inSaintek[3].Avg) / 4)
					inSaintek[4].AvgSchool = Round((inSaintek[0].AvgSchool + inSaintek[1].AvgSchool + inSaintek[2].AvgSchool + inSaintek[3].AvgSchool) / 4)
					for _, maxAvgAll := range maxAvgAll {
						inSaintek[4].Max = Round(maxAvgAll.Score / 4)
					}
					for _, maxAvgSekolah := range maxAvgSekolah {
						inSaintek[4].MaxSchool = Round(maxAvgSekolah.Score / 4)
					}
				}
	
			}	
			for _, subtest := range []string{"sos", "sej", "geo", "eko", "rata"} {
				skor := s.Lider(studentsID,int(testID))
				rataAll := s.repo.RataAll(testID, subtest)
				rataSekolah := s.repo.RataSekolahType(testID, school, subtest)
				maxAll := s.repo.MaxAll(testID, subtest)
				maxSekolah := s.repo.MaxSekolah(testID, school, subtest)
				maxAvgAll := s.repo.MaxAllAverageSoshum(testID)
				maxAvgSekolah := s.repo.MaxAllAverageSchoolSoshum(testID, school)
				switch subtest {
				case "sos":
					inSoshum[0].Subject = "sos"
					inSoshum[0].Score = Round(skor.Sos)
					for _, rataall := range rataAll {
						inSoshum[0].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSoshum[0].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSoshum[0].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSoshum[0].MaxSchool = Round(maxSekolah.Score)
					}
	
				case "sej":
					inSoshum[1].Subject = "sej"
					inSoshum[1].Score = Round(skor.Sej)
					for _, rataall := range rataAll {
						inSoshum[1].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSoshum[1].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSoshum[1].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSoshum[1].MaxSchool = Round(maxSekolah.Score)
					}
				case "geo":
					inSoshum[2].Subject = "geo"
					inSoshum[2].Score = Round(skor.Geo)
					for _, rataall := range rataAll {
						inSoshum[2].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSoshum[2].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSoshum[2].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSoshum[2].MaxSchool = Round(maxSekolah.Score)
					}
				case "eko":
					inSoshum[3].Subject = "eko"
					inSoshum[3].Score = Round(skor.Eko)
					for _, rataall := range rataAll {
						inSoshum[3].Avg = Round(rataall.Score)
					}
					for _, rataSekolah := range rataSekolah {
						inSoshum[3].AvgSchool = Round(rataSekolah.Score)
					}
					for _, maxAll := range maxAll {
						inSoshum[3].Max = Round(maxAll.Score)
					}
					for _, maxSekolah := range maxSekolah {
						inSoshum[3].MaxSchool = Round(maxSekolah.Score)
					}
				case "rata":
					inSoshum[4].Subject = "rata"
					inSoshum[4].Score = Round((inSoshum[0].Score + inSoshum[1].Score + inSoshum[2].Score + inSoshum[3].Score) / 4)
					inSoshum[4].Avg = Round((inSoshum[0].Avg + inSoshum[1].Avg + inSoshum[2].Avg + inSoshum[3].Avg) / 4)
					inSoshum[4].AvgSchool = Round((inSoshum[0].AvgSchool + inSoshum[1].AvgSchool + inSoshum[2].AvgSchool + inSoshum[3].AvgSchool) / 4)
					for _, maxAvgAll := range maxAvgAll {
						inSoshum[4].Max = Round(maxAvgAll.Score / 4)
					}
					for _, maxAvgSekolah := range maxAvgSekolah {
						inSoshum[4].MaxSchool = Round(maxAvgSekolah.Score / 4)
					}
				}
			}
		}
		wibu := s.testRepo.NameandType2(testID)
		if wibu.Types == 1 {
			outTps := append(outTps, inTps[0], inTps[1], inTps[2], inTps[3], inTps[4], inTps[5])
			outSa := append(outSa, inSaintek[0], inSaintek[1], inSaintek[2], inSaintek[3], inSaintek[4])
			rara.Tps = outTps
			rara.Saintek = outSa
		} else if wibu.Types == 2 {
			outTps := append(outTps, inTps[0], inTps[1], inTps[2], inTps[3], inTps[4], inTps[5])
			outSos := append(outSos, inSoshum[0], inSoshum[1], inSoshum[2], inSoshum[3], inSoshum[4])
			rara.Tps = outTps
			rara.Soshum = outSos
		} else if wibu.Types == 3 {
			outTps := append(outTps, inTps[0], inTps[1], inTps[2], inTps[3], inTps[4], inTps[5])
			outSos := append(outSos, inSoshum[0], inSoshum[1], inSoshum[2], inSoshum[3], inSoshum[4])
			outSa := append(outSa, inSaintek[0], inSaintek[1], inSaintek[2], inSaintek[3], inSaintek[4])
			rara.Tps = outTps
			rara.Soshum = outSos
			rara.Saintek = outSa
		} else if wibu.Types == 0 {
			outTps := append(outTps, inTps[0], inTps[1], inTps[2], inTps[3], inTps[4], inTps[5])
			rara.Tps = outTps
		}
		rara.Name = wibu.Name
	}
	return rara
} 
func (s *scoreService) BenarSalah(testID int,StudentID int) entity.Nilai {
	var nilai entity.Nilai
	var check int
	pantek := s.testRepo.NameandType(testID)
	if pantek.Types == 0{
		for _,subtest := range []string{"pu","ppu","pmm","pk","eng"}{
			benar := s.repo.BenarSalah(StudentID,testID,subtest)
			switch subtest {
				case "pu":
					check =0
					var array []int
					array = append(array,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
					for i:=0;i<len(array);i++{
						if array[i]==1{
							check++
						}
					}
					nilai.Pu = float64(check)/float64(16)*float64(100)
				case "ppu":
					check =0
					var array2 []int
					array2 = append(array2,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
					for i:=0;i<len(array2);i++{
						if array2[i]==1{
							check++
						}
					}
					nilai.Ppu = float64(check)/float64(19)*float64(100)
				case "pmm":
					check =0
					var array3 []int
					array3 = append(array3,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
					for i:=0;i<len(array3);i++{
						if array3[i]==1{
							check++
						}
					}
					nilai.Pmm = float64(check)/float64(15)*float64(100)
				case "pk":
					check =0
					var array4 []int
					array4 = append(array4,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
					for i:=0;i<len(array4);i++{
						if array4[i]==1{
							check++
						}
					}
					nilai.Pk= float64(check)/float64(16)*float64(100)
				case "eng":
					check =0
					var array5 []int
					array5 = append(array5,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
					for i:=0;i<len(array5);i++{
						if array5[i]==1{
							check++
						}
					}
					nilai.Eng = float64(check)/float64(10)*float64(100)
			}
		}
	}
	if pantek.Types == 1{
		for _,subtest := range []string{"pu","ppu","pmm","pk","eng","ma","fi","ki","bi"}{
			benar := s.repo.BenarSalah(StudentID,testID,subtest)
			switch subtest {
			case "pu":
				check =0
				var array []int
				array = append(array,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array);i++{
					if array[i]==1{
						check++
					}
				}
				nilai.Pu = float64(check)/float64(16)*float64(100)
			case "ppu":
				check =0
				var array2 []int
				array2 = append(array2,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array2);i++{
					if array2[i]==1{
						check++
					}
				}
				nilai.Ppu = float64(check)/float64(19)*float64(100)
			case "pmm":
				check =0
				var array3 []int
				array3 = append(array3,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array3);i++{
					if array3[i]==1{
						check++
					}
				}
				nilai.Pmm = float64(check)/float64(15)*float64(100)
			case "pk":
				check =0
				var array4 []int
				array4 = append(array4,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array4);i++{
					if array4[i]==1{
						check++
					}
				}
				nilai.Pk= float64(check)/float64(16)*float64(100)
			case "eng":
				check =0
				var array5 []int
				array5 = append(array5,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array5);i++{
					if array5[i]==1{
						check++
					}
				}
				nilai.Eng = float64(check)/float64(10)*float64(100)
		case "ma":
			check =0
			var array6 []int
			array6 = append(array6,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
			for i:=0;i<len(array6);i++{
				if array6[i]==1{
					check++
				}
			}
			nilai.Ma = float64(check)/float64(20)*float64(100)
			case "fi":
				check =0
				var array7 []int
				array7 = append(array7,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array7);i++{
					if array7[i]==1{
						check++
					}
				}
				nilai.Fi = float64(check)/float64(20)*float64(100)

				case "ki":
					check =0
					var array8 []int
					array8 = append(array8,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
					for i:=0;i<len(array8);i++{
						if array8[i]==1{
							check++
						}
					}
					nilai.Ki = float64(check)/float64(20)*float64(100)
					case "bi":
						check =0
						var array9[]int
						array9 = append(array9,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
						for i:=0;i<len(array9);i++{
							if array9[i]==1{
								check++
							}
						}
						nilai.Bi = float64(check)/float64(20)*float64(100)
			}
		}
	}
	if pantek.Types == 2{
		for _,subtest := range []string{"pu","ppu","pmm","pk","eng","sos","sej","geo","eko"}{
			benar := s.repo.BenarSalah(StudentID,testID,subtest)
			switch subtest {
			case "pu":
				check =0
				var array []int
				array = append(array,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array);i++{
					if array[i]==1{
						check++
					}
				}
				nilai.Pu = float64(check)/float64(16)*float64(100)
			case "ppu":
				check =0
				var array2 []int
				array2 = append(array2,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array2);i++{
					if array2[i]==1{
						check++
					}
				}
				nilai.Ppu = float64(check)/float64(19)*float64(100)
			case "pmm":
				check =0
				var array3 []int
				array3 = append(array3,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array3);i++{
					if array3[i]==1{
						check++
					}
				}
				nilai.Pmm = float64(check)/float64(15)*float64(100)
			case "pk":
				check =0
				var array4 []int
				array4 = append(array4,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array4);i++{
					if array4[i]==1{
						check++
					}
				}
				nilai.Pk= float64(check)/float64(16)*float64(100)
			case "eng":
				check =0
				var array5 []int
				array5 = append(array5,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array5);i++{
					if array5[i]==1{
						check++
					}
				}
				nilai.Eng = float64(check)/float64(10)*float64(100)
				case "sos":
					check =0
					var array6 []int
					array6 = append(array6,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
					for i:=0;i<len(array6);i++{
						if array6[i]==1{
							check++
						}
					}
					nilai.Sos = float64(check)/float64(10)*float64(100)
					case "sej":
						check =0
						var array7 []int
						array7 = append(array7,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
						for i:=0;i<len(array7);i++{
							if array7[i]==1{
								check++
							}
						}
						nilai.Sej = float64(check)/float64(20)*float64(100)
						case "geo":
							check =0
							var array8 []int
							array8 = append(array8,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
							for i:=0;i<len(array8);i++{
								if array8[i]==1{
									check++
								}
							}
							nilai.Geo = float64(check)/float64(20)*float64(100)
			
				case "eko":
					check =0
					var array9 []int
					array9 = append(array9,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
					for i:=0;i<len(array9);i++{
						if array9[i]==1{
							check++
						}
					}
					nilai.Eko = float64(check)/float64(20)*float64(100)
			}
		}
	}
	if pantek.Types == 3{
		for _,subtest := range []string{"pu","ppu","pmm","pk","eng","ma","fi","ki","bi","sos","sej","geo","eko"}{
			benar := s.repo.BenarSalah(StudentID,testID,subtest)
			switch subtest {
			case "pu":
				check =0
				var array []int
				array = append(array,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array);i++{
					if array[i]==1{
						check++
					}
				}
				nilai.Pu = float64(check)/float64(16)*float64(100)
			case "ppu":
				check =0
				var array2 []int
				array2 = append(array2,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array2);i++{
					if array2[i]==1{
						check++
					}
				}
				nilai.Ppu = float64(check)/float64(19)*float64(100)
			case "pmm":
				check =0
				var array3 []int
				array3 = append(array3,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array3);i++{
					if array3[i]==1{
						check++
					}
				}
				nilai.Pmm = float64(check)/float64(15)*float64(100)
			case "pk":
				check =0
				var array4 []int
				array4 = append(array4,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array4);i++{
					if array4[i]==1{
						check++
					}
				}
				nilai.Pk= float64(check)/float64(16)*float64(100)
			case "eng":
				check =0
				var array5 []int
				array5 = append(array5,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array5);i++{
					if array5[i]==1{
						check++
					}
				}
				nilai.Eng = float64(check)/float64(10)*float64(100)
		case "ma":
			check =0
			var array6 []int
			array6 = append(array6,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
			for i:=0;i<len(array6);i++{
				if array6[i]==1{
					check++
				}
			}
			nilai.Ma = float64(check)/float64(20)*float64(100)
			case "fi":
				check =0
				var array7 []int
				array7 = append(array7,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array7);i++{
					if array7[i]==1{
						check++
					}
				}
				nilai.Fi = float64(check)/float64(20)*float64(100)

				case "ki":
					check =0
					var array8 []int
					array8 = append(array8,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
					for i:=0;i<len(array8);i++{
						if array8[i]==1{
							check++
						}
					}
					nilai.Ki = float64(check)/float64(20)*float64(100)
					case "bi":
						check =0
						var array9[]int
						array9 = append(array9,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
						for i:=0;i<len(array9);i++{
							if array9[i]==1{
								check++
							}
						}
						nilai.Bi = float64(check)/float64(20)*float64(100)
				case "sos":
					check =0
					var array10[]int
					array10 = append(array10,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
					for i:=0;i<len(array10);i++{
						if array10[i]==1{
							check++
						}
					}
					nilai.Sos = float64(check)/float64(20)*float64(100)
			case "sej":
				check =0
				var array11[]int
				array11 = append(array11,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
				for i:=0;i<len(array11);i++{
					if array11[i]==1{
						check++
					}
				}
				nilai.Sej = float64(check)/float64(20)*float64(100)
		case "geo":
			check =0
			var array12[]int
			array12 = append(array12,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
			for i:=0;i<len(array12);i++{
				if array12[i]==1{
					check++
				}
			}
			nilai.Geo = float64(check)/float64(20)*float64(100)
	case "eko":
		check =0
		var array13[]int
		array13 = append(array13,benar.X1,benar.X2,benar.X3,benar.X4,benar.X5,benar.X6,benar.X7,benar.X8,benar.X9,benar.X10,benar.X11,benar.X12,benar.X13,benar.X14,benar.X15,benar.X16,benar.X17,benar.X18,benar.X19,benar.X20)
		for i:=0;i<len(array13);i++{
			if array13[i]==1{
				check++
			}
		}
		nilai.Eko = float64(check)/float64(20)*float64(100)
}
		}
	}
	return nilai
}
func (s *scoreService) Distribution(TestID int) entity.Distribution {
	var output []entity.RangeTps
	var outputSa []entity.RangeSaintek
	var outputSo []entity.RangeSoshum
	var tps [10]entity.RangeTps
	var didi entity.Distribution
	var saintek [10]entity.RangeSaintek
	var soshum [10]entity.RangeSoshum
	tps[0].Range = "<100"
	tps[1].Range = "101-200"
	tps[2].Range = "201-300"
	tps[3].Range = "301-400"
	tps[4].Range = "401-500"
	tps[5].Range = "501-600"
	tps[6].Range = "601-700"
	tps[7].Range = "701-800"
	tps[8].Range = "801-900"
	tps[9].Range = ">900"
	saintek[0].Range = "<100"
	saintek[1].Range = "101-200"
	saintek[2].Range = "201-300"
	saintek[3].Range = "301-400"
	saintek[4].Range = "401-500"
	saintek[5].Range = "501-600"
	saintek[6].Range = "601-700"
	saintek[7].Range = "701-800"
	saintek[8].Range = "801-900"
	saintek[9].Range = ">900"
	soshum[0].Range = "<100"
	soshum[1].Range = "101-200"
	soshum[2].Range = "201-300"
	soshum[3].Range = "301-400"
	soshum[4].Range = "401-500"
	soshum[5].Range = "501-600"
	soshum[6].Range = "601-700"
	soshum[7].Range = "701-800"
	soshum[8].Range = "801-900"
	soshum[9].Range = ">900"
	pantek := s.testRepo.NameandType(TestID)
	if pantek.Types == 0 {
	avg1, avg2, avg3, avg4, avg5, avg6, avg7, avg8, avg9, avg10 := s.repo.AVGDistributionTps(TestID)
	tps[0].Avg = avg1
	tps[1].Avg = avg2
	tps[2].Avg = avg3
	tps[3].Avg = avg4
	tps[4].Avg = avg5
	tps[5].Avg = avg6
	tps[6].Avg = avg7
	tps[7].Avg = avg8
	tps[8].Avg = avg9
	tps[9].Avg = avg10
	for _, i := range []string{"pu", "ppu", "pmm", "pk", "eng"} {
		value1, value2, value3, value4, value5, value6, value7, value8, value9, value10 := s.repo.Distribution(TestID, i)
		switch i {
		case "pu":
			tps[0].Pu = value1
			tps[1].Pu = value2
			tps[2].Pu = value3
			tps[3].Pu = value4
			tps[4].Pu = value5
			tps[5].Pu = value6
			tps[6].Pu = value7
			tps[7].Pu = value8
			tps[8].Pu = value9
			tps[9].Pu = value10
		case "ppu":
			tps[0].Ppu = value1
			tps[1].Ppu = value2
			tps[2].Ppu = value3
			tps[3].Ppu = value4
			tps[4].Ppu = value5
			tps[5].Ppu = value6
			tps[6].Ppu = value7
			tps[7].Ppu = value8
			tps[8].Ppu = value9
			tps[9].Ppu = value10
		case "pmm":
			tps[0].Pmm = value1
			tps[1].Pmm = value2
			tps[2].Pmm = value3
			tps[3].Pmm = value4
			tps[4].Pmm = value5
			tps[5].Pmm = value6
			tps[6].Pmm = value7
			tps[7].Pmm = value8
			tps[8].Pmm = value9
			tps[9].Pmm = value10
		case "pk":
			tps[0].Pk = value1
			tps[1].Pk = value2
			tps[2].Pk = value3
			tps[3].Pk = value4
			tps[4].Pk = value5
			tps[5].Pk = value6
			tps[6].Pk = value7
			tps[7].Pk = value8
			tps[8].Pk = value9
			tps[9].Pk = value10
		case "eng":
			tps[0].Eng = value1
			tps[1].Eng = value2
			tps[2].Eng = value3
			tps[3].Eng = value4
			tps[4].Eng = value5
			tps[5].Eng = value6
			tps[6].Eng = value7
			tps[7].Eng = value8
			tps[8].Eng = value9
			tps[9].Eng = value10
		}
	}
	} else if pantek.Types == 1{
	avg1, avg2, avg3, avg4, avg5, avg6, avg7, avg8, avg9, avg10 := s.repo.AVGDistributionTps(TestID)
	tps[0].Avg = avg1
	tps[1].Avg = avg2
	tps[2].Avg = avg3
	tps[3].Avg = avg4
	tps[4].Avg = avg5
	tps[5].Avg = avg6
	tps[6].Avg = avg7
	tps[7].Avg = avg8
	tps[8].Avg = avg9
	tps[9].Avg = avg10
	avgSa1, avgSa2, avgSa3, avgSa4, avgSa5, avgSa6, avgSa7, avgSa8, avgSa9, avgSa10 := s.repo.AVGDistributionSaintek(TestID)
	saintek[0].Avg = avgSa1
	saintek[1].Avg = avgSa2
	saintek[2].Avg = avgSa3
	saintek[3].Avg = avgSa4
	saintek[4].Avg = avgSa5
	saintek[5].Avg = avgSa6
	saintek[6].Avg = avgSa7
	saintek[7].Avg = avgSa8
	saintek[8].Avg = avgSa9
	saintek[9].Avg = avgSa10
	for _, i := range []string{"pu", "ppu", "pmm", "pk", "eng"} {
		value1, value2, value3, value4, value5, value6, value7, value8, value9, value10 := s.repo.Distribution(TestID, i)
		switch i {
		case "pu":
			tps[0].Pu = value1
			tps[1].Pu = value2
			tps[2].Pu = value3
			tps[3].Pu = value4
			tps[4].Pu = value5
			tps[5].Pu = value6
			tps[6].Pu = value7
			tps[7].Pu = value8
			tps[8].Pu = value9
			tps[9].Pu = value10
		case "ppu":
			tps[0].Ppu = value1
			tps[1].Ppu = value2
			tps[2].Ppu = value3
			tps[3].Ppu = value4
			tps[4].Ppu = value5
			tps[5].Ppu = value6
			tps[6].Ppu = value7
			tps[7].Ppu = value8
			tps[8].Ppu = value9
			tps[9].Ppu = value10
		case "pmm":
			tps[0].Pmm = value1
			tps[1].Pmm = value2
			tps[2].Pmm = value3
			tps[3].Pmm = value4
			tps[4].Pmm = value5
			tps[5].Pmm = value6
			tps[6].Pmm = value7
			tps[7].Pmm = value8
			tps[8].Pmm = value9
			tps[9].Pmm = value10
		case "pk":
			tps[0].Pk = value1
			tps[1].Pk = value2
			tps[2].Pk = value3
			tps[3].Pk = value4
			tps[4].Pk = value5
			tps[5].Pk = value6
			tps[6].Pk = value7
			tps[7].Pk = value8
			tps[8].Pk = value9
			tps[9].Pk = value10
		case "eng":
			tps[0].Eng = value1
			tps[1].Eng = value2
			tps[2].Eng = value3
			tps[3].Eng = value4
			tps[4].Eng = value5
			tps[5].Eng = value6
			tps[6].Eng = value7
			tps[7].Eng = value8
			tps[8].Eng = value9
			tps[9].Eng = value10
		}
	}
	for _, i := range []string{"ma", "fi", "ki", "bi"} {
		saintek1, saintek2, saintek3, saintek4, saintek5, saintek6, saintek7, saintek8, saintek9, saintek10 := s.repo.Distribution(TestID, i)
		switch i {
		case "ma":
			saintek[0].Ma = saintek1
			saintek[1].Ma = saintek2
			saintek[2].Ma = saintek3
			saintek[3].Ma = saintek4
			saintek[4].Ma = saintek5
			saintek[5].Ma = saintek6
			saintek[6].Ma = saintek7
			saintek[7].Ma = saintek8
			saintek[8].Ma = saintek9
			saintek[9].Ma = saintek10
		case "fi":
			saintek[0].Fi = saintek1
			saintek[1].Fi = saintek2
			saintek[2].Fi = saintek3
			saintek[3].Fi = saintek4
			saintek[4].Fi = saintek5
			saintek[5].Fi = saintek6
			saintek[6].Fi = saintek7
			saintek[7].Fi = saintek8
			saintek[8].Fi = saintek9
			saintek[9].Fi = saintek10
		case "ki":
			saintek[0].Ki = saintek1
			saintek[1].Ki = saintek2
			saintek[2].Ki = saintek3
			saintek[3].Ki = saintek4
			saintek[4].Ki = saintek5
			saintek[5].Ki = saintek6
			saintek[6].Ki = saintek7
			saintek[7].Ki = saintek8
			saintek[8].Ki = saintek9
			saintek[9].Ki = saintek10
		case "bi":
			saintek[0].Bi = saintek1
			saintek[1].Bi = saintek2
			saintek[2].Bi = saintek3
			saintek[3].Bi = saintek4
			saintek[4].Bi = saintek5
			saintek[5].Bi = saintek6
			saintek[6].Bi = saintek7
			saintek[7].Bi = saintek8
			saintek[8].Bi = saintek9
			saintek[9].Bi = saintek10
		}
	}
	}else if pantek.Types == 2 {
		avg1, avg2, avg3, avg4, avg5, avg6, avg7, avg8, avg9, avg10 := s.repo.AVGDistributionTps(TestID)
		tps[0].Avg = avg1
		tps[1].Avg = avg2
		tps[2].Avg = avg3
		tps[3].Avg = avg4
		tps[4].Avg = avg5
		tps[5].Avg = avg6
		tps[6].Avg = avg7
		tps[7].Avg = avg8
		tps[8].Avg = avg9
		tps[9].Avg = avg10
		avgSo1, avgSo2, avgSo3, avgSo4, avgSo5, avgSo6, avgSo7, avgSo8, avgSo9, avgSo10 := s.repo.AVGDistributionSoshum(TestID)
		soshum[0].Avg = avgSo1
		soshum[1].Avg = avgSo2
		soshum[2].Avg = avgSo3
		soshum[3].Avg = avgSo4
		soshum[4].Avg = avgSo5
		soshum[5].Avg = avgSo6
		soshum[6].Avg = avgSo7
		soshum[7].Avg = avgSo8
		soshum[8].Avg = avgSo9
		soshum[9].Avg = avgSo10
		for _, i := range []string{"pu", "ppu", "pmm", "pk", "eng"} {
			value1, value2, value3, value4, value5, value6, value7, value8, value9, value10 := s.repo.Distribution(TestID, i)
			switch i {
			case "pu":
				tps[0].Pu = value1
				tps[1].Pu = value2
				tps[2].Pu = value3
				tps[3].Pu = value4
				tps[4].Pu = value5
				tps[5].Pu = value6
				tps[6].Pu = value7
				tps[7].Pu = value8
				tps[8].Pu = value9
				tps[9].Pu = value10
			case "ppu":
				tps[0].Ppu = value1
				tps[1].Ppu = value2
				tps[2].Ppu = value3
				tps[3].Ppu = value4
				tps[4].Ppu = value5
				tps[5].Ppu = value6
				tps[6].Ppu = value7
				tps[7].Ppu = value8
				tps[8].Ppu = value9
				tps[9].Ppu = value10
			case "pmm":
				tps[0].Pmm = value1
				tps[1].Pmm = value2
				tps[2].Pmm = value3
				tps[3].Pmm = value4
				tps[4].Pmm = value5
				tps[5].Pmm = value6
				tps[6].Pmm = value7
				tps[7].Pmm = value8
				tps[8].Pmm = value9
				tps[9].Pmm = value10
			case "pk":
				tps[0].Pk = value1
				tps[1].Pk = value2
				tps[2].Pk = value3
				tps[3].Pk = value4
				tps[4].Pk = value5
				tps[5].Pk = value6
				tps[6].Pk = value7
				tps[7].Pk = value8
				tps[8].Pk = value9
				tps[9].Pk = value10
			case "eng":
				tps[0].Eng = value1
				tps[1].Eng = value2
				tps[2].Eng = value3
				tps[3].Eng = value4
				tps[4].Eng = value5
				tps[5].Eng = value6
				tps[6].Eng = value7
				tps[7].Eng = value8
				tps[8].Eng = value9
				tps[9].Eng = value10
			}
		}
		for _, i := range []string{"sos", "sej", "geo", "eko"} {
			soshum1, soshum2, soshum3, soshum4, soshum5, soshum6, soshum7, soshum8, soshum9, soshum10 := s.repo.Distribution(TestID, i)
			switch i {
			case "sos":
				soshum[0].Sos = soshum1
				soshum[1].Sos = soshum2
				soshum[2].Sos = soshum3
				soshum[3].Sos = soshum4
				soshum[4].Sos = soshum5
				soshum[5].Sos = soshum6
				soshum[6].Sos = soshum7
				soshum[7].Sos = soshum8
				soshum[8].Sos = soshum9
				soshum[9].Sos = soshum10
			case "sej":
				soshum[0].Sej = soshum1
				soshum[1].Sej = soshum2
				soshum[2].Sej = soshum3
				soshum[3].Sej = soshum4
				soshum[4].Sej = soshum5
				soshum[5].Sej = soshum6
				soshum[6].Sej = soshum7
				soshum[7].Sej = soshum8
				soshum[8].Sej = soshum9
				soshum[9].Sej = soshum10
			case "geo":
				soshum[0].Geo = soshum1
				soshum[1].Geo = soshum2
				soshum[2].Geo = soshum3
				soshum[3].Geo = soshum4
				soshum[4].Geo = soshum5
				soshum[5].Geo = soshum6
				soshum[6].Geo = soshum7
				soshum[7].Geo = soshum8
				soshum[8].Geo = soshum9
				soshum[9].Geo = soshum10
			case "eko":
				soshum[0].Eko = soshum1
				soshum[1].Eko = soshum2
				soshum[2].Eko = soshum3
				soshum[3].Eko = soshum4
				soshum[4].Eko = soshum5
				soshum[5].Eko = soshum6
				soshum[6].Eko = soshum7
				soshum[7].Eko = soshum8
				soshum[8].Eko = soshum9
				soshum[9].Eko = soshum10
			}
		}
	}else if pantek.Types == 3 {
		avg1, avg2, avg3, avg4, avg5, avg6, avg7, avg8, avg9, avg10 := s.repo.AVGDistributionTps(TestID)
		tps[0].Avg = avg1
		tps[1].Avg = avg2
		tps[2].Avg = avg3
		tps[3].Avg = avg4
		tps[4].Avg = avg5
		tps[5].Avg = avg6
		tps[6].Avg = avg7
		tps[7].Avg = avg8
		tps[8].Avg = avg9
		tps[9].Avg = avg10
		avgSa1, avgSa2, avgSa3, avgSa4, avgSa5, avgSa6, avgSa7, avgSa8, avgSa9, avgSa10 := s.repo.AVGDistributionSaintek(TestID)
		saintek[0].Avg = avgSa1
		saintek[1].Avg = avgSa2
		saintek[2].Avg = avgSa3
		saintek[3].Avg = avgSa4
		saintek[4].Avg = avgSa5
		saintek[5].Avg = avgSa6
		saintek[6].Avg = avgSa7
		saintek[7].Avg = avgSa8
		saintek[8].Avg = avgSa9
		saintek[9].Avg = avgSa10
		avgSo1, avgSo2, avgSo3, avgSo4, avgSo5, avgSo6, avgSo7, avgSo8, avgSo9, avgSo10 := s.repo.AVGDistributionSoshum(TestID)
		soshum[0].Avg = avgSo1
		soshum[1].Avg = avgSo2
		soshum[2].Avg = avgSo3
		soshum[3].Avg = avgSo4
		soshum[4].Avg = avgSo5
		soshum[5].Avg = avgSo6
		soshum[6].Avg = avgSo7
		soshum[7].Avg = avgSo8
		soshum[8].Avg = avgSo9
		soshum[9].Avg = avgSo10
		for _, i := range []string{"pu", "ppu", "pmm", "pk", "eng"} {
			value1, value2, value3, value4, value5, value6, value7, value8, value9, value10 := s.repo.Distribution(TestID, i)
			switch i {
			case "pu":
				tps[0].Pu = value1
				tps[1].Pu = value2
				tps[2].Pu = value3
				tps[3].Pu = value4
				tps[4].Pu = value5
				tps[5].Pu = value6
				tps[6].Pu = value7
				tps[7].Pu = value8
				tps[8].Pu = value9
				tps[9].Pu = value10
			case "ppu":
				tps[0].Ppu = value1
				tps[1].Ppu = value2
				tps[2].Ppu = value3
				tps[3].Ppu = value4
				tps[4].Ppu = value5
				tps[5].Ppu = value6
				tps[6].Ppu = value7
				tps[7].Ppu = value8
				tps[8].Ppu = value9
				tps[9].Ppu = value10
			case "pmm":
				tps[0].Pmm = value1
				tps[1].Pmm = value2
				tps[2].Pmm = value3
				tps[3].Pmm = value4
				tps[4].Pmm = value5
				tps[5].Pmm = value6
				tps[6].Pmm = value7
				tps[7].Pmm = value8
				tps[8].Pmm = value9
				tps[9].Pmm = value10
			case "pk":
				tps[0].Pk = value1
				tps[1].Pk = value2
				tps[2].Pk = value3
				tps[3].Pk = value4
				tps[4].Pk = value5
				tps[5].Pk = value6
				tps[6].Pk = value7
				tps[7].Pk = value8
				tps[8].Pk = value9
				tps[9].Pk = value10
			case "eng":
				tps[0].Eng = value1
				tps[1].Eng = value2
				tps[2].Eng = value3
				tps[3].Eng = value4
				tps[4].Eng = value5
				tps[5].Eng = value6
				tps[6].Eng = value7
				tps[7].Eng = value8
				tps[8].Eng = value9
				tps[9].Eng = value10
			}
		}
		for _, i := range []string{"ma", "fi", "ki", "bi"} {
			saintek1, saintek2, saintek3, saintek4, saintek5, saintek6, saintek7, saintek8, saintek9, saintek10 := s.repo.Distribution(TestID, i)
			switch i {
			case "ma":
				saintek[0].Ma = saintek1
				saintek[1].Ma = saintek2
				saintek[2].Ma = saintek3
				saintek[3].Ma = saintek4
				saintek[4].Ma = saintek5
				saintek[5].Ma = saintek6
				saintek[6].Ma = saintek7
				saintek[7].Ma = saintek8
				saintek[8].Ma = saintek9
				saintek[9].Ma = saintek10
			case "fi":
				saintek[0].Fi = saintek1
				saintek[1].Fi = saintek2
				saintek[2].Fi = saintek3
				saintek[3].Fi = saintek4
				saintek[4].Fi = saintek5
				saintek[5].Fi = saintek6
				saintek[6].Fi = saintek7
				saintek[7].Fi = saintek8
				saintek[8].Fi = saintek9
				saintek[9].Fi = saintek10
			case "ki":
				saintek[0].Ki = saintek1
				saintek[1].Ki = saintek2
				saintek[2].Ki = saintek3
				saintek[3].Ki = saintek4
				saintek[4].Ki = saintek5
				saintek[5].Ki = saintek6
				saintek[6].Ki = saintek7
				saintek[7].Ki = saintek8
				saintek[8].Ki = saintek9
				saintek[9].Ki = saintek10
			case "bi":
				saintek[0].Bi = saintek1
				saintek[1].Bi = saintek2
				saintek[2].Bi = saintek3
				saintek[3].Bi = saintek4
				saintek[4].Bi = saintek5
				saintek[5].Bi = saintek6
				saintek[6].Bi = saintek7
				saintek[7].Bi = saintek8
				saintek[8].Bi = saintek9
				saintek[9].Bi = saintek10
			}
		}
		for _, i := range []string{"sos", "sej", "geo", "eko"} {
			soshum1, soshum2, soshum3, soshum4, soshum5, soshum6, soshum7, soshum8, soshum9, soshum10 := s.repo.Distribution(TestID, i)
			switch i {
			case "sos":
				soshum[0].Sos = soshum1
				soshum[1].Sos = soshum2
				soshum[2].Sos = soshum3
				soshum[3].Sos = soshum4
				soshum[4].Sos = soshum5
				soshum[5].Sos = soshum6
				soshum[6].Sos = soshum7
				soshum[7].Sos = soshum8
				soshum[8].Sos = soshum9
				soshum[9].Sos = soshum10
			case "sej":
				soshum[0].Sej = soshum1
				soshum[1].Sej = soshum2
				soshum[2].Sej = soshum3
				soshum[3].Sej = soshum4
				soshum[4].Sej = soshum5
				soshum[5].Sej = soshum6
				soshum[6].Sej = soshum7
				soshum[7].Sej = soshum8
				soshum[8].Sej = soshum9
				soshum[9].Sej = soshum10
			case "geo":
				soshum[0].Geo = soshum1
				soshum[1].Geo = soshum2
				soshum[2].Geo = soshum3
				soshum[3].Geo = soshum4
				soshum[4].Geo = soshum5
				soshum[5].Geo = soshum6
				soshum[6].Geo = soshum7
				soshum[7].Geo = soshum8
				soshum[8].Geo = soshum9
				soshum[9].Geo = soshum10
			case "eko":
				soshum[0].Eko = soshum1
				soshum[1].Eko = soshum2
				soshum[2].Eko = soshum3
				soshum[3].Eko = soshum4
				soshum[4].Eko = soshum5
				soshum[5].Eko = soshum6
				soshum[6].Eko = soshum7
				soshum[7].Eko = soshum8
				soshum[8].Eko = soshum9
				soshum[9].Eko = soshum10
			}
		}
	}
	wibu := s.testRepo.NameandType(TestID)
	if wibu.Types == 1 {
		output := append(output, tps[0], tps[1], tps[2], tps[3], tps[4], tps[5], tps[6], tps[7], tps[8], tps[9])
		outputSa := append(outputSa, saintek[0], saintek[1], saintek[2], saintek[3], saintek[4], saintek[5], saintek[6], saintek[7], saintek[8], saintek[9])
		didi.Tps = output
		didi.Saintek = outputSa
	} else if wibu.Types == 2 {
		output := append(output, tps[0], tps[1], tps[2], tps[3], tps[4], tps[5], tps[6], tps[7], tps[8], tps[9])
		outputSo := append(outputSo, soshum[0], soshum[1], soshum[2], soshum[3], soshum[4], soshum[5], soshum[6], soshum[7], soshum[8], soshum[9])
		didi.Tps = output
		didi.Soshum = outputSo
	} else if wibu.Types == 3 {
		output := append(output, tps[0], tps[1], tps[2], tps[3], tps[4], tps[5], tps[6], tps[7], tps[8], tps[9])
		outputSa := append(outputSa, saintek[0], saintek[1], saintek[2], saintek[3], saintek[4], saintek[5], saintek[6], saintek[7], saintek[8], saintek[9])
		outputSo := append(outputSo, soshum[0], soshum[1], soshum[2], soshum[3], soshum[4], soshum[5], soshum[6], soshum[7], soshum[8], soshum[9])
		didi.Tps = output
		didi.Saintek = outputSa
		didi.Soshum = outputSo
	} else if wibu.Types == 0 {
		output := append(output, tps[0], tps[1], tps[2], tps[3], tps[4], tps[5], tps[6], tps[7], tps[8], tps[9])
		didi.Tps = output
	}

	return didi

}
func (s *scoreService) GetQuartile(TestID int) entity.FQuartile {
	var squartile entity.QSubtests
	var quartile []float64
	for _, subtest := range []string{"pu", "ppu", "pmm", "pk", "eng", "all"} {
		scores := s.repo.FindScoreByTestIDAndType(TestID, subtest)
		sah := len(scores)
		if len(scores) > 0 {
			q := []float64{scores[0].Score, scores[int(float64(sah/4))].Score, scores[int(float64(sah/2))].Score, scores[int(float64(sah*3/4))].Score, scores[sah-1].Score}
			switch subtest {
			case "pu":
				squartile.Pu = q
			case "ppu":
				squartile.Ppu = q
			case "pmm":
				squartile.Pmm = q
			case "pk":
				squartile.Pk = q
			case "eng":
				squartile.Eng = q
			case "all":
				quartile = q
			}
		}
	}

	return entity.FQuartile{quartile, squartile}
}
func (s *scoreService) QuartilTPS(TestID int) []entity.QSubtestsTps {
	var squartile [5]entity.QSubtestsTps
	var quartiletps []entity.QSubtestsTps
	squartile[0].Subject = "PU"
	squartile[1].Subject = "PPU"
	squartile[2].Subject = "PMM"
	squartile[3].Subject = "PK"
	squartile[4].Subject = "ENG"
	for _, subtest := range []string{"pu", "ppu", "pmm", "pk", "eng"} {
		scores := s.repo.Quartile(TestID, subtest)
		sah := len(scores)
		if len(scores) > 0 {
			q := []float64{scores[0].Score, scores[int(float64(sah/4))].Score, scores[int(float64(sah/2))].Score, scores[int(float64(sah*3/4))].Score, scores[sah-1].Score}
			switch subtest {
			case "pu":
				squartile[0].Q0 = q[4]
				squartile[0].Q1 = q[3]
				squartile[0].Q2 = q[2]
				squartile[0].Q3 = q[1]
				squartile[0].Q4 = q[0]
			case "ppu":
				squartile[1].Q0 = q[4]
				squartile[1].Q1 = q[3]
				squartile[1].Q2 = q[2]
				squartile[1].Q3 = q[1]
				squartile[1].Q4 = q[0]
			case "pmm":
				
				squartile[2].Q0 = q[4]
				squartile[2].Q1 = q[3]
				squartile[2].Q2 = q[2]
				squartile[2].Q3 = q[1]
				squartile[2].Q4 = q[0]
			case "pk":
				
				squartile[3].Q0 = q[4]
				squartile[3].Q1 = q[3]
				squartile[3].Q2 = q[2]
				squartile[3].Q3 = q[1]
				squartile[3].Q4 = q[0]
			case "eng":
				
				squartile[4].Q0 = q[4]
				squartile[4].Q1 = q[3]
				squartile[4].Q2 = q[2]
				squartile[4].Q3 = q[1]
				squartile[4].Q4 = q[0]
			}
		}
	}
	quartiletps = append(quartiletps,squartile[0],squartile[1],squartile[2],squartile[3],squartile[4])
	return quartiletps
}
func (s *scoreService) QuartilSaintek(TestID int) []entity.QSubtestsSaintek {
	var squartile [4]entity.QSubtestsSaintek
	var quartilesaintek []entity.QSubtestsSaintek
	squartile[0].Subject = "MA"
	squartile[1].Subject = "FI"
	squartile[2].Subject = "KI"
	squartile[3].Subject = "BI"
	for _, subtest := range []string{"ma", "fi", "ki", "bi"} {
		scores := s.repo.Quartile(TestID, subtest)
		sah := len(scores)
		if len(scores) > 0 {
			q := []float64{scores[0].Score, scores[int(float64(sah/4))].Score, scores[int(float64(sah/2))].Score, scores[int(float64(sah*3/4))].Score, scores[sah-1].Score}
			switch subtest {
			case "ma":
				
				squartile[0].Q0 = q[4]
				squartile[0].Q1 = q[3]
				squartile[0].Q2 = q[2]
				squartile[0].Q3 = q[1]
				squartile[0].Q4 = q[0]
			case "fi":
				
				squartile[1].Q0 = q[4]
				squartile[1].Q1 = q[3]
				squartile[1].Q2 = q[2]
				squartile[1].Q3 = q[1]
				squartile[1].Q4 = q[0]
			case "ki":
				
				squartile[2].Q0 = q[4]
				squartile[2].Q1 = q[3]
				squartile[2].Q2 = q[2]
				squartile[2].Q3 = q[1]
				squartile[2].Q4 = q[0]
			case "bi":
				
				squartile[3].Q0 = q[4]
				squartile[3].Q1 = q[3]
				squartile[3].Q2 = q[2]
				squartile[3].Q3 = q[1]
				squartile[3].Q4 = q[0]
			}
		}
	}
	quartilesaintek = append(quartilesaintek,squartile[0],squartile[1],squartile[2],squartile[3])
	return quartilesaintek
}
func (s *scoreService) QuartilSoshum(TestID int) []entity.QSubtestsSoshum {
	var squartile [4]entity.QSubtestsSoshum
	var quartilesoshum []entity.QSubtestsSoshum
	squartile[0].Subject = "SOS"
	squartile[1].Subject = "SEJ"
	squartile[2].Subject = "EKO"
	squartile[3].Subject = "GEO"
	for _, subtest := range []string{"sos", "sej", "eko", "geo"} {
		scores := s.repo.Quartile(TestID, subtest)
		sah := len(scores)
		if len(scores) > 0 {
			q := []float64{scores[0].Score, scores[int(float64(sah/4))].Score, scores[int(float64(sah/2))].Score, scores[int(float64(sah*3/4))].Score, scores[sah-1].Score}
			switch subtest {
			case "sos":
				
				squartile[0].Q0 = q[4]
				squartile[0].Q1 = q[3]
				squartile[0].Q2 = q[2]
				squartile[0].Q3 = q[1]
				squartile[0].Q4 = q[0]
			case "sej":
				
				squartile[1].Q0 = q[4]
				squartile[1].Q1 = q[3]
				squartile[1].Q2 = q[2]
				squartile[1].Q3 = q[1]
				squartile[1].Q4 = q[0]
			case "eko":
				
				squartile[2].Q0 = q[4]
				squartile[2].Q1 = q[3]
				squartile[2].Q2 = q[2]
				squartile[2].Q3 = q[1]
				squartile[2].Q4 = q[0]
			case "geo":
				
				squartile[3].Q0 = q[4]
				squartile[3].Q1 = q[3]
				squartile[3].Q2 = q[2]
				squartile[3].Q3 = q[1]
				squartile[3].Q4 = q[0]
			}
		}
	}
	quartilesoshum = append(quartilesoshum,squartile[0],squartile[1],squartile[2],squartile[3])
	return quartilesoshum
}
func (s *scoreService) AllQuartil(TestID int) entity.Quartile {
	var i entity.Quartile
	wibu := s.testRepo.NameandType(TestID)
	if wibu.Types == 0 {
		i.Tps = s.QuartilTPS(TestID)
	} else if wibu.Types == 1 {
		i.Tps = s.QuartilTPS(TestID)
		i.Saintek = s.QuartilSaintek(TestID)
	} else if wibu.Types == 2 {
		i.Tps = s.QuartilTPS(TestID)
		i.Soshum = s.QuartilSoshum(TestID)
	} else if wibu.Types == 3 {
		i.Tps = s.QuartilTPS(TestID)
		i.Saintek = s.QuartilSaintek(TestID)
		i.Soshum = s.QuartilSoshum(TestID)
	} else {
	}
	
	
	
	return i
}
func (s *scoreService) GetRank(TestID int, StudentID int) entity.Rank {
	pref := s.testAuthRepo.FindTestAuthByStudentIDAndTestID(StudentID, TestID)
	var (
		rank    entity.Rank
		counter int = 0
	)

	for _, one := range s.repo.GroupScoreByTestID(TestID) {
		counter++
		if int(one.StudentID) == StudentID {
			rank.Public.Rank = counter
		}
	}
	rank.Public.Participant = counter

	if pref.Pref1Uni != 0 && pref.Pref1Prodi != 0 {
		counter = 0
		for _, one := range s.repo.GroupScoreByTestIDAndPref(TestID, int(pref.Pref1Uni), int(pref.Pref1Prodi)) {
			counter++
			if int(one.StudentID) == StudentID {
				rank.Pref1.Rank = counter
			}
		}
		rank.Pref1.Participant = counter
	}

	if pref.Pref2Uni != 0 && pref.Pref2Prodi != 0 {
		counter = 0
		for _, one := range s.repo.GroupScoreByTestIDAndPref(TestID, int(pref.Pref2Uni), int(pref.Pref2Prodi)) {
			counter++
			if int(one.StudentID) == StudentID {
				rank.Pref2.Rank = counter
			}
		}
		rank.Pref2.Participant = counter
	}

	return rank
}

//////////////////////////
// BASIC CRUD OPERATION //
//////////////////////////

func (s *scoreService) FindAllScore() []entity.ScoreMin {
	return s.repo.FindAllScoreMin()
}

func (s *scoreService) FindScoreByID(id int) entity.Score {
	return s.repo.FindScoreByID(id)
}

func (s *scoreService) CreateScore(score entity.Score) entity.Score {
	return s.repo.CreateScore(score)
}

func (s *scoreService) UpdateScore(score entity.Score) entity.Score {
	s.repo.UpdateScore(score)
	return score
}

func (s *scoreService) DeleteScore(score entity.Score) {
	s.repo.DeleteScore(score)
}

func (s *scoreService) Graphic(test entity.Test, score entity.Score) []entity.IkutBattle {
	return s.repo.Graphic(test, score)
}

func (s *scoreService) Jawab(jawab entity.Jawaban) entity.Jawaban{
	return s.repo.Jawab(jawab)
}
func (s *scoreService) AmbilID(id int) uint64  {
	test := s.testRepo.NameandID(id)
	idtest := test.ID
	return idtest
}

func (s *scoreService) BikinArray(id int) []uint64{
	var wibu []uint64
	finished := s.repo.GroupScoreByStudentID(id)
	for _,i := range finished {
		idnya := s.AmbilID(int(i.ID))
		wibu = append(wibu, idnya)
	}
	return wibu
}

// func (s *scoreService) CariElemen(id int,testID int){
// 	array := s.BikinArray(id)
// }

func (s *scoreService) BikinPesan(id int,testID int) {

}
