package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/hanoguz056-hub/jobboard/config"
	"github.com/hanoguz056-hub/jobboard/internal/handler"
	"github.com/hanoguz056-hub/jobboard/internal/middleware"
	"github.com/hanoguz056-hub/jobboard/internal/repository"
	"github.com/hanoguz056-hub/jobboard/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/hanoguz056-hub/jobboard/docs"
)

// @title           Job Board API
// @version         1.0
// @description     Это сервер API для биржи труда.
// @host            localhost:8080
// @BasePath
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.Load()

	pool, err := pgxpool.New(context.Background(), cfg.DBurl)

	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
	os.MkdirAll("./uploads/resumes", os.ModePerm)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE",
	}))

	// User
	userRepo, err := repository.NewUserRepository(context.Background(), pool)
	if err != nil {
    log.Fatalf("failed to initialize user repository: %v", err)
}
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Company
	companyRepo, err := repository.NewCompanyRepository(context.Background(), pool)
	if err != nil {
		log.Fatal("failed to initialize company repository")
	}

	companyService := service.NewCompanyService(companyRepo)
	companyHandler := handler.NewCompanyHandler(companyService)


	// Job
	jobRepo, _ := repository.NewJobRepository(context.Background(), pool)
	jobService := service.NewJobService(jobRepo, companyRepo)
	jobHandler := handler.NewJobHandler(jobService)

	//Applications
	applicationRepo:= repository.NewApplicationRepository(pool)
	applicationService := service.NewApplicationService(applicationRepo)
	applicationHandler := handler.NewApplicationHandler(applicationService)


	api := app.Group("/api/v1")

	// auth

	api.Post("/auth/register", userHandler.Register)
	api.Post("/auth/login", userHandler.Login)

	// public Companies

	api.Get("/companies", companyHandler.GetCompanies)
	api.Get("/companies/:id", companyHandler.GetCompanyByID)
	api.Get("/jobs", jobHandler.GetJobs)
	api.Get("/jobs/:id", jobHandler.GetJobByID)
	

	// protected 

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware)

	protected.Post("/companies", companyHandler.CreateCompany)
	protected.Put("/companies/:id", companyHandler.UpdateCompany)

	protected.Post("/jobs", jobHandler.CreateJob)
	protected.Put("/jobs/:id", jobHandler.UpdateJob)
	protected.Patch("/jobs/:id/status", jobHandler.UpdateJobStatus)
	protected.Delete("/jobs/:id", jobHandler.DeleteJob)
	protected.Post("/jobs/:id/apply", applicationHandler.Apply)
	protected.Patch("/applications/:id/status", applicationHandler.UpdateApplicationStatus)
	protected.Get("/jobs/:id/applications", applicationHandler.GetJobApplications)
	protected.Get("/applications/my", applicationHandler.GetMyApplications)

	// admin
	protected.Get("/admin/users", middleware.RequireRole("admin"), userHandler.GetAllUsers)
	protected.Patch("/admin/users/:id/ban", middleware.RequireRole("admin"), userHandler.BanUser)
	protected.Patch("/admin/jobs/:id/approve", middleware.RequireRole("admin"), jobHandler.ApproveJob)

	app.Get("/swagger/*", swagger.New(swagger.Config{
    URL: "http://localhost:8080/swagger/doc.json",
	}))

	log.Fatal(app.Listen(":" + cfg.Port))
}

