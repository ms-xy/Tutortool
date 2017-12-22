package querytools

import (
	"database/sql"
	"reflect"
	"regexp"
	"strings"
)

var _re_alphanumeric_, _ = regexp.Compile("[^a-zA-Z0-9]+")

/*
Instantiate or test the row scanning utility
*/
func NewRowScanner(rows *sql.Rows, d interface{}) *RowScanner {
	this := &RowScanner{}

	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}

	this.dtype = reflect.TypeOf(d).Elem()
	this.columns = columns
	this.mapping = this.createMapping()
	this.rows = rows
	return this
}

func TestRowScanner(d interface{}, columns []string) []interface{} {
	rowScanner := &RowScanner{}
	rowScanner.dtype = reflect.TypeOf(d).Elem()
	rowScanner.columns = columns
	rowScanner.mapping = rowScanner.createMapping()
	d_inst := rowScanner.instantiateDataStruct()
	return rowScanner.applyMapping(d_inst)
}

/*
Row scanning utility.

Can be used to mostly automate the process of fetching items from a database.
Example:

  results := &<datatype>{}
  scanner := querytools.NewRowScanner(rows, (*<datatype>)(nil))

  for rows.Next {
    result = append(result, scanner.Fetch().(<datatype>))
  }

*/
type RowScanner struct {
	dtype   reflect.Type
	columns []string
	mapping [][]int
	rows    *sql.Rows
}

func (this *RowScanner) Fetch() interface{} {
	d := this.instantiateDataStruct()
	this.rows.Scan(this.applyMapping(d)...)
	return d
}

/*
Normalize the given string to lowercase and remove any non-alphanumeric
characters.
*/
func (this *RowScanner) normalizeKey(name string) string {
	n := _re_alphanumeric_.ReplaceAllString(name, "")
	return strings.ToLower(n)
}

/*
recursively iterate through all anonymous subtypes and add all fields
*/
func (this *RowScanner) createMappingRecursive(t reflect.Type, idx []int) map[string][]int {
	m := map[string][]int{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fidx := append(idx, f.Index...)

		if f.Anonymous {
			// iterate over inherited fields
			_m := this.createMappingRecursive(f.Type, fidx)
			for key, cidx := range _m {
				if midx, exists := m[key]; !exists || len(cidx) < len(midx) {
					m[key] = cidx
				}
			}

		} else {
			ks := f.Name[0]
			if 'A' <= ks && ks <= 'Z' {
				key := this.normalizeKey(f.Name)
				m[key] = fidx
			}
		}
	}

	return m
}

/*
New largely simplified mapping function to be used by the automatic row scanner
*/
func (this *RowScanner) createMapping() [][]int {
	m := this.createMappingRecursive(this.dtype, []int{})
	l := make([][]int, len(this.columns))
	for i, column := range this.columns {
		l[i] = m[this.normalizeKey(column)]
	}
	return l
}

/*
Create a new item of the required type
*/
func (this *RowScanner) instantiateDataStruct() interface{} {
	return reflect.New(this.dtype).Interface()
}

/*
Apply the known mapping to an object instance (preferably created automatically
using instantiateDataStruct)
*/
func (this *RowScanner) applyMapping(d interface{}) []interface{} {
	ptrs := make([]interface{}, len(this.mapping))
	v := reflect.ValueOf(d).Elem()
	for i := 0; i < len(this.mapping); i++ {
		vf := v.FieldByIndex(this.mapping[i])
		if vf.CanAddr() && vf.Addr().CanInterface() {
			ptrs[i] = vf.Addr().Interface()
		}
	}
	return ptrs
}
