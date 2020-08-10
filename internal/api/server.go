package api


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/seinyan/go-rest-api/configs"
	_ "github.com/seinyan/go-rest-api/docs"
	"github.com/seinyan/go-rest-api/internal/controller"
	"github.com/seinyan/go-rest-api/internal/database"
	"github.com/seinyan/go-rest-api/internal/middleware"
	"github.com/seinyan/go-rest-api/internal/migrations"
	"github.com/seinyan/go-rest-api/internal/service"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type Server interface {
	Start()
}

type server struct {}



func setupLogOutput()  {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// @title Swagger TEST Example API
// @version 5.0
// @description DDD This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email admin@seinayn.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
// @BasePath /

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func (s server) Start() {
	//setupLogOutput()

	fmt.Println("=====================")
	fmt.Println("======= CMD START =======")
	fmt.Println("=====================")

	c, err:= configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(c)

	securityService := service.NewSecurityService()

	server := gin.New()

	// validator convert field json
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	server.Use(gin.Recovery(), gin.Logger())

	DBConn, err := database.NewDBConn(c.Database)
	if err != nil {
		fmt.Println(err)
	}

	migrations.Migrate(DBConn)

	fmt.Println(middleware.JWTAuthMiddleware(securityService))

	store := database.NewStore(DBConn)

	securityController := controller.NewSecurityController(store.UserRepository)
	server.POST("/register", securityController.Register)
	server.POST("/login", securityController.Login)
	server.POST("/restore", securityController.Restore)
	server.GET("/user", securityController.GetUser)
	server.POST("/user", securityController.UpdateUser)
	server.POST("/userpass", securityController.UpdateUserPass)

	userController := controller.NewUserController(store.UserRepository)
	v1 := server.Group("/users") // , middleware.JWTAuthMiddleware(JWTService)
	{
		v1.GET("", userController.List)
		v1.GET(":id", userController.Get)
		v1.POST("", userController.Create)
		v1.PUT(":id", userController.Update)
		v1.DELETE(":id", userController.Delete)
	}


	// Redirect to docs page
	server.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently,  "/swagger/index.html")
	})
	// use ginSwagger middleware to serve the API docs
	// http://127.0.0.1:9000/swagger/index.html
	// https://github.com/swaggo/swag/tree/master/example
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	server.Run(c.HTTP.Host+":"+c.HTTP.Port)
	defer DBConn.Close()
}

func NewServer() Server {
	return &server{}
}

