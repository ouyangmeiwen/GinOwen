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

func newAbpnotification(db *gorm.DB) abpnotification {
	_abpnotification := abpnotification{}

	_abpnotification.abpnotificationDo.UseDB(db)
	_abpnotification.abpnotificationDo.UseModel(&model.Abpnotification{})

	tableName := _abpnotification.abpnotificationDo.TableName()
	_abpnotification.ALL = field.NewAsterisk(tableName)
	_abpnotification.ID = field.NewString(tableName, "Id")
	_abpnotification.CreationTime = field.NewTime(tableName, "CreationTime")
	_abpnotification.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_abpnotification.NotificationName = field.NewString(tableName, "NotificationName")
	_abpnotification.Data = field.NewString(tableName, "Data")
	_abpnotification.DataTypeName = field.NewString(tableName, "DataTypeName")
	_abpnotification.EntityTypeName = field.NewString(tableName, "EntityTypeName")
	_abpnotification.EntityTypeAssemblyQualifiedName = field.NewString(tableName, "EntityTypeAssemblyQualifiedName")
	_abpnotification.EntityID = field.NewString(tableName, "EntityId")
	_abpnotification.Severity = field.NewInt64(tableName, "Severity")
	_abpnotification.UserIds = field.NewString(tableName, "UserIds")
	_abpnotification.ExcludedUserIds = field.NewString(tableName, "ExcludedUserIds")
	_abpnotification.TenantIds = field.NewString(tableName, "TenantIds")
	_abpnotification.Discriminator = field.NewString(tableName, "Discriminator")
	_abpnotification.MessageName = field.NewString(tableName, "MessageName")

	_abpnotification.fillFieldMap()

	return _abpnotification
}

type abpnotification struct {
	abpnotificationDo abpnotificationDo

	ALL                             field.Asterisk
	ID                              field.String
	CreationTime                    field.Time
	CreatorUserID                   field.Int64
	NotificationName                field.String
	Data                            field.String
	DataTypeName                    field.String
	EntityTypeName                  field.String
	EntityTypeAssemblyQualifiedName field.String
	EntityID                        field.String
	Severity                        field.Int64
	UserIds                         field.String
	ExcludedUserIds                 field.String
	TenantIds                       field.String
	Discriminator                   field.String
	MessageName                     field.String

	fieldMap map[string]field.Expr
}

func (a abpnotification) Table(newTableName string) *abpnotification {
	a.abpnotificationDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a abpnotification) As(alias string) *abpnotification {
	a.abpnotificationDo.DO = *(a.abpnotificationDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *abpnotification) updateTableName(table string) *abpnotification {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "Id")
	a.CreationTime = field.NewTime(table, "CreationTime")
	a.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	a.NotificationName = field.NewString(table, "NotificationName")
	a.Data = field.NewString(table, "Data")
	a.DataTypeName = field.NewString(table, "DataTypeName")
	a.EntityTypeName = field.NewString(table, "EntityTypeName")
	a.EntityTypeAssemblyQualifiedName = field.NewString(table, "EntityTypeAssemblyQualifiedName")
	a.EntityID = field.NewString(table, "EntityId")
	a.Severity = field.NewInt64(table, "Severity")
	a.UserIds = field.NewString(table, "UserIds")
	a.ExcludedUserIds = field.NewString(table, "ExcludedUserIds")
	a.TenantIds = field.NewString(table, "TenantIds")
	a.Discriminator = field.NewString(table, "Discriminator")
	a.MessageName = field.NewString(table, "MessageName")

	a.fillFieldMap()

	return a
}

func (a *abpnotification) WithContext(ctx context.Context) IAbpnotificationDo {
	return a.abpnotificationDo.WithContext(ctx)
}

func (a abpnotification) TableName() string { return a.abpnotificationDo.TableName() }

func (a abpnotification) Alias() string { return a.abpnotificationDo.Alias() }

func (a *abpnotification) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *abpnotification) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 15)
	a.fieldMap["Id"] = a.ID
	a.fieldMap["CreationTime"] = a.CreationTime
	a.fieldMap["CreatorUserId"] = a.CreatorUserID
	a.fieldMap["NotificationName"] = a.NotificationName
	a.fieldMap["Data"] = a.Data
	a.fieldMap["DataTypeName"] = a.DataTypeName
	a.fieldMap["EntityTypeName"] = a.EntityTypeName
	a.fieldMap["EntityTypeAssemblyQualifiedName"] = a.EntityTypeAssemblyQualifiedName
	a.fieldMap["EntityId"] = a.EntityID
	a.fieldMap["Severity"] = a.Severity
	a.fieldMap["UserIds"] = a.UserIds
	a.fieldMap["ExcludedUserIds"] = a.ExcludedUserIds
	a.fieldMap["TenantIds"] = a.TenantIds
	a.fieldMap["Discriminator"] = a.Discriminator
	a.fieldMap["MessageName"] = a.MessageName
}

func (a abpnotification) clone(db *gorm.DB) abpnotification {
	a.abpnotificationDo.ReplaceDB(db)
	return a
}

type abpnotificationDo struct{ gen.DO }

