package TelegramBot

import "context"

type SendJoke interface {
	sendJoke(userId int, joke string) error
}
type AnekdotProvider interface {
	GetAnekdot(ctx context.Context, category string) (string, error)
}
