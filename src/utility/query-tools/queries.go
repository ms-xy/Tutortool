package querytools

import (
	"database/sql"
	"reflect"
)

/*
Run an insert query on the provided database transaction
*/
func Insert(fm *FieldMap, tx *sql.Tx, d interface{}) (sql.Result, error) {
	i_d := reflect.ValueOf(d)
	columns, query := fm.DatabaseTable.CreateInsertQuery()
	columnMapping := fm.ApplyMapping(fm.CreateMapping(columns), i_d, false)
	return tx.Exec(query, columnMapping...)
}

/*
Run a select query on a provided database transaction
*/
func Select(fm *FieldMap, tx *sql.Tx, columns []string, where WhereClause, orderby OrderByClause) (*sql.Rows, error) {
	_, whereMapping, query := fm.DatabaseTable.CreateSelectQuery(columns, where, orderby)
	return tx.Query(query, whereMapping...)
}

/*
Run a delete query on the provided database transaction
*/
func Delete(fm *FieldMap, tx *sql.Tx, where WhereClause) (sql.Result, error) {
	whereMapping, query := fm.DatabaseTable.CreateDeleteQuery(where)
	return tx.Exec(query, whereMapping...)
}

/*
Fetch all entries of a *sql.Rows object
*/
func FetchAll(fm *FieldMap, rows *sql.Rows) (interface{}, error) {
	columns, err := rows.Columns()
	mapping := fm.CreateMapping(columns)
	if err != nil {
		return nil, err
	}
	ptr := reflect.New(reflect.SliceOf(fm.Type)).Interface()
	slice := reflect.ValueOf(ptr).Elem()
	for rows.Next() {
		d := fm.newInstance()
		err = rows.Scan(fm.ApplyMapping(mapping, d, true)...)
		if err != nil {
			return nil, err
		}
		slice = reflect.Append(slice, d)
	}
	return slice.Interface(), nil
}

/*
Fetch the first entry of a *sql.Rows object
*/
func FetchOne(fm *FieldMap, rows *sql.Rows) (result interface{}, err error) {
	columns, err := rows.Columns()
	mapping := fm.CreateMapping(columns)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		d := fm.newInstance()
		err = rows.Scan(fm.ApplyMapping(mapping, d, true)...)
		result = d.Interface()
	}
	return
}
