package handler

import (
	"context"
	"github.com/ebar-go/ego/component"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"unknown/api"
	"unknown/internal/domain/application"
	"unknown/internal/domain/service"
)

type WebHandlerSuite struct {
	suite.Suite

	handler *WebHandler
}

func (suite *WebHandlerSuite) SetupTest() {
	err := component.DB().Connect("root:123456@tcp(127.0.0.1:3306)/game?charset=utf8mb4&parseTime=True&loc=Local", &gorm.Config{})
	if err != nil {
		suite.T().Fatal(err.Error())
	}
	application.SetSign([]byte("12345678"))

	suite.handler = NewWebHandler(service.Provider().Web())
}

func (suite *WebHandlerSuite) TestLogin() {
	suite.T().Run("registered", func(t *testing.T) {
		resp, err := suite.handler.Login(context.Background(), &api.WebLoginRequest{Phone: "17628280001"})
		if err != nil {
			t.Fatal(err)
		}

		log.Println(resp)
	})
	suite.T().Run("unregistered", func(t *testing.T) {
		phone := strconv.Itoa(rand.Intn(9999999999) + 10000000000)
		resp, err := suite.handler.Login(context.Background(), &api.WebLoginRequest{Phone: phone})
		if err != nil {
			t.Fatal(err)
		}

		log.Println(resp)
	})
}

func TestWebHandlerSuite(t *testing.T) {
	suite.Run(t, new(WebHandlerSuite))
}
