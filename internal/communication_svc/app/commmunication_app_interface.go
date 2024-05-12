package app

import (
	"context"
)

type CommunicationApp interface {
	SendSMS(ctx context.Context, phone_number string, message string) (bool, error)
}
