package repository

import (
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"

	"gorm.io/gorm"
)

type ScoreRepo interface {
	// BASIC CRUD
	FindAllScore() []entity.Score
	FindScoreByID(id int) entity.Score
	CreateScore(score entity.Score) entity.Score
	UpdateScore(score entity.Score)
	DeleteScore(score entity.Score)

	// NOT SO BASIC CRUD
	AnswerQuestion(TestID int, StudentID int, Type string, Mark string, Correct bool)

	// SCORE QUERYS BY ID
	FindScoreByStudentsID(id int) []entity.Score
	FindScoreByTestID(id int) []entity.Score
	StudentScoreOfTest(studentsID int, testID int) []entity.Score

	// SCORE MIN QUERIES BY ID
	FindAllScoreMin() []entity.ScoreMin
	FindScoreMinByID(id int) entity.ScoreMin
	FindScoreMinByStudentsIDAndTestID(studentsID int, testID int) []entity.ScoreMin
	FindScoreMinByStudentsID(id int) []entity.ScoreMin
	BenarSalah(studentID int,testID int,subtest string) entity.BenarSalah 
	Quartile(testID int, subtest string) []entity.ScoreOnly

	// GROUPING QUERIES
	GroupScoreByTestID(testID int) []entity.ScoreOnly
	GroupScoreByStudentID(studentID int) []entity.ScoreOnly
	GroupScoreByTestIDAndStudentID(testID int, studentID int) []entity.ScoreOnly
	GroupScoreByStudentIDAndType(studentID int) []entity.ScoreOnly
	GroupScoreByTestIDAndType(testID int) []entity.ScoreOnly
	FindScoreByTestIDAndType(testID int, subtest string) []entity.ScoreOnly
	GroupScoreByTestIDAndPref(testID int, uniID int, prodiID int) []entity.ScoreOnly
	GroupScoreByTestIDWithType(testID int, subtest string) []entity.ScoreOnly
	GroupScoreByStudentIDLast(studentID int) []entity.ScoreOnly
	Graphic(test entity.Test, score entity.Score) []entity.IkutBattle
	RataAll(testID uint64, subtest string) []entity.ScoreOnly
	RataSekolah(testID int, school string) []entity.ScoreOnly
	RataSekolahType(testID uint64, school string, subtest string) []entity.ScoreOnly
	ScoreOne(testID uint64, studentID int) []entity.ScoreReal
	MaxAll(testID uint64, subtest string) []entity.ScoreOnly
	MaxSekolah(testID uint64, school string, subtest string) []entity.ScoreOnly
	MaxAllAverage(testID uint64) []entity.ScoreOnly
	MaxAllAverageSaintek(testID uint64) []entity.ScoreOnly
	MaxAllAverageSoshum(testID uint64) []entity.ScoreOnly
	MaxAllAverageSchool(testID uint64, school string) []entity.ScoreOnly
	MaxAllAverageSchoolSaintek(testID uint64, school string) []entity.ScoreOnly
	MaxAllAverageSchoolSoshum(testID uint64, school string) []entity.ScoreOnly
	Leaderboard(testID int) []entity.ScoreReal
	TigaBawah(testID int) []entity.ScoreReal
	MaxScore(testid int) []entity.Score
	MaxScoreSchool(testid int,school string) []entity.Score
	ScoreAnak(testID int, studentID int) []entity.ScoreReal
	Distribution(testID int, subtest string) (int64, int64, int64, int64, int64, int64, int64, int64, int64, int64)
	AVGDistributionTps(testID int) (int64, int64, int64, int64, int64, int64, int64, int64, int64, int64)
	AVGDistributionSaintek(testID int) (int64, int64, int64, int64, int64, int64, int64, int64, int64, int64)
	AVGDistributionSoshum(testID int) (int64, int64, int64, int64, int64, int64, int64, int64, int64, int64)
	Jawab(jawab entity.Jawaban) entity.Jawaban
}

