package main

type Move interface {
}

type Game interface {
	PossibleMoves() []Move
	Move(Move) error
	GameOver() bool
	Winner() int32
}

type Bot interface {
	Move(game Game) Move
}
