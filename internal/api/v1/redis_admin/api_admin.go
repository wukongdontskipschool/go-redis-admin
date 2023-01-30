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

func (r *api_admin) Index(ctx *gin.Context) (int, gin.H, string) {
	http_state, gin_H := admin.Get_user_list()

	return http_state, gin_H, ""
}

func (r *api_admin) Store(ctx *gin.Context) (int, gin.H, string) {
	name, _ := ctx.GetPostForm("name")
	pass, _ := ctx.GetPostForm("pass")
	roleId, _ := ctx.GetPostForm("roleId")
	rid, _ := strconv.Atoi(roleId)

	http_state, gin_H := admin.Add_user(name, pass, uint(rid))

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

func (r *api_admin) Migrate(ctx *gin.Context) (int, gin.H, string) {
	admin.Migrate()
	return 200, gin.H{"user": "a", "value": "Info"}, ""
}
