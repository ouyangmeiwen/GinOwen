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

func newDassecuritygatecount(db *gorm.DB) dassecuritygatecount {
	_dassecuritygatecount := dassecuritygatecount{}

	_dassecuritygatecount.dassecuritygatecountDo.UseDB(db)
	_dassecuritygatecount.dassecuritygatecountDo.UseModel(&model.Dassecuritygatecount{})

	tableName := _dassecuritygatecount.dassecuritygatecountDo.TableName()
	_dassecuritygatecount.ALL = field.NewAsterisk(tableName)
	_dassecuritygatecount.ID = field.NewString(tableName, "Id")
	_dassecuritygatecount.CreationTime = field.NewTime(tableName, "CreationTime")
	_dassecuritygatecount.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_dassecuritygatecount.StartTime = field.NewTime(tableName, "StartTime")
	_dassecuritygatecount.EndTime = field.NewTime(tableName, "EndTime")
	_dassecuritygatecount.TerminalID = field.NewString(tableName, "TerminalId")
	_dassecuritygatecount.TerminalCode = field.NewString(tableName, "TerminalCode")
	_dassecuritygatecount.TerminalName = field.NewString(tableName, "TerminalName")
	_dassecuritygatecount.TotalInCount = field.NewInt64(tableName, "TotalInCount")
	_dassecuritygatecount.TotalOutCount = field.NewInt64(tableName, "TotalOutCount")
	_dassecuritygatecount.TenantID = field.NewInt64(tableName, "TenantId")

	_dassecuritygatecount.fillFieldMap()

	return _dassecuritygatecount
}

type dassecuritygatecount struct {
	dassecuritygatecountDo dassecuritygatecountDo

	ALL           field.Asterisk
	ID            field.String
	CreationTime  field.Time
	CreatorUserID field.Int64
	StartTime     field.Time
	EndTime       field.Time
	TerminalID    field.String
	TerminalCode  field.String
	TerminalName  field.String
	TotalInCount  field.Int64
	TotalOutCount field.Int64
	TenantID      field.Int64

	fieldMap map[string]field.Expr
}

func (d dassecuritygatecount) Table(newTableName string) *dassecuritygatecount {
	d.dassecuritygatecountDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dassecuritygatecount) As(alias string) *dassecuritygatecount {
	d.dassecuritygatecountDo.DO = *(d.dassecuritygatecountDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dassecuritygatecount) updateTableName(table string) *dassecuritygatecount {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewString(table, "Id")
	d.CreationTime = field.NewTime(table, "CreationTime")
	d.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	d.StartTime = field.NewTime(table, "StartTime")
	d.EndTime = field.NewTime(table, "EndTime")
	d.TerminalID = field.NewString(table, "TerminalId")
	d.TerminalCode = field.NewString(table, "TerminalCode")
	d.TerminalName = field.NewString(table, "TerminalName")
	d.TotalInCount = field.NewInt64(table, "TotalInCount")
	d.TotalOutCount = field.NewInt64(table, "TotalOutCount")
	d.TenantID = field.NewInt64(table, "TenantId")

	d.fillFieldMap()

	return d
}

func (d *dassecuritygatecount) WithContext(ctx context.Context) IDassecuritygatecountDo {
	return d.dassecuritygatecountDo.WithContext(ctx)
}

func (d dassecuritygatecount) TableName() string { return d.dassecuritygatecountDo.TableName() }

func (d dassecuritygatecount) Alias() string { return d.dassecuritygatecountDo.Alias() }

func (d *dassecuritygatecount) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dassecuritygatecount) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 11)
	d.fieldMap["Id"] = d.ID
	d.fieldMap["CreationTime"] = d.CreationTime
	d.fieldMap["CreatorUserId"] = d.CreatorUserID
	d.fieldMap["StartTime"] = d.StartTime
	d.fieldMap["EndTime"] = d.EndTime
	d.fieldMap["TerminalId"] = d.TerminalID
	d.fieldMap["TerminalCode"] = d.TerminalCode
	d.fieldMap["TerminalName"] = d.TerminalName
	d.fieldMap["TotalInCount"] = d.TotalInCount
	d.fieldMap["TotalOutCount"] = d.TotalOutCount
	d.fieldMap["TenantId"] = d.TenantID
}

func (d dassecuritygatecount) clone(db *gorm.DB) dassecuritygatecount {
	d.dassecuritygatecountDo.ReplaceDB(db)
	return d
}

type dassecuritygatecountDo struct{ gen.DO }

