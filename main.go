package main

import (
	"flag"
	"fmt"
	"memo/dao"
	"memo/model"
	"memo/router"
	"memo/setting"
)

var appConfig = flag.String("conf", "./conf/conf.ini", "app config file")

func main() {
	// 读取配置文件信息
	flag.Parse()
	if err := setting.Init(*appConfig); err != nil {
		panic(err)
	}

	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		panic(err)
	}
	dao.DB.AutoMigrate(&model.Todo{})

	r := router.SetupRouter()

	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}