package handler

import (
	"golang-relational-db/internal/dto"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
	"golang-relational-db/internal/service"
	"golang-relational-db/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// handler untuk HTTP layer
type TeacherHandler struct {
	service *service.TeacherService
	logRepo repository.ActivityLogRepository
}

// constructor
func NewTeacherHandler(s *service.TeacherService, logRepo repository.ActivityLogRepository) *TeacherHandler {
	return &TeacherHandler{service: s, logRepo: logRepo}
}

func (h *TeacherHandler) Create(c *gin.Context) {
	var req dto.CreateTeacherReq

	// bind JSON ke DTO
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErr := utils.FormatValidationError(err)
		c.JSON(400, utils.ErrorResponse("Validation error", validationErr))
		return
	}

	// convert DTO → model
	teacher := model.Teachers{
		Name:  req.Name,
		Email: req.Email,
	}

	err := h.service.CreateTeacher(&teacher)
	if err != nil {
		if err.Error() == "email sudah digunakan" {
			c.JSON(http.StatusConflict, utils.ErrorResponse("konflik data", err.Error()))
			return
		}
		c.JSON(500, utils.ErrorResponse("Terjadi kesalahan pada server", err.Error()))
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Menambahkan data guru " + teacher.Name,
	}

	h.logRepo.Create(log)
	c.JSON(201, utils.SuccessResponse("Guru berhasil ditambahkan", teacher))
}

func (h *TeacherHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	data, err := h.service.GetTeacherByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Data Guru tidak ditemukan", err.Error()))
		return
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Mengambil data guru",
	}

	h.logRepo.Create(log)

	c.JSON(200, utils.SuccessResponse("Data guru berhasil dimuat", data))
}

func (h *TeacherHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	search := c.Query("search")
	sortBy := c.DefaultQuery("sort_by", "id")
	order := c.DefaultQuery("order", "asc")

	data, total, err := h.service.GetAllTeachers(page, limit, search, sortBy, order)
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
		Aktivitas: "Mengambil list data guru",
	}

	h.logRepo.Create(log)

	c.JSON(200, utils.ListResponse("List guru berhasil dimuat", data, meta))
}

func (h *TeacherHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	var req dto.UpdateTeacherReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	teacher := model.Teachers{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}

	err = h.service.UpdateTeacher(id, teacher)
	if err != nil {
		validationErr := utils.FormatValidationError(err)
		c.JSON(400, utils.ErrorResponse("Validation error", validationErr))
		return
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Mengubah data guru",
	}

	h.logRepo.Create(log)

	c.JSON(200, utils.SuccessResponse("Data guru berhasil di update", teacher))
}

func (h *TeacherHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	err = h.service.DeleteTeacher(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Data guru tidak ditemukan", err.Error()))
		return
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Menghapus data guru",
	}

	h.logRepo.Create(log)

	c.JSON(200, utils.SuccessResponse("Data guru berhasil dihapus", nil))
}
