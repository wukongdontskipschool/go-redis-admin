package admin

import (
	"log"
	"net/http"
	"redisadmin/internal/accessControl"

	"github.com/gin-gonic/gin"
)

func Get_user_list() (int, gin.H) {
	sub := accessControl.Enforcer.GetAllSubjects()
	role := accessControl.Enforcer.GetAllRoles()
	allNamedObjects := accessControl.Enforcer.GetAllNamedObjects("p")
	allNamedActions := accessControl.Enforcer.GetAllNamedActions("p")
	log.Println(sub, role, allNamedObjects, allNamedActions)
	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": 1}
}
