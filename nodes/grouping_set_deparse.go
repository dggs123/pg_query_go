// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"
)

func (node GroupingSet) Deparse(ctx DeparseContext) (string, error) {
	var out []string

	switch node.Kind {
	case GROUPING_SET_SIMPLE:
		out = append(out, "SIMPLE")
	case GROUPING_SET_ROLLUP:
		out = append(out, "ROLLUP")
	case GROUPING_SET_CUBE:
		out = append(out, "CUBE")
	case GROUPING_SET_SETS:
		out = append(out, "GROUPING SETS")
	default:
		panic("Not Implemented")
	}

	contentOut, err := node.Content.DeparseList(ctx)
	if err != nil {
		return "", err
	}

	out = append(out, fmt.Sprintf("(%s)", strings.Join(contentOut, ", ")))

	return strings.Join(out, " "), nil
}
