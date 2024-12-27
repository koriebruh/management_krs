package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: false,
	}))

	app.Get("/", hellobg)
	app.Post("api/auth/register", authController.Register)
	app.Post("api/auth/login", authController.Login) //login jwt
	authorized := app.Group("/", conf.JWTAuthMiddleware)
	authorized.Get("api/user", authController.CurrentAcc)

	authorized.Get("api/students/krs-offers", studentStatusController.KrsOffers)                // menampilkan jadwal aktif sem ini
	authorized.Get("api/students/krs-schedule", studentStatusController.KrsSchedule)            // jadwal input krs sem ini
	authorized.Get("api/students/status", studentStatusController.InformationStudent)           // informasi keaktifan mhs
	authorized.Put("api/students/class", studentStatusController.SetClassTime)                  // mengubah jenis kelas
	authorized.Get("api/students/krs", studentStatusController.GetAllKRSPick)                   // menampilkan krs yg sudah di pilih
	authorized.Get("api/students/permit", studentStatusController.InsertKRSPermit)              // nempilkan ijin apkah dia di ijinkan insert di luar jadwal
	authorized.Get("api/students/krs-status", studentStatusController.StatusKRSMhs)             // DONE
	authorized.Get("api/students/scores", studentStatusController.GetAllScores)                 // menampikan semua score
	authorized.Get("api/students/schedule-conflict", studentStatusController.ScheduleConflicts) // menampilkan semua jadwal yg bisa di ambil beserta yg conlfic
	authorized.Get("api/students/schedule-prodi", studentStatusController.KrsOffersProdi)       // menampilkan jadwal yg di tawarkan dari prodi yg fi ambil mhs
	authorized.Post("api/students/schedule/:id", studentStatusController.InsertSchedule)        // menambahkan jdwal berdasarkan id jadwal
	authorized.Get("api/students/log", studentStatusController.GetKrsLog)                       // menampilkan log aktifity yg di lakukan mhs
	authorized.Delete("api/students/krs/:id", studentStatusController.DeleteKrsRecByIdKrs)      // menghapus jadwal yg sudah di tambhakan berdsarkan idkrs
	authorized.Put("api/students/validate", studentStatusController.UpdateValidate)             // memvalidasi mhs

	//KRS yg user pilih itu krs_record
	server := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	if err := app.Listen(server); err != nil {
		log.Fatalf("server terminated %v", err)
	}
}

func hellobg(ctx *fiber.Ctx) error {
	return ctx.SendString("woiii")
}
