package entity

type Todo struct {
	ID     int64  `gorm:"primary_key:auto_increment" json:"-"`
	Name   string `gorm:"type:text" json:"-"`
	Image  string `gorm:"type:text" json:"-"`
	Bounty uint64 `gorm:"type:bigint" json:"-"`
	UserID int64  `gorm:"not null" json:"-"`
	User   User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}