package checkers

import (
	"fmt"
	"testing"
)

func TestNewCheckers(t *testing.T) {
	c := NewCheckers(XO)
	fmt.Println(c.String())
}

func TestFrom64To32(t *testing.T) {
	invalid := []struct{Y, X int8}{{0, 0}, {0, 4}, {6, 2}, {1, 1}, {5, 7}}
	for i := 0; i < len(invalid); i++ {
		pos, err := from64To32(invalid[i].Y, invalid[i].X)
		if err == nil {
			t.Errorf("Should have got UndefinedPosition error, got nil, inputs: (%v, %v), pos: %v", invalid[i].Y,
				invalid[i].X, pos)
		} else if err.Error() != UndefinedPosition {
			t.Errorf("Should have got UndefinedPosition error, got %v, inputs: (%v, %v)", err, invalid[i].Y,
				invalid[i].X)
		}
	}

	valid := []struct{Y, X int8}{{0, 1}, {0, 3}, {1, 2}, {4, 3}, {7, 6}}
	output := []int8{1, 2, 6, 18, 32}
	for i := 0; i < len(invalid); i++ {
		pos, err := from64To32(valid[i].Y, valid[i].X)
		if err != nil {
			t.Errorf("Error: %v, inputs: (%v, %v)", err, valid[i].Y, valid[i].X)
		} else if pos != output[i] {
			t.Errorf("Error: wrong output, inputs: (%v, %v), output: %v, expected: %v", valid[i].Y, valid[i].X,
				pos, output[i])
		}
	}
}

func TestFrom32To64(t *testing.T) {
	positions := []int8{1, 5, 7, 12, 16, 30, 32}
	expected := []struct{Y, X int8}{{0, 1}, {1, 0}, {1, 4}, {2, 7}, {3, 6}, {7, 2}, {7, 6}}
	for i := 0; i < len(positions); i++ {
		y, x := from32To64(positions[i])
		if y != expected[i].Y || x != expected[i].X {
			t.Errorf("Error: wrong output, input: %v, output: (%v, %v), expected: (%v, %v)", positions[i],
				y, x, expected[i].Y, expected[i].X)
		}
	}
}

func TestAreEnemies(t *testing.T) {
	type input struct {
		Type1 int8
		Type2 int8
		Expected bool
	}
	inputs := []input{{WhiteKing, WhiteMan, false}, {WhiteKing, BlackMan, true},
		{WhiteKing, BlackKing, true}, {WhiteMan, WhiteKing, false}, {WhiteMan, BlackMan, true},
		{WhiteMan, WhiteKing, false}, {BlackKing, BlackMan, false}, {BlackKing, WhiteMan, true},
		{BlackKing, WhiteKing, true}, {BlackMan, BlackKing, false}, {BlackMan, WhiteMan, true},
		{BlackMan, WhiteKing, true}, {WhiteKing, WhiteKing, false}, {WhiteMan, WhiteMan, false},
		{BlackKing, BlackKing, false}, {BlackMan, BlackMan, false}}

	for _, in := range inputs {
		out := areEnemies(in.Type1, in.Type2)
		if out != in.Expected {
			t.Errorf("Error: wrong answer, ipnuts: (%v, %v), output: %v, expected: %v", in.Type1, in.Type2, out, in.Expected)
		}
	}
}