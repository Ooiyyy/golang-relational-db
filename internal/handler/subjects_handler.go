package handler

import (
	"golang-relational-db/internal/dto"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/service"
	"golang-relational-db/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubjectsHandler struct {
	service *service.SubjectsService
}

func NewSubjectsHandler(s *service.SubjectsService) *SubjectsHandler {
	return &SubjectsHandler{service: s}
}

func (h *SubjectsHandler) CreateSubjects(c *gin.Context) {
	var req dto.CreateSubjectsReq

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErr := utils.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, utils.ErrorResponse("Validasi error", validationErr))
		return
	}

	subject := model.Subjects{
		Name:      req.Name,
		TeacherID: req.TeacherID,
	}

	err := h.service.CreateSubjects(&subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Terjadi kesalahan server", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse("Kelas berhasil ditambahkan", subject))
}

func (h *SubjectsHandler) GetAllSubjects(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	search := c.Query("search")
	sortBy := c.DefaultQuery("sort_by", "id")
	order := c.DefaultQuery("order_by", "asc")

	data, total, err := h.service.GetAllSubjects(page, limit, search, sortBy, order)
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

	c.JSON(http.StatusOK, utils.ListResponse("list data pelajaran berhasil dimuat", data, meta))
}

func (h *SubjectsHandler) GetSubjectByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	data, err := h.service.GetSubjectByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Data pelajaran tidak ditemukan", err.Error()))
		return
	}

	c.JSON(200, utils.SuccessResponse("Data pelajaran berhasil dimuat", data))
}

func (h *SubjectsHandler) UpdateSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	var req dto.UpdateSubjectsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	subject := model.Subjects{
		Id:        id,
		Name:      req.Name,
		TeacherID: req.TeacherID,
	}

	err = h.service.UpdateSubject(id, subject)
	if err != nil {
		validationErr := utils.FormatValidationError(err)
		c.JSON(400, utils.ErrorResponse("Validation error", validationErr))
		return
	}

	c.JSON(200, utils.SuccessResponse("Data pelajaran berhasil di update", subject))
}

func (h *SubjectsHandler) DeleteSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	err = h.service.DeleteSubject(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Data pelajaran tidak ditemukan", err.Error()))
		return
	}

	c.JSON(200, utils.SuccessResponse("Data pelajaran berhasil dihapus", nil))
}
