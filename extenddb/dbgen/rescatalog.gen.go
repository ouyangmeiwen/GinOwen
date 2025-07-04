// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dbgen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"GINOWEN/extenddb/model"
)

func newRescatalog(db *gorm.DB) rescatalog {
	_rescatalog := rescatalog{}

	_rescatalog.rescatalogDo.UseDB(db)
	_rescatalog.rescatalogDo.UseModel(&model.Rescatalog{})

	tableName := _rescatalog.rescatalogDo.TableName()
	_rescatalog.ALL = field.NewAsterisk(tableName)
	_rescatalog.ID = field.NewInt64(tableName, "Id")
	_rescatalog.Code = field.NewString(tableName, "Code")
	_rescatalog.Name = field.NewString(tableName, "Name")
	_rescatalog.Level = field.NewInt64(tableName, "Level")
	_rescatalog.ParentCode = field.NewString(tableName, "ParentCode")

	_rescatalog.fillFieldMap()

	return _rescatalog
}

type rescatalog struct {
	rescatalogDo rescatalogDo

	ALL        field.Asterisk
	ID         field.Int64
	Code       field.String
	Name       field.String
	Level      field.Int64
	ParentCode field.String

	fieldMap map[string]field.Expr
}

func (r rescatalog) Table(newTableName string) *rescatalog {
	r.rescatalogDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r rescatalog) As(alias string) *rescatalog {
	r.rescatalogDo.DO = *(r.rescatalogDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *rescatalog) updateTableName(table string) *rescatalog {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "Id")
	r.Code = field.NewString(table, "Code")
	r.Name = field.NewString(table, "Name")
	r.Level = field.NewInt64(table, "Level")
	r.ParentCode = field.NewString(table, "ParentCode")

	r.fillFieldMap()

	return r
}

func (r *rescatalog) WithContext(ctx context.Context) IRescatalogDo {
	return r.rescatalogDo.WithContext(ctx)
}

func (r rescatalog) TableName() string { return r.rescatalogDo.TableName() }

func (r rescatalog) Alias() string { return r.rescatalogDo.Alias() }

func (r *rescatalog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *rescatalog) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 5)
	r.fieldMap["Id"] = r.ID
	r.fieldMap["Code"] = r.Code
	r.fieldMap["Name"] = r.Name
	r.fieldMap["Level"] = r.Level
	r.fieldMap["ParentCode"] = r.ParentCode
}

func (r rescatalog) clone(db *gorm.DB) rescatalog {
	r.rescatalogDo.ReplaceDB(db)
	return r
}

type rescatalogDo struct{ gen.DO }

