package handler

import (
	"golang-relational-db/internal/dto"
	"golang-relational-db/internal/service"
	"golang-relational-db/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	service *service.SubjectService
}

func NewSubjectHandler(s *service.SubjectService) *SubjectHandler {
	return &SubjectHandler{service: s}
}

func (h *SubjectHandler) Create(c *gin.Context) {
	var req dto.SubjectReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		validationErr := utils.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErr,
		})
		return
	}

	// Otomatis ambil ID User sang Guru dari token JWT yang sudah dipecah di Middleware sebelumnya!
	req.TeacherID = c.GetInt("id")

	err = h.service.CreateSubject(dto.SubjectReq{
		Name:      req.Name,
		TeacherID: req.TeacherID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Gagal menyimpan data",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Subject berhasil dibuat",
	})
}

func (h *SubjectHandler) GetAll(c *gin.Context) {
	subjects, err := h.service.GetAllSubjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": subjects})
}

func (h *SubjectHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	subject, err := h.service.GetSubjectByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": subject})
}

func (h *SubjectHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	var req dto.SubjectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErr := utils.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	}

	// Otomatis pakai ID guru yang sedang login
	req.TeacherID = c.GetInt("id")

	err = h.service.UpdateSubject(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject berhasil diupdate"})
}

func (h *SubjectHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	err = h.service.DeleteSubject(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject berhasil dihapus"})
}
