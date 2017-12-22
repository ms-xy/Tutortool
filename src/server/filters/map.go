package filters

import (
	// templating engine
	"github.com/flosch/pongo2"
)

func init() {
	pongo2.RegisterFilter("mapGet", filterMapGet)
}

/*
Return the student ID as a string
*/
func filterMapGet(_v *pongo2.Value, p *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	if m, ok := _v.Interface().(map[string]interface{}); ok {
		if k, ok := p.Interface().(string); ok {
			if v, exists := m[k]; exists {
				return pongo2.AsValue(v), nil
			} else {
				return nil, &pongo2.Error{
					OrigError: ENOTFOUND,
					Sender:    "filter:mapGet",
				}
			}
		} else {
			return nil, &pongo2.Error{
				OrigError: EINVALIDTYPE,
				Sender:    "filter:mapGet",
			}
		}
	} else {
		return nil, &pongo2.Error{
			OrigError: EINVALIDTYPE,
			Sender:    "filter:mapGet",
		}
	}
}
