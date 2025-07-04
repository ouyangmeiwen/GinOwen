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

func newSysauditapplog(db *gorm.DB) sysauditapplog {
	_sysauditapplog := sysauditapplog{}

	_sysauditapplog.sysauditapplogDo.UseDB(db)
	_sysauditapplog.sysauditapplogDo.UseModel(&model.Sysauditapplog{})

	tableName := _sysauditapplog.sysauditapplogDo.TableName()
	_sysauditapplog.ALL = field.NewAsterisk(tableName)
	_sysauditapplog.ID = field.NewInt64(tableName, "Id")
	_sysauditapplog.BrowserInfo = field.NewString(tableName, "BrowserInfo")
	_sysauditapplog.ClientIPAddress = field.NewString(tableName, "ClientIpAddress")
	_sysauditapplog.ClientName = field.NewString(tableName, "ClientName")
	_sysauditapplog.CustomData = field.NewString(tableName, "CustomData")
	_sysauditapplog.Exception = field.NewString(tableName, "Exception")
	_sysauditapplog.ExecutionDuration = field.NewInt64(tableName, "ExecutionDuration")
	_sysauditapplog.ExecutionTime = field.NewTime(tableName, "ExecutionTime")
	_sysauditapplog.ImpersonatorTenantID = field.NewInt64(tableName, "ImpersonatorTenantId")
	_sysauditapplog.ImpersonatorUserID = field.NewInt64(tableName, "ImpersonatorUserId")
	_sysauditapplog.MethodName = field.NewString(tableName, "MethodName")
	_sysauditapplog.Parameters = field.NewString(tableName, "Parameters")
	_sysauditapplog.ServiceName = field.NewString(tableName, "ServiceName")
	_sysauditapplog.TenantID = field.NewInt64(tableName, "TenantId")
	_sysauditapplog.UserID = field.NewInt64(tableName, "UserId")
	_sysauditapplog.ReturnValue = field.NewString(tableName, "ReturnValue")

	_sysauditapplog.fillFieldMap()

	return _sysauditapplog
}

type sysauditapplog struct {
	sysauditapplogDo sysauditapplogDo

	ALL                  field.Asterisk
	ID                   field.Int64
	BrowserInfo          field.String
	ClientIPAddress      field.String
	ClientName           field.String
	CustomData           field.String
	Exception            field.String
	ExecutionDuration    field.Int64
	ExecutionTime        field.Time
	ImpersonatorTenantID field.Int64
	ImpersonatorUserID   field.Int64
	MethodName           field.String
	Parameters           field.String
	ServiceName          field.String
	TenantID             field.Int64
	UserID               field.Int64
	ReturnValue          field.String

	fieldMap map[string]field.Expr
}

func (s sysauditapplog) Table(newTableName string) *sysauditapplog {
	s.sysauditapplogDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysauditapplog) As(alias string) *sysauditapplog {
	s.sysauditapplogDo.DO = *(s.sysauditapplogDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysauditapplog) updateTableName(table string) *sysauditapplog {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "Id")
	s.BrowserInfo = field.NewString(table, "BrowserInfo")
	s.ClientIPAddress = field.NewString(table, "ClientIpAddress")
	s.ClientName = field.NewString(table, "ClientName")
	s.CustomData = field.NewString(table, "CustomData")
	s.Exception = field.NewString(table, "Exception")
	s.ExecutionDuration = field.NewInt64(table, "ExecutionDuration")
	s.ExecutionTime = field.NewTime(table, "ExecutionTime")
	s.ImpersonatorTenantID = field.NewInt64(table, "ImpersonatorTenantId")
	s.ImpersonatorUserID = field.NewInt64(table, "ImpersonatorUserId")
	s.MethodName = field.NewString(table, "MethodName")
	s.Parameters = field.NewString(table, "Parameters")
	s.ServiceName = field.NewString(table, "ServiceName")
	s.TenantID = field.NewInt64(table, "TenantId")
	s.UserID = field.NewInt64(table, "UserId")
	s.ReturnValue = field.NewString(table, "ReturnValue")

	s.fillFieldMap()

	return s
}

func (s *sysauditapplog) WithContext(ctx context.Context) ISysauditapplogDo {
	return s.sysauditapplogDo.WithContext(ctx)
}

func (s sysauditapplog) TableName() string { return s.sysauditapplogDo.TableName() }

func (s sysauditapplog) Alias() string { return s.sysauditapplogDo.Alias() }

func (s *sysauditapplog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysauditapplog) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 16)
	s.fieldMap["Id"] = s.ID
	s.fieldMap["BrowserInfo"] = s.BrowserInfo
	s.fieldMap["ClientIpAddress"] = s.ClientIPAddress
	s.fieldMap["ClientName"] = s.ClientName
	s.fieldMap["CustomData"] = s.CustomData
	s.fieldMap["Exception"] = s.Exception
	s.fieldMap["ExecutionDuration"] = s.ExecutionDuration
	s.fieldMap["ExecutionTime"] = s.ExecutionTime
	s.fieldMap["ImpersonatorTenantId"] = s.ImpersonatorTenantID
	s.fieldMap["ImpersonatorUserId"] = s.ImpersonatorUserID
	s.fieldMap["MethodName"] = s.MethodName
	s.fieldMap["Parameters"] = s.Parameters
	s.fieldMap["ServiceName"] = s.ServiceName
	s.fieldMap["TenantId"] = s.TenantID
	s.fieldMap["UserId"] = s.UserID
	s.fieldMap["ReturnValue"] = s.ReturnValue
}

func (s sysauditapplog) clone(db *gorm.DB) sysauditapplog {
	s.sysauditapplogDo.ReplaceDB(db)
	return s
}