type IDassecuritygatecountDo interface {
	gen.SubQuery
	Debug() IDassecuritygatecountDo
	WithContext(ctx context.Context) IDassecuritygatecountDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDassecuritygatecountDo
	Not(conds ...gen.Condition) IDassecuritygatecountDo
	Or(conds ...gen.Condition) IDassecuritygatecountDo
	Select(conds ...field.Expr) IDassecuritygatecountDo
	Where(conds ...gen.Condition) IDassecuritygatecountDo
	Order(conds ...field.Expr) IDassecuritygatecountDo
	Distinct(cols ...field.Expr) IDassecuritygatecountDo
	Omit(cols ...field.Expr) IDassecuritygatecountDo
	Join(table schema.Tabler, on ...field.Expr) IDassecuritygatecountDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDassecuritygatecountDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDassecuritygatecountDo
	Group(cols ...field.Expr) IDassecuritygatecountDo
	Having(conds ...gen.Condition) IDassecuritygatecountDo
	Limit(limit int) IDassecuritygatecountDo
	Offset(offset int) IDassecuritygatecountDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDassecuritygatecountDo
	Unscoped() IDassecuritygatecountDo
	Create(values ...*model.Dassecuritygatecount) error
	CreateInBatches(values []*model.Dassecuritygatecount, batchSize int) error
	Save(values ...*model.Dassecuritygatecount) error
	First() (*model.Dassecuritygatecount, error)
	Take() (*model.Dassecuritygatecount, error)
	Last() (*model.Dassecuritygatecount, error)
	Find() ([]*model.Dassecuritygatecount, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Dassecuritygatecount, err error)
	FindInBatches(result *[]*model.Dassecuritygatecount, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Dassecuritygatecount) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDassecuritygatecountDo
	Assign(attrs ...field.AssignExpr) IDassecuritygatecountDo
	Joins(fields ...field.RelationField) IDassecuritygatecountDo
	Preload(fields ...field.RelationField) IDassecuritygatecountDo
	FirstOrInit() (*model.Dassecuritygatecount, error)
	FirstOrCreate() (*model.Dassecuritygatecount, error)
	FindByPage(offset int, limit int) (result []*model.Dassecuritygatecount, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDassecuritygatecountDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dassecuritygatecountDo) Debug() IDassecuritygatecountDo {
	return d.withDO(d.DO.Debug())
}

func (d dassecuritygatecountDo) WithContext(ctx context.Context) IDassecuritygatecountDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dassecuritygatecountDo) ReadDB() IDassecuritygatecountDo {
	return d.Clauses(dbresolver.Read)
}

func (d dassecuritygatecountDo) WriteDB() IDassecuritygatecountDo {
	return d.Clauses(dbresolver.Write)
}

func (d dassecuritygatecountDo) Clauses(conds ...clause.Expression) IDassecuritygatecountDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dassecuritygatecountDo) Returning(value interface{}, columns ...string) IDassecuritygatecountDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dassecuritygatecountDo) Not(conds ...gen.Condition) IDassecuritygatecountDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dassecuritygatecountDo) Or(conds ...gen.Condition) IDassecuritygatecountDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dassecuritygatecountDo) Select(conds ...field.Expr) IDassecuritygatecountDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dassecuritygatecountDo) Where(conds ...gen.Condition) IDassecuritygatecountDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dassecuritygatecountDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDassecuritygatecountDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dassecuritygatecountDo) Order(conds ...field.Expr) IDassecuritygatecountDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dassecuritygatecountDo) Distinct(cols ...field.Expr) IDassecuritygatecountDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dassecuritygatecountDo) Omit(cols ...field.Expr) IDassecuritygatecountDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dassecuritygatecountDo) Join(table schema.Tabler, on ...field.Expr) IDassecuritygatecountDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dassecuritygatecountDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDassecuritygatecountDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dassecuritygatecountDo) RightJoin(table schema.Tabler, on ...field.Expr) IDassecuritygatecountDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dassecuritygatecountDo) Group(cols ...field.Expr) IDassecuritygatecountDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dassecuritygatecountDo) Having(conds ...gen.Condition) IDassecuritygatecountDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dassecuritygatecountDo) Limit(limit int) IDassecuritygatecountDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dassecuritygatecountDo) Offset(offset int) IDassecuritygatecountDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dassecuritygatecountDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDassecuritygatecountDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dassecuritygatecountDo) Unscoped() IDassecuritygatecountDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dassecuritygatecountDo) Create(values ...*model.Dassecuritygatecount) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dassecuritygatecountDo) CreateInBatches(values []*model.Dassecuritygatecount, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dassecuritygatecountDo) Save(values ...*model.Dassecuritygatecount) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dassecuritygatecountDo) First() (*model.Dassecuritygatecount, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dassecuritygatecount), nil
	}
}

func (d dassecuritygatecountDo) Take() (*model.Dassecuritygatecount, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dassecuritygatecount), nil
	}
}

func (d dassecuritygatecountDo) Last() (*model.Dassecuritygatecount, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dassecuritygatecount), nil
	}
}

func (d dassecuritygatecountDo) Find() ([]*model.Dassecuritygatecount, error) {
	result, err := d.DO.Find()
	return result.([]*model.Dassecuritygatecount), err
}

func (d dassecuritygatecountDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Dassecuritygatecount, err error) {
	buf := make([]*model.Dassecuritygatecount, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dassecuritygatecountDo) FindInBatches(result *[]*model.Dassecuritygatecount, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dassecuritygatecountDo) Attrs(attrs ...field.AssignExpr) IDassecuritygatecountDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dassecuritygatecountDo) Assign(attrs ...field.AssignExpr) IDassecuritygatecountDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dassecuritygatecountDo) Joins(fields ...field.RelationField) IDassecuritygatecountDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dassecuritygatecountDo) Preload(fields ...field.RelationField) IDassecuritygatecountDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dassecuritygatecountDo) FirstOrInit() (*model.Dassecuritygatecount, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dassecuritygatecount), nil
	}
}

func (d dassecuritygatecountDo) FirstOrCreate() (*model.Dassecuritygatecount, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dassecuritygatecount), nil
	}
}

func (d dassecuritygatecountDo) FindByPage(offset int, limit int) (result []*model.Dassecuritygatecount, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dassecuritygatecountDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dassecuritygatecountDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dassecuritygatecountDo) Delete(models ...*model.Dassecuritygatecount) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dassecuritygatecountDo) withDO(do gen.Dao) *dassecuritygatecountDo {
	d.DO = *do.(*gen.DO)
	return d
}
