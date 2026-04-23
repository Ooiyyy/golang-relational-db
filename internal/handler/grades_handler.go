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

type GradesHandler struct {
	service service.GradesService
	logRepo repository.ActivityLogRepository
}

func NewGradesHandler(s service.GradesService, logRepo repository.ActivityLogRepository) *GradesHandler {
	return &GradesHandler{service: s, logRepo: logRepo}
}

func (h *GradesHandler) CreateGrades(c *gin.Context) {
	var req dto.CreateGradesReq

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErr := utils.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("validasi error", validationErr))
		return
	}

	grade := model.Grades{
		StudentID: req.StudentID,
		SubjectID: req.SubjectID,
		Score:     req.Score,
	}

	err := h.service.CreateGrades(&grade)
	if err != nil {
		if err.Error() == "siswa tidak ditemukan" || err.Error() == "mapel tidak ditemukan" {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("not found", err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("terjadi kesalahan pada server", err.Error()))
		return
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Menambah nilai siswa",
	}
	h.logRepo.Create(log)
	c.JSON(http.StatusCreated, utils.SuccessResponse("Nilai berhasil ditambahkan", grade))
}

func (h *GradesHandler) GetGradeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	data, err := h.service.GetGradeByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Data nilai tidak ditemukan", err.Error()))
		return
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Mengambil nilai siswa",
	}
	h.logRepo.Create(log)
	c.JSON(200, utils.SuccessResponse("Data nilai berhasil dimuat", data))
}

func (h *GradesHandler) UpdateGrade(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}
	_, err = h.service.GetGradeByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("not found", err.Error()))
		return
	}

	var req dto.UpdateGradesReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	grade := model.Grades{
		ID:        id,
		StudentID: req.StudentID,
		SubjectID: req.SubjectID,
		Score:     req.Score,
	}

	err = h.service.UpdateGrade(id, grade)
	if err != nil {
		if err.Error() == "siswa tidak ditemukan" || err.Error() == "mapel tidak ditemukan" {
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
		Aktivitas: "Mengubah nilai siswa",
	}
	h.logRepo.Create(log)

	c.JSON(200, utils.SuccessResponse("Data nilai berhasil di update", grade))
}

func (h *GradesHandler) DeleteGrade(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	err = h.service.DeleteGrade(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Data nilai tidak ditemukan", err.Error()))
		return
	}
	ip := c.ClientIP()

	log := model.ActivityLog{
		IP:        ip,
		Aktivitas: "Menghapus nilai siswa",
	}
	h.logRepo.Create(log)

	c.JSON(200, utils.SuccessResponse("Data nilai berhasil dihapus", nil))
}
