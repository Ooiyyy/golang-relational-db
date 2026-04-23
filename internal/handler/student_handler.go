package handler

import (
	"golang-relational-db/internal/dto"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
	"golang-relational-db/internal/service"
	"golang-relational-db/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// handler untuk HTTP layer
type StudentHandler struct {
	service *service.StudentService
	logRepo repository.ActivityLogRepository
}

// constructor
func NewStudentHandler(s *service.StudentService, logRepo repository.ActivityLogRepository) *StudentHandler {
	return &StudentHandler{service: s, logRepo: logRepo}
}

func (h *StudentHandler) Create(c *gin.Context) {
	var req dto.CreateStudentReq

	// bind JSON ke DTO
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErr := utils.FormatValidationError(err)
		c.JSON(400, utils.ErrorResponse("validasi error", validationErr))
		return
	}

	// convert DTO → model
	student := model.Students{
		Name:    req.Name,
		Email:   req.Email,
		ClassID: req.ClassID,
	}

	err := h.service.CreateStudent(&student)
	if err != nil {
		if err.Error() == "kelas tidak ditemukan" {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("not found", err.Error()))
			return
		}
		if err.Error() == "email sudah digunakan" {
			c.JSON(http.StatusConflict, utils.ErrorResponse("konflik data", err.Error()))
			return
		}
		c.JSON(500, utils.ErrorResponse("Terjadi kesalahan pada server", err.Error()))
		return
	}
	utils.TambahLog(h.logRepo, c, "Menambah data siswa")

	c.JSON(201, utils.SuccessResponse("Siswa berhasil ditambahkan", student))
}

func (h *StudentHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	data, err := h.service.GetStudentByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Data siswa tidak ditemukan", err.Error()))
		return
	}
	utils.TambahLog(h.logRepo, c, "Mengambil data siswa")

	c.JSON(200, utils.SuccessResponse("Data siswa berhasil dimuat", data))
}

func (h *StudentHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	search := c.Query("search")
	classID := c.Query("class_id")
	sortBy := c.DefaultQuery("sort_by", "id")
	order := c.DefaultQuery("order", "asc")

	log.Printf("DEBUG : %s %s", sortBy, order)

	data, total, err := h.service.GetAllStudents(page, limit, search, classID, sortBy, order)
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
	utils.TambahLog(h.logRepo, c, "Mengambil list data siswa")

	c.JSON(200, utils.ListResponse("List siswa berhasil dimuat", data, meta))
}

func (h *StudentHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}
	_, err = h.service.GetStudentByID(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("not found", err.Error()))
		return
	}

	var req dto.UpdateStudentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	student := model.Students{
		ID:      id,
		Name:    req.Name,
		Email:   req.Email,
		ClassID: req.ClassID,
	}

	err = h.service.UpdateStudent(id, student)
	if err != nil {
		if err.Error() == "kelas tidak ditemukan" {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("not found", err.Error()))
			return
		}
		validationErr := utils.FormatValidationError(err)
		c.JSON(400, utils.ErrorResponse("Validation error", validationErr))
		return
	}
	utils.TambahLog(h.logRepo, c, "Mengubah data siswa")

	c.JSON(200, utils.SuccessResponse("Data siswa berhasil di update", student))
}

func (h *StudentHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, utils.ErrorResponse("id tidak valid", err.Error()))
		return
	}

	err = h.service.DeleteStudent(id)
	if err != nil {
		c.JSON(404, utils.ErrorResponse("Data siswa tidak ditemukan", err.Error()))
		return
	}
	utils.TambahLog(h.logRepo, c, "Menghapus data siswa")

	c.JSON(200, utils.SuccessResponse("Data siswa berhasil dihapus", nil))
}
