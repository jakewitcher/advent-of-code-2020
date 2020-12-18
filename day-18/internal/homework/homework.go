package homework

import (
	"strconv"
	"unicode"
)

const (
	ADD          = '+'
	MULTIPLY     = '*'
	OPEN_PARENS  = '('
	CLOSE_PARENS = ')'
)

type Operand interface {
	Evaluate() int
}

type Value struct {
	Value int
}

func (v *Value) Evaluate() int {
	return v.Value
}

type Expression struct {
	LeftOperand, RightOperand Operand
	Operator                  rune
}

func (e *Expression) Evaluate() int {
	left, right := e.LeftOperand.Evaluate(), e.RightOperand.Evaluate()
	switch e.Operator {
	case ADD:
		return left + right
	default:
		return left * right
	}
}

func SumEvaluatedExpressions(input []string, precedence Precedence) (int, error) {
	var sum int
	for _, expressionString := range input {
		expression, err := ParseExpression(expressionString, precedence)
		if err != nil {
			return sum, err
		}

		sum += expression.Evaluate()
	}

	return sum, nil
}

func ParseExpression(input string, precedence Precedence) (Operand, error) {
	return precedence.Parse(input)
}

type Precedence interface {
	Parse(input string) (Operand, error)
}

type AdditionHigherPrecedence struct{}

func (p AdditionHigherPrecedence) Parse(input string) (Operand, error) {
	var left Operand
	var operator rune
	var skip int
	for i, r := range input {
		if skip > 0 {
			skip--
			continue
		}

		if r == ' ' {
			continue
		}

		if r == ADD || r == MULTIPLY {
			operator = r
			continue
		}

		if r == OPEN_PARENS {
			openCount, closeCount := 1, 0
			var end int
			for j, s := range input[i+1:] {
				if s == OPEN_PARENS {
					openCount++
				} else if s == CLOSE_PARENS {
					closeCount++
				}

				if openCount == closeCount {
					skip = j + 1
					end = j + i + 1
					break
				}
			}

			if left == nil {
				newLeft, err := p.Parse(input[i+1 : end])
				if err != nil {
					return nil, err
				}

				left = newLeft
			} else {
				if operator == MULTIPLY && end < len(input)-1 && input[end+2] == ADD {
					right, err := p.Parse(input[i:])
					if err != nil {
						return nil, err
					}

					left = &Expression{
						LeftOperand:  left,
						RightOperand: right,
						Operator:     operator,
					}
					break
				} else {
					right, err := p.Parse(input[i+1 : end])
					if err != nil {
						return nil, err
					}

					left = &Expression{
						LeftOperand:  left,
						RightOperand: right,
						Operator:     operator,
					}
				}
			}
			continue
		}

		if unicode.IsDigit(r) {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}

			if left == nil {
				left = &Value{Value: n}
			} else {
				if operator == MULTIPLY && i < len(input)-4 && rune(input[i+2]) == ADD {
					right, err := p.Parse(input[i:])
					if err != nil {
						return nil, err
					}

					left = &Expression{
						LeftOperand:  left,
						RightOperand: right,
						Operator:     operator,
					}

					break
				} else {
					right := &Value{Value: n}
					left = &Expression{
						LeftOperand:  left,
						RightOperand: right,
						Operator:     operator,
					}
				}
			}
		}
	}

	return left, nil
}

type EqualPrecedence struct{}

func (p EqualPrecedence) Parse(input string) (Operand, error) {
	var left Operand
	var operator rune
	var skip int
	for i, r := range input {
		if skip > 0 {
			skip--
			continue
		}

		if r == ' ' {
			continue
		}

		if r == ADD || r == MULTIPLY {
			operator = r
			continue
		}

		if r == OPEN_PARENS {
			openCount, closeCount := 1, 0
			var end int
			for j, s := range input[i+1:] {
				if s == OPEN_PARENS {
					openCount++
				} else if s == CLOSE_PARENS {
					closeCount++
				}

				if openCount == closeCount {
					skip = j + 1
					end = j + i + 1
					break
				}
			}

			if left == nil {
				newLeft, err := p.Parse(input[i+1 : end])
				if err != nil {
					return nil, err
				}

				left = newLeft
			} else {
				right, err := p.Parse(input[i+1 : end])
				if err != nil {
					return nil, err
				}

				left = &Expression{
					LeftOperand:  left,
					RightOperand: right,
					Operator:     operator,
				}
			}
			continue
		}

		if unicode.IsDigit(r) {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, err
			}

			if left == nil {
				left = &Value{Value: n}
			} else {
				right := &Value{Value: n}
				left = &Expression{
					LeftOperand:  left,
					RightOperand: right,
					Operator:     operator,
				}
			}
		}
	}

	return left, nil
}
