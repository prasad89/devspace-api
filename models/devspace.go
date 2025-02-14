package models

// DevSpac struct to hold devspace details
type DevSpace struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	User string `json:"user"`
	Name string `json:"name"`

	UniqueIndex string `gorm:"uniqueIndex:user_name_unique" json:"-"`
}
