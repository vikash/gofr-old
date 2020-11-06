package gofr

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"regexp"
)

type DB struct {
	*sql.DB
}

func (d *DB) Select(ctx context.Context, data interface{}, query string, args ...interface{}) {
	// If context is done, it is not needed
	if ctx.Err() != nil {
		return
	}

	// First confirm that what we got in v is a pointer else it won't be settable
	rvo := reflect.ValueOf(data)
	if rvo.Kind() != reflect.Ptr {
		fmt.Println("We did not get a pointer. data is not settable.")
		return
	}

	// Deference the pointer to the underlying element, if underlying element is a slice, multiple rows are expected.
	// If underlying element is a struct, one row is expected.
	rv := rvo.Elem()

	switch rv.Kind() {
	case reflect.Slice:
		rows, err := d.QueryContext(ctx, query, args...)
		if err != nil {
			fmt.Println(err)
			return
		}

		for rows.Next() {
			val := reflect.New(rv.Type().Elem())
			err = rows.Scan(val.Interface())
			rv = reflect.Append(rv,val.Elem())
		}

		if rvo.Elem().CanSet() {
			rvo.Elem().Set(rv)
		}

		fmt.Println(rvo)
	case reflect.Struct:
		rows, _ := d.QueryContext(ctx, query, args...)

		// Map fields and their indexes by normalised name
		fieldNameIndex := map[string]int{}
		for i := 0; i < rv.Type().NumField(); i++ {
			var name string
			f := rv.Type().Field(i)
			tag := f.Tag.Get("db")
			if tag != "" {
				name = tag
			} else {
				name = ToSnakeCase(f.Name)
			}
			fieldNameIndex[name] = i
		}

		fields := []interface{}{}
		columns, _ := rows.Columns()
		for _, c := range columns {
			if i, ok := fieldNameIndex[c]; ok {
				fields = append(fields, rv.Field(i).Addr().Interface())
			} else {
				var i interface{}
				fields = append(fields, &i)
			}
		}

		for rows.Next() {
			rows.Scan(fields...)
		}

	default:
		fmt.Println("a pointer to", rv.Kind(), "was not expected.")
	}

}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
