package day02

// Bag is a group of colored cubes (in a bag) that a Game is meant to be played with
type Bag Hand

// CanDraw checks if a given Hand could be drawn from the bag
func (b *Bag) CanDraw(h *Hand) bool {
	return b.Red >= h.Red &&
		b.Green >= h.Green &&
		b.Blue >= h.Blue
}

// ExpandToContain will grow the bag to match if the hand cannot be drawn from it
func (b *Bag) ExpandToContain(h *Hand) {
	b.Red = max(b.Red, h.Red)
	b.Green = max(b.Green, h.Green)
	b.Blue = max(b.Blue, h.Blue)
}

// Power determines the bag's 'power', the product of its component counts
func (b *Bag) Power() int {
	return b.Red * b.Green * b.Blue
}
