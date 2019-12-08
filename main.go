package main

import (
	"github.com/xyfll7/login/database"
	"github.com/xyfll7/login/routers"
)

func main() {
	mongoURI := "mongodb://xyf:yangqi7'@localhost:27018,localhost:27019,localhost:27017/?replicaSet=rep&authSource=admin"
	db, err := database.New(mongoURI, "test")
	if err != nil {
		panic(err)
	}
	// 创建唯一索引，这个只执就一次就可以了，用户名和email不得有重复。
	// 这样就利用了数据库的唯一索引做了用户注册重名检查，如果重名，数据库回返回错误，
	// 我们可以利用这个返回的错误，发送给前端提示用户用户名或email已经存在。
	// db.CreateUniqueIndex("admins", "email")
	// db.CreateUniqueIndex("admins", "name")
	defer db.Close()
	r := routers.InitGin(db)
	r.Run(":3001")
}
