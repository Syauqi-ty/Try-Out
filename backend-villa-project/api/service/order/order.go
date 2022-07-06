package service

import (
	"strconv"
	"time"
	"villa-akmali/api/model"
	repocategory "villa-akmali/api/repository/category"
	repository "villa-akmali/api/repository/order"
	repouser "villa-akmali/api/repository/user"
	repoyear "villa-akmali/api/repository/year"

	"github.com/spf13/viper"
)

var timezone string = viper.GetString("timezone")

type OrderService interface {
	CreateOrder(order model.Order) model.Order
	FindAll(pagination model.Pagination) []model.Orderlist
	Update(id int,order model.Order)
	Yearly() model.Query
	Monthly() model.Query
	AllOrder(pagination model.Pagination)([]model.AllOrder,int)
	Weekly() model.Query
	GraphYear()[]model.Query
	GraphMonth()[]model.Query
	GraphWeek()[]model.Query
	SumAll() model.QuerySum
	FindById(id int) model.ById
	Filter(filter model.QueryFilter) []model.AllOrder
}

type orderService struct {
	orderrepo repository.OrderRepo
	userrepo repouser.UserRepo
	yearrepo repoyear.YearRepo
	categoryrepo repocategory.CategoryRepo
}

func NewOrderService(repo repository.OrderRepo,urepo repouser.UserRepo,yrepo repoyear.YearRepo,crepo repocategory.CategoryRepo) OrderService  {
	return &orderService{
		orderrepo: repo,
		userrepo: urepo,
		yearrepo: yrepo,
		categoryrepo: crepo,
	}
}
func (service *orderService) CreateOrder(order model.Order) model.Order{
	data := service.orderrepo.CreateOrder(order)
	var year model.Year
	year.Year = data.CreatedAt.Year()
	service.yearrepo.CreateYear(year)
	return data
}

func (service *orderService) makearray(i int ,j int) []int{
	var week []int
	for i:=i; i<=j; i++ {
		week = append(week, i)
	}
	return week
}

func (service *orderService) FindAll(pagination model.Pagination) []model.Orderlist {
	return service.orderrepo.FindAll(pagination)
}

func (service *orderService) Update(id int,order model.Order)  {
	service.orderrepo.Update(id,order)
}
func (service *orderService) FindById(id int) model.ById {
	var order model.ById
	data := service.orderrepo.FindByID(id)
	uname := service.userrepo.FindByID(data.UserID)
	category := service.categoryrepo.FindByID(data.CategoryID)
	order.ID = data.ID
	order.CreatedAt = data.CreatedAt
	order.Image = data.Image
	if data.Type == "input"{
		order.Type = "Income"
	} else if data.Type == "output"{
		order.Type = "Expense"
	}
	order.Judul = data.Judul
	order.Jumlah = data.Jumlah
	order.Keterangan = data.Keterangan
	order.CategoryName = category.Name
	order.Username = uname.Name
	return order
}
func (service *orderService) SumAll() model.QuerySum {
	var sum model.QuerySum
	input,output := service.orderrepo.SumAll()
	sum.Input = input.Jumlah
	sum.Output = output.Jumlah
	return sum
}

func (service *orderService) Filter(filter model.QueryFilter) []model.AllOrder  {
	var order model.AllOrder
	var arrorder []model.AllOrder
	data := service.orderrepo.FilterQuery(filter)
	if len(data) == 0{
		order.ID = 1
		order.Judul = ""
		order.Username = ""
		order.CategoryName = ""
		order.Jumlah = 0
		arrorder = append(arrorder,order)
	}else{
		for i := range data{
		category := service.categoryrepo.FindByID(data[i].CategoryID)
			user := service.userrepo.FindByID(data[i].UserID)
			order.ID = data[i].ID
			order.CategoryName = category.Name
			order.Username = user.Name
			order.Judul = data[i].Judul
			order.Jumlah = data[i].Jumlah
			arrorder = append(arrorder, order)
		}
	}
	return arrorder
}

