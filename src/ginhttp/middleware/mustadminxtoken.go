package middleware

import (
	"github.com/SuperH-0630/cat-shop-back/src/ginhttp/data"
	"github.com/SuperH-0630/cat-shop-back/src/model"
	"github.com/gin-gonic/gin"
)

func mustAdminErrorData(debugMsgLst ...string) data.Data {
	debugMsg := ""

	if len(debugMsgLst) == 1 {
		debugMsg = debugMsgLst[0]
	}

	if debugMsg == "" {
		debugMsg = "权限不足"
	}

	return data.NewData(data.GlobalErrorCodeMustAdmin, nil, data.DefaultMsg(data.GlobalErrorCodeXTokenExpire), debugMsg)
}

func MustAdminXTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Value("user").(*model.User)
		if !ok {
			c.JSON(200, mustAdminErrorData("用户获取失败"))
			return
		}

		if user.Type == model.NormalUserType {
			c.JSON(200, mustAdminErrorData("普通用户权限不足"))
			return
		}

		c.Next()
	}
}