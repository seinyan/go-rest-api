package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/seinyan/go-rest-api/models"
	"github.com/seinyan/go-rest-api/repository"
	"net/http"
)

type UserController interface {
	Controller
}

type userController struct {
	repo repository.UserRepository
}

func NewUserController(repo repository.UserRepository) UserController {
	return &userController{
		repo: repo,
	}
}


// List godoc
// @Summary User List
// @Description User Lst
// @Tags Users
// @Accept json
// @Produce json
// @Param email query string false "name search by email" Format(email)
// @Success 200 {array} models.User
// @Router /users [get]
// @Security bearerAuth
func (c userController) List(ctx *gin.Context) {
	var items [] models.User
	items = append(items, models.User{})
	ctx.JSON(200, items)
}

// Get godoc
// @Summary User Get
// @Description User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path uint64 true "Id"
// @Success 200 {object}  models.User "ok"
// @Failure 400 {string} string "bad request"
// @Failure 404 {string} string "Not fount"
// @Failure 401 {string} string "error"
// @Failure 500 {string} string "error"
// @Router /users/{id} [get]
// @Security bearerAuth
func (c userController) Get(ctx *gin.Context) {
	id, err := GetPathInt(ctx,"id")
	//if err != nil {
	//	httputil.NewError(ctx, http.StatusBadRequest, err)
	//	return
	//}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := c.repo.Get(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// Create godoc
// @Summary User Create
// @Description do ping
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.User true "Add account"
// @Success 201 {object}  models.User
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /users [post]
// @Security bearerAuth
func (c userController) Create(ctx *gin.Context)  {
	var item models.User
	var err error
	err = ctx.ShouldBindJSON(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err = c.repo.Create(item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, item)
}

// Update godoc
// @Summary User Create
// @Description do ping
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path uint64 true "Id"
// @Param user body models.User true "Add account"
// @Success 200 {object}  models.User
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /users/{id} [put]
// @Security bearerAuth
func (c userController) Update(ctx *gin.Context) {
	var item models.User
	var err error

	id, err := GetPathInt(ctx,"id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctx.ShouldBindJSON(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item.Id = id
	if _, err = c.repo.Get(item.Id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.repo.Update(item)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

// Delete godoc
// @Summary User Delete
// @Description User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path uint64 true "Id"
// @Success 200 {object}  models.User "ok"
// @Failure 404 {string} string "Not fount"
// @Failure 500 {string} string "error"
// @Router /users/{id} [delete]
// @Security bearerAuth
func (c userController) Delete(ctx *gin.Context) {
	id, err := GetPathInt(ctx,"id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := c.repo.Get(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = c.repo.Delete(item)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

