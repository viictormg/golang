package main

import "fmt"

// Sms
// Email

// Sistema para enviar notificaciones
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSnotifiaction struct {
}

func (SMSnotifiaction) SendNotification() {
	fmt.Println("Sending notification SMS")
}

func (SMSnotifiaction) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "twilio"
}

// Email notification

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification Email")
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "Amazon"
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func getNotificationFactory(notifycationType string) (INotificationFactory, error) {
	if notifycationType == "SMS" {
		return &SMSnotifiaction{}, nil
	}
	if notifycationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No notification type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactoty, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactoty)
	sendNotification(emailFactory)

	getMethod(smsFactoty)
	getMethod(emailFactory)
}
