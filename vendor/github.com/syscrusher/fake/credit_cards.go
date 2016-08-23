package fake

import (
	"math/rand"
	"strings"

	"strconv"
)

type creditCard struct {
	vendor   string
	length   int
	prefixes []int
}

// https://en.wikipedia.org/wiki/Payment_card_number#Issuer_identification_number_.28IIN.29
var creditCards = map[string]creditCard{
	"amex":       {"American Express", 15, []int{34, 37}},
	"discover":   {"Discover", 16, []int{6011, 622126, 622925, 644, 649, 65}},
	"mastercard": {"MasterCard", 16, []int{5}},
	"visa":       {"VISA", 16, []int{4}},
}

// CreditCardType returns one of the following credit values:
// VISA, MasterCard, American Express and Discover
func CreditCardType() string {
	n := len(creditCards)
	var vendors []string
	for _, cc := range creditCards {
		vendors = append(vendors, cc.vendor)
	}

	return vendors[rand.Intn(n)]
}

// CreditCardNum generated credit card number according to the vendor's card number rules.
// Currently supports amex, discover, mastercard, and visa.
func CreditCardNum(vendor string) string {
	if vendor != "" {
		vendor = strings.ToLower(vendor)
	} else {
		var vendors []string
		for v := range creditCards {
			vendors = append(vendors, v)
		}
		vendor = vendors[rand.Intn(len(vendors))]
	}
	card := creditCards[vendor]
	prefix := strconv.Itoa(card.prefixes[rand.Intn(len(card.prefixes))])
	num := []rune(prefix)
	for i := 0; i < card.length-len(prefix); i++ {
		num = append(num, genCCDigit(num))
	}
	return string(num)
}

func genCCDigit(num []rune) rune {
	sum := 0
	for i := len(num) - 1; i >= 0; i-- {
		n := int(num[i])
		if i%2 != 0 {
			sum += n
		} else {
			if n*2 > 9 {
				sum += n*2 - 9
			} else {
				sum += n * 2
			}
		}
	}
	return rune(((sum/10+1)*10 - sum) % 10)
}
