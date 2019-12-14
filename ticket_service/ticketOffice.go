package ticket_service

type TicketOffice struct {
	amount  uint64
	tickets []Ticket
}

func newTicketOffice(amount uint64, tickets ...Ticket) TicketOffice {
	return TicketOffice{
		amount:  amount,
		tickets: tickets,
	}
}

func (t TicketOffice) GetTicket() Ticket {
	var ticket Ticket
	ticket, t.tickets = t.tickets[0], t.tickets[1:]
	return ticket
}

func (t TicketOffice) MinusAmount(amount uint64) {
	t.amount -= amount
}

func (t TicketOffice) PlusAmount(amount uint64) {
	t.amount += amount
}
