package user

import (
	"github.com/gin-gonic/gin"
	"main/domain/user/dto"
	"main/domain/user/entity"
	"net/http"
)

type Controller struct {
	model *Repository
}

func NewController() *Controller {
	return &Controller{model: NewUserModel(entity.Repository())}
}

func (con *Controller) Get(c *gin.Context) {
	var user dto.FindOne
	if err := c.BindUri(&user); err != nil {
		c.JSON(400, gin.H{"XXX": err})
		return
	}
	u := con.model.findOne(user)
	c.JSON(http.StatusOK, u)
}

func (con *Controller) GetAll(c *gin.Context) {
	u := con.model.findAll()
	c.JSON(http.StatusOK, u)
}

func (con *Controller) Create(c *gin.Context) {
}

func (con *Controller) Delete(c *gin.Context) {
}

func (con *Controller) Update(c *gin.Context) {
}