func (service *orderService) Yearly() model.Query{
	var query model.Query
	var suminput float64
	var sumoutput float64
	input,output := service.orderrepo.Yearly()
	for i := range input{
		suminput += input[i].Jumlah
		query.Input = suminput
	}
	for t := range output{
		sumoutput += output[t].Jumlah
		query.Output = sumoutput
		query.Waktu = "Tahun " + strconv.Itoa(output[t].CreatedAt.Year())
	}
	return query
}
func (service *orderService) Monthly() model.Query{
	var query model.Query
	var suminput float64
	var sumoutput float64
	input,output := service.orderrepo.Monthly()
	for i := range input{
		suminput += input[i].Jumlah
		query.Input = suminput
	}
	for t := range output{
		sumoutput += output[t].Jumlah
		query.Output = sumoutput
		query.Waktu = "Bulan " + output[t].CreatedAt.Month().String()
	}
	return query
}
func (service *orderService) Weekly() model.Query{
	var query model.Query
	var suminput float64
	var sumoutput float64
	input,output := service.orderrepo.Weekly()
	for i := range input{
		suminput += input[i].Jumlah
		query.Input = suminput
	}
	for t := range output{
		sumoutput += output[t].Jumlah
		query.Output = sumoutput
	}
	query.Waktu = "Minggu ini"
	return query
}

func (service *orderService) AllOrder(pagination model.Pagination) ([]model.AllOrder,int) {
	var order []model.AllOrder
	var modelorder model.AllOrder
	angka := service.orderrepo.Halaman(pagination)
	page := int(angka)
	data := service.FindAll(pagination)
	if len(data) == 0 {
		modelorder.ID = 1
		modelorder.CategoryName = ""
		modelorder.Username = ""
		modelorder.Judul = ""
		modelorder.Jumlah = 0
		order = append(order, modelorder)
	}else{
		for i := range data {
			category := service.categoryrepo.FindByID(data[i].CategoryID)
			user := service.userrepo.FindByID(data[i].UserID)
			modelorder.ID = data[i].ID
			modelorder.CategoryName = category.Name
			modelorder.Username = user.Name
			modelorder.Judul = data[i].Judul
			modelorder.Jumlah = data[i].Jumlah
			order = append(order, modelorder)
		}
	}
	return order,page
}

func (service *orderService) GraphYear()[]model.Query  {
	var query model.Query
	var arrq []model.Query
	tahun := service.yearrepo.FindAll()
	for i := range tahun {
		input,output := service.orderrepo.GraphYear(tahun[i].Year)
		query.Output = output.Jumlah
		query.Input = input.Jumlah
		query.Waktu = strconv.Itoa(tahun[i].Year)
		arrq = append(arrq, query)
	}
	return arrq
}
func (service *orderService) GraphMonth()[]model.Query{
	arraymonth := service.makearray(1,12)
	var query model.Query
	var arrq []model.Query
	loc,_ := time.LoadLocation(timezone)
	data := time.Now().In(loc).Year()
		for z := range arraymonth{
			input,output:= service.orderrepo.GraphMonth(data,arraymonth[z])
			query.Output = output.Jumlah
			query.Input = input.Jumlah
			if input.CreatedAt.Year() == 1{
				query.Waktu = output.CreatedAt.Month().String()
			} else{
				query.Waktu = input.CreatedAt.Month().String()
			}
			if query.Output ==0&& query.Input == 0{
			}else{
				arrq = append(arrq, query)
			}
			}
	result := arrq
	return result
}

func (service *orderService) GraphWeek()[]model.Query{
	var query model.Query
	var arrq []model.Query
	loc,_ := time.LoadLocation(timezone)
	data := time.Now().In(loc).Year()
	array := service.makearray(1,53)
		for z := range array{
			input,output:= service.orderrepo.GraphWeekly(data,array[z])
			query.Output = output.Jumlah
			query.Input = input.Jumlah
			if input.CreatedAt.Year() == 1{
				_,week := output.CreatedAt.ISOWeek()
				query.Waktu =  "Week-"+strconv.Itoa(week)
			} else{
				_,week := input.CreatedAt.ISOWeek()
				query.Waktu = "Week-"+strconv.Itoa(week)
			}
			if query.Output ==0&& query.Input == 0{
			}else{
				arrq = append(arrq, query)
			}
			}
	result := arrq
	return result
}