package router

import (
	"net/http"
	baseApi "redisadmin/internal/api/v1"
	"redisadmin/internal/api/v1/redisAdmin"
	"redisadmin/internal/consts"
	"redisadmin/internal/services/auth"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	// 验证token
	token := ctx.GetHeader("Authorization")
	jwtClaims, err := auth.CheckJwtToken(token)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusUnauthorized, err.Error())
		ctx.Set("test", "test")
	}

	// 设置jwt参数
	ctx.Set(consts.JWT_CLAIMS, jwtClaims)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLFiles("internal/web/weAdmin/index.html", "internal/web/weAdmin/login.html", "internal/web/weAdmin/pages/redisItem/category.html")
	r.StaticFS("/static", http.Dir("./internal/web/weAdmin/static"))
	r.StaticFS("/json", http.Dir("./internal/web/weAdmin/json"))
	r.StaticFS("/pages", http.Dir("./internal/web/weAdmin/pages"))
	r.StaticFS("/lib/layui", http.Dir("./internal/web/weAdmin/lib/layui"))

	r.GET("/index", baseApi.DealRequest(redisAdmin.ApiRedisList.WebIndex))
	r.GET("/index.html", baseApi.DealRequest(redisAdmin.ApiRedisList.WebIndex))
	r.GET("/login.html", baseApi.DealRequest(redisAdmin.ApiRedisList.WebLogin))
	r.PUT("/login", baseApi.DealRequest(redisAdmin.ApiAdmin.Login))
	r.GET("/admin/migrate", baseApi.DealRequest(redisAdmin.ApiAdmin.Migrate))

	v1 := r.Group("/v1")
	v1.Use(Auth)
	{
		v1.GET("/menu", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiMenuList.Index))

		v1.GET("/admin/user", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiAdmin.Index))
		v1.POST("/admin/user", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiAdmin.Store))
		v1.PUT("/admin/user", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiAdmin.Update))
		v1.DELETE("/admin/user/:uId", baseApi.AuthRule(redisAdmin.ApiAdmin.GetUserAuthRule), baseApi.DealRequest(redisAdmin.ApiAdmin.Detele))
		v1.GET("/admin/user/:uId", baseApi.AuthRule(redisAdmin.ApiAdmin.GetUserAuthRule), baseApi.DealRequest(redisAdmin.ApiAdmin.UserInfo))

		v1.GET("/admin/role", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiAdmin.RoleIndex))
		v1.POST("/admin/role", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiAdmin.RoleStore))
		v1.PUT("/admin/role", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiAdmin.RoleUpdate))
		v1.DELETE("/admin/role/:rId", baseApi.AuthRule(redisAdmin.ApiAdmin.GetRoleAuthRule), baseApi.DealRequest(redisAdmin.ApiAdmin.RoleDelete))
		v1.GET("/admin/role/:rId/rule", baseApi.AuthRule(redisAdmin.ApiAdmin.GetRoleAuthRule), baseApi.DealRequest(redisAdmin.ApiAdmin.RoleRule))

		v1.GET("/admin/rule", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiAdmin.RuleIndex))
		v1.POST("/admin/rule", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiAdmin.RuleStore))
		v1.DELETE("/admin/rule", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiAdmin.RuleDelete))

		v1.GET("/redisTypeList", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiRedisList.TypeList))
		v1.POST("/redisTypeList", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiRedisList.TypeStore))
		v1.PUT("/redisTypeList", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiRedisList.TypeUpdate))
		v1.DELETE("/redisTypeList", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiRedisList.TypeDelete))

		v1.GET("/redisList/item", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiRedisList.ItemList))
		v1.POST("/redisList/item", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiRedisList.ItemStore))
		v1.PUT("/redisList/item", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiRedisList.ItemUpdate))
		v1.DELETE("/redisList/item", baseApi.AuthRule(baseApi.GetBaseRule), baseApi.DealRequest(redisAdmin.ApiRedisList.ItemDelete))
		v1.GET("/redisList/itemInfo", baseApi.AuthRule(redisAdmin.ApiRedisList.GetAuthRule), baseApi.DealRequest(redisAdmin.ApiRedisList.ItemInfo))

		v1.GET("/redisItem/indexHtml/:typeId", baseApi.AuthRule(redisAdmin.ApiRedisItem.GetAuthRule), baseApi.DealRequest(redisAdmin.ApiRedisItem.IndexHtml))
		v1.GET("/redisItem/redisList/:typeId", baseApi.AuthRule(redisAdmin.ApiRedisItem.GetAuthRule), baseApi.DealRequest(redisAdmin.ApiRedisItem.GetRdisList))
		v1.GET("/redisType/:typeId/redisItem/:rdId/keys", baseApi.AuthRule(redisAdmin.ApiRedisItem.GetAuthRule), baseApi.DealRequest(redisAdmin.ApiRedisItem.Keys))
		v1.GET("/redisType/:typeId/redisItem/:rdId/val", baseApi.AuthRule(redisAdmin.ApiRedisItem.GetAuthRule), baseApi.DealRequest(redisAdmin.ApiRedisItem.GetVal))
	}

	return r
}
