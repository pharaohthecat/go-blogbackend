package models

type Blog struct{
	Id uint `json:"id"`
	Title string `json:title`
	Description string `json:description`
	Image string `json:image`
	UserID string `json:userid`
	// Foreing Key
	User User `json:"user";gorm:"foreignkey:UserID"`
}