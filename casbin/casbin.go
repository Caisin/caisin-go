package casbin

import (
	_ "embed"
	"gitee.com/Caisin/caisin-go/utils/strutil"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

//go:embed rbac_model.conf
var defMode string

type Casbin struct {
	*casbin.Enforcer
}

const (
	typ_role  = "role"
	typ_user  = "user"
	typ_udeny = "udeny"
	typ_rdeny = "rdeny"
)

func NewCasbin(db *gorm.DB, modeStr string) (*Casbin, error) {
	if strutil.IsBlank(modeStr) {
		modeStr = defMode
	}
	mode, err := model.NewModelFromString(modeStr)
	if err != nil {
		return nil, err
	}
	a, err := gormadapter.NewAdapterByDB(db)
	e, err := casbin.NewEnforcer(mode, a)
	if err != nil {
		return nil, err
	}
	return &Casbin{Enforcer: e}, nil
}

func (e *Casbin) AddRolePolicy(role string, params ...interface{}) (bool, error) {
	ps := append([]any{role, typ_role}, params...)
	return e.Enforcer.AddPolicy(ps...)
}

func (e *Casbin) AddUserPolicy(user string, params ...interface{}) (bool, error) {
	ps := append([]any{user, typ_user}, params...)
	return e.Enforcer.AddPolicy(ps...)
}

func (e *Casbin) AddUserDenyPolicy(user, res, method string) (bool, error) {
	return e.Enforcer.AddPolicy(user, typ_udeny, res, method)
}
func (e *Casbin) AddRoleDenyPolicy(role, res, method string) (bool, error) {
	return e.Enforcer.AddPolicy(role, typ_udeny, res, method)
}

func (e *Casbin) IsDeny(user, res, method string) (bool, error) {
	return e.Enforcer.Enforce(user, typ_udeny, res, method)
}
func (e *Casbin) IsRoleDeny(role, res, method string) (bool, error) {
	return e.Enforcer.Enforce(role, typ_rdeny, res, method)
}

func (e *Casbin) HasRolePer(role, res, method string) (bool, error) {
	isDeny, err := e.IsRoleDeny(role, res, method)
	if err != nil {
		return false, err
	}
	if isDeny {
		return false, nil
	}
	return e.Enforcer.Enforce(role, typ_role, res, method)
}

func (e *Casbin) HasUserPer(user, res, method string) (bool, error) {
	isDeny, err := e.IsDeny(user, res, method)
	if err != nil {
		return false, err
	}
	if isDeny {
		return false, nil
	}
	return e.Enforcer.Enforce(user, typ_user, res, method)
}

func (e *Casbin) HasPer(role, user, res, method string) (bool, error) {
	hasRolePer, err := e.HasRolePer(role, res, method)
	if err != nil {
		return false, err
	}
	if !hasRolePer {
		return e.HasUserPer(user, res, method)
	}
	return hasRolePer, err
}