type IRescatalogDo interface {
	gen.SubQuery
	Debug() IRescatalogDo
	WithContext(ctx context.Context) IRescatalogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRescatalogDo
	Not(conds ...gen.Condition) IRescatalogDo
	Or(conds ...gen.Condition) IRescatalogDo
	Select(conds ...field.Expr) IRescatalogDo
	Where(conds ...gen.Condition) IRescatalogDo
	Order(conds ...field.Expr) IRescatalogDo
	Distinct(cols ...field.Expr) IRescatalogDo
	Omit(cols ...field.Expr) IRescatalogDo
	Join(table schema.Tabler, on ...field.Expr) IRescatalogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRescatalogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRescatalogDo
	Group(cols ...field.Expr) IRescatalogDo
	Having(conds ...gen.Condition) IRescatalogDo
	Limit(limit int) IRescatalogDo
	Offset(offset int) IRescatalogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRescatalogDo
	Unscoped() IRescatalogDo
	Create(values ...*model.Rescatalog) error
	CreateInBatches(values []*model.Rescatalog, batchSize int) error
	Save(values ...*model.Rescatalog) error
	First() (*model.Rescatalog, error)
	Take() (*model.Rescatalog, error)
	Last() (*model.Rescatalog, error)
	Find() ([]*model.Rescatalog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Rescatalog, err error)
	FindInBatches(result *[]*model.Rescatalog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Rescatalog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRescatalogDo
	Assign(attrs ...field.AssignExpr) IRescatalogDo
	Joins(fields ...field.RelationField) IRescatalogDo
	Preload(fields ...field.RelationField) IRescatalogDo
	FirstOrInit() (*model.Rescatalog, error)
	FirstOrCreate() (*model.Rescatalog, error)
	FindByPage(offset int, limit int) (result []*model.Rescatalog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRescatalogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r rescatalogDo) Debug() IRescatalogDo {
	return r.withDO(r.DO.Debug())
}

func (r rescatalogDo) WithContext(ctx context.Context) IRescatalogDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r rescatalogDo) ReadDB() IRescatalogDo {
	return r.Clauses(dbresolver.Read)
}

func (r rescatalogDo) WriteDB() IRescatalogDo {
	return r.Clauses(dbresolver.Write)
}

func (r rescatalogDo) Clauses(conds ...clause.Expression) IRescatalogDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r rescatalogDo) Returning(value interface{}, columns ...string) IRescatalogDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r rescatalogDo) Not(conds ...gen.Condition) IRescatalogDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r rescatalogDo) Or(conds ...gen.Condition) IRescatalogDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r rescatalogDo) Select(conds ...field.Expr) IRescatalogDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r rescatalogDo) Where(conds ...gen.Condition) IRescatalogDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r rescatalogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IRescatalogDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r rescatalogDo) Order(conds ...field.Expr) IRescatalogDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r rescatalogDo) Distinct(cols ...field.Expr) IRescatalogDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r rescatalogDo) Omit(cols ...field.Expr) IRescatalogDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r rescatalogDo) Join(table schema.Tabler, on ...field.Expr) IRescatalogDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r rescatalogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRescatalogDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r rescatalogDo) RightJoin(table schema.Tabler, on ...field.Expr) IRescatalogDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r rescatalogDo) Group(cols ...field.Expr) IRescatalogDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r rescatalogDo) Having(conds ...gen.Condition) IRescatalogDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r rescatalogDo) Limit(limit int) IRescatalogDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r rescatalogDo) Offset(offset int) IRescatalogDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r rescatalogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRescatalogDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r rescatalogDo) Unscoped() IRescatalogDo {
	return r.withDO(r.DO.Unscoped())
}

func (r rescatalogDo) Create(values ...*model.Rescatalog) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r rescatalogDo) CreateInBatches(values []*model.Rescatalog, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r rescatalogDo) Save(values ...*model.Rescatalog) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r rescatalogDo) First() (*model.Rescatalog, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rescatalog), nil
	}
}

func (r rescatalogDo) Take() (*model.Rescatalog, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rescatalog), nil
	}
}

func (r rescatalogDo) Last() (*model.Rescatalog, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rescatalog), nil
	}
}

func (r rescatalogDo) Find() ([]*model.Rescatalog, error) {
	result, err := r.DO.Find()
	return result.([]*model.Rescatalog), err
}

func (r rescatalogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Rescatalog, err error) {
	buf := make([]*model.Rescatalog, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r rescatalogDo) FindInBatches(result *[]*model.Rescatalog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r rescatalogDo) Attrs(attrs ...field.AssignExpr) IRescatalogDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r rescatalogDo) Assign(attrs ...field.AssignExpr) IRescatalogDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r rescatalogDo) Joins(fields ...field.RelationField) IRescatalogDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r rescatalogDo) Preload(fields ...field.RelationField) IRescatalogDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r rescatalogDo) FirstOrInit() (*model.Rescatalog, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rescatalog), nil
	}
}

func (r rescatalogDo) FirstOrCreate() (*model.Rescatalog, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rescatalog), nil
	}
}

func (r rescatalogDo) FindByPage(offset int, limit int) (result []*model.Rescatalog, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r rescatalogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r rescatalogDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r rescatalogDo) Delete(models ...*model.Rescatalog) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *rescatalogDo) withDO(do gen.Dao) *rescatalogDo {
	r.DO = *do.(*gen.DO)
	return r
}
