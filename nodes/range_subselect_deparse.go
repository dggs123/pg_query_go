// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"
)

func (node RangeSubselect) Deparse(ctx DeparseContext) (string, error) {
	var out []string

	queryOut, err := node.Subquery.Deparse(ctx)
	if err != nil {
		return "", nil
	}

	out = append(out, fmt.Sprintf("(%s)", queryOut))
	if node.Alias != nil {
		out = append(out, fmt.Sprintf("AS %s", *node.Alias.Aliasname))
	}
	return strings.Join(out, " "), nil
}
