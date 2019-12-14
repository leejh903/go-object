package ticket_service

type Theater struct {
	ticketSeller *TicketSeller
}

func NewTheater(ticketSeller *TicketSeller) Theater {
	return Theater{ticketSeller}
}

func (t Theater) Enter(audience Audience) {
	t.ticketSeller.SellTo(audience)
}
