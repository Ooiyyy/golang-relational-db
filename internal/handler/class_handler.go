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

type ClassHandler struct {
	service *service.ClassService
	logRepo repository.ActivityLogRepository
}

func NewClassHandler(s *service.ClassService, logRepo repository.ActivityLogRepository) *ClassHandler {
	return &ClassHandler{service: s, logRepo: logRepo}
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	var req dto.CreateClassReq

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErr := utils.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Validasi error", validationErr))
		return
	}

	class := model.Classes{
		Name:      req.Name,
		TeacherID: req.TeacherID,
	}

	err := h.service.CreateClass(&class)
	if err != nil {
		if err.Error() == "guru tidak ditemukan" {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("not found", err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Terjadi kesalahan server", err.Error()))
		return
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Menambah data kelas",
	}
	h.logRepo.Create(log)

	c.JSON(http.StatusCreated, utils.SuccessResponse("Kelas berhasil ditambahkan", class))
}

func (h *ClassHandler) GetAllClass(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	search := c.Query("search")
	sortBy := c.DefaultQuery("sort_by", "id")
	order := c.DefaultQuery("order_by", "asc")

	data, total, err := h.service.GetAllClass(page, limit, search, sortBy, order)
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
		Aktivitas: "Mengambil list data kelas",
	}
	h.logRepo.Create(log)

	c.JSON(http.StatusOK, utils.ListResponse("list data kelas berhasil dimuat", data, meta))
}

func (h *ClassHandler) GetClassByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	data, err := h.service.GetClassByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Data kelas tidak ditemukan", err.Error()))
		return
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Mengambil data kelas",
	}
	h.logRepo.Create(log)

	c.JSON(200, utils.SuccessResponse("Data kelas berhasil dimuat", data))
}

func (h *ClassHandler) UpdateClass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	_, err = h.service.GetClassByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("not found", err.Error()))
		return
	}

	var req dto.UpdateClassReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	class := model.Classes{
		ID:        id,
		Name:      req.Name,
		TeacherID: req.TeacherID,
	}

	err = h.service.UpdateClass(id, class)
	if err != nil {
		if err.Error() == "guru tidak ditemukan" {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("not found", err.Error()))
			return
		}
		validationErr := utils.FormatValidationError(err)
		c.JSON(400, utils.ErrorResponse("Validation error", validationErr))
		return
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Mengubah data kelas",
	}
	h.logRepo.Create(log)

	c.JSON(200, utils.SuccessResponse("Data kelas berhasil di update", class))
}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	err = h.service.DeleteClass(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Data kelas tidak ditemukan", err.Error()))
		return
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Menghapus data kelas",
	}
	h.logRepo.Create(log)

	c.JSON(200, utils.SuccessResponse("Data kelas berhasil dihapus", nil))
}
