package UserProvider

type Domain struct {
	TelegramUser GetTelegramUser
}

func New(TelegramUser GetTelegramUser) *Domain {
	return &Domain{
		TelegramUser: TelegramUser,
	}
}
