package main

import "fmt"

type Notifier interface {
	Notify(message string) string
}

type Email struct{ Address string }
type SMS struct{ Phone string }
type Push struct{ DeviceID string }

func (n Email) Notify(message string) string {
	return fmt.Sprintf("Email on address %s: %s", n.Address, message)
}
func (n SMS) Notify(message string) string {
	return fmt.Sprintf("SMS on phone %s: %s", n.Phone, message)
}
func (n Push) Notify(message string) string {
	return fmt.Sprintf("Push on device %s: %s", n.DeviceID, message)
}

func sendAll(notifiers []Notifier, msg string) {
	for _, n := range notifiers {
		fmt.Println(n.Notify(msg))
	}
}

func main() {
	n := []Notifier{Email{Address: "example@mail.com"}, SMS{Phone: "INotPhone"}, Push{DeviceID: "pc-123"}}
	sendAll(n, "Spam:)")
}
