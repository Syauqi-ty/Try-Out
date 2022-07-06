package misc

import (
	"strconv"
	"strings"
	repository "studybuddy-backend-fast/api/repository/test"
	"time"
)

var repo repository.TestRepo = repository.NewTestRepo()

func GenerateExternalID(studentID int, testID int, method string) string {
	test := repo.FindTestByID(testID)
	if test.ID == 0 {
		return ""
	}
	fslug := strings.Replace(test.Slug, "_", "-", -1)
	fstudentID := strconv.Itoa(studentID)
	now := strconv.Itoa(int(time.Now().Unix()))

	return method + "-" + fslug + "-" + fstudentID + "-" + now
}
