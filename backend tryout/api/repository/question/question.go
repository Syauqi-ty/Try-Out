package repository

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"

	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
	preloader "studybuddy-backend-fast/api/helper/preloader"
)

var timezone string = viper.GetString("timezone")

type QuestionRepo interface {
	// BASIC CRUD
	CreateQuestion(question entity.Question) entity.Question
	FindAllQuestion() []entity.Question
	FindQuestionByID(id int) entity.Question
	UpdateQuestion(question entity.Question) int
	DeleteQuestion(id int)

	// MIN QUERY
	FindAllQuestionMin() []entity.QuestionMin
	FindQuestionMinByID(id int) entity.QuestionMin

	// PANTAT
	FindQuestionsOfTest(TestID int) []entity.QuestionMin
	TemporaryDelete(id int)
}

type questionRepo struct {
	conn *gorm.DB
}

func NewQuestionRepo() QuestionRepo {
	db := connection.Create()
	db.AutoMigrate(&entity.Question{})
	return &questionRepo{conn: db}
}

///////////////////////
// PRATICAL USECASES //
///////////////////////

func (r *questionRepo) FindQuestionsOfTest(TestID int) []entity.QuestionMin {
	var test entity.TestQuestionMin
	r.conn.Table("tests").Preload("Questions", preloader.QuestionMinPreloader).First(&test, TestID)
	questions := test.Questions
	return questions
}

func (r *questionRepo) FindQuestionAndSolutionByID(id int) entity.QuestionAndSolution {
	var q entity.QuestionAndSolution
	r.conn.Table("questions").Select("questions.*, solutions.content").Where("questions.id = ?", id).Joins("solutions ON questions.id = solutions.question_id").First(&q)
	return q
}

///////////////
// MIN QUERY //
///////////////

func (r *questionRepo) FindAllQuestionMin() []entity.QuestionMin {
	var questions []entity.QuestionMin
	r.conn.Model(&entity.Question{}).Find(&questions)
	return questions
}

func (r *questionRepo) FindQuestionMinByID(id int) entity.QuestionMin {
	var q entity.QuestionMin
	r.conn.Model(&entity.Question{}).First(&q, id)
	return q
}

////////////////
// BASIC CRUD //
////////////////

func (r *questionRepo) CreateQuestion(question entity.Question) entity.Question {
	loc, _ := time.LoadLocation(timezone)
	question.CreatedAt = time.Now().In(loc)
	question.UpdatedAt = time.Now().In(loc)
	r.conn.Create(&question)
	return question
}

func (r *questionRepo) FindAllQuestion() []entity.Question {
	var questions []entity.Question
	r.conn.Preload("Creator", preloader.StaffMinPreloader).Preload("LastUpdator", preloader.StaffMinPreloader).Find(&questions)
	return questions
}

func (r *questionRepo) FindQuestionByID(id int) entity.Question {
	var question entity.Question
	r.conn.Preload("UsedFor", preloader.TestBarePreloader).Preload("Creator", preloader.StaffMinPreloader).Preload("LastUpdator", preloader.StaffMinPreloader).First(&question, id)
	return question
}

func (r *questionRepo) UpdateQuestion(question entity.Question) int{
	loc, _ := time.LoadLocation(timezone)
	data := r.conn.Model(&question).Where("id", question.ID).Find(&question)
	if data.RowsAffected == 0{
		return 0
	}else{
		r.conn.Model(&question).Where("id", question.ID).Updates(entity.Question{Name: question.Name, Type: question.Type, Question: question.Question, Duration: question.Duration, CreatorID: question.CreatorID,CreatedAt:question.CreatedAt,UpdatedAt:time.Now().In(loc),DeletedAt:question.DeletedAt,UsedAt:question.UsedAt})
		return 1
	}
}

func (r *questionRepo) DeleteQuestion(id int) {
	var question entity.Question
	r.conn.Where("id = ?",id).Delete(&question)
}
func (r *questionRepo) TemporaryDelete(id int) {
	loc, _ := time.LoadLocation(timezone)
	r.conn.Where("id = ?",id).Updates(entity.Question{DeletedAt:time.Now().In(loc)})
}