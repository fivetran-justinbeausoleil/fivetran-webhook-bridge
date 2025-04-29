package sender

// Sender defines a generic interface for sending events to external systems
type Sender interface {
	Send(payload any) error
}
