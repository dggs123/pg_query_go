// Auto-generated - DO NOT EDIT

package pg_query

import "strings"

func (node CoalesceExpr) Deparse(ctx DeparseContext) (string, error) {
	out := []string{"COALESCE", "("}

	for _, n := range node.Args.Items {
		str, err := n.Deparse(ctx)
		if err != nil {
			return "", err
		}
		out = append(out, str)
	}
	out = append(out, ")")
	return strings.Join(out, ""), nil
}
