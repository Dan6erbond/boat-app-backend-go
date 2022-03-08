package main

import (
	"os"

	"github.com/iris-contrib/swagger/v12"              // swagger middleware for Iris
	"github.com/iris-contrib/swagger/v12/swaggerFiles" // swagger embed files
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
	"github.com/kataras/iris/v12/mvc"
	"openwt.com/boat-app-backend/docs"
	"openwt.com/boat-app-backend/pkg/controllers"
	"openwt.com/boat-app-backend/pkg/database"
	"openwt.com/boat-app-backend/pkg/repositories"
	"openwt.com/boat-app-backend/pkg/security"
	"openwt.com/boat-app-backend/pkg/services"
)

func init() {
	godotenv.Load()
}

// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	app := iris.New()

	jwtUtil := security.NewJWTUtil(os.Getenv("JWT_SECRET"))

	docs.SwaggerInfo.Title = "Boat App"
	docs.SwaggerInfo.Description = "This is a backend written for the Boat App in Golang/Iris."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	swaggerUI := swagger.WrapHandler(
		swaggerFiles.Handler,
		swagger.URL("http://localhost:8080/swagger/doc.json"),
		swagger.DeepLinking(true),
	)

	// Register on http://localhost:8080/swagger
	app.Get("/swagger", swaggerUI)
	// And the wildcard one for index.html, *.js, *.css and e.t.c.
	app.Get("/swagger/{any:path}", swaggerUI)

	db, err := database.GetDatabase()

	if err != nil {
		panic(err)
	}

	boatsRepository := repositories.NewBoatsRepository(db)
	usersRepository := repositories.NewUsersRepository(db)

	api := app.Party("/api")
	{
		boatsApi := api.Party("/boats")
		{
			boatsMvc := mvc.New(boatsApi)
			boatsMvc.Register(services.NewBoatsService(boatsRepository))
			boatsMvc.Handle(new(controllers.BoatsController))
		}

		authApi := api.Party("/auth")
		{
			authMvc := mvc.New(authApi)
			usersService := services.NewUsersService(usersRepository)
			authMvc.Register(services.NewAuthService(usersService, jwtUtil))
			authMvc.Handle(new(controllers.AuthController))
		}
	}

	app.UseRouter(cors.New().AllowOrigin("*").Handler())

	app.Listen(":8080")
}
