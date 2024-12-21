package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"koriebruh/try/conf"
	"koriebruh/try/controller"
	"koriebruh/try/repository"
	"koriebruh/try/service"
	"log"
)

func main() {
	config := conf.GetConfig()
	db := conf.InitDB()
	validate := validator.New()
	rdb := conf.GetRedis()

	//DI INJECT
	cacheRepository := repository.NewRedisCacheRepository(rdb)
	userRepository := repository.NewUserRepository()
	authService := service.NewAuthService(userRepository, db, validate, cacheRepository)
	authController := controller.NewAuthController(authService)

	//3. STATUS
	studentStatusRepository := repository.NewStudentStatusRepository()
	studentStatusServices := service.NewStudentStatusServices(db, studentStatusRepository, validate)
	studentStatusController := controller.NewStudentStatusController(studentStatusServices)

	app := fiber.New()

	app.Get("/", hellobg)
	app.Post("api/auth/register", authController.Register)
	app.Post("api/auth/login", authController.Login) //done
	authorized := app.Group("/", conf.JWTAuthMiddleware)
	authorized.Get("api/user", authController.CurrentAcc)

	authorized.Get("api/students/krs-offers", studentStatusController.KrsOffers)
	authorized.Get("api/students/status", studentStatusController.InformationStudent) //\done
	authorized.Put("api/students/class", studentStatusController.SetClassTime)        // done
	authorized.Get("api/students/krs", studentStatusController.GetAllKRSPick)         // done
	authorized.Get("api/students/permit", studentStatusController.InsertKRSPermit)    // done
	authorized.Get("api/students/krs-status", studentStatusController.StatusKRS)      // done

	server := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	if err := app.Listen(server); err != nil {
		log.Fatalf("server terminated %v", err)
	}
}

func hellobg(ctx *fiber.Ctx) error {
	return ctx.SendString("woiii")
}
