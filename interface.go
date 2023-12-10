package gen

import (
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen/field"
)

type (
	// Condition query condition
	// field.Expr and subquery are expect value
	Condition interface {
		BeCond() interface{}
		CondError() error
	}
)

var (
	_ Condition = (field.Expr)(nil)
	_ Condition = (field.Value)(nil)
	_ Condition = (SubQuery)(nil)
	_ Condition = (Dao)(nil)
)

// SubQuery sub query interface
type SubQuery interface {
	underlyingDB() *gorm.DB
	underlyingDO() *DO

	Condition
}

// Dao CRUD methods
type Dao interface {
	SubQuery
	schema.Tabler
	As(alias string) Dao

	Not(conds ...Condition) Dao
	Or(conds ...Condition) Dao

	Select(columns ...field.Expr) Dao
	Where(conds ...Condition) Dao
	Order(columns ...field.Expr) Dao
	Distinct(columns ...field.Expr) Dao
	Omit(columns ...field.Expr) Dao
	Join(table schema.Tabler, conds ...Condition) Dao
	LeftJoin(table schema.Tabler, conds ...Condition) Dao
	RightJoin(table schema.Tabler, conds ...Condition) Dao
	OuterJoin(table schema.Tabler, conds ...Condition) Dao
	Group(columns ...field.Expr) Dao
	Having(conds ...Condition) Dao
	Limit(limit int) Dao
	Offset(offset int) Dao
	Scopes(funcs ...func(Dao) Dao) Dao
	Unscoped() Dao
	Attrs(attrs ...field.AssignExpr) Dao
	Assign(attrs ...field.AssignExpr) Dao
	Joins(field field.RelationField) Dao
	Preload(field field.RelationField) Dao
	Clauses(conds ...clause.Expression) Dao
	WithCTE(alias string, terms ...SubQuery) Dao
	WithRecursiveCTE(alias string, terms ...SubQuery) Dao

	Create(value interface{}) error
	CreateInBatches(value interface{}, batchSize int) error
	Save(value interface{}) error
	First() (result interface{}, err error)
	Take() (result interface{}, err error)
	Last() (result interface{}, err error)
	Find() (results interface{}, err error)
	FindInBatches(dest interface{}, batchSize int, fc func(tx Dao, batch int) error) error
	FirstOrInit() (result interface{}, err error)
	FirstOrCreate() (result interface{}, err error)
	Update(column field.Expr, value interface{}) ResultInfo
	UpdateSimple(columns ...field.AssignExpr) ResultInfo
	Updates(values interface{}) ResultInfo
	UpdateColumn(column field.Expr, value interface{}) ResultInfo
	UpdateColumns(values interface{}) ResultInfo
	UpdateColumnSimple(columns ...field.AssignExpr) ResultInfo
	Delete(...interface{}) ResultInfo
	Count() (int64, error)
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	Scan(dest interface{}) error
	Pluck(column field.Expr, dest interface{}) error
	ScanRows(rows *sql.Rows, dest interface{}) error
	InsertInto(table schema.Tabler) ResultInfo
	FromValues(alias string, columns []string, values [][]interface{}) Dao
	FromValuesSimple(alias string, dest interface{}) Dao
	OnUniqueConflict(model interface{}, updates []OnConflictUpdate) Dao
	OnUniqueConflictDoNothing(model interface{}) Dao

	AddError(err error) error
}
