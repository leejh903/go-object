package ticket_service

type Ticket struct {
	fee uint64
}

func (t Ticket) GetFee() uint64 {
	return t.fee
}
