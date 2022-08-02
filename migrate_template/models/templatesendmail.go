package models

type address struct {
	Link         string `json:"link"`
	Targetstring string `json:"string"`
}

type Templatesendmail struct {
	//gorm.Model
	Id    int `json:"id"`
	Array []address
}
