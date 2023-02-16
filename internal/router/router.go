package router

import (
	"net/http"
	base_api "redisadmin/internal/api/v1"
	"redisadmin/internal/api/v1/redis_admin"
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

	r.GET("/index", base_api.Deal_request(redis_admin.Api_redis_list.Web_index))
	r.GET("/index.html", base_api.Deal_request(redis_admin.Api_redis_list.Web_index))
	r.GET("/login.html", base_api.Deal_request(redis_admin.Api_redis_list.Web_login))
	r.PUT("/login", base_api.Deal_request(redis_admin.Api_admin.Login))
	r.GET("/admin/migrate", base_api.Deal_request(redis_admin.Api_admin.Migrate))

	v1 := r.Group("/v1")
	v1.Use(Auth)
	{
		v1.GET("/menu", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_menu_list.Index))

		v1.GET("/admin/user", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_admin.Index))
		v1.POST("/admin/user", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_admin.Store))
		v1.PUT("/admin/user", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_admin.Update))
		v1.DELETE("/admin/user/:uId", base_api.AuthRule(redis_admin.Api_admin.GetUserAuthRule), base_api.Deal_request(redis_admin.Api_admin.Detele))
		v1.GET("/admin/user/:uId", base_api.AuthRule(redis_admin.Api_admin.GetUserAuthRule), base_api.Deal_request(redis_admin.Api_admin.UserInfo))

		v1.GET("/admin/role", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_admin.RoleIndex))
		v1.POST("/admin/role", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_admin.RoleStore))
		v1.PUT("/admin/role", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_admin.RoleUpdate))
		v1.DELETE("/admin/role/:rId", base_api.AuthRule(redis_admin.Api_admin.GetRoleAuthRule), base_api.Deal_request(redis_admin.Api_admin.RoleDelete))
		v1.GET("/admin/role/:rId/rule", base_api.AuthRule(redis_admin.Api_admin.GetRoleAuthRule), base_api.Deal_request(redis_admin.Api_admin.RoleRule))

		v1.GET("/admin/rule", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_admin.RuleIndex))
		v1.POST("/admin/rule", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_admin.RuleStore))
		v1.DELETE("/admin/rule", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_admin.RuleDelete))

		v1.GET("/redisTypeList", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_redis_list.TypeList))
		v1.POST("/redisTypeList", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_redis_list.TypeStore))
		v1.PUT("/redisTypeList", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_redis_list.TypeUpdate))
		v1.DELETE("/redisTypeList", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_redis_list.TypeDelete))

		v1.GET("/redisList/item", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_redis_list.ItemList))
		v1.POST("/redisList/item", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_redis_list.ItemStore))
		v1.PUT("/redisList/item", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_redis_list.ItemUpdate))
		v1.DELETE("/redisList/item", base_api.AuthRule(base_api.GetBaseRule), base_api.Deal_request(redis_admin.Api_redis_list.ItemDelete))
		v1.GET("/redisList/itemInfo", base_api.AuthRule(redis_admin.Api_redis_list.GetAuthRule), base_api.Deal_request(redis_admin.Api_redis_list.ItemInfo))

		v1.GET("/redisItem/indexHtml/:typeId", base_api.AuthRule(redis_admin.Api_redis_item.GetAuthRule), base_api.Deal_request(redis_admin.Api_redis_item.IndexHtml))
		v1.GET("/redisItem/redisList/:typeId", base_api.AuthRule(redis_admin.Api_redis_item.GetAuthRule), base_api.Deal_request(redis_admin.Api_redis_item.Get_Rdis_list))
		v1.GET("/redisType/:typeId/redisItem/:rdId/keys", base_api.AuthRule(redis_admin.Api_redis_item.GetAuthRule), base_api.Deal_request(redis_admin.Api_redis_item.Keys))
		v1.GET("/redisType/:typeId/redisItem/:rdId/val", base_api.AuthRule(redis_admin.Api_redis_item.GetAuthRule), base_api.Deal_request(redis_admin.Api_redis_item.Get_val))
	}

	return r
}
