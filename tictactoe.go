type Board struct {
    tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

func (b *Board) put(x, y int, u string) // interface
func (b *Board) get(x, y int) string // interface