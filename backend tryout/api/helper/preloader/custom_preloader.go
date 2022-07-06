package preloader

import (
	"gorm.io/gorm"
	entity "studybuddy-backend-fast/api/entity"
)

func QuestionMinPreloader(db *gorm.DB) *gorm.DB {
	return db.Table("questions").Order(`
	case 
		when type = 'pu' then 1
		when type = 'pk' then 2
		when type = 'ppu' then 3
		when type = 'pmm' then 4
		else 5
	end asc`).Find(&entity.QuestionMin{})
}

func QuestionAndSolutionPreloader(db *gorm.DB) *gorm.DB {
	return db.Table("questions").Select("questions.*, solutions.content").Joins("LEFT JOIN solutions ON solutions.question_id = questions.id").Order(`
	case 
		when type = 'pu' then 1
		when type = 'pk' then 2
		when type = 'ppu' then 3
		when type = 'pmm' then 4
		else 5
	end asc`).Find(&entity.QuestionMin{})
}

func StaffMinPreloader(db *gorm.DB) *gorm.DB {
	return db.Model(&entity.Staff{}).First(&entity.StaffMin{})
}

func TestBarePreloader(db *gorm.DB) *gorm.DB {
	return db.Table("tests").Find(&entity.TestBare{})
}
