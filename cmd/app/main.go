package main

import (
	_ "github.com/dian/bank-api/docs"
	"github.com/dian/bank-api/internal/constant"
	"github.com/dian/bank-api/internal/db"
	"github.com/dian/bank-api/internal/handler"
	"github.com/dian/bank-api/internal/middleware"
	"github.com/dian/bank-api/internal/repository"
	"github.com/dian/bank-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"       // ini untuk swaggerFiles
	"github.com/swaggo/gin-swagger" // tanpa `swaggo` dua kali
	"log"
	"os"
)

func main() {
	_ = godotenv.Load()
	dbConn := db.InitDB()
	userRepo := repository.NewUserRepo(dbConn)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.New()
	r.Use(middleware.RequestIDMiddleware()) // 🆕 PASANG PALING AWAL

	r.Use(middleware.LoggingMiddleware())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", handler.Ping)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(constant.RoleAdmin, constant.RoleUser))

	auth.GET("/users/email/:email", userHandler.GetUserByEmail)

	adminOnly := r.Group("/")
	adminOnly.Use(middleware.AuthMiddleware(constant.RoleAdmin))

	adminOnly.POST("/users", handler.CreateUser)
	adminOnly.GET("/users", handler.GetUsers)
	r.GET("/admin", middleware.AuthMiddleware("admin"), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "welcome admin"})
	})
	r.POST("/login", userHandler.Login)

	port := os.Getenv("APP_PORT")
	log.Fatal(r.Run(":" + port))
}
