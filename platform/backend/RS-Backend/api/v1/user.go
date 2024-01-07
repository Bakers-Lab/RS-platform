package v1

import (
	"RS-Backend/dal/db"
	"RS-Backend/models/dao"
	service "RS-Backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func RegisterUserRoutes(router *gin.RouterGroup, dB db.IDB) {
	api := &UserHandler{service: service.NewUserService(db.NewUserAccesser(dB))}
	router.POST("/register", api.Register)
	router.POST("/login", api.LogIn)
}

func (api *UserHandler) Register(c *gin.Context) {
	var user dao.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, APIError{Error: "invalid input"})
		return
	}

	err := api.service.Register(c, &user)
	if err != nil {

		// 这里你可能需要根据错误类型返回不同的状态码
		c.JSON(http.StatusInternalServerError, APIError{"Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Register successfully"})
}
func (api *UserHandler) LogIn(c *gin.Context) {
	var requestData map[string]interface{}

    if err := c.BindJSON(&requestData); err != nil {
        // 处理绑定 JSON 数据失败的情况
        c.JSON(http.StatusBadRequest, gin.H{"error":  "Invalid input"})
        return
    }
	email, ok := requestData["email"].(string)// 这里需要做类型断言
    if !ok {
        // 处理类型断言失败的情况
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
        return
    }
    password, ok := requestData["password"].(string)
    if !ok {
        // 处理类型断言失败的情况
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
        return
    }

	user, err := api.service.LogIn(c, email, password)
	if err != nil {

		// 这里你可能需要根据错误类型返回不同的状态码
		c.JSON(http.StatusInternalServerError, APIError{"Account or password is incorrect"})
		return
	}
	c.JSON(http.StatusOK, user)
}
//文件上传成功，注意建数据集的时候用的是绝对路径（服务器的root下面）