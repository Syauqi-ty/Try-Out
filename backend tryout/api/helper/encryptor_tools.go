package shalala

import (
	"gopkg.in/hlandau/passlib.v1"
	"gopkg.in/hlandau/passlib.v1/abstract"
)

func Encrypt(password string) string {
	passlib.UseDefaults(passlib.DefaultsLatest)
	passlib.DefaultSchemes = []abstract.Scheme{passlib.DefaultSchemes[3]}

	hash, _ := passlib.Hash(password)

	return hash
}

func Verify(plaintext string, encrypted string) bool {
	passlib.UseDefaults(passlib.DefaultsLatest)
	passlib.DefaultSchemes = []abstract.Scheme{passlib.DefaultSchemes[3]}
	err := passlib.VerifyNoUpgrade(plaintext, encrypted)
	if err != nil {
		return false
	}
	return true
}
