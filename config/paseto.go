package config

import (
	"aidanwoods.dev/go-paseto"
)

const (
	AudienceAndroid = "android"
	AudienceIOS     = "ios"
	AudienceWebsite = "website"
)

var PasetoSecretKey paseto.V4AsymmetricSecretKey
var PasetoPublicKey paseto.V4AsymmetricPublicKey

func InitPasetoPrivateKey() {
	PasetoSecretKey = paseto.NewV4AsymmetricSecretKey()
}

func InitPasetoPublicKey() {
	PasetoPublicKey = PasetoSecretKey.Public()
}
