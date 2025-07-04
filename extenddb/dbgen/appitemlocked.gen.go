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

func newAppitemlocked(db *gorm.DB) appitemlocked {
	_appitemlocked := appitemlocked{}

	_appitemlocked.appitemlockedDo.UseDB(db)
	_appitemlocked.appitemlockedDo.UseModel(&model.Appitemlocked{})

	tableName := _appitemlocked.appitemlockedDo.TableName()
	_appitemlocked.ALL = field.NewAsterisk(tableName)
	_appitemlocked.ID = field.NewString(tableName, "Id")
	_appitemlocked.CreationTime = field.NewTime(tableName, "CreationTime")
	_appitemlocked.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_appitemlocked.LastModificationTime = field.NewTime(tableName, "LastModificationTime")
	_appitemlocked.LastModifierUserID = field.NewInt64(tableName, "LastModifierUserId")
	_appitemlocked.PatronID = field.NewString(tableName, "PatronId")
	_appitemlocked.SerialNo = field.NewString(tableName, "SerialNo")
	_appitemlocked.PatronName = field.NewString(tableName, "PatronName")
	_appitemlocked.PatronBarcode = field.NewString(tableName, "PatronBarcode")
	_appitemlocked.TerminalID = field.NewString(tableName, "TerminalId")
	_appitemlocked.TerminalCode = field.NewString(tableName, "TerminalCode")
	_appitemlocked.TerminalName = field.NewString(tableName, "TerminalName")
	_appitemlocked.TerminalShelfID = field.NewString(tableName, "TerminalShelfId")
	_appitemlocked.TerminalShelfName = field.NewString(tableName, "TerminalShelfName")
	_appitemlocked.ItemID = field.NewString(tableName, "ItemId")
	_appitemlocked.ItemTitle = field.NewString(tableName, "ItemTitle")
	_appitemlocked.ItemBarcode = field.NewString(tableName, "ItemBarcode")
	_appitemlocked.ItemAuthor = field.NewString(tableName, "ItemAuthor")
	_appitemlocked.ItemISBN = field.NewString(tableName, "ItemISBN")
	_appitemlocked.IsEnable = field.NewField(tableName, "IsEnable")
	_appitemlocked.LockStartTime = field.NewTime(tableName, "LockStartTime")
	_appitemlocked.LockEndTime = field.NewTime(tableName, "LockEndTime")
	_appitemlocked.PickUpTime = field.NewTime(tableName, "PickUpTime")
	_appitemlocked.IsCanceled = field.NewField(tableName, "IsCanceled")
	_appitemlocked.TenantID = field.NewInt64(tableName, "TenantId")

	_appitemlocked.fillFieldMap()

	return _appitemlocked
}

type appitemlocked struct {
	appitemlockedDo appitemlockedDo

	ALL                  field.Asterisk
	ID                   field.String
	CreationTime         field.Time
	CreatorUserID        field.Int64
	LastModificationTime field.Time
	LastModifierUserID   field.Int64
	PatronID             field.String
	SerialNo             field.String
	PatronName           field.String
	PatronBarcode        field.String
	TerminalID           field.String
	TerminalCode         field.String
	TerminalName         field.String
	TerminalShelfID      field.String
	TerminalShelfName    field.String
	ItemID               field.String
	ItemTitle            field.String
	ItemBarcode          field.String
	ItemAuthor           field.String
	ItemISBN             field.String
	IsEnable             field.Field
	LockStartTime        field.Time
	LockEndTime          field.Time
	PickUpTime           field.Time
	IsCanceled           field.Field
	TenantID             field.Int64

	fieldMap map[string]field.Expr
}

