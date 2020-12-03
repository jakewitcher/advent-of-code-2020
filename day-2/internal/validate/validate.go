package validate

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type PasswordPolicy struct {
	Character rune
	Min, Max  int
}

type Password string
type Validator func(policy *PasswordPolicy, password Password) bool

func IdentifyValidPasswords(inputs []string, validator Validator) (int, error) {
	var valid int

	for _, input := range inputs {
		password, policy, err := ParsePasswordPolicy(input)
		if err != nil {
			return 0, err
		}

		isValid := validator(policy, password)
		if isValid {
			valid++
		}
	}

	return valid, nil
}

func ApplySledRentalPolicy(policy *PasswordPolicy, password Password) bool {
	var count int

	for _, r := range password {
		if r == policy.Character {
			count++
		}
	}

	return count >= policy.Min && count <= policy.Max
}

func ApplyTobogganCorpPolicy(policy *PasswordPolicy, password Password) bool {
	var count int

	for i, r := range password {
		if (i + 1 == policy.Min || i + 1 == policy.Max) && r == policy.Character {
			count++
		}
	}

	return count == 1
}

func ParsePasswordPolicy(input string) (Password, *PasswordPolicy, error) {
	components := strings.Split(input, " ")

	if len(components) != 3 {
		return "", nil, fmt.Errorf("%s is not a valid password policy", input)
	}

	min, max, err := ParsePolicyMinMax(components[0])
	if err != nil {
		return "", nil, err
	}

	character, err := ParsePolicyCharacter(components[1])
	if err != nil {
		return "", nil, err
	}

	password := Password(components[2])

	return password, &PasswordPolicy{Character: character, Min: min, Max: max}, nil
}

func ParsePolicyMinMax(input string) (int, int, error) {
	var min, max int
	components := strings.Split(input, "-")

	if len(components) != 2 {
		return min, max, fmt.Errorf("%s is not a valid policy Min/Max", input)
	}

	min, err := strconv.Atoi(components[0])
	if err != nil {
		return min, max, err
	}

	max, err = strconv.Atoi(components[1])
	if err != nil {
		return min, max, err
	}

	return min, max, nil
}

func ParsePolicyCharacter(input string) (rune, error) {
	var character rune
	for i, r := range input {
		if i == 0 && !unicode.IsLetter(r) {
			return character, fmt.Errorf("%s is not a valid policy Character", input)
		}

		character = r
		break
	}

	return character, nil
}
