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

func newAbpauditlog(db *gorm.DB) abpauditlog {
	_abpauditlog := abpauditlog{}

	_abpauditlog.abpauditlogDo.UseDB(db)
	_abpauditlog.abpauditlogDo.UseModel(&model.Abpauditlog{})

	tableName := _abpauditlog.abpauditlogDo.TableName()
	_abpauditlog.ALL = field.NewAsterisk(tableName)
	_abpauditlog.ID = field.NewInt64(tableName, "Id")
	_abpauditlog.TenantID = field.NewInt64(tableName, "TenantId")
	_abpauditlog.UserID = field.NewInt64(tableName, "UserId")
	_abpauditlog.ServiceName = field.NewString(tableName, "ServiceName")
	_abpauditlog.MethodName = field.NewString(tableName, "MethodName")
	_abpauditlog.Parameters = field.NewString(tableName, "Parameters")
	_abpauditlog.ReturnValue = field.NewString(tableName, "ReturnValue")
	_abpauditlog.ExecutionTime = field.NewTime(tableName, "ExecutionTime")
	_abpauditlog.ExecutionDuration = field.NewInt64(tableName, "ExecutionDuration")
	_abpauditlog.ClientIPAddress = field.NewString(tableName, "ClientIpAddress")
	_abpauditlog.ClientName = field.NewString(tableName, "ClientName")
	_abpauditlog.BrowserInfo = field.NewString(tableName, "BrowserInfo")
	_abpauditlog.Exception = field.NewString(tableName, "Exception")
	_abpauditlog.ImpersonatorUserID = field.NewInt64(tableName, "ImpersonatorUserId")
	_abpauditlog.ImpersonatorTenantID = field.NewInt64(tableName, "ImpersonatorTenantId")
	_abpauditlog.CustomData = field.NewString(tableName, "CustomData")

	_abpauditlog.fillFieldMap()

	return _abpauditlog
}

type abpauditlog struct {
	abpauditlogDo abpauditlogDo

	ALL                  field.Asterisk
	ID                   field.Int64
	TenantID             field.Int64
	UserID               field.Int64
	ServiceName          field.String
	MethodName           field.String
	Parameters           field.String
	ReturnValue          field.String
	ExecutionTime        field.Time
	ExecutionDuration    field.Int64
	ClientIPAddress      field.String
	ClientName           field.String
	BrowserInfo          field.String
	Exception            field.String
	ImpersonatorUserID   field.Int64
	ImpersonatorTenantID field.Int64
	CustomData           field.String

	fieldMap map[string]field.Expr
}

func (a abpauditlog) Table(newTableName string) *abpauditlog {
	a.abpauditlogDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a abpauditlog) As(alias string) *abpauditlog {
	a.abpauditlogDo.DO = *(a.abpauditlogDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *abpauditlog) updateTableName(table string) *abpauditlog {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "Id")
	a.TenantID = field.NewInt64(table, "TenantId")
	a.UserID = field.NewInt64(table, "UserId")
	a.ServiceName = field.NewString(table, "ServiceName")
	a.MethodName = field.NewString(table, "MethodName")
	a.Parameters = field.NewString(table, "Parameters")
	a.ReturnValue = field.NewString(table, "ReturnValue")
	a.ExecutionTime = field.NewTime(table, "ExecutionTime")
	a.ExecutionDuration = field.NewInt64(table, "ExecutionDuration")
	a.ClientIPAddress = field.NewString(table, "ClientIpAddress")
	a.ClientName = field.NewString(table, "ClientName")
	a.BrowserInfo = field.NewString(table, "BrowserInfo")
	a.Exception = field.NewString(table, "Exception")
	a.ImpersonatorUserID = field.NewInt64(table, "ImpersonatorUserId")
	a.ImpersonatorTenantID = field.NewInt64(table, "ImpersonatorTenantId")
	a.CustomData = field.NewString(table, "CustomData")

	a.fillFieldMap()

	return a
}

func (a *abpauditlog) WithContext(ctx context.Context) IAbpauditlogDo {
	return a.abpauditlogDo.WithContext(ctx)
}

func (a abpauditlog) TableName() string { return a.abpauditlogDo.TableName() }

func (a abpauditlog) Alias() string { return a.abpauditlogDo.Alias() }

func (a *abpauditlog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *abpauditlog) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 16)
	a.fieldMap["Id"] = a.ID
	a.fieldMap["TenantId"] = a.TenantID
	a.fieldMap["UserId"] = a.UserID
	a.fieldMap["ServiceName"] = a.ServiceName
	a.fieldMap["MethodName"] = a.MethodName
	a.fieldMap["Parameters"] = a.Parameters
	a.fieldMap["ReturnValue"] = a.ReturnValue
	a.fieldMap["ExecutionTime"] = a.ExecutionTime
	a.fieldMap["ExecutionDuration"] = a.ExecutionDuration
	a.fieldMap["ClientIpAddress"] = a.ClientIPAddress
	a.fieldMap["ClientName"] = a.ClientName
	a.fieldMap["BrowserInfo"] = a.BrowserInfo
	a.fieldMap["Exception"] = a.Exception
	a.fieldMap["ImpersonatorUserId"] = a.ImpersonatorUserID
	a.fieldMap["ImpersonatorTenantId"] = a.ImpersonatorTenantID
	a.fieldMap["CustomData"] = a.CustomData
}

func (a abpauditlog) clone(db *gorm.DB) abpauditlog {
	a.abpauditlogDo.ReplaceDB(db)
	return a
}

type abpauditlogDo struct{ gen.DO }

