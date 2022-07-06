package service

import (
	entity "studybuddy-backend-fast/api/entity"
	repository "studybuddy-backend-fast/api/repository/solution"
)

type SolutionService interface {
	CreateSolution(solution entity.Solution) entity.Solution
	FindAllSolution() []entity.SolutionMin
	FindSolutionByID(id int) entity.Solution
	FindSolutionByQID(id int) entity.SolutionMin 
}

type solutionService struct {
	repo repository.SolutionRepo
}

func NewSolutionService(repo repository.SolutionRepo) SolutionService {
	return &solutionService{repo}
}

func (s *solutionService) CreateSolution(solution entity.Solution) entity.Solution {
	s.repo.CreateSolution(solution)
	return solution
}

func (s *solutionService) FindAllSolution() []entity.SolutionMin {
	return s.repo.FindAllSolutionMin()
}

func (s *solutionService) FindSolutionByID(id int) entity.Solution {
	return s.repo.FindSolutionByID(id)
}
func (s *solutionService) FindSolutionByQID(id int) entity.SolutionMin  {
	return s.repo.FindSolutionMinByQID(id)
}