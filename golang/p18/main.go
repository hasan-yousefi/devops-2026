package main

import "fmt"

type Notifier interface {
	send(message string)
}

type EmailNotifier struct {
	email string
}
type SMSNotifier struct {
	phoneNumber uint64
}

func (e EmailNotifier) send(message string) {
	fmt.Printf("Send mail to %s \n Detail: %s\n", e.email, message)
}

func (s SMSNotifier) send(message string) {
	fmt.Printf("Send SMS to %d \n Detail: %s\n", s.phoneNumber, message)
}

func Notify(notifier Notifier, message string) {
	notifier.send(message)
}
func main() {
	email := EmailNotifier{email: "hsn.yousefi158@gmail.com"}
	sms := SMSNotifier{phoneNumber: 989123456789}
	Notify(email, "Hello Email")
	Notify(sms, "Hello SMS")
}
