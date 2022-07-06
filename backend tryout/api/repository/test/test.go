package repository

import (
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
	preloader "studybuddy-backend-fast/api/helper/preloader"
	"time"

	"gorm.io/gorm"
)

type TestRepo interface {
	FindAllTest() []entity.Test
	FindTestByID(id int) entity.TestQuestionMin
	FindTestBySlug(slug string) entity.TestQuestionMin
	CreateTest(newTest entity.Test) entity.Test
	UpdateTest(newTestData entity.Test)
	DeleteTest(test entity.Test)

	FindAllTestMin() []entity.TestMin
	FindTestByIDMin(id int) entity.TestMin
	FindTestBySlugMin(slug string) entity.TestMin

	FindAllTestBare() []entity.TestBare
	FindTestByIDBare(id int) entity.TestBare
	FindTestBySlugBare(slug string) entity.TestBare

	FindTestByIDWithSolution(id int) entity.TestQuestionAndSolution
	FindTestBySlugWithSolution(slug string) entity.TestQuestionAndSolution
	FindTestWithFilter(pagination map[string]int) []entity.TestMin
	NameandType(id int) entity.NameType
	NameandType2(id uint64) entity.NameType
	AvailableByID(id int) entity.AvaibleTest

	LastBattle(test entity.Test) entity.Last
	NameandID(id int) entity.BattleIkut
	FindAvaibleBattle() entity.AvaibleTest
	ArrayFindAvaibleBattle() []entity.AvaibleTest
	SoalBattle(ID int,subtest string) []entity.QuestionMin
}

type konak struct {
	connection *gorm.DB
}

func NewTestRepo() TestRepo {
	// koneksi ke db gais
	db := connection.Create()
	db.AutoMigrate(&entity.Test{})

	// end contoh
	return &konak{
		connection: db,
	}
}

func (db *konak) FindAllTest() []entity.Test {
	var tests []entity.Test
	db.connection.Find(&tests)
	return tests
}
func (db *konak) FindAvaibleBattle() entity.AvaibleTest {
	var test entity.AvaibleTest
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	db.connection.Table("tests").Where("endsAt > ?", now).Find(&test)
	return test
}
func (db *konak) ArrayFindAvaibleBattle() []entity.AvaibleTest {
	var test []entity.AvaibleTest
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	db.connection.Table("tests").Where("endsAt > ?", now).Find(&test)
	return test
}
func (db *konak) FindTestWithFilter(pagination map[string]int) []entity.TestMin {
	var tests []entity.TestMin
	db.connection.Model(&entity.Test{}).Offset((pagination["page"] - 1) * pagination["limit"]).Limit(pagination["limit"]).Order("scheduledAt DESC").Find(&tests)
	return tests
}

func (db *konak) FindTestByID(id int) entity.TestQuestionMin {
	var test entity.TestQuestionMin
	db.connection.Table("tests").Preload("Questions", preloader.QuestionMinPreloader).First(&test, id)
	return test
}

func (db *konak) FindTestBySlug(slug string) entity.TestQuestionMin {
	var test entity.TestQuestionMin
	db.connection.Table("tests").Preload("Questions", preloader.QuestionMinPreloader).Where("slug = ?", slug).First(&test)
	return test
}

func (db *konak) FindTestByIDWithSolution(id int) entity.TestQuestionAndSolution {
	var test entity.TestQuestionAndSolution
	db.connection.Table("tests").Preload("Questions", preloader.QuestionAndSolutionPreloader).First(&test, id)
	return test
}

func (db *konak) FindTestBySlugWithSolution(slug string) entity.TestQuestionAndSolution {
	var test entity.TestQuestionAndSolution
	db.connection.Table("tests").Preload("Questions", preloader.QuestionAndSolutionPreloader).Where("slug = ?", slug).First(&test)
	return test
}

func (db *konak) CreateTest(newTest entity.Test) entity.Test {
	db.connection.Create(&newTest)
	return newTest
}

func (db *konak) UpdateTest(newTestData entity.Test) {
	db.connection.Save(&newTestData)
}

func (db *konak) DeleteTest(test entity.Test) {
	db.connection.Delete(&test)
}

func (db *konak) FindAllTestMin() []entity.TestMin {
	var tests []entity.TestMin
	db.connection.Model(&entity.Test{}).Find(&tests)
	return tests
}

func (db *konak) FindTestByIDMin(id int) entity.TestMin {
	var test entity.TestMin
	db.connection.Model(&entity.Test{}).First(&test, id)
	return test
}

func (db *konak) FindTestBySlugMin(slug string) entity.TestMin {
	var test entity.TestMin
	db.connection.Model(&entity.Test{}).Where("slug = ?", slug).First(&test)
	return test
}

func (db *konak) FindAllTestBare() []entity.TestBare {
	var tests []entity.TestBare
	db.connection.Model(&entity.Test{}).Find(&tests)
	return tests
}

func (db *konak) FindTestByIDBare(id int) entity.TestBare {
	var test entity.TestBare
	db.connection.Model(&entity.Test{}).First(&test, id)
	return test
}
func (db *konak) NameandType(id int) entity.NameType {
	var test entity.NameType
	db.connection.Model(&entity.Test{}).First(&test, id)
	return test
}
func (db *konak) NameandType2(id uint64) entity.NameType {
	var test entity.NameType
	db.connection.Model(&entity.Test{}).First(&test, id)
	return test
}
func (db *konak) NameandID(id int) entity.BattleIkut {
	var test entity.BattleIkut
	db.connection.Model(&entity.Test{}).First(&test, id)
	return test
}

func (db *konak) AvailableByID(id int) entity.AvaibleTest {
	var test entity.AvaibleTest
	db.connection.Table("tests").Where("id = ?", id).First(&test)
	return test
}

func (db *konak) FindTestBySlugBare(slug string) entity.TestBare {
	var test entity.TestBare
	db.connection.Model(&entity.Test{}).Where("slug = ?", slug).First(&test)
	return test
}

func (db *konak) LastBattle(test entity.Test) entity.Last {
	var last entity.Last
	db.connection.Last(&test)
	last.Name = test.Name
	last.Image = test.Image
	last.ScheduledAt = test.ScheduledAt
	return last
}

func (db *konak) SoalBattle(ID int,subtest string) []entity.QuestionMin {
	var hehe []entity.QuestionMin
	db.connection.Table("test_question").Select("*,question_id").Joins("left join questions on questions.id = test_question.question_id").Where("test_question.test_id = ? AND questions.type = ?",ID,subtest).Find(&hehe)
	return hehe
}
