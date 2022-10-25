package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var engine *gin.Engine
var flag = "moectf{This_is_a_fake_flag}"

func RouterInit() error {
	gin.SetMode("release")
	engine = gin.Default()
	engine.GET("/welcome", welcomeHandler)
	flagRouter := engine.Group("/find")
	flagRouter.Use(authRequired)
	flagRouter.GET("/", findHandler)
	flagRouter.GET("/flag", flagHandler)
	return nil
}

func Run() error {
	log.Println("Service running at [127.0.0.1:8080]")
	fmt.Println("Type localhost:8080/welcome in your browser and open it")
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		return err
	}
	return nil
}

func authRequired(ctx *gin.Context) {
	password := ctx.Query("password")
	if password == "" {
		ctx.JSON(400, gin.H{
			"message": "password required",
		})
		ctx.Abort()
		return
	}
	if password != "---moeCTF2022---" {
		ctx.JSON(400, gin.H{
			"message": "password incorrect",
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}

func welcomeHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Welcome to MoeCTF2022, can you find the flag?",
	})
	ctx.SetAccepted()
}

func findHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "you are so close to get flag",
	})
	ctx.SetAccepted()
}

func flagHandler(ctx *gin.Context) {
	flag = ctx.Query("flag")
	if flag == "" {
		ctx.JSON(400, gin.H{
			"message": "please input your flag and I will check it",
		})
		ctx.Abort()
		return
	}
	if !check(flag) {
		ctx.JSON(200, gin.H{
			"message": "flag incorrect, please try again",
		})
		ctx.Abort()
		return
	} else {
		ctx.JSON(200, gin.H{
			"message": "congratulations",
		})
	}

	ctx.SetAccepted()
}
