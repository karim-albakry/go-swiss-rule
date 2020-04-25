package go_swiss_rule

import (
	"fmt"
	"github.com/antonmedv/expr"
	"github.com/karim-albakry/go-swiss-rule/utils/errors"
	"reflect"
)

func buildStringQuery(input map[string]interface{}, rules []*Rule) (result string) {
	for _, rule := range rules {
		for index, condition := range rule.Conditions {
			if value, ok := input[condition.Key]; ok {
				valueType := reflect.
					TypeOf(value).
					String()
				if valueType == "string" {
					result += fmt.
						Sprintf("(('%v') %s '%v') ",
							value, condition.Operator,
							condition.Value)
				} else {
					result += fmt.
						Sprintf("(%v %s %v) ",
							value,
							condition.Operator,
							condition.Value)
				}
				if index !=
					(len(rule.Conditions)-1) && condition.Joint != "" {
					result += fmt.
						Sprintf(" %v ", condition.Joint)
				}
			}
		}
	}
	return
}

func invokeActions(rules []*Rule) error {
	for _, rule := range rules {
		for _, action := range rule.Actions {
			if err := action.Fire(); err != nil {
				return err
			}
		}
	}
	return nil
}

func EvalAndInvoke(input map[string]interface{}, rules []*Rule) error {
	if len(rules) < 0 {
		return errors.SimpleError("Insufficient rules.")
	}
	if len(input) < 0 {
		return errors.SimpleError("Insufficient input.")
	}
	if constraints := buildStringQuery(input, rules); constraints != "" {
		result, exprError := expr.Eval(constraints, input)
		if exprError != nil {
			return exprError
		}
		if result == true {
			err := invokeActions(rules)
			if err != nil {
				return err
			}
		} else {
			return errors.SimpleError("Invalid Input")
		}
	}
	return errors.SimpleError("No conditions to execute, review your conditions keys.")
}
