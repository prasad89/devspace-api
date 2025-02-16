package models

// DevSpac struct to hold devspace details
type Devspace struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Owner string `gorm:"uniqueIndex:owner_name_unique"`
	Name  string `gorm:"uniqueIndex:owner_name_unique"`
}

// DevSpac struct to hold devspace details
type DevspaceResponse struct {
	Name string `json:"name"`
}
