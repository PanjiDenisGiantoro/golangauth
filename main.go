package main

import (
	"auth2/handler"
	"auth2/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:W@rung01@tcp(127.0.0.1:3306)/founding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	userService := user.Newservice(userRepository)

	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	router.Run(":8083")

	//fmt.Println("Connected to database")
	//
	//var users []user.User
	//db.Find(&users)
	//
	//for _, user := range users {
	//	fmt.Println(user.Name)
	//	fmt.Println(user.Email)
	//	fmt.Println("===============")
	//}
	//
}

//func handler(c *gin.Context) {
//	dsn := "root:W@rung01@tcp(127.0.0.1:3306)/founding?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	var users []user.User
//	db.Find(&users)
//
//	c.JSON(http.StatusOK, users)
//
//	//handler
//	//service
//	//repository
//	//db
//
//}
