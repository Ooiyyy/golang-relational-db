package middleware

import (
	"golang-relational-db/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RoleCheck adalah middleware untuk memisahkan hak akses antara guru dan siswa
func RoleCheck(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role")

		// Jika token tidak memiliki role
		if userRole == "" {
			utils.SendErrorResponse(c, http.StatusForbidden, "Akses ditolak", "Role di token tidak terbaca atau token dibuat sebelum fitur role aktif")
			c.Abort()
			return
		}

		// Mengecek apakah userRole termasuk di daftar yang diperbolehkan di rute ini
		isAllowed := false
		for _, role := range allowedRoles {
			// Toleransi penulisan Guru, guru, GURU
			if strings.ToLower(userRole) == strings.ToLower(role) {
				isAllowed = true
				break
			}
		}

		// Tolak akses apabila datanya tidak sesuai (bukan guru, dst)
		if !isAllowed {
			utils.SendErrorResponse(c, http.StatusForbidden, "Akses ditolak", "Tingkat hak akses anda tidak mencukupi untuk menggunakan fitur ini")
			c.Abort()
			return
		}

		c.Next()
	}
}
