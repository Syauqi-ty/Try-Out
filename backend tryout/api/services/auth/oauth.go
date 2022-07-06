package service

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	entity "studybuddy-backend-fast/api/entity"
	"time"
)

type OAuthResp struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

var (
	oauthStateString  string        = "pantat-kuda"
	timezone          string        = viper.GetString("timezone")
	googleOauthConfig oauth2.Config = oauth2.Config{
		RedirectURL:  viper.GetString("server.name") + "/api/v2/oauth/callback",
		ClientID:     viper.GetString("oauth.client_id"),
		ClientSecret: viper.GetString("oauth.client_secret"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
)

func GetAuthURL() string {
	return googleOauthConfig.AuthCodeURL(oauthStateString)
}

func HandleGoogleCallback(state string, code string) entity.Student {
	content, err := GetUserInfo(state, code)
	if err != nil {
		fmt.Println(err.Error())
		return entity.Student{}
	}

	return studentRepo.FindOneByUsernameOrEmail(content.Email)
}

func GetUserInfo(state string, code string) (OAuthResp, error) {
	var contents OAuthResp

	if state != oauthStateString {
		return contents, fmt.Errorf("Error: Invalid state string")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return contents, fmt.Errorf("Code exchange failed: %s", err.Error())
	}

	fmt.Println(token.AccessToken)

	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return contents, fmt.Errorf("Failed getting user info: %s", err.Error())
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&contents); err != nil {
		return contents, fmt.Errorf("Error: Failed to decode json response")
	}

	return contents, nil
}

func VerifyAccessToken(accessToken string, provider string) (string, entity.Student) {
	var contents OAuthResp
	var res *http.Response
	var err error

	if provider == "facebook" {
		res, err = http.Get("https://graph.facebook.com/v2.12/me?fields=email&access_token=" + accessToken)
	} else if provider == "google" {
		res, err = http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	} else {
		return "", entity.Student{}
	}

	if err != nil {
		return "", entity.Student{}
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&contents); err != nil {
		return "", entity.Student{}
	}

	student := studentRepo.FindOneByUsernameOrEmail(contents.Email)
	if student.ID == 0 {
		return "", entity.Student{}
	}

	loc, _ := time.LoadLocation(timezone)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().In(loc).Add(time.Hour).Unix(),
		"orig_iat": time.Now().In(loc).Unix(),
		"status":   "student",
		"user_id":  student.ID,
	})

	if tokenString, err := token.SignedString([]byte(viper.GetString("secret"))); err == nil {
		return tokenString, student
	}
	return "", entity.Student{}
}
