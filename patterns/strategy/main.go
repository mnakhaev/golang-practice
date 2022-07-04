package strategy

func main() {
	product := "vehicle"
	payWay := 3

	var payment Payment
	switch payWay {
	case 1:
		payment = NewCardPayment("12345", "1234")
	case 2:
		payment = NewPaypalPayment()
	case 3:
		payment = NewQIWIPayment()
	}
	processOrder(product, payment)
}

func processOrder(product string, payment Payment) {
	// implementation
	err := payment.Pay()
	if err != nil {
		return
	}
}

// ----

type Payment interface {
	Pay() error
}

// -----

type cardPayment struct {
	cardNumber string
	cvv        string
}

func NewCardPayment(cardNumber string, cvv string) *cardPayment {
	return &cardPayment{cardNumber: cardNumber, cvv: cvv}
}

func (p *cardPayment) Pay() error {
	// implementation
	return nil
}

type paypalPayment struct {
}

func NewPaypalPayment() *paypalPayment {
	return &paypalPayment{}
}

func (p *paypalPayment) Pay() error {
	// implementation
	return nil
}

type qiwiPayment struct {
}

func NewQIWIPayment() *qiwiPayment {
	return &qiwiPayment{}
}

func (p *qiwiPayment) Pay() error {
	// implementation
	return nil
}
