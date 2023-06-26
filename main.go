package main

import (
	"memo/dao"
	"memo/model"
	"memo/router"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	dao.DB.AutoMigrate(&model.Todo{})

	r := router.SetupRouter()

	r.Run(":10000")
}