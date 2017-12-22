package filters

import (
	// templating engine
	"github.com/flosch/pongo2"

	// config
	"github.com/ms-xy/Tutortool/src/configuration"
)

func init() {
	pongo2.RegisterFilter("getTaskName", filterGetTaskName)
}

/*
Return the student ID as a string
*/
func filterGetTaskName(v *pongo2.Value, p *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	if taskID, ok := v.Interface().(uint); ok {
		if task, exists := configuration.Tasks[taskID]; exists {
			return pongo2.AsSafeValue(task.Name), nil
		} else {
			return nil, &pongo2.Error{
				OrigError: ENOTFOUND,
				Sender:    "filter:getTaskName",
			}
		}
	} else {
		return nil, &pongo2.Error{
			OrigError: EINVALIDTYPE,
			Sender:    "filter:getTaskName",
		}
	}
}
