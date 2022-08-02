package controllers

import (
	"errors"
	"myapp/database"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TemplateRepo struct {
	Db *gorm.DB
}

func New() *TemplateRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Template{})
	return &TemplateRepo{Db: db}
}

//create Template
func (repository *TemplateRepo) CreateTemplate(c *gin.Context) {
	var template models.Template
	c.BindJSON(&template)
	err := models.CreateTemplate(repository.Db, &template)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, template)
}

//get Templates
func (repository *TemplateRepo) GetTemplates(c *gin.Context) {
	var template []models.Template
	err := models.GetTemplates(repository.Db, &template)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, template)
}

//get Template by id
func (repository *TemplateRepo) GetTemplate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var template models.Template
	err := models.GetTemplate(repository.Db, &template, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, template)
}

// update user
func (repository *TemplateRepo) UpdateTemplate(c *gin.Context) {
	var template models.Template
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetTemplate(repository.Db, &template, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&template)
	err = models.UpdateTemplate(repository.Db, &template)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, template)
}

// delete Template
func (repository *TemplateRepo) DeleteTemplate(c *gin.Context) {
	var template models.Template
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteTemplate(repository.Db, &template, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Template deleted successfully"})
}
