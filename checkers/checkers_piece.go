package checkers

import "fmt"

type Piece struct {
	Type int8
	Position int8
}

func pieceToChar(p *Piece, style int8) (string, error) {
	if p == nil {
		if style == XO {
			return " ", nil
		} else if style == Classical {
			return "  ", nil
		} else {
			return "", fmt.Errorf(UndefinedStyle)
		}
	}

	if style == Classical {
		if p.Type == Empty {
			return "  ", nil
		} else if p.Type == BlackMan {
			return "⛂", nil
		} else if p.Type == WhiteMan {
			return "⛀", nil
		} else if p.Type == BlackKing {
			return "⛃", nil
		} else if p.Type == WhiteKing {
			return "⛁", nil
		}
	} else if style == XO {
		if p.Type == Empty {
			return " ", nil
		} else if p.Type == BlackMan {
			return "x", nil
		} else if p.Type == WhiteMan {
			return "o", nil
		} else if p.Type == BlackKing {
			return "X", nil
		} else if p.Type == WhiteKing {
			return "O", nil
		}
	} else {
		return "", fmt.Errorf(UndefinedStyle)
	}

	return "", fmt.Errorf(UndefinedPiece)
}