type IAbpnotificationDo interface {
	gen.SubQuery
	Debug() IAbpnotificationDo
	WithContext(ctx context.Context) IAbpnotificationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAbpnotificationDo
	Not(conds ...gen.Condition) IAbpnotificationDo
	Or(conds ...gen.Condition) IAbpnotificationDo
	Select(conds ...field.Expr) IAbpnotificationDo
	Where(conds ...gen.Condition) IAbpnotificationDo
	Order(conds ...field.Expr) IAbpnotificationDo
	Distinct(cols ...field.Expr) IAbpnotificationDo
	Omit(cols ...field.Expr) IAbpnotificationDo
	Join(table schema.Tabler, on ...field.Expr) IAbpnotificationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAbpnotificationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAbpnotificationDo
	Group(cols ...field.Expr) IAbpnotificationDo
	Having(conds ...gen.Condition) IAbpnotificationDo
	Limit(limit int) IAbpnotificationDo
	Offset(offset int) IAbpnotificationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAbpnotificationDo
	Unscoped() IAbpnotificationDo
	Create(values ...*model.Abpnotification) error
	CreateInBatches(values []*model.Abpnotification, batchSize int) error
	Save(values ...*model.Abpnotification) error
	First() (*model.Abpnotification, error)
	Take() (*model.Abpnotification, error)
	Last() (*model.Abpnotification, error)
	Find() ([]*model.Abpnotification, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Abpnotification, err error)
	FindInBatches(result *[]*model.Abpnotification, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Abpnotification) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAbpnotificationDo
	Assign(attrs ...field.AssignExpr) IAbpnotificationDo
	Joins(fields ...field.RelationField) IAbpnotificationDo
	Preload(fields ...field.RelationField) IAbpnotificationDo
	FirstOrInit() (*model.Abpnotification, error)
	FirstOrCreate() (*model.Abpnotification, error)
	FindByPage(offset int, limit int) (result []*model.Abpnotification, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAbpnotificationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a abpnotificationDo) Debug() IAbpnotificationDo {
	return a.withDO(a.DO.Debug())
}

func (a abpnotificationDo) WithContext(ctx context.Context) IAbpnotificationDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a abpnotificationDo) ReadDB() IAbpnotificationDo {
	return a.Clauses(dbresolver.Read)
}

func (a abpnotificationDo) WriteDB() IAbpnotificationDo {
	return a.Clauses(dbresolver.Write)
}

func (a abpnotificationDo) Clauses(conds ...clause.Expression) IAbpnotificationDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a abpnotificationDo) Returning(value interface{}, columns ...string) IAbpnotificationDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a abpnotificationDo) Not(conds ...gen.Condition) IAbpnotificationDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a abpnotificationDo) Or(conds ...gen.Condition) IAbpnotificationDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a abpnotificationDo) Select(conds ...field.Expr) IAbpnotificationDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a abpnotificationDo) Where(conds ...gen.Condition) IAbpnotificationDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a abpnotificationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAbpnotificationDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a abpnotificationDo) Order(conds ...field.Expr) IAbpnotificationDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a abpnotificationDo) Distinct(cols ...field.Expr) IAbpnotificationDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a abpnotificationDo) Omit(cols ...field.Expr) IAbpnotificationDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a abpnotificationDo) Join(table schema.Tabler, on ...field.Expr) IAbpnotificationDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a abpnotificationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAbpnotificationDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a abpnotificationDo) RightJoin(table schema.Tabler, on ...field.Expr) IAbpnotificationDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a abpnotificationDo) Group(cols ...field.Expr) IAbpnotificationDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a abpnotificationDo) Having(conds ...gen.Condition) IAbpnotificationDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a abpnotificationDo) Limit(limit int) IAbpnotificationDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a abpnotificationDo) Offset(offset int) IAbpnotificationDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a abpnotificationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAbpnotificationDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a abpnotificationDo) Unscoped() IAbpnotificationDo {
	return a.withDO(a.DO.Unscoped())
}

func (a abpnotificationDo) Create(values ...*model.Abpnotification) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a abpnotificationDo) CreateInBatches(values []*model.Abpnotification, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a abpnotificationDo) Save(values ...*model.Abpnotification) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a abpnotificationDo) First() (*model.Abpnotification, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpnotification), nil
	}
}

func (a abpnotificationDo) Take() (*model.Abpnotification, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpnotification), nil
	}
}

func (a abpnotificationDo) Last() (*model.Abpnotification, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpnotification), nil
	}
}

func (a abpnotificationDo) Find() ([]*model.Abpnotification, error) {
	result, err := a.DO.Find()
	return result.([]*model.Abpnotification), err
}

func (a abpnotificationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Abpnotification, err error) {
	buf := make([]*model.Abpnotification, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a abpnotificationDo) FindInBatches(result *[]*model.Abpnotification, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a abpnotificationDo) Attrs(attrs ...field.AssignExpr) IAbpnotificationDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a abpnotificationDo) Assign(attrs ...field.AssignExpr) IAbpnotificationDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a abpnotificationDo) Joins(fields ...field.RelationField) IAbpnotificationDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a abpnotificationDo) Preload(fields ...field.RelationField) IAbpnotificationDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a abpnotificationDo) FirstOrInit() (*model.Abpnotification, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpnotification), nil
	}
}

func (a abpnotificationDo) FirstOrCreate() (*model.Abpnotification, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpnotification), nil
	}
}

func (a abpnotificationDo) FindByPage(offset int, limit int) (result []*model.Abpnotification, count int64, err error) {
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

func (a abpnotificationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a abpnotificationDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a abpnotificationDo) Delete(models ...*model.Abpnotification) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *abpnotificationDo) withDO(do gen.Dao) *abpnotificationDo {
	a.DO = *do.(*gen.DO)
	return a
}
