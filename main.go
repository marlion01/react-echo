package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)          //repositoryのインスタンスを生成
	userUsecase := usecase.NewUserUsecase(userRepository)       //usecaseのインスタンスを生成
	userController := controller.NewUserController(userUsecase) //controllerのインスタンスを生成
	e := router.NewRouter(userController)                       //routerのインスタンスを生成
	e.Logger.Fatal(e.Start(":8080"))                            //サーバーを起動
}
