package abstractions

type IEmailService interface {
	Send(email string, subject string, body string) error
}
