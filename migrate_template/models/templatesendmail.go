package models

type address struct {
	Key string `json:"key"`
	//Targetstring string `json:"string"`
}

type Templatesendmail struct {
	//gorm.Model
	Id    int       `json:"id"`
	Array []address `json:"array"`
}
