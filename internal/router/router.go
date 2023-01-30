package router

import (
	"net/http"
	"redisadmin/internal/accessControl"
	base_api "redisadmin/internal/api/v1"
	"redisadmin/internal/api/v1/redis_admin"

	"github.com/gin-gonic/gin"
)

func Auth(ctx *gin.Context) {
	//获取请求的URI
	obj := ctx.Request.URL.RequestURI()
	//获取请求方法
	act := ctx.Request.Method
	//获取用户的角色
	sub := "role_1"

	if ok, _ := accessControl.Authorize(sub, obj, act); !ok {
		// panic(err)
		ctx.Abort()
		ctx.JSON(http.StatusForbidden, gin.H{"msg": "没有权限"})
	}

	// ctx.Next()
	// ctx.JSON(300, gin.H{"user": "aa", "value": "bb"})
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

	v1 := r.Group("/v1")
	v1.Use(Auth)
	{
		v1.GET("/menu", base_api.Deal_request(redis_admin.Api_menu_list.Index))

		v1.GET("/admin/user", base_api.Deal_request(redis_admin.Api_admin.Index))
		v1.POST("/admin/user", base_api.Deal_request(redis_admin.Api_admin.Store))

		v1.GET("/admin/role", base_api.Deal_request(redis_admin.Api_admin.RoleIndex))
		v1.POST("/admin/role", base_api.Deal_request(redis_admin.Api_admin.RoleStore))

		v1.GET("/admin/rule", base_api.Deal_request(redis_admin.Api_admin.RuleIndex))
		v1.POST("/admin/rule", base_api.Deal_request(redis_admin.Api_admin.RuleStore))
		v1.GET("/admin/migrate", base_api.Deal_request(redis_admin.Api_admin.Migrate))

		v1.GET("/redisList", base_api.Deal_request(redis_admin.Api_redis_list.Index))
		v1.POST("/redisList", base_api.Deal_request(redis_admin.Api_redis_list.Store))

		v1.GET("/redisItem/indexHtml/:typeId", base_api.Deal_request(redis_admin.Api_redis_item.IndexHtml))
		v1.GET("/redisItem/redisList/:typeId", base_api.Deal_request(redis_admin.Api_redis_item.Get_Rdis_list))
		v1.GET("/redisItem/keys/:id", base_api.Deal_request(redis_admin.Api_redis_item.Keys))
		v1.GET("/redisItem/getVal/:id/:key", base_api.Deal_request(redis_admin.Api_redis_item.Get_val))
	}

	return r
}

// Route::get('/photos', [PhotosController::class, 'index'])->name('photos.index');
// Route::get('/photos/create', [PhotosController::class, 'create'])->name('photos.create');
// Route::post('/photos', [PhotosController::class, 'store'])->name('photos.store');
// Route::get('/photos/{photo}', [PhotosController::class, 'show'])->name('photos.show');
// Route::get('/photos/{photo}/edit', [PhotosController::class, 'edit'])->name('photos.edit');
// Route::put('/photos/{photo}', [PhotosController::class, 'update'])->name('photos.update');
// Route::delete('/photos/{photo}', [PhotosController::class, 'destroy'])->name('photos.destroy');
