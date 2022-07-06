package controller

import (
	"fmt"
	"strconv"
	entity "studybuddy-backend-fast/api/entity"
	mapper "studybuddy-backend-fast/api/helper/querystring"
	service "studybuddy-backend-fast/api/services/student"
	testAuthService "studybuddy-backend-fast/api/services/testauth"

	"time"

	tokenoi "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

type StudentController interface {
	FindAll(ctx *gin.Context)
	FindOneByID(ctx *gin.Context)
	CreateStudent(ctx *gin.Context)
	UpdateStudent(ctx *gin.Context)
	DeleteStudent(ctx *gin.Context)
	LoginSpaceBaru(ctx *gin.Context)

	TestAuth(ctx *gin.Context)
	FindTestAuth(ctx *gin.Context)
	Forget(ctx *gin.Context)
}

type studentController struct {
	service       service.StudentService
	serveTestAuth testAuthService.TestAuthService
}

type deleteUserBody struct {
	id int
}

type TestAuthBody struct {
	UsernameTO string `json:"username_to"`
	PasswordTO string `json:"password_to"`
}

func NewStudentController(service service.StudentService, service2 testAuthService.TestAuthService) StudentController {
	return &studentController{
		service:       service,
		serveTestAuth: service2,
	}
}

func (c *studentController) FindTestAuth(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	testID, _ := strconv.Atoi(ctx.Param("id"))
	res := c.serveTestAuth.FindTestAuth(int(claims["user_id"].(float64)), testID)
	ctx.JSON(200, gin.H{"msg": "Test Auth found", "data": res})
}

func (c *studentController) TestAuth(ctx *gin.Context) {
	var req TestAuthBody
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(400, gin.H{"msg": "Bad request, missing fields"})
	} else {
		if token, student, test := c.serveTestAuth.TestAuthentication(req.UsernameTO, req.PasswordTO); student.ID != 0 {
			ctx.JSON(200, gin.H{
				"data": map[string]interface{}{
					"student": student.ID,
					"test":    test.ID,
				},
				"code":  200,
				"token": token,
			})
		} else {
			ctx.JSON(400, gin.H{"msg": "Invalid credentials"})
		}
	}
}
func (c *studentController) LoginSpaceBaru(ctx *gin.Context)  {
	var timezone string = viper.GetString("timezone")
	loc, _ := time.LoadLocation(timezone)
	var req entity.LoginSpace
	err := ctx.ShouldBind(&req)
	data := c.serveTestAuth.LoginBattleSpace(req)
	token := tokenoi.NewWithClaims(tokenoi.SigningMethodHS256, tokenoi.MapClaims{
		"exp":      time.Now().In(loc).Add(72 * time.Hour).Unix(),
		"orig_iat": time.Now().In(loc).Unix(),
		"status":   "student",
		"user_id": data.UserID,
	})
	tokenString, _ := token.SignedString([]byte(viper.GetString("secret")))
	if err != nil {
		ctx.JSON(400, gin.H{"msg": "Bad request, missing fields"})
	} else{
		if data.UserID == 0{
			ctx.JSON(400,gin.H{"msg":"Yang bener dong"})
		} else{
			ctx.JSON(200,gin.H{"data":data,"token":tokenString})
		}
	}
}

func (c *studentController) FindAll(ctx *gin.Context) {
	qstring, pag := mapper.GetStudentMapper(ctx)
	pagenum := strconv.Itoa(pag["page"])
	limnum := strconv.Itoa(pag["limit"])

	res := c.service.FindAllStudent(qstring, pag)
	ctx.JSON(200, gin.H{"msg": "On page " + pagenum + " with limit " + limnum, "data": res})
}

func (c *studentController) FindOneByID(ctx *gin.Context) {
	param := ctx.Param("id")
	if id, err := strconv.Atoi(param); err == nil {
		res := c.service.FindByID(id)
		ctx.JSON(200, gin.H{"msg": "Query successful", "data": res})
	} else {
		ctx.JSON(400, gin.H{"msg": "Bad request, please check the docs"})
	}
}

func (c *studentController) CreateStudent(ctx *gin.Context) {
	var newUser entity.Student
	ctx.ShouldBind(&newUser)
	if res := c.service.CreateStudent(newUser); res.ID == 0 {
		ctx.JSON(202, gin.H{"msg": "Credentials taken"})
	} else {
		ctx.JSON(200, gin.H{"msg": "Student registered", "data": res})
	}
}

func (c *studentController) UpdateStudent(ctx *gin.Context) {
	var newUserData entity.Student
	fmt.Println(newUserData)
	ctx.ShouldBindJSON(&newUserData)
	res := c.service.UpdateStudent(newUserData)
	ctx.JSON(200, gin.H{"msg": "Student data updated", "data": res})
}

func (c *studentController) DeleteStudent(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	var user entity.Student
	ctx.ShouldBind(&user)
	if claims["status"] != "admin" || uint64(claims["user_id"].(float64)) != user.ID {
		ctx.JSON(403, gin.H{"msg": "You are not allowed to delete this account"})
	} else {
		c.service.DeleteStudent(user)
		ctx.JSON(200, gin.H{"msg": "Student deleted"})
	}
}

func (c *studentController) Forget(ctx *gin.Context){
	var user entity.Student
	ctx.ShouldBindJSON(&user)
	if c.service.Forget(user)==0{
		ctx.JSON(200,c.service.Forget(user))
	} else if  c.service.Forget(user)==2 {
		ctx.JSON(200,c.service.Forget(user))
	}else{
		ctx.JSON(200,c.service.Forget(user))
	} 
}