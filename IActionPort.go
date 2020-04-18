package go_swiss_rule

type IAction interface {
	Fire() error
}
