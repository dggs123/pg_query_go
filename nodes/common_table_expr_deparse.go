// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"
)

func (node CommonTableExpr) Deparse(ctx DeparseContext) (string, error) {
	var out []string
	out = append(out, *node.Ctename)

	aliasOut, err := node.Aliascolnames.DeparseList(ctx)
	if err != nil {
		return "", nil
	}
	if len(aliasOut) != 0 {
		out = append(out, fmt.Sprintf("(%s)", strings.Join(aliasOut, ",")))
	}

	out = append(out, "AS")

	queryOut, err := node.Deparse(ctx)
	if err != nil {
		return "", nil
	}
	out = append(out, fmt.Sprintf("(%s)", queryOut))

	return strings.Join(out, " "), nil
}
