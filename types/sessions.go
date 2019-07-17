package types

type Session struct {
	ID             int
	SubscriptionID int
	Uploaded       int
	Downloaded     int
	Status         string
	AddedAt        string
	UpdatedAt      string
}
