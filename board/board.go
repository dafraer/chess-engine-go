package board

type Board struct {
	b    [64]int8
	side bool
}

func New() *Board {
	return &Board{}
}

func (b *Board) String() string {
	return "im a board"
}
