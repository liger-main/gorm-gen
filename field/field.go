package field

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"gorm.io/gorm/clause"
)

// ScanValuer interface for Field
type ScanValuer interface {
	Scan(src interface{}) error   // sql.Scanner
	Value() (driver.Value, error) // driver.Valuer
}

// Field a standard field struct
type Field struct{ expr }

// Eq judge equal
func (field Field) Eq(value driver.Valuer) Expr {
	return expr{e: clause.Eq{Column: field.RawExpr(), Value: value}}
}

// Neq judge not equal
func (field Field) Neq(value driver.Valuer) Expr {
	return expr{e: clause.Neq{Column: field.RawExpr(), Value: value}}
}

// In ...
func (field Field) In(values ...driver.Valuer) Expr {
	return expr{e: clause.IN{Column: field.RawExpr(), Values: field.toSlice(values...)}}
}

// Gt ...
func (field Field) Gt(value driver.Valuer) Expr {
	return expr{e: clause.Gt{Column: field.RawExpr(), Value: value}}
}

// Gte ...
func (field Field) Gte(value driver.Valuer) Expr {
	return expr{e: clause.Gte{Column: field.RawExpr(), Value: value}}
}

// Lt ...
func (field Field) Lt(value driver.Valuer) Expr {
	return expr{e: clause.Lt{Column: field.RawExpr(), Value: value}}
}

// Lte ...
func (field Field) Lte(value driver.Valuer) Expr {
	return expr{e: clause.Lte{Column: field.RawExpr(), Value: value}}
}

// Like ...
func (field Field) Like(value driver.Valuer) Expr {
	return expr{e: clause.Like{Column: field.RawExpr(), Value: value}}
}

// Value ...
func (field Field) Value(value driver.Valuer) AssignExpr {
	return field.value(value)
}

// Sum ...
func (field Field) Sum() Field {
	return Field{field.sum()}
}

// IfNull ...
func (field Field) IfNull(value driver.Valuer) Expr {
	return field.ifNull(value)
}

func (field Field) IsDistinctFrom(other Expr) Expr {
	return field.isDistinctFrom(other)
}

func (field Field) MD5() Expr {
	return String{expr{e: clause.Expr{
		SQL:  "MD5(?)",
		Vars: []interface{}{field.RawExpr()},
	}}}
}

func (field Field) ArrayAppend(other Expr) Expr {
	return String{expr{e: clause.Expr{
		SQL:  "ARRAY_APPEND(?, ?)",
		Vars: []interface{}{field.RawExpr(), other.RawExpr()},
	}}}
}

func (field Field) ArrayRemove(values ...string) AssignExpr {
	if len(values) == 0 {
		return field
	}
	valuesPlaceholder := strings.TrimSuffix(strings.Repeat("?,", len(values)), ",")
	vars := make([]interface{}, len(values)+1)
	vars[0] = field.RawExpr()
	for i, v := range values {
		vars[i+1] = v
	}
	return field.setE(clause.Eq{Column: field.col.Name, Value: clause.Expr{
		SQL:  fmt.Sprintf("ARRAY(SELECT UNNEST(?) EXCEPT SELECT UNNEST(ARRAY[%s]))", valuesPlaceholder),
		Vars: vars,
	}})
}

// Field ...
func (field Field) Field(value []interface{}) Expr {
	return field.field(value)
}

func (field Field) toSlice(values ...driver.Valuer) []interface{} {
	slice := make([]interface{}, len(values))
	for i, v := range values {
		slice[i] = v
	}
	return slice
}
