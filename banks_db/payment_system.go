package banks_db

import "regexp"

var rawPaymentSystems = map[string]string{
	"electron":           `^(4026|417500|4405|4508|4844|4913|4917)\d+$`,
	"maestro":            `^(5018|5020|5038|5612|5893|6304|6759|6761|6762|6763|0604|6390)\d+$`,
	"dankort":            `^(5019)\d+$`,
	"interpayment":       `^(636)\d+$`,
	"unionpay":           `^(62|88)\d+$`,
	"visa":               `^4[0-9]{12}(?:[0-9]{3})?`,
	"mastercard":         `^5[1-5][0-9]{14}`,
	"amex":               `^3[47][0-9]{13}`,
	"diners":             `^3(?:0[0-5]|[68][0-9])[0-9]{11}`,
	"discover":           `^6(?:011|5[0-9]{2})[0-9]{12}`,
	"jcb":                `^(?:2131|1800|35\d{3})\d{11}$`,
	"forbrugsforeningen": `^(600)\d+$`,
	"mir":                `^220[0-4][0-9][0-9]\d{10}$`,
}

var paymentSystems = map[string]*regexp.Regexp{}

func init() {
	for paymentSystem, rawRe := range rawPaymentSystems {
		re, err := regexp.Compile(rawRe)
		if err != nil {
			panic(err)
		}
		paymentSystems[paymentSystem] = re
	}
}

func FindPaymentSystem(creditCard string) *string {
	for paymentSystem, re := range paymentSystems {
		if re.MatchString(creditCard) {
			return &paymentSystem
		}
	}
	return nil
}
