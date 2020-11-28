package checkers

type Move struct {
	From int8
	To int8
	CapturedPieces []Piece
}
