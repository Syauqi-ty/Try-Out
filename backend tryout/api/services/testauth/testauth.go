package service

import (
	"fmt"
	entity "studybuddy-backend-fast/api/entity"
	studentRepository "studybuddy-backend-fast/api/repository/student"
	testRepository "studybuddy-backend-fast/api/repository/test"
	repository "studybuddy-backend-fast/api/repository/testauth"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var timezone string = viper.GetString("timezone")

type TestAuthService interface {
	FindTestAuth(StudentID int, TestID int) entity.TestAuth
	TestAuthentication(UsernameTO string, PasswordTO string) (string, entity.Student, entity.TestQuestionMin)
	ChangePref(TestAuth entity.SetPrefBody)
	LoginBattleSpace(login entity.LoginSpace) entity.ResponLogin
}

type testAuthService struct {
	repo        repository.TestAuthRepo
	studentRepo studentRepository.StudentRepo
	testRepo    testRepository.TestRepo
}

func NewTestAuthService(repo repository.TestAuthRepo, studentRepo studentRepository.StudentRepo, testRepo testRepository.TestRepo) TestAuthService {
	return &testAuthService{repo, studentRepo, testRepo}
}

func (s *testAuthService) FindTestAuth(StudentID int, TestID int) entity.TestAuth {
	return s.repo.FindTestAuthByStudentIDAndTestID(StudentID, TestID)
}

func (s *testAuthService) TestAuthentication(UsernameTO string, PasswordTO string) (string, entity.Student, entity.TestQuestionMin) {
	auth := s.repo.FindTestAuthByCreds(UsernameTO, PasswordTO)
	student := s.studentRepo.FindOneById(int(auth.StudentID))
	test := s.testRepo.FindTestByID(int(auth.TestID))

	loc, _ := time.LoadLocation(timezone)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().In(loc).Add(2 * time.Hour).Unix(),
		"orig_iat": time.Now().In(loc).Unix(),
		"status":   "student",
		"user_id":  auth.StudentID,
	})

	// Sign and get the complete encoded token as a string using the secret
	if tokenString, err := token.SignedString([]byte(viper.GetString("secret"))); err == nil {
		return tokenString, student, test
	}
	return "", entity.Student{}, entity.TestQuestionMin{}
}

func (s *testAuthService) ChangePref(TestAuth entity.SetPrefBody) {
	fmt.Println(TestAuth)
	testAuth := map[string]interface{}{
		"test_id":     TestAuth.TestID,
		"student_id":  TestAuth.StudentID,
		"pref1_uni":   0,
		"pref1_prodi": 0,
		"pref2_uni":   0,
		"pref2_prodi": 0,
	}

	if TestAuth.Pref1Uni != 0 {
		testAuth["pref1_uni"] = TestAuth.Pref1Uni
	}
	if TestAuth.Pref2Uni != 0 {
		testAuth["pref2_uni"] = TestAuth.Pref2Uni
	}
	if TestAuth.Pref1Prodi != 0 {
		testAuth["pref1_prodi"] = TestAuth.Pref1Prodi
	}
	if TestAuth.Pref2Prodi != 0 {
		testAuth["pref2_prodi"] = TestAuth.Pref2Prodi
	}
	s.repo.UpdateTestAuth(testAuth)
}

func (s *testAuthService) LoginBattleSpace(login entity.LoginSpace) entity.ResponLogin {
	var hehe entity.ResponLogin
	data := s.repo.LoginSpace(login)
	if data.StudentID == 0 || data.TestID == 0{
		hehe.UserID = 0
	}else{
		studentdata:= s.studentRepo.FindOneById(int(data.StudentID))
		testdata := s.testRepo.FindTestByID(int(data.TestID))
		hehe.UserID = int(studentdata.ID)
		hehe.TestID = int(data.TestID)
		hehe.Name = studentdata.Name
		hehe.DataTest.Name = testdata.Name
		hehe.DataTest.Image = testdata.Image
		hehe.DataTest.StartedAt = testdata.ScheduledAt
		hehe.DataTest.EndsAt = testdata.EndsAt
		if testdata.Types == 1{
			hehe.DataTest.Subtest = []string{"pu","pk","ppu","pmm","eng","ma","fi","ki","bi"}
		}else if testdata.Types == 2{
			hehe.DataTest.Subtest = []string{"pu","pk","ppu","pmm","eng","geo","sej","eko","sos"}
		}else{
			hehe.DataTest.Subtest = []string{"pu","pk","ppu","pmm","eng"}
		}
	}
	return hehe
}