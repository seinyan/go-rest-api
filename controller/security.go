package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/seinyan/go-rest-api/models"
	"github.com/seinyan/go-rest-api/repository"
	"github.com/seinyan/go-rest-api/service"
	"net/http"
	"strings"
)

type SecurityController interface {
	Register(ctx *gin.Context)
	Restore(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type securityController struct {
	repo repository.UserRepository
	loginService service.LoginService
	JWTService service.JWTService
}

func NewSecurityController(repo repository.UserRepository) SecurityController {
	return &securityController{
		repo: repo,
		loginService: service.NewLoginService(),
		JWTService: service.NewJWTService(),
	}
}

// Create godoc
// @Summary User Create
// @Description do ping
// @Tags Security
// @Accept  x-www-form-urlencoded
// @Produce x-www-form-urlencoded
// @Consumes multipart/form-data
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 201 {object}  models.User
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /login [post]
func (c securityController) Login(ctx *gin.Context)  {

	item, err := c.repo.GetByUsername(ctx.PostForm("username"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Bad credentials.",
		})
		return
	}
	item.Password = ctx.PostForm("password")
	if res := c.loginService.IsAuthenticated(item); res == true {
		// return new token
		token, err := c.JWTService.GenerateToken(string(item.Id), item.Email, "ADMIN")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}

	ctx.JSON(http.StatusUnauthorized, gin.H{
		"error": "Bad credentials.",
	})
}


// Create godoc
// @Summary User Create
// @Description do ping
// @Tags Security
// @Accept  json
// @Produce  json
// @Param user body models.User true "Add account"
// @Success 201 {object}  models.User
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /register [post]
func (c securityController) Register(ctx *gin.Context)  {
	var item models.User
	var err error
	err = ctx.ShouldBindJSON(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.PasswordHash, _ = c.loginService.GeneratePasswordHash(item.Password)
	item.Email = strings.ToLower(item.Email)
	item, err = c.repo.Register(item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, item)
}

// Create godoc
// @Summary User Create
// @Description do ping
// @Tags Security
// @Accept  json
// @Produce  json
// @Param user body models.User true "Add account"
// @Success 201 {object}  models.User
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /restore [post]
func (c securityController) Restore(ctx *gin.Context)  {
	var item models.User
	var err error
	err = ctx.ShouldBindJSON(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.PasswordHash, _ = c.loginService.GeneratePasswordHash(item.Password)
	item.Email = strings.ToLower(item.Email)
	item, err = c.repo.Register(item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, item)
}

