package module

type Category struct {
	ID   int    `gorm:"primay_key;aoto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
