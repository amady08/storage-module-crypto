package services

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

/*InitStripe initializes stripe service with our API keys*/
func InitStripe(stripeKey string) error {
	stripe.Key = stripeKey
	return nil
}

/*CreateCharge creates a charge with stripe and returns the charge*/
func CreateCharge(costInDollars float64, stripeToken string, accountID string) (*stripe.Charge, error) {
	cost := int64(costInDollars * 100)

	params := &stripe.ChargeParams{
		Amount:              stripe.Int64(cost),
		Currency:            stripe.String(string(stripe.CurrencyUSD)),
		Description:         stripe.String("File Storage"),
		StatementDescriptor: stripe.String("File Storage"),
	}
	params.SetSource(stripeToken)
	params.AddMetadata("accountID", accountID)
	return charge.New(params)
}

/*CheckChargeStatus checks the status of a charge*/
func CheckChargeStatus(chargeID string) (string, error) {
	c, err := charge.Get(chargeID, nil)
	if err != nil {
		return "", err
	}
	return c.Status, nil
}

/*CheckChargePaid checks that a charge has been paid*/
func CheckChargePaid(chargeID string) (bool, error) {
	c, err := charge.Get(chargeID, nil)
	if err != nil {
		return false, err
	}
	return c.Paid, nil
}

/*CheckChargeAmount returns the amount of a charge*/
func CheckChargeAmount(chargeID string) (float64, error) {
	c, err := charge.Get(chargeID, nil)
	if err != nil {
		return 0, err
	}
	return float64(c.Amount) / 100.00, nil
}
