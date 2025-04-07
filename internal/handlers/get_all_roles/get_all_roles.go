package getallroles

import (
	rolemodel "github.com/KusakinDev/Catering-Auth-Service/internal/models/role_model"
	"github.com/gin-gonic/gin"
)

func GetAllRoles(c *gin.Context) (int, []rolemodel.Role) {

	var role rolemodel.Role
	roles, err := role.GetAllRoles()
	if err != nil {
		return 404, []rolemodel.Role{}
	}
	return 200, roles
}
