package service

import (
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	scorerepo "studybuddy-backend-fast/api/repository/score"
	repository "studybuddy-backend-fast/api/repository/test"
)

type TestService interface {
	FindAllTest(pagination map[string]int) []entity.TestMin
	FindTest(param string) entity.TestQuestionMin
	CreateTest(newTest entity.Test) entity.Test
	UpdateTest(newTestData entity.Test)
	DeleteTest(test entity.Test)

	FindAllTestMin() []entity.TestMin
	FindTestMin(param string) entity.TestMin
	FindTestFull(param string) entity.TestQuestionAndSolution
	LastBattle(test entity.Test) entity.Last
	BattleIkut(id int) []entity.IkutTest
	AllAvailableTest(id int) []entity.AvaibleTest
	CheckTest(id int) entity.AvaibleTest
	removeDuplicatesTest(elements []entity.AvaibleTest) []entity.AvaibleTest
	Available(id int) []entity.AvaibleBattle
	SoalBattle(ID int, subtest string) entity.QuestionSpace
}

type testService struct {
	testRepo repository.TestRepo
	repo     scorerepo.ScoreRepo
}

func NewTestService(repo repository.TestRepo, score scorerepo.ScoreRepo) TestService {
	return &testService{
		testRepo: repo,
		repo:     score,
	}
}

func (s *testService) FindAllTest(pagination map[string]int) []entity.TestMin {
	return s.testRepo.FindTestWithFilter(pagination)
}

func (s *testService) FindTestByID(id int) entity.TestQuestionMin {
	return s.testRepo.FindTestByID(id)
}

func (s *testService) FindTest(param string) entity.TestQuestionMin {
	id, err := strconv.Atoi(param)
	if err != nil {
		return s.testRepo.FindTestBySlug(param)
	}
	return s.testRepo.FindTestByID(id)
}

func (s *testService) FindTestFull(param string) entity.TestQuestionAndSolution {
	id, err := strconv.Atoi(param)
	if err != nil {
		return s.testRepo.FindTestBySlugWithSolution(param)
	}
	return s.testRepo.FindTestByIDWithSolution(id)
}

func (s *testService) CreateTest(newTest entity.Test) entity.Test {
	return s.testRepo.CreateTest(newTest)
}

func (s *testService) removeDuplicatesTest(elements []entity.AvaibleTest) []entity.AvaibleTest {
	encountered := map[entity.AvaibleTest]bool{}
	result := []entity.AvaibleTest{}
	for key, _ := range elements {
		if encountered[elements[key]] != true {
			encountered[elements[key]] = true
			result = append(result, elements[key])
		}

	}
	return result
}

func (s *testService) CheckTest(id int) entity.AvaibleTest {
	return s.testRepo.AvailableByID(id)
}

func (s *testService) AllAvailableTest(id int) []entity.AvaibleTest{
	var test []entity.AvaibleTest
	
	alltest := s.testRepo.ArrayFindAvaibleBattle()
	isPaid := s.repo.GroupScoreByStudentID(id)
	test = append(test, alltest...)
	for _, t := range isPaid {
		fscore := s.CheckTest(int(t.TestID))
		test = append(test, fscore)
	}
	result := s.removeDuplicatesTest(test)
	return result
}
func (s *testService) Available(id int) []entity.AvaibleBattle {
	var battle []entity.AvaibleBattle
	var avaible entity.AvaibleBattle
	result := s.AllAvailableTest(id)
	for i:=0; i < len(result); i++ {
		avaible.Name = result[i].Name
		avaible.BattleID = result[i].ID
		avaible.Image = result[i].Image
		avaible.Jumlahsoal = result[i].Pu + result[i].Ppu + result[i].Pk + result[i].Pmm + result[i].Eng + result[i].Ma + result[i].Fi + result[i].Ki + result[i].Bi + result[i].Sos + result[i].Sej + result[i].Geo + result[i].Eko
		jadwal := result[i].ScheduledAt.UTC()
		start := result[i].ScheduledAt
		end := result[i].EndsAt
		y,m,d := jadwal.Date()
		durasi := end.Sub(start)
		menit := durasi.Minutes()
		avaible.Durasi = menit
		avaible.ScheduledAt = strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d)
		h := result[i].ScheduledAt.Hour()
		minute := result[i].ScheduledAt.Minute()
		avaible.Waktu = strconv.Itoa(h) + ":" + strconv.Itoa(minute)
		avaible.Types = result[i].Types
		avaible.Dibeli = false
		battle = append(battle,avaible)
	}
	for i:= 0; i< len(battle); i++ {
		isPaid := s.repo.GroupScoreByStudentID(id)
		for j:=0; j< len(isPaid);j++{
		fscore := s.CheckTest(int(isPaid[j].TestID))
		if battle[i].Name ==  fscore.Name{
			battle[i].Dibeli = true
		}
		if battle[i].Types == 1{
			battle[i].Durasi = 180
		}else if battle[i].Types == 2{
			battle[i].Durasi = 180
		}else if battle[i].Types == 0 {
			battle[i].Durasi = 105
		}
		}		
	}
	return battle
}

func (s *testService) IkutTest(id int) entity.IkutTest {
	var i entity.IkutTest
	test := s.testRepo.NameandID(id)
	i.Name = test.Name
	i.BattleID = test.ID
	return i
}

func (s *testService) BattleIkut(id int) []entity.IkutTest {
	var fscores []entity.IkutTest
	finishedTests := s.repo.GroupScoreByStudentID(id)
	for _, t := range finishedTests {
		fscore := s.IkutTest(int(t.TestID))
		fscores = append(fscores, fscore)
	}
	return fscores
}

func (s *testService) UpdateTest(newTestData entity.Test) {
	s.testRepo.UpdateTest(newTestData)
}

func (s *testService) DeleteTest(test entity.Test) {
	s.testRepo.DeleteTest(test)
}

func (s *testService) FindAllTestMin() []entity.TestMin {
	return s.testRepo.FindAllTestMin()
}

func (s *testService) FindTestMin(param string) entity.TestMin {
	if id, err := strconv.Atoi(param); err == nil {
		return s.testRepo.FindTestByIDMin(id)
	}
	return s.testRepo.FindTestBySlugMin(param)
}
func (s *testService) LastBattle(test entity.Test) entity.Last {
	return s.testRepo.LastBattle(test)
}

func(s *testService) SoalBattle(ID int, subtest string) entity.QuestionSpace {
	var hehe entity.QuestionSpace
	var huhu entity.Pertanyaan
	var soal []entity.Pertanyaan
	sum := 0
	data := s.testRepo.SoalBattle(ID,subtest)
	for i := 0; i<len(data);i++{
		huhu.Urutan = i + 1
		huhu.Tests = data[i].Question
		huhu.Jawab = data[i].Answer
		sum += data[i].Duration
		hehe.Waktu = sum
		soal = append(soal, huhu)
	}
	hehe.Type = subtest
	hehe.Tests = soal
	return hehe
}
