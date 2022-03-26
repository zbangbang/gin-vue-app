package model

type Category struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" form:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt MyTime `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt MyTime `json:"updated_at" gorm:"type:timestamp"`
}
