package checkers

import "fmt"

func from64To32(y, x int8) (int8, error) {
	if (x + y) % 2 == 0 {
		return -1, fmt.Errorf(UndefinedPosition)
	}

	return y * 4 + x / 2 + 1, nil
}

func from32To64(i int8) (y, x int8) {
	y = (i - 1) / 4
	if y % 2 == 0 {
		x = ((i - 1) % 4) * 2 + 1
	} else {
		x = ((i - 1) % 4) * 2
	}
	return
}

func isInBoard(y, x int8) bool {
	return 0 <= y && y <= 7 && 0 <= x && x <= 7
}

func areEnemies(type1, type2 int8) bool {
	if ((type1 == WhiteMan || type1 == WhiteKing) && (type2 == BlackMan || type2 == BlackKing)) ||
		((type1 == BlackMan || type1 == BlackKing) && (type2 == WhiteMan || type2 == WhiteKing)) {
		return true
	}

	return false
}