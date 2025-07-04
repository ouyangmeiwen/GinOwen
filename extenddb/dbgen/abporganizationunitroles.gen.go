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

func newAbporganizationunitrole(db *gorm.DB) abporganizationunitrole {
	_abporganizationunitrole := abporganizationunitrole{}

	_abporganizationunitrole.abporganizationunitroleDo.UseDB(db)
	_abporganizationunitrole.abporganizationunitroleDo.UseModel(&model.Abporganizationunitrole{})

	tableName := _abporganizationunitrole.abporganizationunitroleDo.TableName()
	_abporganizationunitrole.ALL = field.NewAsterisk(tableName)
	_abporganizationunitrole.ID = field.NewInt64(tableName, "Id")
	_abporganizationunitrole.CreationTime = field.NewTime(tableName, "CreationTime")
	_abporganizationunitrole.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_abporganizationunitrole.TenantID = field.NewInt64(tableName, "TenantId")
	_abporganizationunitrole.RoleID = field.NewInt64(tableName, "RoleId")
	_abporganizationunitrole.OrganizationUnitID = field.NewInt64(tableName, "OrganizationUnitId")
	_abporganizationunitrole.IsDeleted = field.NewField(tableName, "IsDeleted")

	_abporganizationunitrole.fillFieldMap()

	return _abporganizationunitrole
}

type abporganizationunitrole struct {
	abporganizationunitroleDo abporganizationunitroleDo

	ALL                field.Asterisk
	ID                 field.Int64
	CreationTime       field.Time
	CreatorUserID      field.Int64
	TenantID           field.Int64
	RoleID             field.Int64
	OrganizationUnitID field.Int64
	IsDeleted          field.Field

	fieldMap map[string]field.Expr
}

func (a abporganizationunitrole) Table(newTableName string) *abporganizationunitrole {
	a.abporganizationunitroleDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a abporganizationunitrole) As(alias string) *abporganizationunitrole {
	a.abporganizationunitroleDo.DO = *(a.abporganizationunitroleDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *abporganizationunitrole) updateTableName(table string) *abporganizationunitrole {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "Id")
	a.CreationTime = field.NewTime(table, "CreationTime")
	a.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	a.TenantID = field.NewInt64(table, "TenantId")
	a.RoleID = field.NewInt64(table, "RoleId")
	a.OrganizationUnitID = field.NewInt64(table, "OrganizationUnitId")
	a.IsDeleted = field.NewField(table, "IsDeleted")

	a.fillFieldMap()

	return a
}

func (a *abporganizationunitrole) WithContext(ctx context.Context) IAbporganizationunitroleDo {
	return a.abporganizationunitroleDo.WithContext(ctx)
}

func (a abporganizationunitrole) TableName() string { return a.abporganizationunitroleDo.TableName() }

func (a abporganizationunitrole) Alias() string { return a.abporganizationunitroleDo.Alias() }

func (a *abporganizationunitrole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *abporganizationunitrole) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 7)
	a.fieldMap["Id"] = a.ID
	a.fieldMap["CreationTime"] = a.CreationTime
	a.fieldMap["CreatorUserId"] = a.CreatorUserID
	a.fieldMap["TenantId"] = a.TenantID
	a.fieldMap["RoleId"] = a.RoleID
	a.fieldMap["OrganizationUnitId"] = a.OrganizationUnitID
	a.fieldMap["IsDeleted"] = a.IsDeleted
}

func (a abporganizationunitrole) clone(db *gorm.DB) abporganizationunitrole {
	a.abporganizationunitroleDo.ReplaceDB(db)
	return a
}

type abporganizationunitroleDo struct{ gen.DO }

type IAbporganizationunitroleDo interface {
	gen.SubQuery
	Debug() IAbporganizationunitroleDo
	WithContext(ctx context.Context) IAbporganizationunitroleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAbporganizationunitroleDo
	Not(conds ...gen.Condition) IAbporganizationunitroleDo
	Or(conds ...gen.Condition) IAbporganizationunitroleDo
	Select(conds ...field.Expr) IAbporganizationunitroleDo
	Where(conds ...gen.Condition) IAbporganizationunitroleDo
	Order(conds ...field.Expr) IAbporganizationunitroleDo
	Distinct(cols ...field.Expr) IAbporganizationunitroleDo
	Omit(cols ...field.Expr) IAbporganizationunitroleDo
	Join(table schema.Tabler, on ...field.Expr) IAbporganizationunitroleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAbporganizationunitroleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAbporganizationunitroleDo
	Group(cols ...field.Expr) IAbporganizationunitroleDo
	Having(conds ...gen.Condition) IAbporganizationunitroleDo
	Limit(limit int) IAbporganizationunitroleDo
	Offset(offset int) IAbporganizationunitroleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAbporganizationunitroleDo
	Unscoped() IAbporganizationunitroleDo
	Create(values ...*model.Abporganizationunitrole) error
	CreateInBatches(values []*model.Abporganizationunitrole, batchSize int) error
	Save(values ...*model.Abporganizationunitrole) error
	First() (*model.Abporganizationunitrole, error)
	Take() (*model.Abporganizationunitrole, error)
	Last() (*model.Abporganizationunitrole, error)
	Find() ([]*model.Abporganizationunitrole, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Abporganizationunitrole, err error)
	FindInBatches(result *[]*model.Abporganizationunitrole, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Abporganizationunitrole) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAbporganizationunitroleDo
	Assign(attrs ...field.AssignExpr) IAbporganizationunitroleDo
	Joins(fields ...field.RelationField) IAbporganizationunitroleDo
	Preload(fields ...field.RelationField) IAbporganizationunitroleDo
	FirstOrInit() (*model.Abporganizationunitrole, error)
	FirstOrCreate() (*model.Abporganizationunitrole, error)
	FindByPage(offset int, limit int) (result []*model.Abporganizationunitrole, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAbporganizationunitroleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a abporganizationunitroleDo) Debug() IAbporganizationunitroleDo {
	return a.withDO(a.DO.Debug())
}

func (a abporganizationunitroleDo) WithContext(ctx context.Context) IAbporganizationunitroleDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a abporganizationunitroleDo) ReadDB() IAbporganizationunitroleDo {
	return a.Clauses(dbresolver.Read)
}

