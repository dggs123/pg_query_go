// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"
)

func (node RowExpr) Deparse(ctx DeparseContext) (string, error) {
	out, err := node.Args.DeparseList(ctx)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("(%s)", strings.Join(out, ", ")), nil
}
