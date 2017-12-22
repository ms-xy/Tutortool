package filters

import (
	// templating engine
	"github.com/flosch/pongo2"

	// utility
	"strconv"

	// models
	"github.com/ms-xy/Tutortool/src/database/models"

	"fmt"
	"reflect"
)

func init() {
	pongo2.RegisterFilter("get_student_id", filterGetStudentId)
}

/*
Return the student ID as a string
*/
func filterGetStudentId(_v *pongo2.Value, p *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	if student, ok := _v.Interface().(*models.Student); ok {
		return pongo2.AsSafeValue(strconv.FormatUint(uint64(student.ID), 10)), nil
	} else {
		fmt.Println(reflect.TypeOf(_v.Interface()))
		return nil, &pongo2.Error{
			OrigError: EINVALIDTYPE,
			Sender:    "filter:get_student_id",
		}
	}
}
