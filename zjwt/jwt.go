package zjwt

import (
	"time"

	"gitee.com/sienectagv/gozk/zutils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
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

type PermissionsClaims struct {
	jwt.StandardClaims
	Permissons []string `json:"permissions"`
}

// func (rc *RoleClaims) Valid() error {
// 	err := rc.StandardClaims.Valid()
// 	// zlogger.Error(err)
// 	return err
// }

func TokenOfPermissons(user string, key []byte, perms []string) (string, error) {
	claim := &PermissionsClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  user,
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).UTC().Unix(),
		},
		Permissons: perms,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claim)
	return token.SignedString(key)
}

func ParseTokenOfPermissions(key []byte, tokenStr string) (user string, perms []string, err error) {
	permsClaim := &PermissionsClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, permsClaim, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrWrongMethod
		}
		return key, nil
	})
	if err != nil {
		return permsClaim.Audience, permsClaim.Permissons, errors.Wrap(err, "parse token err")
	}
	if !token.Valid {
		return permsClaim.Audience, permsClaim.Permissons, ErrInvalid
	}
	exp := time.Unix(permsClaim.ExpiresAt, 0)
	if exp.Before(time.Now().UTC()) {
		return permsClaim.Audience, permsClaim.Permissons, ErrExpired
	}
	return permsClaim.Audience, permsClaim.Permissons, nil
}

func ParseTokenOfPermissionsUpExp(key []byte, tokenStr string) (user, newToken string, perms []string, err error) {
	user, perms, err = ParseTokenOfPermissions(key, tokenStr)
	if err != nil {
		return
	}
	newToken, err = TokenOfPermissons(user, key, perms)
	return
}

func ParseTokenOfRoles(key []byte, tokenStr string) (user string, roles []uint, err error) {
	rc := &RoleClaims{}
	// zlogger.Info(tokenStr)
	token, err := jwt.ParseWithClaims(tokenStr, rc, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrWrongMethod
		}
		return key, nil
	})
	if err != nil {
		return rc.Audience, rc.RoleIDs, errors.Wrap(err, "parse token err")
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

func TokenOfRoles(user string, key []byte, roles []uint) (string, error) {
	claim := &RoleClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  user,
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).UTC().Unix(),
		},
		RoleIDs: roles,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claim)
	return token.SignedString(key)
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
