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

func newAbptenant(db *gorm.DB) abptenant {
	_abptenant := abptenant{}

	_abptenant.abptenantDo.UseDB(db)
	_abptenant.abptenantDo.UseModel(&model.Abptenant{})

	tableName := _abptenant.abptenantDo.TableName()
	_abptenant.ALL = field.NewAsterisk(tableName)
	_abptenant.ID = field.NewInt64(tableName, "Id")
	_abptenant.CreationTime = field.NewTime(tableName, "CreationTime")
	_abptenant.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_abptenant.LastModificationTime = field.NewTime(tableName, "LastModificationTime")
	_abptenant.LastModifierUserID = field.NewInt64(tableName, "LastModifierUserId")
	_abptenant.IsDeleted = field.NewField(tableName, "IsDeleted")
	_abptenant.DeleterUserID = field.NewInt64(tableName, "DeleterUserId")
	_abptenant.DeletionTime = field.NewTime(tableName, "DeletionTime")
	_abptenant.TenancyName = field.NewString(tableName, "TenancyName")
	_abptenant.Name = field.NewString(tableName, "Name")
	_abptenant.ConnectionString = field.NewString(tableName, "ConnectionString")
	_abptenant.IsActive = field.NewField(tableName, "IsActive")
	_abptenant.EditionID = field.NewInt64(tableName, "EditionId")
	_abptenant.SubscriptionEndDateUtc = field.NewTime(tableName, "SubscriptionEndDateUtc")
	_abptenant.IsInTrialPeriod = field.NewField(tableName, "IsInTrialPeriod")
	_abptenant.CustomCSSID = field.NewString(tableName, "CustomCssId")
	_abptenant.LogoID = field.NewString(tableName, "LogoId")
	_abptenant.LogoFileType = field.NewString(tableName, "LogoFileType")
	_abptenant.SubscriptionPaymentType = field.NewInt64(tableName, "SubscriptionPaymentType")

	_abptenant.fillFieldMap()

	return _abptenant
}

type abptenant struct {
	abptenantDo abptenantDo

	ALL                     field.Asterisk
	ID                      field.Int64
	CreationTime            field.Time
	CreatorUserID           field.Int64
	LastModificationTime    field.Time
	LastModifierUserID      field.Int64
	IsDeleted               field.Field
	DeleterUserID           field.Int64
	DeletionTime            field.Time
	TenancyName             field.String
	Name                    field.String
	ConnectionString        field.String
	IsActive                field.Field
	EditionID               field.Int64
	SubscriptionEndDateUtc  field.Time
	IsInTrialPeriod         field.Field
	CustomCSSID             field.String
	LogoID                  field.String
	LogoFileType            field.String
	SubscriptionPaymentType field.Int64

	fieldMap map[string]field.Expr
}

func (a abptenant) Table(newTableName string) *abptenant {
	a.abptenantDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a abptenant) As(alias string) *abptenant {
	a.abptenantDo.DO = *(a.abptenantDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *abptenant) updateTableName(table string) *abptenant {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "Id")
	a.CreationTime = field.NewTime(table, "CreationTime")
	a.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	a.LastModificationTime = field.NewTime(table, "LastModificationTime")
	a.LastModifierUserID = field.NewInt64(table, "LastModifierUserId")
	a.IsDeleted = field.NewField(table, "IsDeleted")
	a.DeleterUserID = field.NewInt64(table, "DeleterUserId")
	a.DeletionTime = field.NewTime(table, "DeletionTime")
	a.TenancyName = field.NewString(table, "TenancyName")
	a.Name = field.NewString(table, "Name")
	a.ConnectionString = field.NewString(table, "ConnectionString")
	a.IsActive = field.NewField(table, "IsActive")
	a.EditionID = field.NewInt64(table, "EditionId")
	a.SubscriptionEndDateUtc = field.NewTime(table, "SubscriptionEndDateUtc")
	a.IsInTrialPeriod = field.NewField(table, "IsInTrialPeriod")
	a.CustomCSSID = field.NewString(table, "CustomCssId")
	a.LogoID = field.NewString(table, "LogoId")
	a.LogoFileType = field.NewString(table, "LogoFileType")
	a.SubscriptionPaymentType = field.NewInt64(table, "SubscriptionPaymentType")

	a.fillFieldMap()

	return a
}

func (a *abptenant) WithContext(ctx context.Context) IAbptenantDo {
	return a.abptenantDo.WithContext(ctx)
}

func (a abptenant) TableName() string { return a.abptenantDo.TableName() }

func (a abptenant) Alias() string { return a.abptenantDo.Alias() }

func (a *abptenant) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *abptenant) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 19)
	a.fieldMap["Id"] = a.ID
	a.fieldMap["CreationTime"] = a.CreationTime
	a.fieldMap["CreatorUserId"] = a.CreatorUserID
	a.fieldMap["LastModificationTime"] = a.LastModificationTime
	a.fieldMap["LastModifierUserId"] = a.LastModifierUserID
	a.fieldMap["IsDeleted"] = a.IsDeleted
	a.fieldMap["DeleterUserId"] = a.DeleterUserID
	a.fieldMap["DeletionTime"] = a.DeletionTime
	a.fieldMap["TenancyName"] = a.TenancyName
	a.fieldMap["Name"] = a.Name
	a.fieldMap["ConnectionString"] = a.ConnectionString
	a.fieldMap["IsActive"] = a.IsActive
	a.fieldMap["EditionId"] = a.EditionID
	a.fieldMap["SubscriptionEndDateUtc"] = a.SubscriptionEndDateUtc
	a.fieldMap["IsInTrialPeriod"] = a.IsInTrialPeriod
	a.fieldMap["CustomCssId"] = a.CustomCSSID
	a.fieldMap["LogoId"] = a.LogoID
	a.fieldMap["LogoFileType"] = a.LogoFileType
	a.fieldMap["SubscriptionPaymentType"] = a.SubscriptionPaymentType
}

