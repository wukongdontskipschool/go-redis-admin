package redisAdmin

import (
	baseApi "redisadmin/internal/api/v1"
	"redisadmin/internal/consts"
	"redisadmin/internal/jobs/admin"
	"redisadmin/internal/services/auth"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ApiAdmin = &apiAdmin{}

type apiAdmin struct {
	baseApi.BaseApi
}

func (b *apiAdmin) GetUserAuthRule(ctx *gin.Context) string {
	return "/v1/admin/user"

}
func (b *apiAdmin) GetRoleAuthRule(ctx *gin.Context) string {
	return "/v1/admin/role"
}

func (r *apiAdmin) Login(ctx *gin.Context) (int, gin.H, string) {
	user := ctx.PostForm("user")
	pass := ctx.PostForm("pass")
	httpState, gin_H := admin.Login(user, pass)

	return httpState, gin_H, ""
}

func (r *apiAdmin) Index(ctx *gin.Context) (int, gin.H, string) {
	httpState, gin_H := admin.GetUserList()

	return httpState, gin_H, ""
}

func (r *apiAdmin) UserInfo(ctx *gin.Context) (int, gin.H, string) {
	rId := ctx.Param("uId")
	httpState, gin_H := admin.UserInfo(rId)

	return httpState, gin_H, ""
}

func (r *apiAdmin) Store(ctx *gin.Context) (int, gin.H, string) {
	name := ctx.PostForm("name")
	pass := ctx.PostForm("pass")
	roleId := ctx.PostForm("roleId")
	rid, _ := strconv.Atoi(roleId)

	httpState, gin_H := admin.AddUser(name, pass, uint(rid))

	return httpState, gin_H, ""
}

func (r *apiAdmin) Update(ctx *gin.Context) (int, gin.H, string) {
	uId := ctx.PostForm("uId")
	name := ctx.PostForm("name")
	pass := ctx.PostForm("pass")
	roleId := ctx.PostForm("roleId")
	rid, _ := strconv.Atoi(roleId)
	uIdInt, _ := strconv.Atoi(uId)
	jwt, _ := ctx.Get(consts.JWT_CLAIMS)
	jwtObj := jwt.(*auth.JwtClaims)

	httpState, gin_H := admin.UpdateUser(uint(uIdInt), name, pass, uint(rid), jwtObj)

	return httpState, gin_H, ""
}

func (r *apiAdmin) RoleIndex(ctx *gin.Context) (int, gin.H, string) {
	httpState, gin_H := admin.GetRoleList()
	return httpState, gin_H, ""
}

func (r *apiAdmin) RoleStore(ctx *gin.Context) (int, gin.H, string) {
	name := ctx.PostForm("name")
	ids := ctx.PostFormMap("ruleIds")

	httpState, gin_H := admin.AddRole(name, ids)

	return httpState, gin_H, ""
}

func (r *apiAdmin) Detele(ctx *gin.Context) (int, gin.H, string) {
	rId := ctx.Param("uId")

	httpState, gin_H := admin.Delete(rId)

	return httpState, gin_H, ""
}

func (r *apiAdmin) RoleUpdate(ctx *gin.Context) (int, gin.H, string) {
	rId := ctx.PostForm("rId")
	name := ctx.PostForm("name")
	ids := ctx.PostFormMap("ruleIds")

	httpState, gin_H := admin.RoleUpdate(rId, name, ids)

	return httpState, gin_H, ""
}

func (r *apiAdmin) RoleDelete(ctx *gin.Context) (int, gin.H, string) {
	rId := ctx.Param("rId")

	httpState, gin_H := admin.RoleDelete(rId)

	return httpState, gin_H, ""
}

func (r *apiAdmin) RuleIndex(ctx *gin.Context) (int, gin.H, string) {
	httpState, gin_H := admin.GetRuleList()
	return httpState, gin_H, ""
}

func (r *apiAdmin) RuleStore(ctx *gin.Context) (int, gin.H, string) {
	rule, _ := ctx.GetPostForm("rule")
	act, _ := ctx.GetPostForm("act")
	desc, _ := ctx.GetPostForm("desc")

	httpState, gin_H := admin.AddRule(rule, act, desc)

	return httpState, gin_H, ""
}

func (r *apiAdmin) RuleDelete(ctx *gin.Context) (int, gin.H, string) {
	rId := ctx.Query("id")

	httpState, gin_H := admin.RuleDelete(rId)

	return httpState, gin_H, ""
}

func (r *apiAdmin) RoleRule(ctx *gin.Context) (int, gin.H, string) {
	rId := ctx.Param("rId")

	httpState, gin_H := admin.RoleRule(rId)

	return httpState, gin_H, ""
}

func (r *apiAdmin) Migrate(ctx *gin.Context) (int, gin.H, string) {
	admin.Migrate()
	return 200, gin.H{"user": "a", "value": "Info"}, ""
}
