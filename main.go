package main

import (
	"final-project-4/config"
	"final-project-4/httpserver/controllers"
	"final-project-4/httpserver/repositories"
	"final-project-4/httpserver/routers"
	"final-project-4/httpserver/services"
	"final-project-4/utils"

	"final-project-4/docs"

	"github.com/gin-gonic/gin" // swagger embed files
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host                       localhost:3030
// @BasePath                   /api
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
func main() {

	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal("Environment Variables not found")
	// }
	app := gin.Default()
	appRoute := app.Group("/api")
	db, _ := config.Connect()

	authService := utils.NewAuthHelper(utils.Constants.JWT_SECRET_KEY)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService, authService)

	productRepo := repositories.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService, authService)

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService)

	transactionRepo := repositories.NewTransactionHistoryRepository(db)
	transactionService := services.NewTransactionHistoryService(
		transactionRepo,
		productRepo,
		userRepository,
		categoryRepository,
	)
	transactionController := controllers.NewTransactionHistoryController(transactionService)

	routers.UserRouter(appRoute, userController, authService)
	routers.ProductRouter(appRoute, productController, authService)
	routers.CategoryRouter(appRoute, categoryController, authService)
	routers.TransactionHistoryRouter(appRoute, transactionController, authService)

	docs.SwaggerInfo.Title = "Hacktiv8 final-project-4 API"
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// log.Println("Generating Swagger")
	// path, err := os.Getwd()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(path) // for example /home/user
	// cmd, err := exec.Command("swag", "fmt").Output()
	// log.Println(cmd)
	// log.Println("Swagger Generated")

	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.Run()
}