func (a appitemlocked) Table(newTableName string) *appitemlocked {
	a.appitemlockedDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a appitemlocked) As(alias string) *appitemlocked {
	a.appitemlockedDo.DO = *(a.appitemlockedDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *appitemlocked) updateTableName(table string) *appitemlocked {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewString(table, "Id")
	a.CreationTime = field.NewTime(table, "CreationTime")
	a.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	a.LastModificationTime = field.NewTime(table, "LastModificationTime")
	a.LastModifierUserID = field.NewInt64(table, "LastModifierUserId")
	a.PatronID = field.NewString(table, "PatronId")
	a.SerialNo = field.NewString(table, "SerialNo")
	a.PatronName = field.NewString(table, "PatronName")
	a.PatronBarcode = field.NewString(table, "PatronBarcode")
	a.TerminalID = field.NewString(table, "TerminalId")
	a.TerminalCode = field.NewString(table, "TerminalCode")
	a.TerminalName = field.NewString(table, "TerminalName")
	a.TerminalShelfID = field.NewString(table, "TerminalShelfId")
	a.TerminalShelfName = field.NewString(table, "TerminalShelfName")
	a.ItemID = field.NewString(table, "ItemId")
	a.ItemTitle = field.NewString(table, "ItemTitle")
	a.ItemBarcode = field.NewString(table, "ItemBarcode")
	a.ItemAuthor = field.NewString(table, "ItemAuthor")
	a.ItemISBN = field.NewString(table, "ItemISBN")
	a.IsEnable = field.NewField(table, "IsEnable")
	a.LockStartTime = field.NewTime(table, "LockStartTime")
	a.LockEndTime = field.NewTime(table, "LockEndTime")
	a.PickUpTime = field.NewTime(table, "PickUpTime")
	a.IsCanceled = field.NewField(table, "IsCanceled")
	a.TenantID = field.NewInt64(table, "TenantId")

	a.fillFieldMap()

	return a
}

func (a *appitemlocked) WithContext(ctx context.Context) IAppitemlockedDo {
	return a.appitemlockedDo.WithContext(ctx)
}

func (a appitemlocked) TableName() string { return a.appitemlockedDo.TableName() }

func (a appitemlocked) Alias() string { return a.appitemlockedDo.Alias() }

func (a *appitemlocked) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *appitemlocked) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 25)
	a.fieldMap["Id"] = a.ID
	a.fieldMap["CreationTime"] = a.CreationTime
	a.fieldMap["CreatorUserId"] = a.CreatorUserID
	a.fieldMap["LastModificationTime"] = a.LastModificationTime
	a.fieldMap["LastModifierUserId"] = a.LastModifierUserID
	a.fieldMap["PatronId"] = a.PatronID
	a.fieldMap["SerialNo"] = a.SerialNo
	a.fieldMap["PatronName"] = a.PatronName
	a.fieldMap["PatronBarcode"] = a.PatronBarcode
	a.fieldMap["TerminalId"] = a.TerminalID
	a.fieldMap["TerminalCode"] = a.TerminalCode
	a.fieldMap["TerminalName"] = a.TerminalName
	a.fieldMap["TerminalShelfId"] = a.TerminalShelfID
	a.fieldMap["TerminalShelfName"] = a.TerminalShelfName
	a.fieldMap["ItemId"] = a.ItemID
	a.fieldMap["ItemTitle"] = a.ItemTitle
	a.fieldMap["ItemBarcode"] = a.ItemBarcode
	a.fieldMap["ItemAuthor"] = a.ItemAuthor
	a.fieldMap["ItemISBN"] = a.ItemISBN
	a.fieldMap["IsEnable"] = a.IsEnable
	a.fieldMap["LockStartTime"] = a.LockStartTime
	a.fieldMap["LockEndTime"] = a.LockEndTime
	a.fieldMap["PickUpTime"] = a.PickUpTime
	a.fieldMap["IsCanceled"] = a.IsCanceled
	a.fieldMap["TenantId"] = a.TenantID
}

func (a appitemlocked) clone(db *gorm.DB) appitemlocked {
	a.appitemlockedDo.ReplaceDB(db)
	return a
}

type appitemlockedDo struct{ gen.DO }

type IAppitemlockedDo interface {
	gen.SubQuery
	Debug() IAppitemlockedDo
	WithContext(ctx context.Context) IAppitemlockedDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAppitemlockedDo
	Not(conds ...gen.Condition) IAppitemlockedDo
	Or(conds ...gen.Condition) IAppitemlockedDo
	Select(conds ...field.Expr) IAppitemlockedDo
	Where(conds ...gen.Condition) IAppitemlockedDo
	Order(conds ...field.Expr) IAppitemlockedDo
	Distinct(cols ...field.Expr) IAppitemlockedDo
	Omit(cols ...field.Expr) IAppitemlockedDo
	Join(table schema.Tabler, on ...field.Expr) IAppitemlockedDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAppitemlockedDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAppitemlockedDo
	Group(cols ...field.Expr) IAppitemlockedDo
	Having(conds ...gen.Condition) IAppitemlockedDo
	Limit(limit int) IAppitemlockedDo
	Offset(offset int) IAppitemlockedDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAppitemlockedDo
	Unscoped() IAppitemlockedDo
	Create(values ...*model.Appitemlocked) error
	CreateInBatches(values []*model.Appitemlocked, batchSize int) error
	Save(values ...*model.Appitemlocked) error
	First() (*model.Appitemlocked, error)
	Take() (*model.Appitemlocked, error)
	Last() (*model.Appitemlocked, error)
	Find() ([]*model.Appitemlocked, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Appitemlocked, err error)
	FindInBatches(result *[]*model.Appitemlocked, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Appitemlocked) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAppitemlockedDo
	Assign(attrs ...field.AssignExpr) IAppitemlockedDo
	Joins(fields ...field.RelationField) IAppitemlockedDo
	Preload(fields ...field.RelationField) IAppitemlockedDo
	FirstOrInit() (*model.Appitemlocked, error)
	FirstOrCreate() (*model.Appitemlocked, error)
	FindByPage(offset int, limit int) (result []*model.Appitemlocked, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAppitemlockedDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a appitemlockedDo) Debug() IAppitemlockedDo {
	return a.withDO(a.DO.Debug())
}

func (a appitemlockedDo) WithContext(ctx context.Context) IAppitemlockedDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a appitemlockedDo) ReadDB() IAppitemlockedDo {
	return a.Clauses(dbresolver.Read)
}

