package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"myapp/database"
	"myapp/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go/v4"
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

//-------------------------sendmail Temple

// Your available domain names can be found here:
// (https://app.mailgun.com/app/domains)
var yourDomain string = "***s*andbox1b4f79d8ff50472487ffdd638292d793.mailgun.org" // e.g. mg.yourcompany.com

// You can find the Private API Key in your Account Menu, under "Settings":
// (https://app.mailgun.com/app/account/security)
var privateAPIKey string = "***4*7c46f002b5121fb6408d78aab109051-1b3a03f6-07cd747d" //"210edf5ee4c770ebcba833e747bc5eea-1b3a03f6-5e0578d0"

func (repository *TemplateRepo) SendMailTemplate(c *gin.Context) {

	var templatesendmail models.Templatesendmail
	c.BindJSON(&templatesendmail)
	mailrequest := templatesendmail
	id := mailrequest.Id

	var template models.Template

	err := models.GetTemplate(repository.Db, &template, id)
	htmldb := template.Html

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// targetstringrequest := mailrequest.Array[0].Targetstring
	// linkrequest := mailrequest.Array[0].Link
	linkrequest := mailrequest.Array[0].Key
	targetstringrequest := mailrequest.Array[1].Key

	newstring := strings.Replace(htmldb, "<##link##>", linkrequest, -1)
	newstring2 := strings.Replace(newstring, "<##string##>", targetstringrequest, -1)

	///////////////////////////PHAN MAILGUN
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	//When you have an EU-domain, you must specify the endpoint:
	//mg.SetAPIBase("https://api.eu.mailgun.net/v3")
	sender := "hoangkimanhduc01031999@gmail.com"
	//sender := "Mailgun Sandbox<postmaster@sandboxb35142e463314da09b5baa7fd0878ffb.mailgun.org>"
	subject := "Hello Hoang Kim Anh Duc"
	body := newstring2
	println(body)
	recipient := "duc0103999@gmail.com"

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, idt, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %s Resp: %s\n", idt, resp)
	//////////////////////////
	c.JSON(http.StatusOK, gin.H{"message": "Template sendmail successfully"})
}
