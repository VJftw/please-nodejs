package please

import (
	"fmt"

	"github.com/bazelbuild/buildtools/build"
	"github.com/bazelbuild/buildtools/edit"
)

// Rule represents a Please Build Rule.
type Rule struct {
	*build.Rule

	Target *Target
}

// NewRule returns a new Please Build Rule for the given Target, kind and attrs.
func NewRule(
	target *Target,
	kind string,
	attrs map[string]interface{},
) (*Rule, error) {
	rule, _ := edit.ExprToRule(&build.CallExpr{
		X:    &build.Ident{Name: kind},
		List: []build.Expr{},
	}, kind)

	rule.SetAttr("name", &build.StringExpr{Value: target.Name})

	for attrName, val := range attrs {
		switch typedVal := val.(type) {
		case string:
			if typedVal != "" {
				rule.SetAttr(attrName, &build.StringExpr{Value: typedVal})
			}
		case []string:
			listExpr := &build.ListExpr{}
			for _, s := range typedVal {
				if s != "" {
					listExpr.List = append(listExpr.List, &build.StringExpr{Value: s})
				}
			}
			if len(listExpr.List) > 0 {
				rule.SetAttr(attrName, listExpr)
			}
			// TODO: add other types (bool, int)
		default:
			return nil, fmt.Errorf("unsupported type '%T' for '%s'", typedVal, attrName)
		}
	}

	return &Rule{
		Rule:   rule,
		Target: target,
	}, nil
}
