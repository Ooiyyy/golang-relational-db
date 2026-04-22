package handler

import (
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
	"golang-relational-db/internal/service"
	"golang-relational-db/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// handler untuk HTTP layer
type StudentRelationalHandler struct {
	service *service.StudentRelationalService
	logRepo repository.ActivityLogRepository
}

// constructor
func NewStudentRelationalHandler(s *service.StudentRelationalService, logRepo repository.ActivityLogRepository) *StudentRelationalHandler {
	return &StudentRelationalHandler{service: s, logRepo: logRepo}
}

func (h *StudentRelationalHandler) GetAllDetails(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	search := c.Query("search")
	classID := c.Query("class_id")
	sortBy := c.DefaultQuery("sort_by", "id")
	order := c.DefaultQuery("order", "asc")

	data, total, err := h.service.GetStudentsDetail(page, limit, search, classID, sortBy, order)
	if err != nil {
		c.JSON(500, utils.ErrorResponse("Terjadi kesalahan pada server", err.Error()))
		return
	}
	totalPage := (total + limit - 1) / limit

	meta := utils.Meta{
		Page:      page,
		Limit:     limit,
		TotalData: total,
		TotalPage: totalPage,
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Mengambil list detail data siswa",
	}
	h.logRepo.Create(log)

	c.JSON(200, utils.ListResponse("List detail siswa berhasil dimuat", data, meta))
}

func (h *StudentRelationalHandler) GetStudentsGrade(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	search := c.Query("search")
	classID := c.Query("class_id")
	sortBy := c.DefaultQuery("sort_by", "id")
	order := c.DefaultQuery("order", "asc")

	data, total, err := h.service.GetStudentsGrade(page, limit, search, classID, sortBy, order)
	if err != nil {
		c.JSON(500, utils.ErrorResponse("Terjadi kesalahan pada server", err.Error()))
		return
	}
	totalPage := (total + limit - 1) / limit

	meta := utils.Meta{
		Page:      page,
		Limit:     limit,
		TotalData: total,
		TotalPage: totalPage,
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Mengambil list nilai siswa",
	}
	h.logRepo.Create(log)

	c.JSON(200, utils.ListResponse("List nilai siswa berhasil dimuat", data, meta))
}

func (h *StudentRelationalHandler) GetStudentsAvg(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	search := c.Query("search")
	classID := c.Query("class_id")
	sortBy := c.DefaultQuery("sort_by", "id")
	order := c.DefaultQuery("order", "asc")

	data, total, err := h.service.GetStudensAvg(page, limit, search, classID, sortBy, order)
	if err != nil {
		c.JSON(500, utils.ErrorResponse("Terjadi kesalahan pada server", err.Error()))
		return
	}
	totalPage := (total + limit - 1) / limit

	meta := utils.Meta{
		Page:      page,
		Limit:     limit,
		TotalData: total,
		TotalPage: totalPage,
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Mengambil list rata-rata nilai siswa",
	}
	h.logRepo.Create(log)

	c.JSON(200, utils.ListResponse("List rata-rata siswa berhasil dimuat", data, meta))
}
