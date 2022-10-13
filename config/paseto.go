package config

import (
	"aidanwoods.dev/go-paseto"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	AudienceAndroid = "android"
	AudienceIOS     = "ios"
	AudienceWebsite = "website"
)

var PasetoSecretKey paseto.V4AsymmetricSecretKey
var PasetoPublicKey paseto.V4AsymmetricPublicKey

func InitPasetoPrivateKey() {
	var err error

	PasetoSecretKey, err = paseto.NewV4AsymmetricSecretKeyFromHex(os.Getenv("PASETO_SECRET_KEY"))
	if err != nil {
		logrus.Panicf("Error Initializing Paseto Private Key, Err: %v", err)
	}
}

func InitPasetoPublicKey() {
	var err error

	PasetoPublicKey, err = paseto.NewV4AsymmetricPublicKeyFromHex(os.Getenv("PASETO_PUBLIC_KEY"))
	if err != nil {
		logrus.Panicf("Error Initializing Paseto Public Key, Err: %v", err)
	}
}
