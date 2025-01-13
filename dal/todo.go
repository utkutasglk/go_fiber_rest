package dal

type Todo struct {
	ID int
	Title string
	Completed bool `gorm:"default:false"`
} 