func (a appitemlockedDo) WriteDB() IAppitemlockedDo {
	return a.Clauses(dbresolver.Write)
}

func (a appitemlockedDo) Clauses(conds ...clause.Expression) IAppitemlockedDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a appitemlockedDo) Returning(value interface{}, columns ...string) IAppitemlockedDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a appitemlockedDo) Not(conds ...gen.Condition) IAppitemlockedDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a appitemlockedDo) Or(conds ...gen.Condition) IAppitemlockedDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a appitemlockedDo) Select(conds ...field.Expr) IAppitemlockedDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a appitemlockedDo) Where(conds ...gen.Condition) IAppitemlockedDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a appitemlockedDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAppitemlockedDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a appitemlockedDo) Order(conds ...field.Expr) IAppitemlockedDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a appitemlockedDo) Distinct(cols ...field.Expr) IAppitemlockedDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a appitemlockedDo) Omit(cols ...field.Expr) IAppitemlockedDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a appitemlockedDo) Join(table schema.Tabler, on ...field.Expr) IAppitemlockedDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a appitemlockedDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAppitemlockedDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a appitemlockedDo) RightJoin(table schema.Tabler, on ...field.Expr) IAppitemlockedDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a appitemlockedDo) Group(cols ...field.Expr) IAppitemlockedDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a appitemlockedDo) Having(conds ...gen.Condition) IAppitemlockedDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a appitemlockedDo) Limit(limit int) IAppitemlockedDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a appitemlockedDo) Offset(offset int) IAppitemlockedDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a appitemlockedDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAppitemlockedDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a appitemlockedDo) Unscoped() IAppitemlockedDo {
	return a.withDO(a.DO.Unscoped())
}

func (a appitemlockedDo) Create(values ...*model.Appitemlocked) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a appitemlockedDo) CreateInBatches(values []*model.Appitemlocked, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a appitemlockedDo) Save(values ...*model.Appitemlocked) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a appitemlockedDo) First() (*model.Appitemlocked, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Appitemlocked), nil
	}
}

func (a appitemlockedDo) Take() (*model.Appitemlocked, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Appitemlocked), nil
	}
}

func (a appitemlockedDo) Last() (*model.Appitemlocked, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Appitemlocked), nil
	}
}

func (a appitemlockedDo) Find() ([]*model.Appitemlocked, error) {
	result, err := a.DO.Find()
	return result.([]*model.Appitemlocked), err
}

func (a appitemlockedDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Appitemlocked, err error) {
	buf := make([]*model.Appitemlocked, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a appitemlockedDo) FindInBatches(result *[]*model.Appitemlocked, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a appitemlockedDo) Attrs(attrs ...field.AssignExpr) IAppitemlockedDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a appitemlockedDo) Assign(attrs ...field.AssignExpr) IAppitemlockedDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a appitemlockedDo) Joins(fields ...field.RelationField) IAppitemlockedDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a appitemlockedDo) Preload(fields ...field.RelationField) IAppitemlockedDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a appitemlockedDo) FirstOrInit() (*model.Appitemlocked, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Appitemlocked), nil
	}
}

func (a appitemlockedDo) FirstOrCreate() (*model.Appitemlocked, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Appitemlocked), nil
	}
}

func (a appitemlockedDo) FindByPage(offset int, limit int) (result []*model.Appitemlocked, count int64, err error) {
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

func (a appitemlockedDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a appitemlockedDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a appitemlockedDo) Delete(models ...*model.Appitemlocked) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *appitemlockedDo) withDO(do gen.Dao) *appitemlockedDo {
	a.DO = *do.(*gen.DO)
	return a
}
