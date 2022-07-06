package repository

import (
	"pampam/backend/tuqa/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AlamatRepo interface {
	Save(alamat entity.Alamat)
	Update(alamat entity.Alamat)
	Delete(alamat entity.Alamat)
	FindAlamat(alamat entity.Alamat) []entity.Alamat
}

type databaseA struct {
	connection *gorm.DB
}

func NewAlamatRepo() AlamatRepo {
	dsn := "user=postgres password=250330 dbname=pampam port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to konak")
	}
	db.AutoMigrate(&entity.Alamat{})
	return &databaseA{
		connection: db,
	}
}

func (db *databaseA) Save(alamat entity.Alamat) {
	db.connection.Create(&alamat)
}
func (db *databaseA) Update(alamat entity.Alamat) {
	hehe := alamat.Jalan
	hihi := alamat.Kabupaten
	huhu := alamat.Kecamatan
	hoho := alamat.Kelurahan
	haha := alamat.Koordinat
	db.connection.Model(&alamat).Where("merge_id = ?", alamat.Merge_id).Updates(entity.Alamat{Jalan: hehe, Kelurahan: hoho, Kecamatan: huhu, Kabupaten: hihi, Koordinat: haha})
}
func (db *databaseA) Delete(alamat entity.Alamat) {
	db.connection.Model(&alamat).Where("merge_id = ?", alamat.Merge_id).Delete(&alamat)
}
func (db *databaseA) FindAlamat(alamat entity.Alamat) []entity.Alamat {
	var alamatarray []entity.Alamat
	db.connection.Model(&alamat).Where("merge_id = ?", alamat.Merge_id).Find(&alamatarray)
	return alamatarray
}
