package validators

import (
	"time"
)

////////////////////
//  REQUEST BODYs //
////////////////////

type RegistrationBody struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Email        string `json:"email" binding:"email,required"`
	Name         string `json:"name" binding:"required,lte=50"`
	Phone        string `json:"phone" binding:"required,lte=14"`
	StudentClass string `json:"student_class"`
}

type AuthBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TestAuthBody struct {
	UsernameTO string `json:"username_to" binding:"required"`
	PasswordTO string `json:"password_to" binding:"required"`
}

type UpdateStudentBody struct {
	ID           uint64 `json:"id" binding:"required"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email" binding:"email"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	StudentClass string `json:"student_class"`
	Target       int    `json:"target"`
}

type DeleteBody struct {
	ID uint64 `json:"id" binding:"required"`
}

type StaffRegistrationBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"email,required"`
	Name     string `json:"name" binding:"required,lte=50"`
	Phone    string `json:"phone" binding:"required,lte=14"`
	Access   string `json:"access" binding:"required"`
	Ma       bool   `json:"ma"`
	Fi       bool   `json:"fi"`
	Ki       bool   `json:"ki"`
	Bi       bool   `json:"bi"`
	Sos      bool   `json:"sos"`
	Sej      bool   `json:"sej"`
	Geo      bool   `json:"geo"`
	Eko      bool   `json:"eko"`
	Pu       bool   `json:"pu"`
	Ppu      bool   `json:"ppu"`
	Pk       bool   `json:"pk"`
	Pmm      bool   `json:"pmm"`
	Eng      bool   `json:"eng"`
}

type UpdateStaffBody struct {
	ID       uint64 `json:"id" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Access   string `json:"access"`
	Ma       bool   `json:"ma"`
	Fi       bool   `json:"fi"`
	Ki       bool   `json:"ki"`
	Bi       bool   `json:"bi"`
	Sos      bool   `json:"sos"`
	Sej      bool   `json:"sej"`
	Geo      bool   `json:"geo"`
	Eko      bool   `json:"eko"`
	Pu       bool   `json:"pu"`
	Ppu      bool   `json:"ppu"`
	Pk       bool   `json:"pk"`
	Pmm      bool   `json:"pmm"`
	Eng      bool   `json:"eng"`
}

type CreateTestBody struct {
	Name        string    `json:"name" binding:"required"`
	Price       int       `json:"price"`
	Image       string    `json:"image"`
	Ma          int       `json:"ma"`
	Fi          int       `json:"fi"`
	Ki          int       `json:"ki"`
	Bi          int       `json:"bi"`
	Pu          int       `json:"pu"`
	Ppu         int       `json:"ppu"`
	Pmm         int       `json:"pmm"`
	Pk          int       `json:"pk"`
	Eng         int       `json:"eng"`
	Sos         int       `json:"sos"`
	Sej         int       `json:"sej"`
	Eko         int       `json:"eko"`
	Geo         int       `json:"geo"`
	ScheduledAt time.Time `json:"scheduledAt" binding:"required"`
	EndsAt      time.Time `json:"endsAt" binding:"required"`
}

type UpdateTestBody struct {
	ID          uint64    `json:"id" binding:"required"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Image       string    `json:"image"`
	Ma          int       `json:"ma"`
	Fi          int       `json:"fi"`
	Ki          int       `json:"ki"`
	Bi          int       `json:"bi"`
	Pu          int       `json:"pu"`
	Ppu         int       `json:"ppu"`
	Pmm         int       `json:"pmm"`
	Pk          int       `json:"pk"`
	Eng         int       `json:"eng"`
	Sos         int       `json:"sos"`
	Sej         int       `json:"sej"`
	Eko         int       `json:"eko"`
	Geo         int       `json:"geo"`
	ScheduledAt time.Time `json:"scheduledAt"`
	EndsAt      time.Time `json:"endsAt"`
}

type CreateSolutionBody struct {
	Content    string `json:"content" binding:"required"`
	CreatorID  uint64 `json:"creator_id" binding:"required"`
	QuestionID uint64 `json:"question_id" binding:"required"`
}

type UpdateSolutionBody struct {
	ID         uint64 `json:"id" binding:"required"`
	Content    string `json:"content"`
	CreatorID  uint64 `json:"creator_id"`
	QuestionID uint64 `json:"question_id"`
}

type CreateQuestionBody struct {
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required"`
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

type UpdateQuestionBody struct {
	ID       uint64 `json:"id" binding:"required"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
