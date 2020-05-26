// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"
)

func (node RangeFunction) Deparse(ctx DeparseContext) (string, error) {

	var out []string

	functions, err := node.Functions.DeparseList(ctx)
	if err != nil {
		return "", nil
	}

	out = append(out, fmt.Sprintf("%s", strings.Join(functions, ", ")))

	if node.Alias != nil {
		out = append(out, fmt.Sprintf("AS %s", *node.Alias.Aliasname))
	}
	return strings.Join(out, " "), nil
}
