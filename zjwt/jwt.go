package zjwt

import (

	// "github.com/dgrijalva/jwt-go"

	"time"

	"gitee.com/sienectagv/gozk/zutils"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	JwtHeader = "Authorization"
)

var Secret = []byte("Wuzhikun@2021$zjhz")

var (
	ErrWrongMethod = zutils.NewError(1, "the method is wrong")
	ErrInvalid     = zutils.NewError(2, "invalid token")
	ErrExpired     = zutils.NewError(3, "the token is expired")
)

type RoleClaims struct {
	jwt.StandardClaims
	RoleIDs []uint `json:"rid"`
}

// func (rc *RoleClaims) Valid() error {
// 	err := rc.StandardClaims.Valid()
// 	// zlogger.Error(err)
// 	return err
// }

func TokenOfRoles(user string, key []byte, roles []uint) (string, error) {
	claim := &RoleClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  user,
			ExpiresAt: time.Now().Add(3 * time.Hour).UTC().Unix(),
		},
		RoleIDs: roles,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claim)
	return token.SignedString(key)
}

func ParseTokenOfRoles(key []byte, tokenStr string) (user string, roles []uint, err error) {
	rc := &RoleClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, rc, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrWrongMethod
		}
		return key, nil
	})
	if err != nil {
		return rc.Audience, rc.RoleIDs, err
	}
	if !token.Valid {
		return rc.Audience, rc.RoleIDs, ErrInvalid
	}
	// judge ExpiresAt
	exp := time.Unix(rc.ExpiresAt, 0)
	if exp.Before(time.Now().UTC()) {
		return rc.Audience, rc.RoleIDs, ErrExpired
	}
	return rc.Audience, rc.RoleIDs, nil
}

// 解析token并更新
func ParseTokenOfRolesUpExp(key []byte, tokenStr string) (user string, roles []uint, newToken string, err error) {
	user, roles, err = ParseTokenOfRoles(key, tokenStr)
	if nil != err {
		return
	}
	newToken, err = TokenOfRoles(user, key, roles)
	return
}
