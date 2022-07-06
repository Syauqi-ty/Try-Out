package controller

import (
	"fmt"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	entity "studybuddy-backend-fast/api/entity"
	service "studybuddy-backend-fast/api/services/payment"
)

type PaymentController interface {
	FindAllPayment(c *gin.Context)
	RequestEwalletPayment(c *gin.Context)
	OVOCallback(c *gin.Context)
	DanaCallback(c *gin.Context)
	LinkAjaCallback(c *gin.Context)
}

type paymentController struct {
	service service.PaymentService
}

func NewPaymentController(service service.PaymentService) PaymentController {
	return &paymentController{service}
}

func (s *paymentController) FindAllPayment(c *gin.Context) {
	res := s.service.FindAllPayment()
	c.JSON(200, gin.H{"msg": "Query successful", "data": res})
}

func (s *paymentController) RequestEwalletPayment(c *gin.Context) {
	var req entity.PaymentRequest
	claims := jwt.ExtractClaims(c)

	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{"msg": "Bad request, nothing requested"})
	} else {
		req.StudentID = uint64(claims["user_id"].(float64))
		switch req.Method {
		case "ovo":
			if res, err := s.service.RequestOVOPayment(req); err != nil {
				c.JSON(500, gin.H{"msg": "Something went wrong"})
			} else {
				c.JSON(200, gin.H{"msg": "Request sent", "data": res})
			}
		case "dana":
			if res, err := s.service.RequestDanaPayment(req); err != nil {
				c.JSON(500, gin.H{"msg": "Something went wrong"})
			} else {
				c.JSON(200, gin.H{"msg": "Request sent", "data": res})
			}
		case "linkaja":
			if res, err := s.service.RequestLinkajaPayment(req); err != nil {
				c.JSON(500, gin.H{"msg": "Something went wrong"})
			} else {
				c.JSON(200, gin.H{"msg": "Request sent", "data": res})
			}
		}
	}
}

func (s *paymentController) OVOCallback(c *gin.Context) {
	var header entity.CallbackHeader
	var payload entity.OVOCallback

	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(400, gin.H{"msg": "Invalid Source"})
	} else {
		if err := c.ShouldBind(&payload); err != nil {
			c.JSON(400, gin.H{"msg": "Invalid body"})
		} else {
			s.service.OVOCallback(payload)
			c.JSON(200, gin.H{"msg": "Callback recieved"})
		}
	}
}

func (s *paymentController) DanaCallback(c *gin.Context) {
	var header entity.CallbackHeader
	var payload entity.DanaCallback

	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(400, gin.H{"msg": "Invalid Source"})
	} else {
		if err := c.ShouldBind(&payload); err != nil {
			c.JSON(400, gin.H{"msg": "Invalid body"})
		} else {
			s.service.DanaCallback(payload)
			c.JSON(200, gin.H{"msg": "Callback recieved"})
		}
	}
}

func (s *paymentController) LinkAjaCallback(c *gin.Context) {
	var header entity.CallbackHeader
	var payload entity.LinkAjaCallback

	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(400, gin.H{"msg": "Invalid Source"})
	} else {
		if err := c.ShouldBind(&payload); err != nil {
			c.JSON(400, gin.H{"msg": "Invalid body"})
		} else {
			s.service.LinkAjaCallback(payload)
			c.JSON(200, gin.H{"msg": "Callback recieved"})
		}
	}
}
