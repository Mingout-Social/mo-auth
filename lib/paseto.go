package lib

import (
	"mo-auth/config"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateToken(userID primitive.ObjectID, audience string) string {
	token := paseto.NewToken()
	token.SetIssuer("com.mingout")
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(730 * time.Hour)) // 1 Month
	token.SetAudience(audience)

	token.SetString("user-id", userID.Hex())

	return token.V4Sign(config.PasetoSecretKey, nil)
}

func VerifyToken(signed string, audience string) (primitive.ObjectID, error) {
	var userID primitive.ObjectID

	parser := paseto.NewParser()
	parser.AddRule(paseto.IssuedBy("com.mingout"))
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.ValidAt(time.Now()))
	parser.AddRule(paseto.ForAudience(audience))

	token, err := parser.ParseV4Public(config.PasetoPublicKey, signed, nil)
	if err != nil {
		logrus.Errorf("User Token Verification Failed: %v", err)
		return userID, err
	}

	id, err := token.GetString("user-id")
	if err != nil {
		logrus.Errorf("Malformed Token Received - Token: %v, Err: %v", token, err)
		return userID, err
	}

	userID, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		logrus.Errorf("Malformed ID in Auth Token - UserID: %v, Err: %v", id, err)
		return userID, err
	}

	return userID, nil
}
