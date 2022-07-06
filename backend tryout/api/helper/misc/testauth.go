package misc

import (
	"github.com/google/uuid"
	"strings"
	entity "studybuddy-backend-fast/api/entity"
	scoreRepository "studybuddy-backend-fast/api/repository/score"
	testAuthRepository "studybuddy-backend-fast/api/repository/testauth"
)

var (
	scoreRepo    scoreRepository.ScoreRepo       = scoreRepository.NewScoreRepo()
	testAuthRepo testAuthRepository.TestAuthRepo = testAuthRepository.NewTestAuthRepo()
)

func CreateTestUser(TestID int, UserID int) {
	res := strings.Split(uuid.New().String(), "-")
	username := res[0]
	password := res[4]

	authData := entity.TestAuth{
		UsernameTO: username,
		PasswordTO: password,
		StudentID:  uint64(UserID),
		TestID:     uint64(TestID),
	}
	testAuthRepo.CreateTestAuth(authData)

	scoreData := entity.Score{
		StudentID: uint64(UserID),
		TestID:    uint64(TestID),
	}

	for _, subtest := range []string{"pu", "ppu", "pmm", "pk", "eng"} {
		scoreData.Type = subtest
		scoreRepo.CreateScore(scoreData)
	}
}
