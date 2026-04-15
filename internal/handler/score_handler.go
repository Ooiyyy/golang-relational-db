package handler

import (
	"golang-relational-db/internal/dto"
	"golang-relational-db/internal/service"
	"golang-relational-db/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ScoreHandler struct {
	service *service.ScoreService
}

func NewScoreHandler(s *service.ScoreService) *ScoreHandler {
	return &ScoreHandler{service: s}
}

func (h *ScoreHandler) Create(c *gin.Context) {
	var req dto.ScoreReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		validationErr := utils.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Gagal menyimpan data",
			"message": validationErr,
		})
		return
	}

	err = h.service.Create(dto.ScoreReq{
		SubjectID: req.SubjectID,
		StudentID: req.StudentID,
		Score:     req.Score,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Gagal menyimpan data",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Nilai berhasil disimpan",
	})
}

func (h *ScoreHandler) GetAll(c *gin.Context) {
	scores, err := h.service.GetAllScores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Gagal menyimpan data",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": scores})
}

func (h *ScoreHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	score, err := h.service.GetScoreByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": score})
}

func (h *ScoreHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	var updateReq struct {
		Score int `json:"score" binding:"required"`
	}

	if err := c.ShouldBindJSON(&updateReq); err != nil {
		validationErr := utils.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	}

	req := dto.ScoreReq{
		Score: updateReq.Score,
	}

	err = h.service.UpdateScore(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Nilai berhasil diupdate"})
}

func (h *ScoreHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID harus berupa angka"})
		return
	}

	err = h.service.DeleteScore(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Nilai berhasil dihapus"})
}
