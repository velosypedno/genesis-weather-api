package services

import (
	"fmt"

	"github.com/velosypedno/genesis-weather-api/internal/models"
)

type DebugEmailService struct{}

func NewDebugEmailService() *DebugEmailService {
	return &DebugEmailService{}
}

func (d *DebugEmailService) SendConfirmationEmail(subscription models.Subscription) error {
	fmt.Printf("Subscription: %s, link: 127.0.0.1:8080/api/confirm/%s", subscription.Email, subscription.Token)
	return nil
}

func (d *DebugEmailService) SendWeatherEmail(subscription models.Subscription, weather models.Weather) error {
	fmt.Println("Email:", subscription.Email)
	fmt.Println("Weather:", weather)
	return nil
}
