package querytools

import (
	"fmt"
	"strings"
)

/*
Database field and table support structures
*/
type DatabaseTable struct {
	Name       string
	Columns    []*DatabaseTableColumn
	PrimaryKey *DatabaseTableColumn
}

type DatabaseTableColumn struct {
	Name          string
	IsPrimaryKey  bool
	AutoIncrement bool
}

type FilterFunction func(*DatabaseTableColumn) bool

type WhereClause map[string]interface{}

type OrderByOrientation int
type OrderByClause map[string]OrderByOrientation

const (
	OrderByAsc OrderByOrientation = iota
	OrderByDesc
)

func (this OrderByOrientation) String() string {
	switch this {
	case OrderByAsc:
		return "ASC"
	case OrderByDesc:
		return "DESC"
	default:
		panic("unknown order-by orientation")
	}
}

/*
Helper functions for creating tables and columns
*/
func newDatabaseTable(name string) *DatabaseTable {
	return &DatabaseTable{
		Name:       name,
		Columns:    []*DatabaseTableColumn{},
		PrimaryKey: nil,
	}
}

func newDatabaseTableColumn(name, key string) *DatabaseTableColumn {
	field := &DatabaseTableColumn{}
	field.Name = name
	if key != "" {
		for _, item := range strings.Split(key, ",") {
			if item == "primary" {
				field.IsPrimaryKey = true

			} else if item == "autoincrement" {
				field.AutoIncrement = true
			}
		}
	}
	return field
}

/*
DatabaseTable members
*/
func (this *DatabaseTable) AddColumn(column *DatabaseTableColumn) {
	if this.PrimaryKey != nil && column.IsPrimaryKey {
		panic("multiple primary keys defined for table '" + this.Name + "'")
	}
	this.Columns = append(this.Columns, column)
	if column.IsPrimaryKey {
		this.PrimaryKey = column
	}
}

func (this *DatabaseTable) GetColumns(fn FilterFunction) []string {
	result := []string{}
	for _, column := range this.Columns {
		if fn(column) {
			result = append(result, column.Name)
		}
	}
	return result
}

func (this *DatabaseTable) CreatePlaceholders(n int) []string {
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = "?"
	}
	return result
}

/*
Helper functions to assemble clauses
*/
func (this *DatabaseTable) AssembleWhereClause(where WhereClause) ([]interface{}, string) {
	// tackle empty where clause
	if where == nil || len(where) == 0 {
		return []interface{}{}, ""
	}

	mapping, columns := []interface{}{}, []string{}
	for key, val := range where {
		columns = append(columns, fmt.Sprintf("%s=?", key))
		mapping = append(mapping, val)
	}

	return mapping, fmt.Sprintf("WHERE %s", strings.Join(columns, " AND "))
}

func (this *DatabaseTable) AssembleOrderByClause(orderby OrderByClause) string {
	if orderby == nil || len(orderby) == 0 {
		return ""
	}
	parts := []string{}
	for column, orientation := range orderby {
		parts = append(parts, fmt.Sprintf("%s %s", column, orientation.String()))
	}
	return fmt.Sprintf("ORDER BY %s", strings.Join(parts, ", "))
}

/*
Create a default insert query - assumes that all fields except autoincrement are
to be set.
*/
func (this *DatabaseTable) CreateInsertQuery() ([]string, string) {
	columns := this.GetColumns(func(c *DatabaseTableColumn) bool {
		return !c.AutoIncrement
	})

	placeholders := this.CreatePlaceholders(len(columns))

	return columns, fmt.Sprintf(
		"INSERT OR REPLACE INTO %s (%s) VALUES (%s)",
		this.Name, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
}

/*
Create a select query
*/
func (this *DatabaseTable) CreateSelectQuery(columns []string, where WhereClause, orderby OrderByClause) ([]string, []interface{}, string) {
	if columns == nil {
		columns = this.GetColumns(func(c *DatabaseTableColumn) bool { return true })
	}

	whereMapping, whereString := this.AssembleWhereClause(where)
	orderbyString := this.AssembleOrderByClause(orderby)

	return columns, whereMapping, fmt.Sprintf(
		"SELECT %s FROM %s %s %s",
		strings.Join(columns, ", "), this.Name, whereString, orderbyString)
}

/*
Create a delete query
*/
func (this *DatabaseTable) CreateDeleteQuery(where WhereClause) ([]interface{}, string) {
	whereMapping, whereString := this.AssembleWhereClause(where)
	return whereMapping, fmt.Sprintf(
		"DELETE FROM %s %s",
		this.Name, whereString)
}
