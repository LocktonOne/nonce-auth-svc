package util

import (
	"fmt"
	"regexp"
	"strings"
)

const LayoutISO8601Date = "2006-01-02"

var AddressRegexp = regexp.MustCompile("^(0x)?[0-9a-fA-F]{40}$")
var SignatureRegexp = regexp.MustCompile("^0x[0-9a-fA-F]+$")

func NonceToTermsMessage(nonce, hash string) string {
	return fmt.Sprintf(
		"User Service Test Authentication\n\n"+
			"By signing this message you're also agreeing to the Terms of Service\n"+
			"Terms of Service hash: %s\n\n"+
			"Below is your one-time nonce to make your authentication secure\n%s",
		hash, nonce)
}

func NonceToMessage(nonce string) string {
	return fmt.Sprintf(
		"User Service Test Authentication\n\n"+
			"Below is your one-time nonce to make your authentication secure\n%s",
		nonce,
	)
}
func PrefixNonceMessage(message string) string {
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	return prefixedMessage
}

func StringSliceToLowercase(inp []string) []string {
	var newSlice []string
	for _, str := range inp {
		newSlice = append(newSlice, strings.ToLower(str))
	}
	return newSlice
}