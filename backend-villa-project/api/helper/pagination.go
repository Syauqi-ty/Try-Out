package helper

import (
	"strconv"
	"villa-akmali/api/model"

	"github.com/gin-gonic/gin"
)


func Pagination(c *gin.Context) model.Pagination{
	limit := 100
	page := 1
	sort := "id desc"
	query := c.Request.URL.Query()
	for key,value := range query {
		queryvalue := value[len(value)-1]
		switch key {
		case "limit" :
			limit,_ = strconv.Atoi(queryvalue)
			break
		case "page" :
			page,_ = strconv.Atoi(queryvalue)
			break
		case "sort" :
			sort = queryvalue
			break
		}
	}
	return model.Pagination{
		Limit: limit,
		Page : page,
		Sort: sort,
	}
}  