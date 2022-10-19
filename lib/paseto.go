package lib

import (
	"strconv"
	"time"

	"github.com/Mingout-Social/mo-auth/config"

	"aidanwoods.dev/go-paseto"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateToken(userID primitive.ObjectID, audience string, oldUserId int) string {
	token := paseto.NewToken()
	token.SetIssuer("com.mingout")
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(730 * time.Hour)) // 1 Month
	token.SetAudience(audience)

	token.SetString("user-id", userID.Hex())
	token.SetString("old-user-id", strconv.Itoa(oldUserId))

	return token.V4Sign(config.PasetoSecretKey, nil)
}

func VerifyToken(signed string, audience string) (primitive.ObjectID, string, error) {
	var userID primitive.ObjectID
	var oldUserId string

	parser := paseto.NewParser()
	parser.AddRule(paseto.IssuedBy("com.mingout"))
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.ValidAt(time.Now()))
	parser.AddRule(paseto.ForAudience(audience))

	token, err := parser.ParseV4Public(config.PasetoPublicKey, signed, nil)
	if err != nil {
		logrus.Errorf("User Token Verification Failed: %v", err)
		return userID, oldUserId, err
	}

	id, err := token.GetString("user-id")
	if err != nil {
		logrus.Errorf("Malformed Token Received - Token: %v, Err: %v", token, err)
		return userID, oldUserId, err
	}

	oldUserId, err = token.GetString("old-user-id")
	if err != nil {
		logrus.Errorf("Malformed Token Received - Token: %v, Err: %v", token, err)
		return userID, oldUserId, err
	}

	userID, err = primitive.ObjectIDFromHex(id)

	if err != nil {
		logrus.Errorf("Malformed ID in Auth Token - UserID: %v, Err: %v", id, err)
		return userID, oldUserId, err
	}

	return userID, oldUserId, nil
}
