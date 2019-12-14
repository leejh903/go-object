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
