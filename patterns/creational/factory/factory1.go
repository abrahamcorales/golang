package factory

func Main() {
	sendType := "email"
	processorNotification := NewNotifer(sendType)
	processorNotification.Send("hola")
}

type Notifier interface {
	Send(message string) string
}

type EmailNotifier struct {
}

type SMSNotifier struct {
}

func (e *EmailNotifier) Send(message string) string {
	return "send email"
}

func (e *SMSNotifier) Send(message string) string {
	return "send SMS"
}

func NewNotifer(kind string) Notifier {
	switch kind {
	case "email":
		return &EmailNotifier{}
	case "sms":
		return &SMSNotifier{}
	default:
		return nil

	}

}
