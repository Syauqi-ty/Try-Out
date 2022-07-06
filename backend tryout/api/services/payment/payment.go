package service

import (
	"fmt"
	"github.com/spf13/viper"
	xendit "github.com/xendit/xendit-go"
	ewallet "github.com/xendit/xendit-go/ewallet"
	"strconv"
	"strings"
	entity "studybuddy-backend-fast/api/entity"
	misc "studybuddy-backend-fast/api/helper/misc"
	repository "studybuddy-backend-fast/api/repository/payment"
	testRepository "studybuddy-backend-fast/api/repository/test"
	"time"
)

var timezone string = viper.GetString("timezone")

type PaymentService interface {
	FindAllPayment() []entity.Payment
	FindAllStudentPayment(id int) []entity.Payment
	FindPaymentByID(id int) entity.Payment
	RequestOVOPayment(request entity.PaymentRequest) (*xendit.EWallet, error)
	RequestLinkajaPayment(request entity.PaymentRequest) (*xendit.EWallet, error)
	RequestDanaPayment(request entity.PaymentRequest) (*xendit.EWallet, error)
	OVOCallback(content entity.OVOCallback)
	DanaCallback(content entity.DanaCallback)
	LinkAjaCallback(content entity.LinkAjaCallback)
}

type paymentService struct {
	repo     repository.PaymentRepo
	testRepo testRepository.TestRepo
}

func NewPaymentService(repo repository.PaymentRepo, trep testRepository.TestRepo) PaymentService {
	return &paymentService{repo, trep}
}

func (s *paymentService) RequestOVOPayment(request entity.PaymentRequest) (*xendit.EWallet, error) {
	test := s.testRepo.FindTestByID(int(request.TestID))
	extID := misc.GenerateExternalID(int(request.StudentID), int(request.TestID), request.Method)
	loc, _ := time.LoadLocation(timezone)

	xendit.Opt.SecretKey = viper.GetString("xendit.key")

	data := ewallet.CreatePaymentParams{
		ExternalID:  extID,
		Amount:      float64(test.Price),
		Phone:       request.Phone,
		EWalletType: xendit.EWalletTypeOVO,
		CallbackURL: "https://" + viper.GetString("server.name") + "/api/v2/ovo/callback",
		RedirectURL: "https://" + viper.GetString("server.name") + "/api/v2/ovo/callback",
	}

	res, err := ewallet.CreatePayment(&data)
	if err != nil {
		return &xendit.EWallet{}, err
	}

	s.repo.CreatePayment(entity.Payment{
		ExternalID: extID,
		PayerID:    request.StudentID,
		TestID:     request.TestID,
		Amount:     test.Price,
		Method:     request.Method,
		Status:     "PENDING",
		CreatedAt:  time.Now().In(loc),
		UpdatedAt:  time.Now().In(loc),
	})

	return res, nil
}

func (s *paymentService) RequestDanaPayment(request entity.PaymentRequest) (*xendit.EWallet, error) {
	test := s.testRepo.FindTestByID(int(request.TestID))
	extID := misc.GenerateExternalID(int(request.StudentID), int(request.TestID), request.Method)
	loc, _ := time.LoadLocation(timezone)

	xendit.Opt.SecretKey = viper.GetString("xendit.key")

	data := ewallet.CreatePaymentParams{
		ExternalID:  extID,
		Amount:      float64(test.Price),
		Phone:       request.Phone,
		EWalletType: xendit.EWalletTypeDANA,
		CallbackURL: "https://" + viper.GetString("server.name") + "/api/v2/dana/callback",
		RedirectURL: "https://" + viper.GetString("server.name") + "/market",
	}

	res, err := ewallet.CreatePayment(&data)
	if err != nil {
		return &xendit.EWallet{}, err
	}

	s.repo.CreatePayment(entity.Payment{
		ExternalID: extID,
		PayerID:    request.StudentID,
		TestID:     request.TestID,
		Amount:     test.Price,
		Method:     request.Method,
		Status:     "PENDING",
		CreatedAt:  time.Now().In(loc),
		UpdatedAt:  time.Now().In(loc),
	})

	return res, nil
}