type database struct {
	conn *gorm.DB
}

func NewScoreRepo() ScoreRepo {
	db := connection.Create()
	db.AutoMigrate(&entity.Score{})
	return &database{conn: db}
}

type Yuhu struct {
	Name string
	Type int
}

///////////////
// GROUPINGS //
///////////////

func (d *database) GroupScoreByTestIDAndStudentID(testID int, studentID int) []entity.ScoreOnly {
	var score []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("id, student_id, test_id, type, AVG(score) as score").Where("test_id = ?", testID).Group("student_id").Having("student_id = ?", studentID).Find(&score)
	return score
}

func (d *database) GroupScoreByTestIDAndSchool(testID int, studentID int, school string) []entity.ScoreOnly {
	var score []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("id, student_id, test_id, type, AVG(score) as score").Joins("LEFT JOIN students on students.id = scores.student_id").Where("scores.test_id = ? AND students.school = ?", testID, school).Group("scores.student_id").Find(&score)
	return score
}

func (d *database) GroupScoreByStudentID(studentID int) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*, AVG(score) AS score").Where("student_id = ?", studentID).Group("test_id").Find(&scores)
	return scores
}

func (d *database) GroupScoreByStudentIDAndType(studentID int) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*, AVG(score) AS score").Where("student_id = ?", studentID).Group("type").Find(&scores)
	return scores
}

func (d *database) GroupScoreByTestIDAndType(testID int) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*, AVG(score) AS score").Where("test_id = ?", testID).Group("type").Find(&scores)
	return scores
}

func (d *database) GroupScoreByStudentIDLast(studentID int) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*, AVG(score) AS score").Where("student_id = ?", studentID).Group("test_id").Order("test_id desc").Limit(1).Find(&scores)
	return scores
}

