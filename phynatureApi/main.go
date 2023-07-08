package main

import (
	"fmt"
	"log"
	"phynatureApi/config"
	"phynatureApi/controller"
	"phynatureApi/entity"
	"phynatureApi/repository"
	"phynatureApi/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	db := setupDB(cfg)
	migrateDB(db)

	articleRepo := repository.NewArticleRepository(db)
	articleService := service.NewArticleService(articleRepo)
	articleController := controller.NewArticleController(articleService)

	router := gin.Default()

	router.POST("/article", articleController.CreateArticle)
	router.GET("/article", articleController.GetAllArticle)

	log.Printf("Starting %s server on port %s\n", cfg.AppName, cfg.AppPort)
	err := router.Run(":" + cfg.AppPort)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

}

func setupDB(cfg config.Config) *gorm.DB {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	)

	db, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

func migrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Article{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
