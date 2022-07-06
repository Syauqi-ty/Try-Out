package auth

import (
	"fmt"
	parentRepository "studybuddy-backend-fast/api/repository/parent"
	staffRepository "studybuddy-backend-fast/api/repository/staff"
	studentRepository "studybuddy-backend-fast/api/repository/student"
	service "studybuddy-backend-fast/api/services/auth"
	validators "studybuddy-backend-fast/api/validators"
	config "studybuddy-backend-fast/config"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	tok "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	studentRepo studentRepository.StudentRepo = studentRepository.NewStudentRepo()
	staffRepo   staffRepository.StaffRepo     = staffRepository.NewStaffRepo()
	parentRepo  parentRepository.ParentRepo   = parentRepository.NewParentRepo()
)

func isAllowed(status string, access string) bool {
	a := map[string][]string{
		"student-access":   []string{"student", "admin"},
		"parent-access":    []string{"parent", "admin"},
		"guru-access":      []string{"guru", "admin"},
		"qualcon-access":   []string{"qualcon", "admin"},
		"staff-access":     []string{"admin", "qualcon", "guru"},
		"logged-in-access": []string{"admin", "qualcon", "guru", "student"},
		"admin-access":     []string{"admin"},
	}
	for i := 0; i < len(a[access]); i++ {
		if a[access][i] == status {
			return true
		}
	}
	return false
}

// Auth related structs
type Payload struct {
	UserID uint64
	Status string
}

// Custom auth handler

func New(access string) (*jwt.GinJWTMiddleware, error) {
	config.Init()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(viper.GetString(`secret`)),
		Timeout:     time.Hour,
		MaxRefresh:  30 * 24 * time.Hour,
		IdentityKey: "user_id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*Payload); ok {
				return jwt.MapClaims{
					"user_id": v.UserID,
					"status":  v.Status,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &Payload{
				UserID: uint64(claims["user_id"].(float64)),
				Status: claims["status"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var creds validators.AuthBody
			if err := c.ShouldBindJSON(&creds); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			fmt.Println(creds)

			ok, user := service.Auth(creds)
			if ok {
				if user.Staff.ID != 0 {
					return &Payload{UserID: user.Staff.ID, Status: user.Staff.Access}, nil
				} else if user.Parent.ID != 0 {
					return &Payload{UserID: user.Parent.ID, Status: "parent"}, nil
				} else {
					return &Payload{UserID: user.Student.ID, Status: "student"}, nil
				}
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*Payload); ok && isAllowed(v.Status, access) {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"msg": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		LoginResponse: func(c *gin.Context, code int, tokenString string, expire time.Time) {
			claims := tok.MapClaims{}
			tok.ParseWithClaims(tokenString, claims, func(token *tok.Token) (interface{}, error) {
				return []byte(viper.GetString("secret")), nil
			})

			c.JSON(200, gin.H{
				"data": map[string]interface{}{
					"user_id": int(claims["user_id"].(float64)),
					"status":  claims["status"].(string),
				},
				"code":   200,
				"token":  tokenString,
				"expire": expire.Format(time.RFC3339),
			})

		},
		TimeFunc: time.Now,
	})
	return authMiddleware, err
}
