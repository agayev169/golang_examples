package checkers

import "fmt"

const (
	Empty = iota
	BlackMan
	WhiteMan
	BlackKing
	WhiteKing
)

const (
	Black = iota
	White
)

const (
	XO = iota
	Classical
)

const (
	UndefinedStyle    = "undefined style"
	UndefinedPiece    = "undefined piece"
	UndefinedPosition = "undefined position"
	ImpossibleMove    = "impossible move"
)

type Checkers struct {
	Board      [8][8]*Piece
	PrintStyle int8
	Moves      []Move
	Side       int8
}

func NewCheckers(style int8) Checkers {
	c := Checkers{}
	c.Side = Black
	c.PrintStyle = style

	for i := 0; i < 8; i++ {
		for j := 0; j < 4; j++ {
			if i%2 == 0 {
				c.Board[i][j*2] = nil
			} else {
				c.Board[i][j*2+1] = nil
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if i%2 == 0 {
				pos, err := from64To32(int8(i), int8(j*2+1))
				if err != nil {
					panic(UndefinedPosition)
				}
				c.Board[i][j*2+1] = &Piece{BlackMan, pos}
			} else {
				pos, err := from64To32(int8(i), int8(j*2))
				if err != nil {
					panic(UndefinedPosition)
				}
				c.Board[i][j*2] = &Piece{BlackMan, pos}
			}
		}
	}

	for i := 5; i < 8; i++ {
		for j := 0; j < 4; j++ {
			if i%2 == 0 {
				pos, err := from64To32(int8(i), int8(j*2+1))
				if err != nil {
					panic(UndefinedPosition)
				}
				c.Board[i][j*2+1] = &Piece{WhiteMan, pos}
			} else {
				pos, err := from64To32(int8(i), int8(j*2))
				if err != nil {
					panic(UndefinedPosition)
				}
				c.Board[i][j*2] = &Piece{WhiteMan, pos}
			}
		}
	}

	return c
}

func (c Checkers) PossibleMoves() []Move {
	var simpleMoves []Move
	var jumpMoves []Move

	for i := int8(0); i < 8; i++ {
		for j := int8(0); j < 8; j++ {
			currentJumpMoves, currentSimpleMoves := c.internalPossibleMoves(i, j)
			jumpMoves = append(jumpMoves, currentJumpMoves...)
			simpleMoves = append(simpleMoves, currentSimpleMoves...)
		}
	}

	if len(jumpMoves) > 0 {
		return jumpMoves
	}

	return simpleMoves
}

func (c *Checkers) Move(m Move) error {
	if !c.isPossibleMove(m) {
		return fmt.Errorf(ImpossibleMove)
	}

	c.Moves = append(c.Moves, m)

	fromY, fromX := from32To64(m.From)
	toY, toX := from32To64(m.To)

	c.Board[fromY][fromX].Position = m.To
	c.Board[toY][toX] = c.Board[fromY][fromX]
	c.Board[fromY][fromX] = nil

	for _, cp := range m.CapturedPieces {
		y, x := from32To64(cp.Position)
		c.Board[y][x] = nil
	}

	c.Side = (c.Side + 1) % 2

	return nil
}

func (c Checkers) isPossibleMove(m Move) bool {
	pms := c.PossibleMoves()
	for _, pm := range pms {
		if pm.From == m.From && pm.To == m.To {
			return true
		}
	}

	return false
}

func (c Checkers) internalPossibleMoves(y, x int8) (jumpMoves, simpleMoves []Move) {
	piece := c.Board[y][x]
	if piece == nil {
		return
	}

	if ((piece.Type == BlackKing || piece.Type == BlackMan) && c.Side == White) ||
		(piece.Type == WhiteKing || piece.Type == WhiteMan) && c.Side == Black {
		return
	}

	jumpMoves = append(jumpMoves, c.possibleJumpMoves(piece)...)
	if len(jumpMoves) > 0 {
		return
	}

	simpleMoves = append(simpleMoves, c.possibleSimpleMoves(piece)...)

	return
}

func (c Checkers) possibleJumpMoves(piece *Piece) []Move {
	// TODO: Add multiple jumps
	var res []Move

	y, x := from32To64(piece.Position)
	if piece.Type == BlackMan || piece.Type == BlackKing || piece.Type == WhiteKing {
		m1 := c.checkJumpMoveTo(piece, y+2, x-2)
		if m1 {
			res = append(res, c.newJumpMove(y, x, y+2, x-2))
		}

		m2 := c.checkJumpMoveTo(piece, y+2, x+2)
		if m2 {
			res = append(res, c.newJumpMove(y, x, y+2, x+2))
		}
	}

	if piece.Type == WhiteMan || piece.Type == WhiteKing || piece.Type == BlackKing {
		m1 := c.checkJumpMoveTo(piece, y-2, x-2)
		if m1 {
			res = append(res, c.newJumpMove(y, x, y-2, x-2))
		}

		m2 := c.checkJumpMoveTo(piece, y-2, x+2)
		if m2 {
			res = append(res, c.newJumpMove(y, x, y-2, x+2))
		}
	}

	return res
}

func (c Checkers) checkJumpMoveTo(piece *Piece, toY, toX int8) bool {
	y, x := from32To64(piece.Position)
	if isInBoard(toY, toX) {
		capturedPiece := c.Board[(y+toY)/2][(x+toX)/2]
		if capturedPiece != nil &&
			areEnemies(piece.Type, capturedPiece.Type) && c.Board[toY][toX] == nil{
			return true
		}
	}

	return false
}

func (c Checkers) newJumpMove(fromY, fromX, toY, toX int8) Move {
	from, err := from64To32(fromY, fromX)
	if err != nil {
		panic(err)
	}
	to, err := from64To32(toY, toX)
	if err != nil {
		panic(err)
	}

	return Move{
		From:           from,
		To:             to,
		CapturedPieces: []Piece{*c.Board[(fromY+toY)/2][(fromX+toX)/2]},
	}
}

func (c Checkers) possibleSimpleMoves(piece *Piece) []Move {
	var res []Move

	y, x := from32To64(piece.Position)
	if piece.Type == BlackMan || piece.Type == BlackKing || piece.Type == WhiteKing {
		m1 := c.checkSimpleMoveTo(y+1, x-1)
		if m1 {
			res = append(res, newSimpleMove(y, x, y+1, x-1))
		}

		m2 := c.checkSimpleMoveTo(y+1, x+1)
		if m2 {
			res = append(res, newSimpleMove(y, x, y+1, x+1))
		}
	}
	if piece.Type == WhiteMan || piece.Type == WhiteKing || piece.Type == BlackKing {
		m1 := c.checkSimpleMoveTo(y-1, x-1)
		if m1 {
			res = append(res, newSimpleMove(y, x, y-1, x-1))
		}

		m2 := c.checkSimpleMoveTo(y-1, x+1)
		if m2 {
			res = append(res, newSimpleMove(y, x, y-1, x+1))
		}
	}

	return res
}

func newSimpleMove(fromY, fromX, toY, toX int8) Move {
	from, err := from64To32(fromY, fromX)
	if err != nil {
		panic(err)
	}
	to, err := from64To32(toY, toX)
	if err != nil {
		panic(err)
	}

	return Move{
		From:           from,
		To:             to,
		CapturedPieces: nil,
	}
}

func (c Checkers) checkSimpleMoveTo(toY, toX int8) bool {
	if isInBoard(toY, toX) {
		if c.Board[toY][toX] == nil {
			return true
		} else {
			return false
		}
	}

	return false
}

func (c Checkers) String() string {
	out := ""

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			piece, err := pieceToChar(c.Board[i][j], c.PrintStyle)
			if err != nil {
				panic(err)
			}

			out += "|" + piece
		}
		out += "|"
		if i != 7 {
			out += "\n"
		}
	}

	return out
}
