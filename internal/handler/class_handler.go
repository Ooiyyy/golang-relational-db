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

type ClassHandler struct {
	service *service.ClassService
}

func NewClassHandler(s *service.ClassService) *ClassHandler {
	return &ClassHandler{service: s}
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
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Terjadi kesalahan server", err.Error()))
		return
	}

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

	c.JSON(http.StatusOK, utils.ListResponse("list data kelas berhasil dimuat", data, meta))
}