func (d *database) RataSekolahType(testID uint64, school string, subtest string) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*,test_id, AVG(score) AS score").Joins("LEFT JOIN students on students.id = scores.student_id").Where("test_id = ? AND type = ? AND students.school=?", testID, subtest, school).Order("score desc").Find(&scores)
	return scores
}
func (d *database) ScoreOne(testID uint64, studentID int) []entity.ScoreReal {
	var scores []entity.ScoreReal
	d.conn.Table("scores").Select("id,test_id,student_id,AVG(score) as score").Where("test_id=? AND student_id=?", testID, studentID).Find(&scores)
	return scores
}
func (d *database) MaxAllAverage(testID uint64) []entity.ScoreOnly {
	subtest := [5]string{"pu", "ppu", "pk", "pmm", "eng"}
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*,AVG(score) AS score, max(score) AS score").Where("test_id=? AND type IN (?)", testID, subtest).Group("student_id").Order("score desc").First(&scores)
	return scores
}
func (d *database) MaxAllAverageSaintek(testID uint64) []entity.ScoreOnly {
	subtest := [4]string{"ma", "fi", "ki", "bi"}
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*,AVG(score) AS score, max(score) AS score").Where("test_id=? AND type IN (?)", testID, subtest).Group("student_id").Order("score desc").First(&scores)
	return scores
}
func (d *database) MaxAllAverageSoshum(testID uint64) []entity.ScoreOnly {
	subtest := [4]string{"sos", "sej", "geo", "eko"}
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*,AVG(score) AS score, max(score) AS score").Where("test_id=? AND type IN (?)", testID, subtest).Group("student_id").Order("score desc").First(&scores)
	return scores
}
func (d *database) MaxAllAverageSchool(testID uint64, school string) []entity.ScoreOnly {
	subtest := [5]string{"pu", "ppu", "pk", "pmm", "eng"}
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*,AVG(score) AS score, max(score) AS score").Joins("LEFT JOIN students on students.id = scores.student_id").Where("test_id=? AND type IN (?) AND students.school=?", testID, subtest, school).Group("student_id").Find(&scores)
	return scores
}
func (d *database) MaxAllAverageSchoolSaintek(testID uint64, school string) []entity.ScoreOnly {
	subtest := [4]string{"ma", "fi", "ki", "bi"}
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*, AVG(score) AS score, max(score) AS score").Joins("LEFT JOIN students on students.id = scores.student_id").Where("test_id=? AND type IN (?) AND students.school=?", testID, subtest, school).Group("student_id").Find(&scores)
	return scores
}
func (d *database) MaxAllAverageSchoolSoshum(testID uint64, school string) []entity.ScoreOnly {
	subtest := [4]string{"sos", "sej", "geo", "eko"}
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*,AVG(score) AS score, max(score) AS score").Joins("LEFT JOIN students on students.id = scores.student_id").Where("test_id=? AND type IN (?) AND students.school=?", testID, subtest, school).Group("student_id").Find(&scores)
	return scores
}
func (d *database) RataSekolah(testID int, school string) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*, AVG(score) AS score").Joins("LEFT JOIN students on students.id = scores.student_id").Where("test_id = ? AND students.school=?", testID, school).Group("type").Find(&scores)
	return scores
}
func (d *database) MaxSekolah(testID uint64, school string, subtest string) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*, max(score) AS score").Joins("LEFT JOIN students on students.id = scores.student_id").Where("test_id = ? AND type = ? AND students.school=?", testID, subtest, school).Order("score desc").Find(&scores)
	return scores
}
func (d *database) RataAll(testID uint64, subtest string) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*, AVG(score) AS score").Joins("LEFT JOIN students on students.id = scores.student_id").Where("test_id = ? AND type = ?", testID, subtest).Order("score desc").Find(&scores)
	return scores
}
func (d *database) MaxAll(testID uint64, subtest string) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*, max(score) AS score").Joins("LEFT JOIN students on students.id = scores.student_id").Where("test_id = ? AND type = ?", testID, subtest).Order("score desc").Find(&scores)

	return scores
}
func (d *database) Leaderboard(testID int) []entity.ScoreReal {
	var scores []entity.ScoreReal
	d.conn.Table("scores").Select("id,test_id,student_id,AVG(score) as score").Where("test_id=?", testID).Group("student_id").Order("AVG(score) desc").Limit(3).Find(&scores)
	return scores
}
func (d *database) TigaBawah(testID int) []entity.ScoreReal {
	var scores []entity.ScoreReal
	d.conn.Table("scores").Select("id,test_id,student_id,AVG(score) as score").Where("test_id=? AND score !=?", testID, 0).Group("student_id").Order("AVG(score)").Limit(3).Find(&scores)
	return scores
}
func (d *database) ScoreAnak(testID int, studentID int) []entity.ScoreReal {
	var scores []entity.ScoreReal
	d.conn.Table("scores").Select("id,test_id,student_id,AVG(score) as score").Where("test_id=? AND student_id=?", testID, studentID).Find(&scores)
	return scores
}

func (d *database) GroupScoreByTestID(testID int) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*, AVG(score) as score").Where("test_id = ?", testID).Group("student_id").Order("score desc").Find(&scores)
	return scores
}

func (d *database) GroupScoreByTestIDWithType(testID int, subtest string) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	d.conn.Model(&entity.Score{}).Select("*, AVG(score) as score").Where("test_id = ? AND type = ?", testID, subtest).Group("student_id").Order("score desc").Find(&scores)
	return scores
}

func (d *database) GroupScoreByTestIDAndPref(testID int, uniID int, prodiID int) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	d.conn.Table("scores").Select("scores.*, AVG(scores.score) as score").Joins("LEFT JOIN test_auths ON ( scores.student_id = test_auths.student_id AND scores.test_id = test_auths.test_id )").Where("scores.test_id = ? AND ((test_auths.pref1_uni = ? AND test_auths.pref1_prodi = ?) OR (test_auths.pref2_uni = ? AND test_auths.pref2_prodi = ?))", testID, uniID, prodiID, uniID, prodiID).Group("scores.student_id").Order("scores.score desc").Find(&scores)

	return scores
}

