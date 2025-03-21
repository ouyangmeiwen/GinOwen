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

func newHangfirestate(db *gorm.DB) hangfirestate {
	_hangfirestate := hangfirestate{}

	_hangfirestate.hangfirestateDo.UseDB(db)
	_hangfirestate.hangfirestateDo.UseModel(&model.Hangfirestate{})

	tableName := _hangfirestate.hangfirestateDo.TableName()
	_hangfirestate.ALL = field.NewAsterisk(tableName)
	_hangfirestate.ID = field.NewInt64(tableName, "Id")
	_hangfirestate.JobID = field.NewInt64(tableName, "JobId")
	_hangfirestate.Name = field.NewString(tableName, "Name")
	_hangfirestate.Reason = field.NewString(tableName, "Reason")
	_hangfirestate.CreatedAt = field.NewTime(tableName, "CreatedAt")
	_hangfirestate.Data = field.NewString(tableName, "Data")

	_hangfirestate.fillFieldMap()

	return _hangfirestate
}

type hangfirestate struct {
	hangfirestateDo hangfirestateDo

	ALL       field.Asterisk
	ID        field.Int64
	JobID     field.Int64
	Name      field.String
	Reason    field.String
	CreatedAt field.Time
	Data      field.String

	fieldMap map[string]field.Expr
}

func (h hangfirestate) Table(newTableName string) *hangfirestate {
	h.hangfirestateDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hangfirestate) As(alias string) *hangfirestate {
	h.hangfirestateDo.DO = *(h.hangfirestateDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hangfirestate) updateTableName(table string) *hangfirestate {
	h.ALL = field.NewAsterisk(table)
	h.ID = field.NewInt64(table, "Id")
	h.JobID = field.NewInt64(table, "JobId")
	h.Name = field.NewString(table, "Name")
	h.Reason = field.NewString(table, "Reason")
	h.CreatedAt = field.NewTime(table, "CreatedAt")
	h.Data = field.NewString(table, "Data")

	h.fillFieldMap()

	return h
}

func (h *hangfirestate) WithContext(ctx context.Context) IHangfirestateDo {
	return h.hangfirestateDo.WithContext(ctx)
}

func (h hangfirestate) TableName() string { return h.hangfirestateDo.TableName() }

func (h hangfirestate) Alias() string { return h.hangfirestateDo.Alias() }

func (h *hangfirestate) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hangfirestate) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 6)
	h.fieldMap["Id"] = h.ID
	h.fieldMap["JobId"] = h.JobID
	h.fieldMap["Name"] = h.Name
	h.fieldMap["Reason"] = h.Reason
	h.fieldMap["CreatedAt"] = h.CreatedAt
	h.fieldMap["Data"] = h.Data
}

func (h hangfirestate) clone(db *gorm.DB) hangfirestate {
	h.hangfirestateDo.ReplaceDB(db)
	return h
}

type hangfirestateDo struct{ gen.DO }

type IHangfirestateDo interface {
	gen.SubQuery
	Debug() IHangfirestateDo
	WithContext(ctx context.Context) IHangfirestateDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHangfirestateDo
	Not(conds ...gen.Condition) IHangfirestateDo
	Or(conds ...gen.Condition) IHangfirestateDo
	Select(conds ...field.Expr) IHangfirestateDo
	Where(conds ...gen.Condition) IHangfirestateDo
	Order(conds ...field.Expr) IHangfirestateDo
	Distinct(cols ...field.Expr) IHangfirestateDo
	Omit(cols ...field.Expr) IHangfirestateDo
	Join(table schema.Tabler, on ...field.Expr) IHangfirestateDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHangfirestateDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHangfirestateDo
	Group(cols ...field.Expr) IHangfirestateDo
	Having(conds ...gen.Condition) IHangfirestateDo
	Limit(limit int) IHangfirestateDo
	Offset(offset int) IHangfirestateDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHangfirestateDo
	Unscoped() IHangfirestateDo
	Create(values ...*model.Hangfirestate) error
	CreateInBatches(values []*model.Hangfirestate, batchSize int) error
	Save(values ...*model.Hangfirestate) error
	First() (*model.Hangfirestate, error)
	Take() (*model.Hangfirestate, error)
	Last() (*model.Hangfirestate, error)
	Find() ([]*model.Hangfirestate, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Hangfirestate, err error)
	FindInBatches(result *[]*model.Hangfirestate, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Hangfirestate) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHangfirestateDo
	Assign(attrs ...field.AssignExpr) IHangfirestateDo
	Joins(fields ...field.RelationField) IHangfirestateDo
	Preload(fields ...field.RelationField) IHangfirestateDo
	FirstOrInit() (*model.Hangfirestate, error)
	FirstOrCreate() (*model.Hangfirestate, error)
	FindByPage(offset int, limit int) (result []*model.Hangfirestate, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHangfirestateDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hangfirestateDo) Debug() IHangfirestateDo {
	return h.withDO(h.DO.Debug())
}

func (h hangfirestateDo) WithContext(ctx context.Context) IHangfirestateDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hangfirestateDo) ReadDB() IHangfirestateDo {
	return h.Clauses(dbresolver.Read)
}

