package handler

import (
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
	"golang-relational-db/internal/service"
	"golang-relational-db/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityLogHandler struct {
	service *service.LogService
	logRepo repository.ActivityLogRepository
}

func NewActivityLogHandler(s *service.LogService, logRepo repository.ActivityLogRepository) *ActivityLogHandler {
	return &ActivityLogHandler{service: s, logRepo: logRepo}
}

// func (h *ActivityLogHandler) CreateClass(c *gin.Context) {
// 	var req dto.CreateClassReq

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		validationErr := utils.FormatValidationError(err)
// 		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Validasi error", validationErr))
// 		return
// 	}

// 	class := model.Classes{
// 		Name:      req.Name,
// 		TeacherID: req.TeacherID,
// 	}

// 	err := h.service.CreateClass(&class)
// 	if err != nil {
// 		if err.Error() == "guru tidak ditemukan" {
// 			c.JSON(http.StatusNotFound, utils.ErrorResponse("not found", err.Error()))
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Terjadi kesalahan server", err.Error()))
// 		return
// 	}
// 	ip := c.ClientIP()

// 	log := model.ActivityLog{
// 		IP:        ip,
// 		Aktivitas: "Menambah data kelas",
// 	}
// 	h.logRepo.Create(log)

// 	c.JSON(http.StatusCreated, utils.SuccessResponse("Kelas berhasil ditambahkan", class))
// }

func (h *ActivityLogHandler) GetAllLog(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	search := c.Query("search")
	sortBy := c.DefaultQuery("sort_by", "id")
	order := c.DefaultQuery("order", "asc")

	log.Printf("DEBUG : %s %s", sortBy, order)

	data, total, err := h.service.GetAllLog(page, limit, search, sortBy, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Terjadi kesalahan server", err.Error()))
		return
	}
	totalPage := (total + limit - 1) / limit

	meta := utils.Meta{
		Page:      totalPage,
		Limit:     limit,
		TotalData: total,
		TotalPage: totalPage,
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Mengambil list data activity log",
	}
	h.logRepo.Create(log)

	c.JSON(http.StatusOK, utils.ListResponse("list activity log berhasil dimuat", data, meta))
}

func (h *ActivityLogHandler) GetLogByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	data, err := h.service.GetLogByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Data activity log tidak ditemukan", err.Error()))
		return
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Mengambil data activity log",
	}
	h.logRepo.Create(log)

	c.JSON(200, utils.SuccessResponse("Data activity log berhasil dimuat", data))
}
