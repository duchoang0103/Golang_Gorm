package models

import (
	"gorm.io/gorm"
)

type Template struct {
	//gorm.Model
	ID            uint32 `gorm:"primaryKey;autoIncrement" json:"id"`
	Template_code string `gorm:"not null" json:"template_code"`
	Dept_id       string `gorm:" null" json:"dept_id"`
	Name          string `gorm:: null" json:"name"`
	Html          string `gorm:" null" json:"html"`
}

//create a template
func CreateTemplate(db *gorm.DB, Template *Template) (err error) {
	err = db.Create(Template).Error
	if err != nil {
		return err
	}
	return nil
}

//get template
func GetTemplates(db *gorm.DB, Template *[]Template) (err error) {
	err = db.Find(Template).Error
	if err != nil {
		return err
	}
	return nil
}

//get Template by id
func GetTemplate(db *gorm.DB, Template *Template, id int) (err error) {
	err = db.Where("id = ?", id).First(Template).Error
	if err != nil {
		return err
	}
	return nil
}

//update Template
func UpdateTemplate(db *gorm.DB, Template *Template) (err error) {
	db.Save(Template)
	return nil
}

//delete Template
func DeleteTemplate(db *gorm.DB, Template *Template, id int) (err error) {
	db.Where("id = ?", id).Delete(Template)
	return nil
}
