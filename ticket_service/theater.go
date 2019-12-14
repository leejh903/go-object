package ticket_service

type Theater struct {
	ticketSeller *TicketSeller
}

func NewTheater(ticketSeller *TicketSeller) Theater {
	return Theater{ticketSeller}
}

func (t Theater) Enter(audience Audience) {
	if audience.GetBag().HasInvitation() {
		ticket := t.ticketSeller.GetTicketOffice().GetTicket()
		audience.GetBag().SetTicket(&ticket)
	} else {
		ticket := t.ticketSeller.ticketOffice.GetTicket()
		audience.GetBag().MinusAmount(ticket.GetFee())
		t.ticketSeller.GetTicketOffice().PlusAmount(ticket.GetFee())
		audience.GetBag().SetTicket(&ticket)
	}
}
