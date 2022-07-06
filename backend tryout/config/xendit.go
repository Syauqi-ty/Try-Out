package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/naufalihsan/artaka-payment/wrapper"
	"github.com/xendit/xendit-go/client"
)

var (
	secretKey        string
	authorizationKey string
	xenditCli        *client.API
	qrisCli          *wrapper.API
)

func AuthorizeXendit() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	secret := os.Getenv("xnd_development_DUZflxQtZ9sBgILME2ooOLgHhvgoK9xDUnV3RTWSE1jd7NOel8QvfSyWOq0OUOv")
	authorization := os.Getenv("AUTHORIZATION")
	secretKey = secret
	authorizationKey = authorization
	xenditCli = client.New(secretKey)
	qrisCli = wrapper.New(secretKey)

}

func GetXenditSecretKey() string {
	return secretKey
}

func GetXenditAuthorizationKey() string {
	return authorizationKey
}

func GetXenditCli() *client.API {
	return xenditCli
}

func GetQRISCli() *wrapper.API {
	return qrisCli
}
