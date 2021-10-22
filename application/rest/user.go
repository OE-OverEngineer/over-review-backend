package rest

import (
	"net/http"
	"strconv"

	"github.com/OE-OverEngineer/over-review-backend/domain/user"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userSrv user.UserService
}

type UserHandler interface {
	getAllUser(*gin.Context)
	getUser(*gin.Context)
}

func NewUserHandler(userSrv user.UserService) UserHandler {
	return &userHandler{
		userSrv: userSrv,
	}
}

func Routes(route *gin.Engine, h UserHandler) {
	r := route.Group("/user")
	{
		r.GET("", h.getAllUser)
		r.GET("/:id", h.getUser)
	}
}

func (h *userHandler) getAllUser(c *gin.Context) {
	resp, err := h.userSrv.GetAllUser()
	if err != nil {
		c.AbortWithError(404, err)
	}
	c.JSON(http.StatusOK, resp)
}
func (h *userHandler) getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	resp, err := h.userSrv.GetUser(id)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}
	if resp == nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(http.StatusOK, resp)
}
