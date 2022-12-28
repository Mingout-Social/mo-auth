package config

import (
	"os"

	"aidanwoods.dev/go-paseto"
	"github.com/sirupsen/logrus"
)

const (
	AudienceAndroid = "android"
	AudienceIOS     = "ios"
	AudienceWebsite = "website"
)

var PasetoSecretKey paseto.V4AsymmetricSecretKey
var PasetoPublicKey paseto.V4AsymmetricPublicKey

func InitPasetoPublicKey() {
	var err error

	PasetoPublicKey, err = paseto.NewV4AsymmetricPublicKeyFromHex(os.Getenv("PASETO_PUBLIC_KEY"))
	if err != nil {
		logrus.Panicf("Error Initializing Paseto Public Key, Err: %v", err)
	}
}
