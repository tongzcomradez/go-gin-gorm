package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tongzcomradez/go-gorm/db"
	"github.com/tongzcomradez/go-gorm/model"
	"github.com/tongzcomradez/go-gorm/service"
)

// GetUser abc
func GetUser(c *gin.Context) {
	db, _ := db.ConnectDB()
	userRepo := service.NewUserRepo(db)
	users := model.User{}
	userRepo.GetProfile(&users)
	// profiles := []model.Profile{
	// 	{
	// 		Name: "A",
	// 	},
	// 	{
	// 		Name: "B",
	// 	},
	// }
	// user := model.User{}
	// user := model.User{
	// 	Name:     "Test Tongz",
	// 	Profiles: profiles,
	// }
	// userRepo.Create(&user)
	c.JSON(200, users)
}
