package repository

import (
	"math"
	"strings"
	"time"
	"villa-akmali/api/connection"
	"villa-akmali/api/model"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)


var timezone string = viper.GetString("timezone")

type OrderRepo interface {
	CreateOrder(order model.Order) model.Order
	FindAll(pagination model.Pagination) []model.Orderlist
	Update(id int ,order model.Order)
	Yearly()([]model.Orderlist,[]model.Orderlist2)
	Monthly()([]model.Orderlist,[]model.Orderlist2)
	Weekly()([]model.Orderlist,[]model.Orderlist2)
	GraphYear(year int) (model.Orderlist,model.Orderlist2)
	GraphMonth(year int,month int) (model.Orderlist,model.Orderlist2)
	GraphWeekly(year int,week int) (model.Orderlist,model.Orderlist2)
	SumAll()(model.Orderlist,model.Orderlist2)
	FindByID(id int) model.Orderlist
	Halaman(pagination model.Pagination) float64
	FilterQuery(filter model.QueryFilter) []model.Orderlist
}

type database struct {
	connection *gorm.DB
}

func NewOrderRepo() OrderRepo {
	db := connection.Create()
	db.AutoMigrate(&model.Order{})
	return &database{connection:db}
}

func (db *database) CreateOrder(order model.Order) model.Order {
	var orderkosong model.Order
	loc,_ := time.LoadLocation(timezone)
	order.CreatedAt = time.Now().In(loc)
	buatorder := db.connection.Create(&order)
	if buatorder.Error != nil {
		return orderkosong
	} else {
		return order
	}
}
func Round2(val float64) (float64,float64){
	intpart, div := math.Modf(val)
	return intpart,div
}
func (db *database) Halaman(pagination model.Pagination) float64 {
	var order []model.Orderlist
	db.connection.Table("orders").Find(&order)
	page := float64(len(order))/float64(pagination.Limit)
	depan,belakang := Round2(page)
	if belakang != 0{
		depan = depan +1
	} else if depan == 0{
		depan = depan +1
	}
	return depan
}
func (db *database) FindAll(pagination model.Pagination) []model.Orderlist {
	var order []model.Orderlist
	angka := db.Halaman(pagination)
	if pagination.Page > int(angka){
		pagination.Page = int(angka)
	}
	offset := (pagination.Page -1) * pagination.Limit
	db.connection.Table("orders").Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&order)
	return order
}

func (db *database) FindByID(id int) model.Orderlist{
	var order model.Orderlist
	data := db.connection.Table("orders").Where("id=?",id).Find(&order)
	if data.Error != nil{
		return order
	}else{
		return order
	}
}
func (db *database) Update(id int ,order model.Order){
	image :=order.Image
	jumlah := order.Jumlah
	catid :=order.CategoryID
	keterangan := order.Keterangan
	types := order.Type
	db.connection.Table("orders").Where("id = ?",id).Updates(model.Order{Keterangan: keterangan,CategoryID: catid,Jumlah:jumlah,Image:image,Type:types})
}
func (db *database) SumAll()(model.Orderlist,model.Orderlist2) {
	var order model.Orderlist
	var order2 model.Orderlist2
	db.connection.Table("orders").Where("type = ?","input").Select("sum(jumlah) as jumlah").Find(&order)
	db.connection.Table("orders").Where("type = ?","output").Select("sum(jumlah) as jumlah").Find(&order2)
	return order,order2
}

func (db *database) Yearly()([]model.Orderlist,[]model.Orderlist2) {
	var order []model.Orderlist
	var order2 []model.Orderlist2
	loc,_ := time.LoadLocation(timezone)
	year := time.Now().In(loc).Year()
	db.connection.Table("orders").Where("type = ? AND EXTRACT(YEAR FROM created_at) BETWEEN ? AND ?","input",year,year+1).Find(&order)
	db.connection.Table("orders").Where("type = ? AND EXTRACT(YEAR FROM created_at) BETWEEN ? AND ?","output",year,year+1).Find(&order2)
	return order,order2
}

func (db *database) Monthly()([]model.Orderlist,[]model.Orderlist2) {
	var order []model.Orderlist
	var order2 []model.Orderlist2
	loc,_ := time.LoadLocation(timezone)
	month := int(time.Now().In(loc).Month())
	db.connection.Table("orders").Where("type = ? AND EXTRACT(MONTH FROM created_at) BETWEEN ? AND ?","input",month,month+1).Find(&order)
	db.connection.Table("orders").Where("type = ? AND EXTRACT(MONTH FROM created_at) BETWEEN ? AND ?","output",month,month+1).Find(&order2)
	return order,order2
}

