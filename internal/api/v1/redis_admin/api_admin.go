package redis_admin

import (
	base_api "redisadmin/internal/api/v1"
	"redisadmin/internal/jobs/admin"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Api_admin = &api_admin{}

type api_admin struct {
	base_api.Base_api
}

func (b *api_admin) GetUserAuthRule(ctx *gin.Context) string {
	return "/v1/admin/user"

}
func (b *api_admin) GetRoleAuthRule(ctx *gin.Context) string {
	return "/v1/admin/role"
}

func (r *api_admin) Login(ctx *gin.Context) (int, gin.H, string) {
	user := ctx.PostForm("user")
	pass := ctx.PostForm("pass")
	http_state, gin_H := admin.Login(user, pass)

	return http_state, gin_H, ""
}

func (r *api_admin) Index(ctx *gin.Context) (int, gin.H, string) {
	http_state, gin_H := admin.Get_user_list()

	return http_state, gin_H, ""
}

func (r *api_admin) UserInfo(ctx *gin.Context) (int, gin.H, string) {
	rId := ctx.Param("uId")
	http_state, gin_H := admin.UserInfo(rId)

	return http_state, gin_H, ""
}

func (r *api_admin) Store(ctx *gin.Context) (int, gin.H, string) {
	name := ctx.PostForm("name")
	pass := ctx.PostForm("pass")
	roleId := ctx.PostForm("roleId")
	rid, _ := strconv.Atoi(roleId)

	http_state, gin_H := admin.Add_user(name, pass, uint(rid))

	return http_state, gin_H, ""
}

func (r *api_admin) Update(ctx *gin.Context) (int, gin.H, string) {
	uId := ctx.PostForm("uId")
	name := ctx.PostForm("name")
	pass := ctx.PostForm("pass")
	roleId := ctx.PostForm("roleId")
	rid, _ := strconv.Atoi(roleId)
	uIdInt, _ := strconv.Atoi(uId)

	http_state, gin_H := admin.UpdateUser(uint(uIdInt), name, pass, uint(rid))

	return http_state, gin_H, ""
}

func (r *api_admin) RoleIndex(ctx *gin.Context) (int, gin.H, string) {
	http_state, gin_H := admin.Get_role_list()
	return http_state, gin_H, ""
}

func (r *api_admin) RoleStore(ctx *gin.Context) (int, gin.H, string) {
	name := ctx.PostForm("name")
	ids := ctx.PostFormMap("ruleIds")

	http_state, gin_H := admin.Add_role(name, ids)

	return http_state, gin_H, ""
}

func (r *api_admin) Detele(ctx *gin.Context) (int, gin.H, string) {
	rId := ctx.Param("uId")

	http_state, gin_H := admin.Delete(rId)

	return http_state, gin_H, ""
}

func (r *api_admin) RoleUpdate(ctx *gin.Context) (int, gin.H, string) {
	rId := ctx.PostForm("rId")
	name := ctx.PostForm("name")
	ids := ctx.PostFormMap("ruleIds")

	http_state, gin_H := admin.RoleUpdate(rId, name, ids)

	return http_state, gin_H, ""
}

func (r *api_admin) RoleDelete(ctx *gin.Context) (int, gin.H, string) {
	rId := ctx.Param("rId")

	http_state, gin_H := admin.RoleDelete(rId)

	return http_state, gin_H, ""
}

func (r *api_admin) RuleIndex(ctx *gin.Context) (int, gin.H, string) {
	http_state, gin_H := admin.Get_rule_list()
	return http_state, gin_H, ""
}

func (r *api_admin) RuleStore(ctx *gin.Context) (int, gin.H, string) {
	rule, _ := ctx.GetPostForm("rule")
	act, _ := ctx.GetPostForm("act")
	desc, _ := ctx.GetPostForm("desc")

	http_state, gin_H := admin.Add_rule(rule, act, desc)

	return http_state, gin_H, ""
}

func (r *api_admin) RuleDelete(ctx *gin.Context) (int, gin.H, string) {
	rId := ctx.Query("id")

	http_state, gin_H := admin.RuleDelete(rId)

	return http_state, gin_H, ""
}

func (r *api_admin) RoleRule(ctx *gin.Context) (int, gin.H, string) {
	rId := ctx.Param("rId")

	http_state, gin_H := admin.RoleRule(rId)

	return http_state, gin_H, ""
}

func (r *api_admin) Migrate(ctx *gin.Context) (int, gin.H, string) {
	admin.Migrate()
	return 200, gin.H{"user": "a", "value": "Info"}, ""
}