func (s *paymentService) RequestLinkajaPayment(request entity.PaymentRequest) (*xendit.EWallet, error) {
	test := s.testRepo.FindTestByID(int(request.TestID))
	extID := misc.GenerateExternalID(int(request.StudentID), int(request.TestID), request.Method)
	loc, _ := time.LoadLocation(timezone)

	xendit.Opt.SecretKey = viper.GetString("xendit.key")

	data := ewallet.CreatePaymentParams{
		ExternalID:  extID,
		Amount:      float64(test.Price),
		Phone:       request.Phone,
		EWalletType: xendit.EWalletTypeLINKAJA,
		CallbackURL: "https://" + viper.GetString("server.name") + "/api/v2/linkaja/callback",
		RedirectURL: "https://" + viper.GetString("server.name") + "/market",
		Items: []ewallet.Item{
			ewallet.Item{
				ID:       strconv.Itoa(int(request.TestID)),
				Name:     test.Name,
				Price:    float64(test.Price),
				Quantity: 1,
			},
		},
	}

	res, err := ewallet.CreatePayment(&data)
	if err != nil {
		return &xendit.EWallet{}, err
	}
	fmt.Printf("created payment: %+v\n", res)

	s.repo.CreatePayment(entity.Payment{
		ExternalID: extID,
		PayerID:    request.StudentID,
		TestID:     request.TestID,
		Amount:     test.Price,
		Method:     request.Method,
		Status:     "PENDING",
		CreatedAt:  time.Now().In(loc),
		UpdatedAt:  time.Now().In(loc),
	})

	return res, nil
}

func (s *paymentService) OVOCallback(content entity.OVOCallback) {
	var updatedData entity.Payment
	parsed := strings.Split(content.ExternalID, "-")
	slug := strings.Join(parsed[1:len(parsed)-3], "_")
	test := s.testRepo.FindTestBySlug(slug)
	studentID, _ := strconv.Atoi(parsed[len(parsed)-2])

	loc, _ := time.LoadLocation(timezone)
	timenow := time.Now().In(loc)

	if content.FailureCode == "" {
		updatedData = entity.Payment{
			ExternalID:  content.ExternalID,
			Status:      content.Status,
			FailureCode: content.FailureCode,
		}
	} else {
		updatedData = entity.Payment{
			ExternalID: content.ExternalID,
			Status:     content.Status,
			FinishedAt: timenow,
		}

		misc.CreateTestUser(int(test.ID), studentID)
	}
	s.repo.UpdatePaymentByExternalID(content.ExternalID, updatedData)
}

func (s *paymentService) DanaCallback(content entity.DanaCallback) {
	var updatedData entity.Payment

	loc, _ := time.LoadLocation(timezone)
	timenow := time.Now().In(loc)

	if content.Status != "PAID" {
		updatedData = entity.Payment{
			ExternalID: content.ExternalID,
			Status:     content.Status,
		}
	} else {
		updatedData = entity.Payment{
			ExternalID: content.ExternalID,
			Status:     content.Status,
			FinishedAt: timenow,
		}
	}
	s.repo.UpdatePaymentByExternalID(content.ExternalID, updatedData)
}

func (s *paymentService) LinkAjaCallback(content entity.LinkAjaCallback) {
	var updatedData entity.Payment

	loc, _ := time.LoadLocation(timezone)
	timenow := time.Now().In(loc)

	if content.Status != "SUCCESS_COMPLETED" {
		updatedData = entity.Payment{
			ExternalID: content.ExternalID,
			Status:     content.Status,
		}
	} else {
		updatedData = entity.Payment{
			ExternalID: content.ExternalID,
			Status:     content.Status,
			FinishedAt: timenow,
		}
	}
	s.repo.UpdatePaymentByExternalID(content.ExternalID, updatedData)
}

func (s *paymentService) FindAllPayment() []entity.Payment {
	return s.repo.FindAllPayment()
}

func (s *paymentService) FindAllStudentPayment(id int) []entity.Payment {
	return s.repo.FindPaymentOfPayer(id)
}

func (s *paymentService) FindPaymentByID(id int) entity.Payment {
	return s.repo.FindPaymentByID(id)
}
