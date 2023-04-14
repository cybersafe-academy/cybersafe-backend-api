package validation

import (
	"github.com/asaskevich/govalidator"
)

func Config() {

	govalidator.TagMap["cpf"] = govalidator.Validator(IsCPF)
	govalidator.TagMap["cnpj"] = govalidator.Validator(IsCNPJ)

}
