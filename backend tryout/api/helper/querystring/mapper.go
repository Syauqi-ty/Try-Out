package querystring

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetStudentMapper(c *gin.Context) (map[string]interface{}, map[string]int) {
	q := make(map[string]interface{})
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	for _, param := range []string{"school", "student_class"} {
		if c.DefaultQuery(param, "") != "" {
			q[param] = c.DefaultQuery(param, "")
		}
	}

	return q, map[string]int{"page": page, "limit": limit}
}

func GetStaffMapper(c *gin.Context) (map[string]interface{}, map[string]int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	access := c.DefaultQuery("access", "")

	if access == "" {
		return map[string]interface{}{}, map[string]int{"page": page, "limit": limit}
	}
	return map[string]interface{}{"access": access}, map[string]int{"page": page, "limit": limit}
}

func GetParentMapper(c *gin.Context) (map[string]interface{}, map[string]int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	school := c.DefaultQuery("school", "")

	if school == "" {
		return map[string]interface{}{}, map[string]int{"page": page, "limit": limit}
	}
	return map[string]interface{}{"school": school}, map[string]int{"page": page, "limit": limit}
}

func GetTestMapper(c *gin.Context) map[string]int {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	return map[string]int{"page": page, "limit": limit}
}

func GetQuestionMapper(c *gin.Context) (map[string]interface{}, map[string]int) {
	q := make(map[string]interface{})
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	for _, param := range []string{"type", "creator_id"} {
		if c.DefaultQuery(param, "") == "" {
			if param == "creator_id" {
				q[param] = c.DefaultQuery(param, c.DefaultQuery(param, ""))
			} else {
				q[param] = c.DefaultQuery(param, c.DefaultQuery(param, ""))
			}
		}
	}

	return q, map[string]int{"page": page, "limit": limit}
}

func GetSolutionMapper(c *gin.Context) (map[string]interface{}, map[string]int) {
	q := make(map[string]interface{})
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	for _, param := range []string{"type", "creator_id"} {
		if c.DefaultQuery(param, "") == "" {
			if param == "creator_id" {
				q[param] = c.DefaultQuery(param, c.DefaultQuery(param, ""))
			} else {
				q[param] = c.DefaultQuery(param, c.DefaultQuery(param, ""))
			}
		}
	}

	return q, map[string]int{"page": page, "limit": limit}
}
