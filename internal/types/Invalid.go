package types

type Valider interface {
	InValid() bool
	InValidErr() error
}