func (d *database) FindScoreByTestIDAndType(testID int, subtest string) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	if subtest == "all" {
		d.conn.Model(&entity.Score{}).Select("*, AVG(score) as score").Where("test_id = ?", testID).Group("student_id").Order("score desc").Find(&scores)
	} else {
		d.conn.Model(&entity.Score{}).Where("test_id = ? AND type = ?", testID, subtest).Order("score desc").Find(&scores)
	}
	return scores
}
func (d *database) Quartile(testID int, subtest string) []entity.ScoreOnly {
	var scores []entity.ScoreOnly
	if subtest == "all" {
		d.conn.Model(&entity.Score{}).Select("*, AVG(score) as score").Where("test_id = ?", testID).Group("student_id").Order("score desc").Find(&scores)
	} else {
		d.conn.Model(&entity.Score{}).Where("test_id = ? AND type = ? AND score IS NOT NULL AND score != ?", testID, subtest,0).Order("score desc").Find(&scores)
	}
	return scores
}
func (d *database) Distribution(testID int, subtest string) (int64, int64, int64, int64, int64, int64, int64, int64, int64, int64) {
	var scores []entity.Score
	result1 := d.conn.Table("scores").Where("test_id=? AND type=? AND score<? AND score!=?", testID, subtest, 100,0).Find(&scores)
	result2 := d.conn.Table("scores").Where("test_id=? AND type=? AND score BETWEEN ? AND ?", testID, subtest, 101, 200).Find(&scores)
	result3 := d.conn.Table("scores").Where("test_id=? AND type=? AND score BETWEEN ? AND ?", testID, subtest, 201, 300).Find(&scores)
	result4 := d.conn.Table("scores").Where("test_id=? AND type=? AND score BETWEEN ? AND ?", testID, subtest, 301, 400).Find(&scores)
	result5 := d.conn.Table("scores").Where("test_id=? AND type=? AND score BETWEEN ? AND ?", testID, subtest, 401, 500).Find(&scores)
	result6 := d.conn.Table("scores").Where("test_id=? AND type=? AND score BETWEEN ? AND ?", testID, subtest, 501, 600).Find(&scores)
	result7 := d.conn.Table("scores").Where("test_id=? AND type=? AND score BETWEEN ? AND ?", testID, subtest, 601, 700).Find(&scores)
	result8 := d.conn.Table("scores").Where("test_id=? AND type=? AND score BETWEEN ? AND ?", testID, subtest, 701, 800).Find(&scores)
	result9 := d.conn.Table("scores").Where("test_id=? AND type=? AND score BETWEEN ? AND ?", testID, subtest, 801, 900).Find(&scores)
	result10 := d.conn.Table("scores").Where("test_id=? AND type=? AND score>?", testID, subtest, 900).Find(&scores)
	leng1 := result1.RowsAffected
	leng2 := result2.RowsAffected
	leng3 := result3.RowsAffected
	leng4 := result4.RowsAffected
	leng5 := result5.RowsAffected
	leng6 := result6.RowsAffected
	leng7 := result7.RowsAffected
	leng8 := result8.RowsAffected
	leng9 := result9.RowsAffected
	leng10 := result10.RowsAffected
	return leng1, leng2, leng3, leng4, leng5, leng6, leng7, leng8, leng9, leng10
}
func (d *database) AVGDistributionTps(testID int) (int64, int64, int64, int64, int64, int64, int64, int64, int64, int64) {
	var scores []entity.Score
	subtest := [5]string{"pu", "ppu", "pk", "pmm", "eng"}
	result1 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?) AND score != ?", testID, subtest,0).Having("AVG(score) < ? AND AVG(score) != ?", 100,0).Find(&scores)
	result2 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 101, 200).Find(&scores)
	result3 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 201, 300).Find(&scores)
	result4 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 301, 400).Find(&scores)
	result5 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 401, 500).Find(&scores)
	result6 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 501, 600).Find(&scores)
	result7 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 601, 700).Find(&scores)
	result8 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 701, 800).Find(&scores)
	result9 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 801, 900).Find(&scores)
	result10 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) >?", 900).Find(&scores)
	leng1 := result1.RowsAffected
	leng2 := result2.RowsAffected
	leng3 := result3.RowsAffected
	leng4 := result4.RowsAffected
	leng5 := result5.RowsAffected
	leng6 := result6.RowsAffected
	leng7 := result7.RowsAffected
	leng8 := result8.RowsAffected
	leng9 := result9.RowsAffected
	leng10 := result10.RowsAffected
	return leng1, leng2, leng3, leng4, leng5, leng6, leng7, leng8, leng9, leng10
}
func (d *database) AVGDistributionSaintek(testID int) (int64, int64, int64, int64, int64, int64, int64, int64, int64, int64) {
	var scores []entity.Score
	subtest := [4]string{"ma", "fi", "ki", "bi"}
	result1 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) < ? AND AVG(score) != ?", 100,0).Find(&scores)
	result2 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 101, 200).Find(&scores)
	result3 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 201, 300).Find(&scores)
	result4 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 301, 400).Find(&scores)
	result5 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 401, 500).Find(&scores)
	result6 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 501, 600).Find(&scores)
	result7 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 601, 700).Find(&scores)
	result8 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 701, 800).Find(&scores)
	result9 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 801, 900).Find(&scores)
	result10 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) >?", 900).Find(&scores)
	leng1 := result1.RowsAffected
	leng2 := result2.RowsAffected
	leng3 := result3.RowsAffected
	leng4 := result4.RowsAffected
	leng5 := result5.RowsAffected
	leng6 := result6.RowsAffected
	leng7 := result7.RowsAffected
	leng8 := result8.RowsAffected
	leng9 := result9.RowsAffected
	leng10 := result10.RowsAffected
	return leng1, leng2, leng3, leng4, leng5, leng6, leng7, leng8, leng9, leng10
}
func (d *database) AVGDistributionSoshum(testID int) (int64, int64, int64, int64, int64, int64, int64, int64, int64, int64) {
	var scores []entity.Score
	subtest := [4]string{"sos", "sej", "geo", "eko"}
	result1 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) < ? AND AVG(score) != ?", 100,0).Find(&scores)
	result2 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 101, 200).Find(&scores)
	result3 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 201, 300).Find(&scores)
	result4 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 301, 400).Find(&scores)
	result5 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 401, 500).Find(&scores)
	result6 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 501, 600).Find(&scores)
	result7 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 601, 700).Find(&scores)
	result8 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 701, 800).Find(&scores)
	result9 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) BETWEEN ? AND ?", 801, 900).Find(&scores)
	result10 := d.conn.Select("AVG(score) as score").Group("student_id").Where("test_id=? AND type IN (?)", testID, subtest).Having("AVG(score) >?", 900).Find(&scores)
	leng1 := result1.RowsAffected
	leng2 := result2.RowsAffected
	leng3 := result3.RowsAffected
	leng4 := result4.RowsAffected
	leng5 := result5.RowsAffected
	leng6 := result6.RowsAffected
	leng7 := result7.RowsAffected
	leng8 := result8.RowsAffected
	leng9 := result9.RowsAffected
	leng10 := result10.RowsAffected
	return leng1, leng2, leng3, leng4, leng5, leng6, leng7, leng8, leng9, leng10
}
func (d *database) BenarSalah(studentID int,testID int,subtest string) entity.BenarSalah {
	var benarsalah entity.BenarSalah
	d.conn.Table("scores").Where("student_id=? AND test_id=? AND type=?",studentID,testID,subtest).Find(&benarsalah)
	return benarsalah
}

