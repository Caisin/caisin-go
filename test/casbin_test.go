package test

import (
	_ "embed"
	"gitee.com/Caisin/caisin-go/casbin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

func TestCasbin(t *testing.T) {

	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go-admin"))
	e, _ := casbin.NewCasbin(db, "")
	//e.EnableAutoSave(true)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	e.LoadPolicy()

	// Modify the policy.
	e.AddRolePolicy("admin", "/api/v1/:id", "GET") //管理员有所有权限
	e.AddUserPolicy("caisin", "/api/v1/login", "GET")
	e.AddUserPolicy("caisin", "/api/v1/schema/:id", "GET")
	e.AddUserPolicy("caisin", "/api/v1/schema/:id", "POST")
	e.AddUserDenyPolicy("caisin", "/api/v1/schema/123", "POST")

	per, err := e.HasPer("admin", "caisin", "/api/v1/schema/123", "POST")
	println(per, err)
	per, err = e.HasPer("admin", "caisin", "/api/v1/schema/456", "POST")
	println(per, err)
	per, err = e.HasPer("admin", "caisin", "/api/v1/haha", "GET")
	println(per, err)
	// Save the policy back to DB.
	//e.SavePolicy()

}

func TestTimer(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		println("每隔一秒执行一次")
	}
}
