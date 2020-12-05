package validate

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Height struct {
	Measurement int
	Unit        string
}

type PassportField struct {
	Name, Value string
}

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         Height
	HairColor      string
	EyeColor       string
	PassportId     string
	CountryId      string
}

func (p Passport) IsValid() bool {
	return p.isValidBirthYear() && p.isValidIssueYear() && p.isValidExpirationYear() &&
		p.isValidHeight() && p.isValidHairColor() && p.isValidEyeColor() && p.isValidPassportId()
}

func (p Passport) isValidBirthYear() bool {
	return p.BirthYear >= 1920 && p.BirthYear <= 2002
}

func (p Passport) isValidIssueYear() bool {
	return p.IssueYear >= 2010 && p.IssueYear <= 2020
}

func (p Passport) isValidExpirationYear() bool {
	return p.ExpirationYear >= 2020 && p.ExpirationYear <= 2030
}

func (p Passport) isValidHeight() bool {
	if p.Height.Unit == "cm" {
		return p.Height.Measurement >= 150 && p.Height.Measurement <= 193
	} else if p.Height.Unit == "in" {
		return p.Height.Measurement >= 59 && p.Height.Measurement <= 76
	}

	return false
}

func (p Passport) isValidHairColor() bool {
	if !strings.HasPrefix(p.HairColor, "#") {
		return false
	}

	hcl := strings.TrimPrefix(p.HairColor, "#")
	if len(hcl) != 6 {
		return false
	}

	_, err := hex.DecodeString(hcl)

	return err == nil
}

func (p Passport) isValidEyeColor() bool {
	eyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, eyeColor := range eyeColors {
		if eyeColor == p.EyeColor {
			return true
		}
	}

	return false
}

func (p Passport) isValidPassportId() bool {
	if len(p.PassportId) != 9 {
		return false
	}

	for _, r := range p.PassportId {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func AllPassports(inputs []string) (int, error) {
	var valid int
	for _, input := range inputs {
		passport, err := ParsePassport(input)
		if err != nil {
			return valid, err
		}

		if passport.IsValid() {
			valid++
		}
	}

	return valid, nil
}

func ParsePassport(input string) (*Passport, error) {
	passport := Passport{}

	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		components := strings.Split(line, " ")
		for _, component := range components {
			field, err := ParsePassportField(component)
			if err != nil {
				return nil, err
			}

			err = SetPassportField(&passport, field)
			if err != nil {
				return &passport, err
			}
		}
	}

	return &passport, nil
}

func SetPassportField(passport *Passport, field *PassportField) error {
	switch field.Name {
	case "byr":
		birthYear, err := strconv.Atoi(field.Value)
		if err != nil {
			return fmt.Errorf("invalid passport field %s", field.Name)
		}
		passport.BirthYear = birthYear

	case "iyr":
		issueYear, err := strconv.Atoi(field.Value)
		if err != nil {
			return fmt.Errorf("invalid passport field %s", field.Name)
		}
		passport.IssueYear = issueYear

	case "eyr":
		expirationYear, err := strconv.Atoi(field.Value)
		if err != nil {
			return fmt.Errorf("invalid passport field %s", field.Name)
		}
		passport.ExpirationYear = expirationYear

	case "hgt":
		var unit string
		value := field.Value
		if strings.HasSuffix(value, "in") {
			unit = "in"
			value = strings.TrimSuffix(value, "in")
		} else if strings.HasSuffix(value, "cm") {
			unit = "cm"
			value = strings.TrimSuffix(value, "cm")
		}

		measurement, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("invalid passport field %s", field.Name)
		}
		passport.Height = Height{Measurement: measurement, Unit: unit}

	case "hcl":
		passport.HairColor = field.Value

	case "ecl":
		passport.EyeColor = field.Value

	case "pid":
		passport.PassportId = field.Value

	case "cid":
		passport.CountryId = field.Value

	default:
		return fmt.Errorf("invalid passport field %s", field.Name)
	}

	return nil
}

func ParsePassportField(input string) (*PassportField, error) {
	var name, value string
	components := strings.Split(input, ":")

	if len(components) != 2 {
		return nil, fmt.Errorf("not a valid passport field %s", input)
	}

	name = components[0]
	value = components[1]

	return &PassportField{Name: name, Value: value}, nil
}
