package repository

import (
	"pampam/backend/tuqa/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DeviceRepo interface {
	Save(device entity.Device)
	Update(device entity.Device)
	Delete(device entity.Device)
	FindAll() []entity.Device
	Getvalvestatus(device entity.Device) int
	UpdateB(device entity.Device)
	GetBaterai(device entity.Device) string
}

type database struct {
	connection *gorm.DB
}

func NewDeviceRepo() DeviceRepo {
	dsn := "user=postgres password=250330 dbname=pampam port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to konak")
	}
	db.AutoMigrate(&entity.Device{})
	return &database{
		connection: db,
	}
}

func (db *database) Save(device entity.Device) {
	db.connection.Create(&device)
}
func (db *database) Update(device entity.Device) {
	pantat := device.Valve_status
	if pantat == 0 || pantat == 1 {
		db.connection.Model(&device).Where("merge_id = ?", device.Merge_id).Update("valve_status", pantat)
	}
}
func (db *database) Delete(device entity.Device) {
	db.connection.Delete(&device)
}
func (db *database) FindAll() []entity.Device {
	var device []entity.Device
	db.connection.Set("gorm:auto_preload", true).Find(&device)
	return device
}
func (db *database) Getvalvestatus(device entity.Device) int {
	db.connection.Find(&device)
	hehe := device.Valve_status
	return hehe
}
func (db *database) UpdateB(device entity.Device) {
	hehe := device.Indikator_baterai
	db.connection.Model(&device).Where("merge_id = ?", device.Merge_id).Update("indikator_baterai", hehe)
}
func (db *database) GetBaterai(device entity.Device) string {
	db.connection.Model(&device).Where("merge_id = ?", device.Merge_id).Find(&device)
	hehe := device.Indikator_baterai
	return hehe
}