func (db *database) Weekly()([]model.Orderlist,[]model.Orderlist2) {
	var order []model.Orderlist
	var order2 []model.Orderlist2
	loc,_ := time.LoadLocation(timezone)
	_,week := time.Now().In(loc).ISOWeek()
	db.connection.Table("orders").Where("type = ? AND EXTRACT(WEEK FROM created_at) BETWEEN ? AND ?","input",week,week+1).Find(&order)
	db.connection.Table("orders").Where("type = ? AND EXTRACT(WEEK FROM created_at) BETWEEN ? AND ?","output",week,week+1).Find(&order2)
	return order,order2
}

func (db *database) GraphYear(year int) (model.Orderlist,model.Orderlist2) {
	var order model.Orderlist
	var order2 model.Orderlist2
	db.connection.Table("orders").Where("type = ? AND EXTRACT(YEAR FROM created_at) BETWEEN ? AND ?","input",year,year+1).Select("created_at,sum(jumlah) as jumlah").Find(&order)
	db.connection.Table("orders").Where("type = ? AND EXTRACT(YEAR FROM created_at) BETWEEN ? AND ?","output",year,year+1).Select("created_at,sum(jumlah) as jumlah").Find(&order2)
	return order,order2
}

func (db *database) GraphMonth(year int,month int) (model.Orderlist,model.Orderlist2) {
	var arr1 model.Orderlist
	var arr2 model.Orderlist2
	db.connection.Table("orders").Where("type = ? AND EXTRACT(YEAR FROM created_at) BETWEEN ? AND ? AND EXTRACT(MONTH FROM created_at) = ?","input",year,year+1,month).Group("EXTRACT(MONTH FROM created_at)").Select("created_at,sum(jumlah) as jumlah").Find(&arr1)
	db.connection.Table("orders").Where("type = ? AND EXTRACT(YEAR FROM created_at) BETWEEN ? AND ? AND EXTRACT(MONTH FROM created_at) = ?","output",year,year+1,month).Group("EXTRACT(MONTH FROM created_at)").Select("created_at,sum(jumlah) as jumlah").Find(&arr2)
	return arr1,arr2
}

func (db *database) GraphWeekly(year int,week int) (model.Orderlist,model.Orderlist2) {
	var arr1 model.Orderlist
	var arr2 model.Orderlist2
	db.connection.Table("orders").Where("type = ? AND EXTRACT(YEAR FROM created_at) BETWEEN ? AND ? AND WEEKOFYEAR(created_at) = ?","input",year,year+1,week).Group("WEEKOFYEAR(created_at)").Select("created_at,sum(jumlah) as jumlah").Find(&arr1)
	db.connection.Table("orders").Where("type = ? AND EXTRACT(YEAR FROM created_at) BETWEEN ? AND ? AND WEEKOFYEAR(created_at) = ?","output",year,year+1,week).Group("WEEKOFYEAR(created_at)").Select("created_at,sum(jumlah) as jumlah").Find(&arr2)
	return arr1,arr2
}

func (db *database) FilterQuery(filter model.QueryFilter) []model.Orderlist {
	var array []model.Orderlist
	str := filter.StartedAt
	str2 := filter.EndsAt
	if len(strings.TrimSpace(str)) == 0 || len(strings.TrimSpace(str2)) == 0{

	}else{
		split := strings.Split(str,"/")
		split2 := strings.Split(str2,"/")
		mulai := split[2] + "-"+split[1] + "-"+split[0]
		selesai := split2[2] + "-"+split2[1] + "-"+split2[0]
		if len(strings.TrimSpace(filter.Category)) == 0{
			db.connection.Table("orders").Where("DATE(created_at) BETWEEN ? AND ?",mulai,selesai).Order("id desc").Find(&array)
		} else if filter.Category == "input"{
			db.connection.Table("orders").Where("type = ? AND DATE(created_at) BETWEEN ? AND ?","input",mulai,selesai).Order("id desc").Find(&array)
		}else{
			db.connection.Table("orders").Where("type = ? AND DATE(created_at) BETWEEN ? AND ?","output",mulai,selesai).Order("id desc").Find(&array)
		}
	}
	return array
}