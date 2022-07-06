package repository

import (
	"gorm.io/gorm"
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"
)

type SolutionRepo interface {
	// BASIC CRUD
	CreateSolution(solution entity.Solution)
	FindAllSolution() []entity.Solution
	FindAllSolutionMin() []entity.SolutionMin
	FindSolutionByID(id int) entity.Solution
	FindSolutionMinByID(id int) entity.SolutionMin
	UpdateSolution(solution entity.Solution)
	DeleteSolution(solution entity.Solution)
	FindSolutionMinByQID(id int) entity.SolutionMin 
}

type database struct {
	conn *gorm.DB
}

func NewSolutionRepo() SolutionRepo {
	db := connection.Create()
	db.AutoMigrate(&entity.Solution{})
	return &database{db}
}

////////////////
// BASIC CRUD //
////////////////

func (d *database) CreateSolution(solution entity.Solution) {
	d.conn.Create(&solution)
}

func (d *database) FindAllSolution() []entity.Solution {
	var solutions []entity.Solution
	d.conn.Find(&solutions)
	return solutions
}

func (d *database) FindAllSolutionMin() []entity.SolutionMin {
	var solutions []entity.SolutionMin
	d.conn.Model(&entity.Solution{}).Find(&solutions)
	return solutions
}

func (d *database) FindSolutionByID(id int) entity.Solution {
	var solution entity.Solution
	d.conn.First(&solution, id)
	return solution
}

func (d *database) FindSolutionMinByID(id int) entity.SolutionMin {
	var solution entity.SolutionMin
	d.conn.Model(&entity.Solution{}).First(&solution, id)
	return solution
}
func (d *database) FindSolutionMinByQID(id int) entity.SolutionMin {
	var solution entity.SolutionMin
	d.conn.Table("solutions").Where("question_id = ?",id).First(&solution)
	return solution
}

func (d *database) UpdateSolution(solution entity.Solution) {
	d.conn.Save(&solution)
}

func (d *database) DeleteSolution(solution entity.Solution) {
	d.conn.Delete(&solution)
}
