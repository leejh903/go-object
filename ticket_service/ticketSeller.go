package ticket_service

type TicketSeller struct {
	ticketOffice *TicketOffice
}

func NewTicketSeller(ticketOffice *TicketOffice) TicketSeller {
	return TicketSeller{ticketOffice}
}

func (t TicketSeller) SellTo(audience Audience) {
	ticket := t.ticketOffice.GetTicket()
	t.ticketOffice.plusAmount(audience.Buy(ticket))
}
