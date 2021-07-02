package zjwt

import (
	"testing"

	"gitee.com/sienectagv/gozk/zjwt"
)

func TestJwtAll(t *testing.T) {
	j, err := zjwt.TokenOfRoles("wuzk", zjwt.Secret, []uint{1, 2, 3})
	// t.Log(j, err)
	if err != nil {
		t.Log(err)
	} else {
		user, roles, err := zjwt.ParseTokenOfRoles(zjwt.Secret, j)
		t.Log(user, roles, err)
	}
}

func TestJwtUpExp(t *testing.T) {
	j, err := zjwt.TokenOfRoles("wuzk", zjwt.Secret, []uint{1, 2, 3})
	// t.Log(j, err)
	if err != nil {
		t.Log(err)
	} else {
		user, roles, nt, err := zjwt.ParseTokenOfRolesUpExp(zjwt.Secret, j)
		t.Log(user, roles, nt, err)
	}
}