////////////////
// BASIC CRUD //
////////////////

func (d *database) FindAllScore() []entity.Score {
	var scores []entity.Score
	d.conn.Find(&scores)
	return scores
}

func (d *database) FindScoreByID(id int) entity.Score {
	var score entity.Score
	d.conn.First(&score, id)
	return score
}

func (d *database) CreateScore(score entity.Score) entity.Score {
	d.conn.Create(score)
	return score
}

func (d *database) UpdateScore(score entity.Score) {
	d.conn.Save(score)
}

func (d *database) DeleteScore(score entity.Score) {
	d.conn.Delete(score)
}
func (d *database) MaxScore(testid int) []entity.Score{
	var score []entity.Score
	d.conn.Table("scores").Select("AVG(score) as score").Where("test_id=?",testid).Group("student_id").Order("score desc").Limit(1).Find(&score)
	return score
}
func (d *database) MaxScoreSchool(testid int,school string) []entity.Score{
	var score []entity.Score
	d.conn.Table("scores").Select("AVG(score) as score").Joins("LEFT JOIN students on students.id = scores.student_id").Where("test_id=? AND school=?",testid,school).Group("student_id").Order("score desc").Limit(1).Find(&score)
	return score
}

func (d *database) FindScoreByStudentsID(id int) []entity.Score {
	var scores []entity.Score
	d.conn.Where("students_id = ?", id).Find(&scores)
	return scores
}