func (a abptenant) clone(db *gorm.DB) abptenant {
	a.abptenantDo.ReplaceDB(db)
	return a
}

type abptenantDo struct{ gen.DO }

type IAbptenantDo interface {
	gen.SubQuery
	Debug() IAbptenantDo
	WithContext(ctx context.Context) IAbptenantDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAbptenantDo
	Not(conds ...gen.Condition) IAbptenantDo
	Or(conds ...gen.Condition) IAbptenantDo
	Select(conds ...field.Expr) IAbptenantDo
	Where(conds ...gen.Condition) IAbptenantDo
	Order(conds ...field.Expr) IAbptenantDo
	Distinct(cols ...field.Expr) IAbptenantDo
	Omit(cols ...field.Expr) IAbptenantDo
	Join(table schema.Tabler, on ...field.Expr) IAbptenantDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAbptenantDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAbptenantDo
	Group(cols ...field.Expr) IAbptenantDo
	Having(conds ...gen.Condition) IAbptenantDo
	Limit(limit int) IAbptenantDo
	Offset(offset int) IAbptenantDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAbptenantDo
	Unscoped() IAbptenantDo
	Create(values ...*model.Abptenant) error
	CreateInBatches(values []*model.Abptenant, batchSize int) error
	Save(values ...*model.Abptenant) error
	First() (*model.Abptenant, error)
	Take() (*model.Abptenant, error)
	Last() (*model.Abptenant, error)
	Find() ([]*model.Abptenant, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Abptenant, err error)
	FindInBatches(result *[]*model.Abptenant, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Abptenant) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAbptenantDo
	Assign(attrs ...field.AssignExpr) IAbptenantDo
	Joins(fields ...field.RelationField) IAbptenantDo
	Preload(fields ...field.RelationField) IAbptenantDo
	FirstOrInit() (*model.Abptenant, error)
	FirstOrCreate() (*model.Abptenant, error)
	FindByPage(offset int, limit int) (result []*model.Abptenant, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAbptenantDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a abptenantDo) Debug() IAbptenantDo {
	return a.withDO(a.DO.Debug())
}

func (a abptenantDo) WithContext(ctx context.Context) IAbptenantDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a abptenantDo) ReadDB() IAbptenantDo {
	return a.Clauses(dbresolver.Read)
}

func (a abptenantDo) WriteDB() IAbptenantDo {
	return a.Clauses(dbresolver.Write)
}

func (a abptenantDo) Clauses(conds ...clause.Expression) IAbptenantDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a abptenantDo) Returning(value interface{}, columns ...string) IAbptenantDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a abptenantDo) Not(conds ...gen.Condition) IAbptenantDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a abptenantDo) Or(conds ...gen.Condition) IAbptenantDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a abptenantDo) Select(conds ...field.Expr) IAbptenantDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a abptenantDo) Where(conds ...gen.Condition) IAbptenantDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a abptenantDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAbptenantDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a abptenantDo) Order(conds ...field.Expr) IAbptenantDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a abptenantDo) Distinct(cols ...field.Expr) IAbptenantDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a abptenantDo) Omit(cols ...field.Expr) IAbptenantDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a abptenantDo) Join(table schema.Tabler, on ...field.Expr) IAbptenantDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a abptenantDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAbptenantDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a abptenantDo) RightJoin(table schema.Tabler, on ...field.Expr) IAbptenantDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a abptenantDo) Group(cols ...field.Expr) IAbptenantDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a abptenantDo) Having(conds ...gen.Condition) IAbptenantDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a abptenantDo) Limit(limit int) IAbptenantDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a abptenantDo) Offset(offset int) IAbptenantDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a abptenantDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAbptenantDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a abptenantDo) Unscoped() IAbptenantDo {
	return a.withDO(a.DO.Unscoped())
}

func (a abptenantDo) Create(values ...*model.Abptenant) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a abptenantDo) CreateInBatches(values []*model.Abptenant, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a abptenantDo) Save(values ...*model.Abptenant) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a abptenantDo) First() (*model.Abptenant, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abptenant), nil
	}
}

func (a abptenantDo) Take() (*model.Abptenant, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abptenant), nil
	}
}

func (a abptenantDo) Last() (*model.Abptenant, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abptenant), nil
	}
}

func (a abptenantDo) Find() ([]*model.Abptenant, error) {
	result, err := a.DO.Find()
	return result.([]*model.Abptenant), err
}

func (a abptenantDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Abptenant, err error) {
	buf := make([]*model.Abptenant, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a abptenantDo) FindInBatches(result *[]*model.Abptenant, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a abptenantDo) Attrs(attrs ...field.AssignExpr) IAbptenantDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a abptenantDo) Assign(attrs ...field.AssignExpr) IAbptenantDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a abptenantDo) Joins(fields ...field.RelationField) IAbptenantDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a abptenantDo) Preload(fields ...field.RelationField) IAbptenantDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a abptenantDo) FirstOrInit() (*model.Abptenant, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abptenant), nil
	}
}

func (a abptenantDo) FirstOrCreate() (*model.Abptenant, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abptenant), nil
	}
}

func (a abptenantDo) FindByPage(offset int, limit int) (result []*model.Abptenant, count int64, err error) {
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

func (a abptenantDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a abptenantDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a abptenantDo) Delete(models ...*model.Abptenant) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *abptenantDo) withDO(do gen.Dao) *abptenantDo {
	a.DO = *do.(*gen.DO)
	return a
}
