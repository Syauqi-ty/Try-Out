package repository

import (
	connection "studybuddy-backend-fast/api/connection"
	entity "studybuddy-backend-fast/api/entity"

	"gorm.io/gorm"
)

type PrefRepo interface {
	FindAllUni() []entity.Uni
	FindAllUniMin() []entity.UniMin
	FindUniByID(id int) entity.Uni
	FindUniMinByID(id int) entity.UniMin
	FindUniByUniID(id int) entity.Uni
	FindUniMinByUniID(id int) entity.UniMin

	FindProdiOfUni(ProdiID int, UniID int) entity.Prodi
}

type database struct {
	conn *gorm.DB
}

func NewPrefRepo() PrefRepo {
	db := connection.Create()
	db.AutoMigrate(&entity.Prodi{})
	db.AutoMigrate(&entity.Uni{})
	return &database{conn: db}
}

func (d *database) FindAllUni() []entity.Uni {
	var unis []entity.Uni
	d.conn.Preload("Sci").Preload("Soc").Find(&unis)
	return unis
}

func (d *database) FindAllUniMin() []entity.UniMin {
	var unis []entity.UniMin
	d.conn.Model(&entity.Uni{}).Find(&unis)
	return unis
}

func (d *database) FindUniByID(id int) entity.Uni {
	var uni entity.Uni
	d.conn.Preload("Sci", "type = ?", "sci").Preload("Soc", "type = ?", "soc").First(&uni, id)
	return uni
}

func (d *database) FindUniByUniID(id int) entity.Uni {
	var uni entity.Uni
	d.conn.Preload("Sci", "type = ?", "sci").Preload("Soc", "type = ?", "soc").Where("uni_id = ?", id).First(&uni)
	return uni
}

func (d *database) FindUniMinByID(id int) entity.UniMin {
	var uni entity.UniMin
	d.conn.Model(&entity.Uni{}).First(&uni, id)
	return uni
}

func (d *database) FindUniMinByUniID(id int) entity.UniMin {
	var uni entity.UniMin
	d.conn.Model(&entity.Uni{}).Where("uni_id = ?", id).First(&uni)
	return uni
}

func (d *database) FindProdiOfUni(ProdiID int, UniID int) entity.Prodi {
	var prodi entity.Prodi
	d.conn.Where("prodi_id = ? AND uni_id = ?", ProdiID, UniID).First(&prodi)
	return prodi
}
