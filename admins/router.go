package admins

import "github.com/gin-gonic/gin"

func loginAdmin(c *gin.Context) {
	c.String(200, "Login admin!")
}

func updateAdminData(c *gin.Context) {
	c.String(200, "updateAdminData!")
}

func AdminRouter(r *gin.Engine) {
	r.POST("/admin/login", loginAdmin)
	r.PATCH("/admin", updateAdminData)
}
