package validate

import "day-4/internal/validate"

var parsePassportFieldTestCases = []struct {
	input    string
	expected validate.PassportField
	success  bool
}{
	{
		input: "eyr:2027",
		expected: validate.PassportField{
			Name:  "eyr",
			Value: "2027",
		},
		success: true,
	},
	{
		input: "eyr:2027:123",
		expected: validate.PassportField{
			Name:  "-1",
			Value: "-1",
		},
		success: false,
	},
}

var parsePassportTestCases = []struct {
	input    string
	expected validate.Passport
	success  bool
}{
	{
		input: "eyr:2025 iyr:2011\r\nbyr:1980\r\nhcl:#fffffd cid:129 pid:420023864\r\nhgt:150cm\r\necl:brn",
		expected: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		success: true,
	},
	{
		input: "eyr:2025 iyr:2011\r\nabug!:1980\r\nhcl:#fffffd cid:129 pid:420023864\r\nhgt:150cm\r\necl:brn",
		expected: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		success: false,
	},
}

var parsePassportIsValid = []struct {
	passport validate.Passport
	expected bool
}{
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: true,
	},
	{
		passport: validate.Passport{
			BirthYear:      1919,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      2005,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2009,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2022,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2019,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2035,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: ""},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 135, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 200, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 55, Unit: "in"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 82, Unit: "in"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffgfd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffgf19d",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "meep",
			PassportId:     "420023864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "12345678910",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "4200x3864",
			CountryId:      "129",
		},
		expected: false,
	},
	{
		passport: validate.Passport{
			BirthYear:      1980,
			IssueYear:      2011,
			ExpirationYear: 2025,
			Height:         validate.Height{Measurement: 150, Unit: "cm"},
			HairColor:      "#fffffd",
			EyeColor:       "brn",
			PassportId:     "420023864",
			CountryId:      "",
		},
		expected: true,
	},
}
