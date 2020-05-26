// Auto-generated - DO NOT EDIT

package pg_query

import "fmt"

func (node A_Indirection) Deparse(ctx DeparseContext) (string, error) {
	// deparse args

	arg, err := node.Arg.Deparse(ctx)
	if err != nil {
		return "", err
	}

	var indices string

	for _, n := range node.Indirection.Items {
		str, err := n.Deparse(ctx)
		if err != nil {
			return "", err
		}
		indices = indices + str
	}

	return fmt.Sprintf("%s%s", arg, indices), nil
}