func (a abporganizationunitroleDo) WriteDB() IAbporganizationunitroleDo {
	return a.Clauses(dbresolver.Write)
}

func (a abporganizationunitroleDo) Clauses(conds ...clause.Expression) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a abporganizationunitroleDo) Returning(value interface{}, columns ...string) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a abporganizationunitroleDo) Not(conds ...gen.Condition) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a abporganizationunitroleDo) Or(conds ...gen.Condition) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a abporganizationunitroleDo) Select(conds ...field.Expr) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a abporganizationunitroleDo) Where(conds ...gen.Condition) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a abporganizationunitroleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAbporganizationunitroleDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a abporganizationunitroleDo) Order(conds ...field.Expr) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a abporganizationunitroleDo) Distinct(cols ...field.Expr) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a abporganizationunitroleDo) Omit(cols ...field.Expr) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a abporganizationunitroleDo) Join(table schema.Tabler, on ...field.Expr) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a abporganizationunitroleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAbporganizationunitroleDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a abporganizationunitroleDo) RightJoin(table schema.Tabler, on ...field.Expr) IAbporganizationunitroleDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a abporganizationunitroleDo) Group(cols ...field.Expr) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a abporganizationunitroleDo) Having(conds ...gen.Condition) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a abporganizationunitroleDo) Limit(limit int) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a abporganizationunitroleDo) Offset(offset int) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a abporganizationunitroleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a abporganizationunitroleDo) Unscoped() IAbporganizationunitroleDo {
	return a.withDO(a.DO.Unscoped())
}

func (a abporganizationunitroleDo) Create(values ...*model.Abporganizationunitrole) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a abporganizationunitroleDo) CreateInBatches(values []*model.Abporganizationunitrole, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a abporganizationunitroleDo) Save(values ...*model.Abporganizationunitrole) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a abporganizationunitroleDo) First() (*model.Abporganizationunitrole, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abporganizationunitrole), nil
	}
}

func (a abporganizationunitroleDo) Take() (*model.Abporganizationunitrole, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abporganizationunitrole), nil
	}
}

func (a abporganizationunitroleDo) Last() (*model.Abporganizationunitrole, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abporganizationunitrole), nil
	}
}

func (a abporganizationunitroleDo) Find() ([]*model.Abporganizationunitrole, error) {
	result, err := a.DO.Find()
	return result.([]*model.Abporganizationunitrole), err
}

func (a abporganizationunitroleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Abporganizationunitrole, err error) {
	buf := make([]*model.Abporganizationunitrole, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a abporganizationunitroleDo) FindInBatches(result *[]*model.Abporganizationunitrole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a abporganizationunitroleDo) Attrs(attrs ...field.AssignExpr) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a abporganizationunitroleDo) Assign(attrs ...field.AssignExpr) IAbporganizationunitroleDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a abporganizationunitroleDo) Joins(fields ...field.RelationField) IAbporganizationunitroleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a abporganizationunitroleDo) Preload(fields ...field.RelationField) IAbporganizationunitroleDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a abporganizationunitroleDo) FirstOrInit() (*model.Abporganizationunitrole, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abporganizationunitrole), nil
	}
}

func (a abporganizationunitroleDo) FirstOrCreate() (*model.Abporganizationunitrole, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abporganizationunitrole), nil
	}
}

func (a abporganizationunitroleDo) FindByPage(offset int, limit int) (result []*model.Abporganizationunitrole, count int64, err error) {
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

func (a abporganizationunitroleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a abporganizationunitroleDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a abporganizationunitroleDo) Delete(models ...*model.Abporganizationunitrole) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *abporganizationunitroleDo) withDO(do gen.Dao) *abporganizationunitroleDo {
	a.DO = *do.(*gen.DO)
	return a
}
