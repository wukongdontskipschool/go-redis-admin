package baseApi

import (
	"fmt"
	"net/http"
	"redisadmin/internal/accessControl"
	"redisadmin/internal/consts"
	"redisadmin/internal/services/auth"

	"github.com/gin-gonic/gin"
)

type ApiFunc func(*gin.Context) (int, gin.H, string)
type AuthRuleFunc func(*gin.Context) string

type BaseApi struct {
}

func (b *BaseApi) GetAuthRule(ctx *gin.Context) string {
	return ctx.Request.URL.EscapedPath()
}

var GetBaseRule AuthRuleFunc = func(ctx *gin.Context) string {
	return ctx.Request.URL.EscapedPath()
}

func AuthRule(rulefunc AuthRuleFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtClaimsIt, err := ctx.Get(consts.JWT_CLAIMS)
		if !err {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "没有用户信息"})
		}

		jwtClaims := jwtClaimsIt.(*auth.JwtClaims)

		//获取请求的URI
		obj := rulefunc(ctx)

		//获取请求方法
		act := ctx.Request.Method
		//获取用户的角色
		// sub := "role_1"
		sub := fmt.Sprintf("role_%d", jwtClaims.RId)

		if ok, _ := accessControl.Authorize(sub, obj, act); !ok {
			// panic(err)
			ctx.Abort()
			ctx.JSON(http.StatusForbidden, gin.H{"msg": "没有权限"})
		}
	}
}

func DealRequest(apiFunc ApiFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		deal(apiFunc, ctx)
	}
}

func deal(apiFunc ApiFunc, ctx *gin.Context) {
	httpState, gin_H, template := apiFunc(ctx)
	if template == "" {
		ctx.JSON(httpState, gin_H)
	} else {
		ctx.HTML(httpState, template, gin_H)
	}
}

/**
用js和html按以下要求编写代码
把json格式的字符串转为带空格缩进的格式写入到div标签内
同纬度的key补充空格数相同 子级维度的key填充空格数比父的多一个
可以获取原json字符串
*/
