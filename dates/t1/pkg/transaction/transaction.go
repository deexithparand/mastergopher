package transaction

// type  struct{

// }

type transaction interface {
	transfer() error
}

func TransferMethod(t transaction) error {
	return t.transfer()
}