func (h hangfirestateDo) WriteDB() IHangfirestateDo {
	return h.Clauses(dbresolver.Write)
}

func (h hangfirestateDo) Clauses(conds ...clause.Expression) IHangfirestateDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hangfirestateDo) Returning(value interface{}, columns ...string) IHangfirestateDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hangfirestateDo) Not(conds ...gen.Condition) IHangfirestateDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hangfirestateDo) Or(conds ...gen.Condition) IHangfirestateDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hangfirestateDo) Select(conds ...field.Expr) IHangfirestateDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hangfirestateDo) Where(conds ...gen.Condition) IHangfirestateDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hangfirestateDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IHangfirestateDo {
	return h.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (h hangfirestateDo) Order(conds ...field.Expr) IHangfirestateDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hangfirestateDo) Distinct(cols ...field.Expr) IHangfirestateDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hangfirestateDo) Omit(cols ...field.Expr) IHangfirestateDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hangfirestateDo) Join(table schema.Tabler, on ...field.Expr) IHangfirestateDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hangfirestateDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHangfirestateDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hangfirestateDo) RightJoin(table schema.Tabler, on ...field.Expr) IHangfirestateDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hangfirestateDo) Group(cols ...field.Expr) IHangfirestateDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hangfirestateDo) Having(conds ...gen.Condition) IHangfirestateDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hangfirestateDo) Limit(limit int) IHangfirestateDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hangfirestateDo) Offset(offset int) IHangfirestateDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hangfirestateDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHangfirestateDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hangfirestateDo) Unscoped() IHangfirestateDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hangfirestateDo) Create(values ...*model.Hangfirestate) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hangfirestateDo) CreateInBatches(values []*model.Hangfirestate, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hangfirestateDo) Save(values ...*model.Hangfirestate) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hangfirestateDo) First() (*model.Hangfirestate, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Hangfirestate), nil
	}
}

func (h hangfirestateDo) Take() (*model.Hangfirestate, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Hangfirestate), nil
	}
}

func (h hangfirestateDo) Last() (*model.Hangfirestate, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Hangfirestate), nil
	}
}

func (h hangfirestateDo) Find() ([]*model.Hangfirestate, error) {
	result, err := h.DO.Find()
	return result.([]*model.Hangfirestate), err
}

func (h hangfirestateDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Hangfirestate, err error) {
	buf := make([]*model.Hangfirestate, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hangfirestateDo) FindInBatches(result *[]*model.Hangfirestate, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hangfirestateDo) Attrs(attrs ...field.AssignExpr) IHangfirestateDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hangfirestateDo) Assign(attrs ...field.AssignExpr) IHangfirestateDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hangfirestateDo) Joins(fields ...field.RelationField) IHangfirestateDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hangfirestateDo) Preload(fields ...field.RelationField) IHangfirestateDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hangfirestateDo) FirstOrInit() (*model.Hangfirestate, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Hangfirestate), nil
	}
}

func (h hangfirestateDo) FirstOrCreate() (*model.Hangfirestate, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Hangfirestate), nil
	}
}

func (h hangfirestateDo) FindByPage(offset int, limit int) (result []*model.Hangfirestate, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h hangfirestateDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hangfirestateDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hangfirestateDo) Delete(models ...*model.Hangfirestate) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hangfirestateDo) withDO(do gen.Dao) *hangfirestateDo {
	h.DO = *do.(*gen.DO)
	return h
}
