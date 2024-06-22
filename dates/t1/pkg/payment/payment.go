package payment

// UpiDetails struct is now exported
type UpiDetails struct {
	UpiID string // Field name starts with uppercase to make it exported
}

// TransactionDetails struct is now exported
type TransactionDetails struct {
	Amount        int32      // Field name starts with uppercase to make it exported
	Currency      string     // Field name starts with uppercase to make it exported
	PaymentMode   string     // Field name starts with uppercase to make it exported
	PaymentDetail UpiDetails // Field name starts with uppercase to make it exported
}

// PaymentStruct struct is now exported
type PaymentStruct struct {
	Transactions []TransactionDetails // Field name starts with uppercase to make it exported
}
