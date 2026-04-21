package handler

import (
	"golang-relational-db/internal/dto"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/service"
	"golang-relational-db/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GradesHandler struct {
	service service.GradesService
}

func NewGradesHandler(s service.GradesService) *GradesHandler {
	return &GradesHandler{service: s}
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
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("terjadi kesalahan pada server", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.SuccessResponse("Nilai berhasil ditambahkan", grade))
}
