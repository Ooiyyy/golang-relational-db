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

	c.JSON(http.StatusOK, utils.ListResponse("list data kelas berhasil dimuat", data, meta))
}