func (d *database) FindScoreByTestID(id int) []entity.Score {
	var scores []entity.Score
	d.conn.Where("test_id = ?", id).Find(&scores)
	return scores
}

func (d *database) StudentScoreOfTest(studentsID int, testID int) []entity.Score {
	var scores []entity.Score
	d.conn.Where("students_id = ? AND test_id = ?", studentsID, testID).Find(&scores)
	return scores
}

func (d *database) AnswerQuestion(TestID int, StudentID int, Type string, Mark string, Correct bool) {
	d.conn.Model(&entity.Score{}).Where("test_id = ? AND student_id = ? AND type = ?", TestID, StudentID, Type).Update(Mark, Correct)
}

/////////////////////
// MINIMIZED READS //
/////////////////////

func (d *database) FindScoreMinByStudentsID(id int) []entity.ScoreMin {
	var scores []entity.ScoreMin
	d.conn.Model(&entity.Score{}).Where("students_id = ?", id).Find(&scores)
	return scores
}

func (d *database) FindScoreMinByStudentsIDAndTestID(studentsID int, testID int) []entity.ScoreMin {
	var scores []entity.ScoreMin
	d.conn.Model(&entity.Score{}).Where("student_id = ? AND test_id = ?", studentsID, testID).Find(&scores)
	return scores
}

func (d *database) FindAllScoreMin() []entity.ScoreMin {
	var scores []entity.ScoreMin
	d.conn.Model(&entity.Score{}).Find(&scores)
	return scores
}

func (d *database) FindScoreMinByID(id int) entity.ScoreMin {
	var score entity.ScoreMin
	d.conn.Model(&entity.Score{}).First(&score, id)
	return score
}
func (d *database) Graphic(test entity.Test, score entity.Score) []entity.IkutBattle {
	var yuhu []entity.IkutBattle
	d.conn.Table("scores").Select("score").Where("student_id=? AND type IN ?", 2, []string{"pu", "ppu", "pk", "pmm", "eng"}).Find(&score)
	d.conn.Table("tests").Select("name", "types").Find(&yuhu)
	return yuhu
}

func (d *database) Jawab(jawab entity.Jawaban) entity.Jawaban{
	d.conn.Exec("UPDATE scores SET x"+jawab.Urutan+" = ? WHERE student_id = ? AND test_id = ? AND type = ?",jawab.Jawab,jawab.UserID,jawab.TestID,jawab.Type)
	return jawab
}
