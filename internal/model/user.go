package model

// User struct
type User struct {
	ID          uint64 `gorm:"primaryKey"`
	Shortname   string
	Description string
	Email       string `gorm:"uniqueIndex;not null" json:"email"`
	Password    string `gorm:"not null" json:"password"`
	Name        string
	SecondName  string
	BirthDay    string
	CreatedAt   time.time
	UpdatedAt   time.time
	DeletedAt   *time.time
	ImageUrl    *string
	Username    string `gorm:"uniqueIndex;not null" json:"username"`
	Names       string `json:"names"`
}
