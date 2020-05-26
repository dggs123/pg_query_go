// Auto-generated - DO NOT EDIT

package pg_query

import "fmt"

func (node A_Indices) Deparse(ctx DeparseContext) (string, error) {
	out := make([]string, 0)

	if node.Lidx != nil {
		if str, err := node.Lidx.Deparse(ctx); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	if node.Uidx != nil {
		if str, err := node.Lidx.Deparse(ctx); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	var result string
	if node.IsSlice {
		if node.Lidx != nil {
			result = fmt.Sprintf("[:%s]", out[0])
		} else if node.Uidx != nil {
			result = fmt.Sprintf("[%s:]", out[0])
		} else {
			result = fmt.Sprintf("[%s:%s]", out[0], out[1])
		}
	} else {
		result = fmt.Sprintf("[%s]", out[0])
	}
	return result, nil
}
