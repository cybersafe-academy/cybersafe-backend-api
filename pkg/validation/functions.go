package validation

import (
	"cybersafe-backend-api/pkg/helpers"
)

func IsCPF(doc string) bool {

	cpf := helpers.RemoveAllSpecialChars(doc)

	if len(cpf) != 11 {
		return false
	}

	isAllDigitsSame := true
	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			isAllDigitsSame = false
			break
		}
	}
	if isAllDigitsSame {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}

	remainder := sum % 11
	if remainder < 2 {
		remainder = 0
	} else {
		remainder = 11 - remainder
	}
	if cpf[9]-'0' != byte(remainder) {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	remainder = sum % 11
	if remainder < 2 {
		remainder = 0
	} else {
		remainder = 11 - remainder
	}
	if cpf[10]-'0' != byte(remainder) {
		return false
	}

	return true
}

func IsCNPJ(doc string) bool {

	cnpj := helpers.RemoveAllSpecialChars(doc)

	if len(cnpj) != 14 {
		return false
	}

	isAllDigitsSame := true
	for i := 1; i < 14; i++ {
		if cnpj[i] != cnpj[0] {
			isAllDigitsSame = false
			break
		}
	}
	if isAllDigitsSame {
		return false
	}

	sum := 0
	for i := 0; i < 12; i++ {
		digit := int(cnpj[i] - '0')
		if i < 4 {
			digit *= 5 - i
		} else {
			digit *= 13 - i
		}
		sum += digit
	}
	remainder := sum % 11
	var firstDigit byte
	if remainder < 2 {
		firstDigit = '0'
	} else {
		firstDigit = byte(11 - remainder + '0')
	}
	if cnpj[12] != firstDigit {
		return false
	}

	sum = 0
	for i := 0; i < 13; i++ {
		digit := int(cnpj[i] - '0')
		if i < 5 {
			digit *= 6 - i
		} else {
			digit *= 14 - i
		}
		sum += digit
	}

	remainder = sum % 11
	var secondDigit byte
	if remainder < 2 {
		secondDigit = '0'
	} else {
		secondDigit = byte(11 - remainder + '0')
	}
	if cnpj[13] != secondDigit {
		return false
	}

	return true

}