type sysauditapplogDo struct{ gen.DO }

type ISysauditapplogDo interface {
	gen.SubQuery
	Debug() ISysauditapplogDo
	WithContext(ctx context.Context) ISysauditapplogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysauditapplogDo
	Not(conds ...gen.Condition) ISysauditapplogDo
	Or(conds ...gen.Condition) ISysauditapplogDo
	Select(conds ...field.Expr) ISysauditapplogDo
	Where(conds ...gen.Condition) ISysauditapplogDo
	Order(conds ...field.Expr) ISysauditapplogDo
	Distinct(cols ...field.Expr) ISysauditapplogDo
	Omit(cols ...field.Expr) ISysauditapplogDo
	Join(table schema.Tabler, on ...field.Expr) ISysauditapplogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysauditapplogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysauditapplogDo
	Group(cols ...field.Expr) ISysauditapplogDo
	Having(conds ...gen.Condition) ISysauditapplogDo
	Limit(limit int) ISysauditapplogDo
	Offset(offset int) ISysauditapplogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysauditapplogDo
	Unscoped() ISysauditapplogDo
	Create(values ...*model.Sysauditapplog) error
	CreateInBatches(values []*model.Sysauditapplog, batchSize int) error
	Save(values ...*model.Sysauditapplog) error
	First() (*model.Sysauditapplog, error)
	Take() (*model.Sysauditapplog, error)
	Last() (*model.Sysauditapplog, error)
	Find() ([]*model.Sysauditapplog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Sysauditapplog, err error)
	FindInBatches(result *[]*model.Sysauditapplog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Sysauditapplog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysauditapplogDo
	Assign(attrs ...field.AssignExpr) ISysauditapplogDo
	Joins(fields ...field.RelationField) ISysauditapplogDo
	Preload(fields ...field.RelationField) ISysauditapplogDo
	FirstOrInit() (*model.Sysauditapplog, error)
	FirstOrCreate() (*model.Sysauditapplog, error)
	FindByPage(offset int, limit int) (result []*model.Sysauditapplog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysauditapplogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysauditapplogDo) Debug() ISysauditapplogDo {
	return s.withDO(s.DO.Debug())
}

func (s sysauditapplogDo) WithContext(ctx context.Context) ISysauditapplogDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysauditapplogDo) ReadDB() ISysauditapplogDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysauditapplogDo) WriteDB() ISysauditapplogDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysauditapplogDo) Clauses(conds ...clause.Expression) ISysauditapplogDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysauditapplogDo) Returning(value interface{}, columns ...string) ISysauditapplogDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysauditapplogDo) Not(conds ...gen.Condition) ISysauditapplogDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysauditapplogDo) Or(conds ...gen.Condition) ISysauditapplogDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysauditapplogDo) Select(conds ...field.Expr) ISysauditapplogDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysauditapplogDo) Where(conds ...gen.Condition) ISysauditapplogDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysauditapplogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysauditapplogDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysauditapplogDo) Order(conds ...field.Expr) ISysauditapplogDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysauditapplogDo) Distinct(cols ...field.Expr) ISysauditapplogDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysauditapplogDo) Omit(cols ...field.Expr) ISysauditapplogDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysauditapplogDo) Join(table schema.Tabler, on ...field.Expr) ISysauditapplogDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysauditapplogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysauditapplogDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysauditapplogDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysauditapplogDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysauditapplogDo) Group(cols ...field.Expr) ISysauditapplogDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysauditapplogDo) Having(conds ...gen.Condition) ISysauditapplogDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysauditapplogDo) Limit(limit int) ISysauditapplogDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysauditapplogDo) Offset(offset int) ISysauditapplogDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysauditapplogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysauditapplogDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysauditapplogDo) Unscoped() ISysauditapplogDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysauditapplogDo) Create(values ...*model.Sysauditapplog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysauditapplogDo) CreateInBatches(values []*model.Sysauditapplog, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysauditapplogDo) Save(values ...*model.Sysauditapplog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysauditapplogDo) First() (*model.Sysauditapplog, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysauditapplog), nil
	}
}

func (s sysauditapplogDo) Take() (*model.Sysauditapplog, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysauditapplog), nil
	}
}

func (s sysauditapplogDo) Last() (*model.Sysauditapplog, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysauditapplog), nil
	}
}

func (s sysauditapplogDo) Find() ([]*model.Sysauditapplog, error) {
	result, err := s.DO.Find()
	return result.([]*model.Sysauditapplog), err
}

func (s sysauditapplogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Sysauditapplog, err error) {
	buf := make([]*model.Sysauditapplog, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysauditapplogDo) FindInBatches(result *[]*model.Sysauditapplog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysauditapplogDo) Attrs(attrs ...field.AssignExpr) ISysauditapplogDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysauditapplogDo) Assign(attrs ...field.AssignExpr) ISysauditapplogDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysauditapplogDo) Joins(fields ...field.RelationField) ISysauditapplogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysauditapplogDo) Preload(fields ...field.RelationField) ISysauditapplogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysauditapplogDo) FirstOrInit() (*model.Sysauditapplog, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysauditapplog), nil
	}
}

func (s sysauditapplogDo) FirstOrCreate() (*model.Sysauditapplog, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysauditapplog), nil
	}
}

func (s sysauditapplogDo) FindByPage(offset int, limit int) (result []*model.Sysauditapplog, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysauditapplogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysauditapplogDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysauditapplogDo) Delete(models ...*model.Sysauditapplog) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysauditapplogDo) withDO(do gen.Dao) *sysauditapplogDo {
	s.DO = *do.(*gen.DO)
	return s
}
