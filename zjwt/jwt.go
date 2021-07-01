package zjwt

import (
	"time"

	// "github.com/dgrijalva/jwt-go"
	jwt "github.com/dgrijalva/jwt-go"
)

var Secret = "Wuzhikun@2021$zjhz"

type RoleClaims struct {
	jwt.StandardClaims
	// User string `json:"use"`
	// ExpiresAt uint   `json:"exp"` //UTC time
	RoleIDs []uint `json:"rid"`
}

// func (c *RoleClaims) Valid() error {
// 	if len(c.User) > 0 && len(c.RoleIDs) > 0 {
// 		return nil
// 	} else {
// 		return zutils.NewError(1, "invalid in user or roles")
// 	}
// }

type Takon struct {
	*jwt.Token
}

func TokenOfRoles(user, key string, roles []uint) (string, error) {
	claim := &RoleClaims{StandardClaims: jwt.StandardClaims{
		audience:  user,
		ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
	},
		RoleIDs: roles,
	}
	tc := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	return tc.SignedString(key)
}
