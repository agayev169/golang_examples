package main

type RandomBot struct {
}

func (b RandomBot) Move(g Game) Move {
	pm := g.PossibleMoves()
	if len(pm) == 0 {
		return nil
	}

	return pm[0]
}
