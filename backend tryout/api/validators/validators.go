package validators

import (
	"github.com/gin-gonic/gin"
)

func StudentRegistration(c *gin.Context) error {
	var res RegistrationBody
	if err := c.ShouldBindJSON(&res); err != nil {
		return err
	}
	return nil
}

func UpdateStudent(c *gin.Context) error {
	var res UpdateStudentBody
	if err := c.ShouldBindJSON(&res); err != nil {
		return err
	}
	return nil
}

func DeleteRequest(c *gin.Context) error {
	var res DeleteBody
	if err := c.ShouldBindJSON(&res); err != nil {
		return err
	}
	return nil
}

func AuthRequest(c *gin.Context) error {
	var res AuthBody
	if err := c.ShouldBindJSON(&res); err != nil {
		return err
	}
	return nil
}

func TestAuthRequest(c *gin.Context) error {
	var res TestAuthBody
	if err := c.ShouldBindJSON(&res); err != nil {
		return err
	}
	return nil
}

func StaffRegistration(c *gin.Context) error {
	var res StaffRegistrationBody
	if err := c.ShouldBindJSON(res); err != nil {
		return err
	}
	return nil
}

func UpdateStaff(c *gin.Context) error {
	var res UpdateStaffBody
	if err := c.ShouldBindJSON(&res); err != nil {
		return err
	}
	return nil
}

func CreateTest(c *gin.Context) error {
	var res CreateTestBody
	if err := c.ShouldBindJSON(&res); err != nil {
		return err
	}
	return nil
}

func UpdateTest(c *gin.Context) error {
	var res UpdateTestBody
	if err := c.ShouldBindJSON(&res); err != nil {
		return err
	}
	return nil
}

func CreateQuestion(c *gin.Context) error {
	var res CreateQuestionBody
	if err := c.ShouldBindJSON(res); err != nil {
		return err
	}
	return nil
}

func UpdateQuestion(c *gin.Context) error {
	var res UpdateQuestionBody
	if err := c.ShouldBindJSON(&res); err != nil {
		return err
	}
	return nil
}

func CreateSolution(c *gin.Context) error {
	var res CreateSolutionBody
	if err := c.ShouldBindJSON(&res); err != nil {
		return err
	}
	return nil
}

func UpdateSolution(c *gin.Context) error {
	var res UpdateSolutionBody
	if err := c.ShouldBindJSON(&res); err != nil {
		return err
	}
	return nil
}
