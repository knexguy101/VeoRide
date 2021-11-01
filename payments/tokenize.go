package payments

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/token"
)

const STRIPE_API_KEY = "pk_live_H7TSU4oS7ArzxbRDpK3u8yet"

func init() {
	stripe.Key = STRIPE_API_KEY
}

func TokenizePayment(card, month, year, cvv string) (*stripe.Token, error) {
	params := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number: stripe.String(card),
			ExpMonth: stripe.String(month),
			ExpYear: stripe.String(year),
			CVC: stripe.String(cvv),
		},
	}
	return token.New(params)
}
