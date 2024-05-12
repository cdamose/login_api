package repository

import (
	"context"
)

type CommunicationRepository interface {
	SendSMS(ctx context.Context, phone_number string, message string) error
}
