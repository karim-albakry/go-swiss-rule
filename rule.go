package go_swiss_rule

type Rule struct {
	EventType  string
	Conditions []Condition
	Actions    []IAction
}
