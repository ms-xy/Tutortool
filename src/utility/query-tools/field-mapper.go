package querytools

import (
	"github.com/ms-xy/Tutortool/src/utility/reflect-tools"
	"reflect"
	"strings"
	"sync"
)

var lock sync.Mutex
var instance *FieldMapper

/*
Instantiate or test the row scanning utility
*/
func NewFieldMapper() *FieldMapper {
	this := &FieldMapper{
		Type2Map: make(map[reflect.Type]*FieldMap),
		Db2Map:   make(map[string]*FieldMap),
	}
	return this
}

func GetFieldMapperInstance() *FieldMapper {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = NewFieldMapper()
		}
	}
	return instance
}

func TestFieldMapper() []interface{} {
	return nil
}

/*
Field Mapper
Creates mappings for use with the RowScanner and other purposes regarding
database operations.
*/
type FieldMapper struct {
	Type2Map map[reflect.Type]*FieldMap
	Db2Map   map[string]*FieldMap
}

func (this *FieldMapper) Register(d interface{}) *FieldMap {
	fm := (&FieldMap{}).init(d)
	this.Type2Map[fm.Type] = fm
	this.Db2Map[fm.DatabaseTable.Name] = fm
	return fm
}

func (this *FieldMapper) GetByType(d interface{}) (*FieldMap, bool) {
	fm, exists := this.Type2Map[reflect.TypeOf(d)]
	return fm, exists
}

func (this *FieldMapper) GetByTableName(tn string) (*FieldMap, bool) {
	fm, exists := this.Db2Map[tn]
	return fm, exists
}

/*
FieldMap support struct
Provides a mapping for a single type.
*/
type FieldMap struct {
	Type reflect.Type

	// QueryInsert          string
	// MappingInsertColumns [][]int
	// QuerySelect          string
	// MappingSelectColumns [][]int
	// MappingSelectWhere   [][]int
	// QueryDelete          string
	// MappingDeleteWhere   [][]int

	Db2Struct map[string][]int
	// Struct2Db map[[]int]string

	DatabaseTable *DatabaseTable
}

func (this *FieldMap) init(d interface{}) *FieldMap {
	this.Db2Struct = make(map[string][]int)
	//this.Struct2Db = make(map[[]int]string)
	this.Type = reflect.TypeOf(d)

	// create table definition (required for automatic query generation later)
	this.DatabaseTable = newDatabaseTable(strings.ToLower(this.Type.Name()))

	// create a map fieldname -> index, with minimal indices lengths
	m := reflecttools.MapStructRecursively(this.Type, []int{})

	// take this map and convert it to the required field mapping
	for _, idx := range m {
		field := this.Type.FieldByIndex(idx)

		// only correctly tagged fields are respected
		if dbfield := field.Tag.Get("dbfield"); dbfield != "" {

			this.Db2Struct[dbfield] = idx
			//this.Struct2Db[idx] = dbfield

			// add column to table
			this.DatabaseTable.AddColumn(newDatabaseTableColumn(dbfield,
				field.Tag.Get("dbkey")))

		} else if dblink := field.Tag.Get("dblink"); dblink != "" {
			// could be a dblink specification instead ...
			// though ... how to treat them properly without risk of failure?
		}
	}

	// pre-create standard queries ::: TODO remove?
	// var columns, whereColumns []string
	// where := WhereClause{}
	// if this.DatabaseTable.PrimaryKey != nil {
	// 	where = append(where, this.DatabaseTable.PrimaryKey.Name)
	// }
	// // insert
	// columns, this.QueryInsert = this.DatabaseTable.CreateInsertQuery()
	// this.MappingInsertColumns = this.CreateMapping(columns)
	// // select
	// columns, whereColumns, this.QuerySelect = this.DatabaseTable.CreateSelectQuery(nil, where)
	// this.MappingSelectColumns = this.CreateMapping(columns)
	// this.MappingSelectWhere = this.CreateMapping(whereColumns)
	// // delete
	// whereColumns, this.QueryDelete = this.DatabaseTable.CreateDeleteQuery(where)
	// this.MappingDeleteWhere = this.CreateMapping(whereColumns)

	return this
}

/*
Create a new item of the required type
*/
func (this *FieldMap) newInstance() reflect.Value {
	return reflect.New(this.Type)
}

/*
Apply the requested mapping (based on the given columns) to an instance of the
underlying object (represented by it's reflect.Value)
*/
func (this *FieldMap) ApplyMapping(mapping [][]int, d reflect.Value, write bool) []interface{} {
	ptrs := make([]interface{}, len(mapping))
	v := d.Elem()
	for i, idx := range mapping {
		vf := v.FieldByIndex(idx)
		if write {
			ptrs[i] = vf.Addr().Interface()
		} else {
			ptrs[i] = vf.Interface()
		}
	}
	return ptrs
}

/*
Create a field mapping based on the provided columns
*/
func (this *FieldMap) CreateMapping(columns []string) [][]int {
	idxs := make([][]int, len(columns))
	for i, column := range columns {
		idxs[i] = this.Db2Struct[column]
	}
	return idxs
}
