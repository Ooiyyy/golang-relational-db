package utils

import (
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"

	"github.com/gin-gonic/gin"
)

func TambahLog(r repository.ActivityLogRepository, c *gin.Context, aktivitas string) {

	newLog := model.ActivityLog{
		IP:        c.ClientIP(),
		Aktivitas: aktivitas,
	}
	_ = r.Create(newLog)

}
