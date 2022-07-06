package entity


type Feedback struct{
	ID uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Name string `json:"name"`
	Email string `json:"email"`
	Message string `json:"message"`
}