type IAbpauditlogDo interface {
	gen.SubQuery
	Debug() IAbpauditlogDo
	WithContext(ctx context.Context) IAbpauditlogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAbpauditlogDo
	Not(conds ...gen.Condition) IAbpauditlogDo
	Or(conds ...gen.Condition) IAbpauditlogDo
	Select(conds ...field.Expr) IAbpauditlogDo
	Where(conds ...gen.Condition) IAbpauditlogDo
	Order(conds ...field.Expr) IAbpauditlogDo
	Distinct(cols ...field.Expr) IAbpauditlogDo
	Omit(cols ...field.Expr) IAbpauditlogDo
	Join(table schema.Tabler, on ...field.Expr) IAbpauditlogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAbpauditlogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAbpauditlogDo
	Group(cols ...field.Expr) IAbpauditlogDo
	Having(conds ...gen.Condition) IAbpauditlogDo
	Limit(limit int) IAbpauditlogDo
	Offset(offset int) IAbpauditlogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAbpauditlogDo
	Unscoped() IAbpauditlogDo
	Create(values ...*model.Abpauditlog) error
	CreateInBatches(values []*model.Abpauditlog, batchSize int) error
	Save(values ...*model.Abpauditlog) error
	First() (*model.Abpauditlog, error)
	Take() (*model.Abpauditlog, error)
	Last() (*model.Abpauditlog, error)
	Find() ([]*model.Abpauditlog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Abpauditlog, err error)
	FindInBatches(result *[]*model.Abpauditlog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Abpauditlog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAbpauditlogDo
	Assign(attrs ...field.AssignExpr) IAbpauditlogDo
	Joins(fields ...field.RelationField) IAbpauditlogDo
	Preload(fields ...field.RelationField) IAbpauditlogDo
	FirstOrInit() (*model.Abpauditlog, error)
	FirstOrCreate() (*model.Abpauditlog, error)
	FindByPage(offset int, limit int) (result []*model.Abpauditlog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAbpauditlogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a abpauditlogDo) Debug() IAbpauditlogDo {
	return a.withDO(a.DO.Debug())
}

func (a abpauditlogDo) WithContext(ctx context.Context) IAbpauditlogDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a abpauditlogDo) ReadDB() IAbpauditlogDo {
	return a.Clauses(dbresolver.Read)
}

func (a abpauditlogDo) WriteDB() IAbpauditlogDo {
	return a.Clauses(dbresolver.Write)
}

func (a abpauditlogDo) Clauses(conds ...clause.Expression) IAbpauditlogDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a abpauditlogDo) Returning(value interface{}, columns ...string) IAbpauditlogDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a abpauditlogDo) Not(conds ...gen.Condition) IAbpauditlogDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a abpauditlogDo) Or(conds ...gen.Condition) IAbpauditlogDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a abpauditlogDo) Select(conds ...field.Expr) IAbpauditlogDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a abpauditlogDo) Where(conds ...gen.Condition) IAbpauditlogDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a abpauditlogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAbpauditlogDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a abpauditlogDo) Order(conds ...field.Expr) IAbpauditlogDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a abpauditlogDo) Distinct(cols ...field.Expr) IAbpauditlogDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a abpauditlogDo) Omit(cols ...field.Expr) IAbpauditlogDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a abpauditlogDo) Join(table schema.Tabler, on ...field.Expr) IAbpauditlogDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a abpauditlogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAbpauditlogDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a abpauditlogDo) RightJoin(table schema.Tabler, on ...field.Expr) IAbpauditlogDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a abpauditlogDo) Group(cols ...field.Expr) IAbpauditlogDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a abpauditlogDo) Having(conds ...gen.Condition) IAbpauditlogDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a abpauditlogDo) Limit(limit int) IAbpauditlogDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a abpauditlogDo) Offset(offset int) IAbpauditlogDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a abpauditlogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAbpauditlogDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a abpauditlogDo) Unscoped() IAbpauditlogDo {
	return a.withDO(a.DO.Unscoped())
}

func (a abpauditlogDo) Create(values ...*model.Abpauditlog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a abpauditlogDo) CreateInBatches(values []*model.Abpauditlog, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a abpauditlogDo) Save(values ...*model.Abpauditlog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a abpauditlogDo) First() (*model.Abpauditlog, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpauditlog), nil
	}
}

func (a abpauditlogDo) Take() (*model.Abpauditlog, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpauditlog), nil
	}
}

func (a abpauditlogDo) Last() (*model.Abpauditlog, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpauditlog), nil
	}
}

func (a abpauditlogDo) Find() ([]*model.Abpauditlog, error) {
	result, err := a.DO.Find()
	return result.([]*model.Abpauditlog), err
}

func (a abpauditlogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Abpauditlog, err error) {
	buf := make([]*model.Abpauditlog, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a abpauditlogDo) FindInBatches(result *[]*model.Abpauditlog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a abpauditlogDo) Attrs(attrs ...field.AssignExpr) IAbpauditlogDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a abpauditlogDo) Assign(attrs ...field.AssignExpr) IAbpauditlogDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a abpauditlogDo) Joins(fields ...field.RelationField) IAbpauditlogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a abpauditlogDo) Preload(fields ...field.RelationField) IAbpauditlogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a abpauditlogDo) FirstOrInit() (*model.Abpauditlog, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpauditlog), nil
	}
}

func (a abpauditlogDo) FirstOrCreate() (*model.Abpauditlog, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpauditlog), nil
	}
}

func (a abpauditlogDo) FindByPage(offset int, limit int) (result []*model.Abpauditlog, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a abpauditlogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a abpauditlogDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a abpauditlogDo) Delete(models ...*model.Abpauditlog) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *abpauditlogDo) withDO(do gen.Dao) *abpauditlogDo {
	a.DO = *do.(*gen.DO)
	return a
}
