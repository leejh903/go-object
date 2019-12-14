package ticket_service

type TicketSeller struct {
	ticketOffice *TicketOffice
}

func NewTicketSeller(ticketOffice *TicketOffice) TicketSeller {
	return TicketSeller{ticketOffice}
}

func (t TicketSeller) GetTicketOffice() TicketOffice {
	return *t.ticketOffice
}
