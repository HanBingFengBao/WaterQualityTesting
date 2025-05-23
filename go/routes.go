// @Title  routes
// @Description  程序的路由均集中在这个文件里
package main

import (
	"lianjiang/controller"
	"lianjiang/middleware"

	"github.com/gin-gonic/gin"
)

// @title    CollectRoute
// @description   给gin引擎挂上路由监听
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func CollectRoute(r *gin.Engine) *gin.Engine {
	// TODO 添加中间件
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())

	// TODO 用户的注册路由
	r.POST("/regist", controller.Register)

	// TODO 用户的邮箱验证
	r.GET("/verify/:id", controller.VerifyEmail)

	// TODO 用户找回密码
	r.PUT("/security", controller.Security)

	// TODO 用户更改密码
	r.PUT("/updatepass", middleware.AuthMiddleware(), controller.UpdatePass)

	// TODO 用户的登录路由
	r.POST("/login", controller.Login)

	// TODO 用户的个人信息路由
	r.GET("/personal", middleware.AuthMiddleware(), controller.PersonalPage)

	// 获取用户列表
	r.GET("/users/:start/:end", middleware.AuthMiddleware(), controller.Users)

	// 删除用户
	r.DELETE("/user/:id", middleware.AuthMiddleware(), controller.DeleteUser)

	// 删除用户
	r.DELETE("/users", middleware.AuthMiddleware(), controller.DeleteUsers)

	// 更新用户信息（用户名、等级）
	r.PUT("/user/:id", middleware.AuthMiddleware(), controller.UpdateUser)

	// TODO 文件上传
	r.POST("/upload/:system", middleware.AuthMiddleware(), controller.Upload)

	// TODO 文件列表
	r.GET("/files", middleware.AuthMiddleware(), controller.List)

	// TODO 文件下载
	r.GET("/download", middleware.AuthMiddleware(), controller.Download)

	// TODO 文件删除
	r.DELETE("/file", middleware.AuthMiddleware(), controller.DeleteFile)

	// TODO 数据获取
	r.GET("/data/:name/:system", middleware.AuthMiddleware(), controller.ShowData)

	// TODO 获取一对多的行字段
	r.GET("/data/rowall/:key/:name", middleware.AuthMiddleware(), controller.ShowRowAllData)

	// TODO 获取一对一的行字段
	r.GET("/data/rowone/:key/:name", middleware.AuthMiddleware(), controller.ShowRowOneData)

	// TODO 数据删除
	r.DELETE("/data/:time/:start/:end", middleware.AuthMiddleware(), controller.DeleteData)

	// TODO 数据恢复
	r.PUT("/data/:start/:end", middleware.AuthMiddleware(), controller.RecoverData)

	// TODO 查看用户的文件上传、删除记录
	r.GET("/history/file/:start/:end", middleware.AuthMiddleware(), controller.FileHistory)

	// 删除用户的数据上传、删除记录
	r.DELETE("/history/file", middleware.AuthMiddleware(), controller.DeleteFileHistory)

	// 查看用户的数据上传、删除记录
	r.GET("/history/data/:start/:end", middleware.AuthMiddleware(), controller.DataHistory)

	// 删除用户的数据上传、删除记录
	r.DELETE("/history/data", middleware.AuthMiddleware(), controller.DeleteDataHistory)

	// TODO 查看用户的映射上传、删除记录
	r.GET("/history/map/:start/:end", middleware.AuthMiddleware(), controller.MapHistory)

	// 删除用户的数据上传、删除记录
	r.DELETE("/history/map", middleware.AuthMiddleware(), controller.DeleteMapHistory)

	// TODO 查看映射主键
	r.GET("/map/:id", middleware.AuthMiddleware(), controller.ShowMapKeys)

	// TODO 查看映射键的值
	r.GET("/map/:id/:key", middleware.AuthMiddleware(), controller.ShowMapValue)

	// TODO 通过同名键值创建映射
	r.PUT("/map/:id", middleware.AuthMiddleware(), controller.CreateMapKey)

	// TODO 更新映射键值对
	r.PUT("/map/:id/:key", middleware.AuthMiddleware(), controller.CreateMapValue)

	// TODO 删除映射
	r.DELETE("/map/:id/:key", middleware.AuthMiddleware(), controller.DeleteMapKey)

	// TODO 查看映射备份
	r.GET("/map/backup/:id/:start/:end", middleware.AuthMiddleware(), controller.BackupMap)

	// TODO 预测
	r.GET("/forecast", middleware.AuthMiddleware(), controller.Forecast)
	return r
}
