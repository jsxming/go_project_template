package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_project_template/global"
	"go_project_template/internal/model"
	"go_project_template/pkg"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		panic("setting error!")
	}
	err = setupDBEngine()
	if err != nil {
		panic("DB init error!")
	}
}
func setupDBEngine() error {
	var err error
	global.DB, err = model.NewDbEngine()
	if err != nil {
		return err
	}
	return nil
}

func setupSetting() error {
	s, err := pkg.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout = time.Second
	global.ServerSetting.WriteTimeout = time.Second
	fmt.Println(global.ServerSetting.HttpPort)
	fmt.Println(global.ServerSetting.RunMode)
	return nil
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
