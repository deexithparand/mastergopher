package transaction

type Transaction interface {
	Transfer() bool
}
