package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/seinyan/go-rest-api/internal/models"
	"github.com/seinyan/go-rest-api/internal/repository"
	"github.com/seinyan/go-rest-api/internal/service"
	"net/http"
)

type SecurityController interface {
	Register(ctx *gin.Context)
	Restore(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	UpdateUserPass(ctx *gin.Context)
}

type securityController struct {
	repo            repository.UserRepository
	securityService service.SecurityService
}

func NewSecurityController(repo repository.UserRepository) SecurityController {
	return &securityController{
		repo:            repo,
		securityService: service.NewSecurityService(),
	}
}


// Login godoc
// @Summary Login
// @Description do ping
// @Tags Security
// @Accept  x-www-form-urlencoded
// @Produce x-www-form-urlencoded
// @Consumes multipart/form-data
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {object} TokenResponse
// @Failure 400 {string} string "Bad credentials"
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
	if res := c.securityService.IsAuthenticated(item); res == true {
		// return new token
		token, err := c.securityService.GenerateToken(string(item.ID), item.Email, "ADMIN")
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

// Register godoc
// @Summary Register
// @Description Register User
// @Tags Security
// @Accept  x-www-form-urlencoded
// @Produce x-www-form-urlencoded
// @Consumes multipart/form-data
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {object} TokenResponse "ok
// @Failure 400 {object} ErrorResponse "error"
// @Failure 409 {object} MessageResponse "user exists"
// @Router /register [post]
func (c securityController) Register(ctx *gin.Context)  {
	var err error

	userRegister := models.UserRegister{
		Username: ctx.PostForm("username"),
		Password: ctx.PostForm("password"),
	}

	if err = ctx.ShouldBind(&userRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	// if user exists
	_, err = c.repo.GetByUsername(userRegister.Username)
	if err == nil {
		ctx.JSON(http.StatusConflict, &MessageResponse{
			Code: http.StatusConflict,
			Message: "user exists",
		})
		return
	}

	passwordHash, err := c.securityService.GeneratePasswordHash(userRegister.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	user := models.User{
		Email:        userRegister.Username,
		PasswordHash: passwordHash,
		Person:       models.Person{},
	}

	err = c.repo.Register(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	token, err := c.securityService.GenerateToken(string(user.ID), user.Email, "ADMIN")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,  NewErrorResponse(http.StatusInternalServerError, err))
	}

	ctx.JSON(http.StatusOK, &TokenResponse{
		Code: http.StatusOK,
		Token: token,
	})
}






// Restore godoc
// @Summary Restore
// @Description do ping
// @Tags Security
// @Accept  x-www-form-urlencoded
// @Produce x-www-form-urlencoded
// @Consumes multipart/form-data
// @Param username formData string false "username"
// @Success 200 {object} MessageResponse
// @Failure 400 {string} string "Bad credentials"
// @Failure 404 {string} string "Not fount"
// @Router /restore [post]
func (c securityController) Restore(ctx *gin.Context)  {

}



// Get godoc
// @Summary User Get
// @Description User
// @Tags Security
// @Accept json
// @Produce json
// @Param id path uint64 true "Id"
// @Success 200 {object}  models.User "ok"
// @Failure 400 {string} string "bad request"
// @Failure 404 {string} string "Not fount"
// @Failure 401 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /user [get]
// @Security bearerAuth
func (c securityController) GetUser(ctx *gin.Context) {

}

// Update godoc
// @Summary Update User
// @Description do ping
// @Tags Security
// @Accept  json
// @Produce  json
// @Param id path uint64 true "Id"
// @Param user body models.User true "Add account"
// @Success 200 {object}  models.User
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /user [put]
// @Security bearerAuth
func (c securityController) UpdateUser(ctx *gin.Context) {

}

// Update godoc
// @Summary Update User Pass
// @Description do ping
// @Tags Security
// @Accept  json
// @Produce  json
// @Param id path uint64 true "Id"
// @Param user body models.User true "Add account"
// @Success 200 {object}  models.User
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /userpass [put]
// @Security bearerAuth
func (c securityController) UpdateUserPass(ctx *gin.Context) {

}
