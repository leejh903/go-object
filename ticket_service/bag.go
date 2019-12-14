package ticket_service

type Bag struct {
	amount     uint64
	invitation *Invitation
	ticket     *Ticket
}

func NewBag(amount uint64) Bag {
	return Bag{
		amount: amount,
	}
}

func NewBagWithInvitation(invitation Invitation, amount uint64) Bag {
	return Bag{
		invitation: &invitation,
		amount:     amount,
	}
}

func (b Bag) Hold(ticket Ticket) uint64 {
	if b.hasInvitation() {
		b.setTicket(&ticket)
		return 0
	} else {
		b.setTicket(&ticket)
		b.minusAmount(ticket.GetFee())
		return ticket.GetFee()
	}
}

func (b Bag) hasInvitation() bool {
	return b.invitation != nil
}

func (b Bag) setTicket(ticket *Ticket) {
	b.ticket = ticket
}

func (b Bag) minusAmount(amount uint64) {
	b.amount -= amount
}

