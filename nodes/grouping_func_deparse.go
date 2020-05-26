// Auto-generated - DO NOT EDIT

package pg_query

import "strings"

func (node GroupingFunc) Deparse(ctx DeparseContext) (string, error) {
	out := []string{"GROUPING", "("}

	argOut, err := node.Args.DeparseList(ctx)
	if err != nil {
		return "", err
	}
	out = append(out, strings.Join(argOut, ", "))
	out = append(out, ")")

	return strings.Join(out, " "), nil
}
