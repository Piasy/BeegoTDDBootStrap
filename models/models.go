package models

import (
	"github.com/astaxie/beego/orm"
)

// CONTRACT: all parameter validation is done by controller.
func init() {
	orm.RegisterModel(new(User), new(Verification))
}

const DB_UNIQUE_CONFLICT_TRY int = 3
