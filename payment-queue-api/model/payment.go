package model

type ChannelType string

const (
	Card     ChannelType = "card"
	Transfer ChannelType = "transfer"
	Debit    ChannelType = "debit"
)

type Payment struct {
	ID      int
	Amount  float64
	Channel ChannelType
}
