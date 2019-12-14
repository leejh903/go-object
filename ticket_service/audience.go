package ticket_service

type Audience struct {
	bag *Bag
}

func NewAudience(bag Bag) Audience {
	return Audience{&bag}
}

func (a Audience) GetBag() Bag {
	return *a.bag
}

// return paid amount of money
func (a Audience) Buy(ticket Ticket) uint64 {
	return a.bag.Hold(ticket)
}
