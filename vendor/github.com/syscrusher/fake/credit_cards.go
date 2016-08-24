package fake

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type creditCard struct {
	vendor   string
	length   int
	prefixes [][]int
}

// https://en.wikipedia.org/wiki/Payment_card_number#Issuer_identification_number_.28IIN.29
var creditCards = map[string]creditCard{
	"amex":       {"American Express", 15, [][]int{[]int{3, 4}, []int{3, 7}}},
	"discover":   {"Discover", 16, [][]int{[]int{6, 0, 1, 1}, []int{6, 2, 2, 1, 2, 6}, []int{6, 2, 2, 9, 2, 5}, []int{6, 4, 4}, []int{6, 4, 9}, []int{6, 5}}},
	"mastercard": {"MasterCard", 16, [][]int{[]int{5}}},
	"visa":       {"VISA", 16, [][]int{[]int{4}}},
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
	if vendor == "" {
		var vendors []string
		for v := range creditCards {
			vendors = append(vendors, v)
		}
		vendor = vendors[rand.Intn(len(vendors))]
	}
	vendor = strings.ToLower(vendor)

	card := creditCards[vendor]
	prefix := card.prefixes[rand.Intn(len(card.prefixes))]

	num := generateWithPrefix(card.length, prefix)

	var final string
	for i := 0; i < len(num); i++ {
		final = final + strconv.Itoa(num[i])
	}

	return final
}

// Remainder of file adapted from github.com/joeljunstrom/go-luhn
// License: WTFPL

// generateWithPrefix creates and returns a string of the length of the argument targetSize
// but prefixed with the second argument.
// The returned string is valid according to the Luhn algorithm.
func generateWithPrefix(size int, prefix []int) []int {
	size = size - 1 - len(prefix)

	random := append(prefix, randomIntSlice(size)...)
	controlDigit := generateControlDigit(random)

	return append(random, controlDigit)
}

func randomIntSlice(size int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = rand.Intn(9)
	}

	return result
}

func generateControlDigit(luhnString []int) int {
	controlDigit := calculateChecksum(luhnString, true) % 10

	if controlDigit != 0 {
		controlDigit = 10 - controlDigit
	}

	return controlDigit
}

func calculateChecksum(luhnString []int, double bool) int {
	checksum := 0

	for i := len(luhnString) - 1; i > -1; i-- {
		n := luhnString[i]

		if double {
			n = n * 2
		}
		double = !double

		if n >= 10 {
			n = n - 9
		}

		checksum += n
	}

	return checksum
}
