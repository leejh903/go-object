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

func (b Bag) HasInvitation() bool {
	return b.invitation != nil
}

func (b Bag) HasTicket() bool {
	return b.ticket != nil
}

func (b Bag) SetTicket(ticket *Ticket) {
	b.ticket = ticket
}

func (b Bag) MinusAmount(amount uint64) {
	b.amount -= amount
}

func (b Bag) PlusAmount(amount uint64) {
	b.amount += amount
}
