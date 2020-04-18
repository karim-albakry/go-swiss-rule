package go_swiss_rule

import (
	"fmt"
	"github.com/antonmedv/expr"
	"go-swiss-rule/utils/errors"
	"reflect"
)

func buildStringQuery(rules []*Rule, input map[string]interface{}) (result string) {
	for _, rule := range rules {
		for index, condition := range rule.Conditions {
			valueType := reflect.
				TypeOf(input[condition.Key]).
				String()
			if valueType == "string" {
				result += fmt.
					Sprintf("(('%v') %s '%v') ",
						input[condition.Key], condition.Operator,
						condition.Value)
			} else {
				result += fmt.
					Sprintf("(%v %s %v) ",
						input[condition.Key],
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
	if constraints := buildStringQuery(rules, input); constraints != "" {
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
	} else {

	}
	return nil
}
