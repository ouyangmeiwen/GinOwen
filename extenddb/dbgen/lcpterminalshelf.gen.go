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

func newLcpterminalshelf(db *gorm.DB) lcpterminalshelf {
	_lcpterminalshelf := lcpterminalshelf{}

	_lcpterminalshelf.lcpterminalshelfDo.UseDB(db)
	_lcpterminalshelf.lcpterminalshelfDo.UseModel(&model.Lcpterminalshelf{})

	tableName := _lcpterminalshelf.lcpterminalshelfDo.TableName()
	_lcpterminalshelf.ALL = field.NewAsterisk(tableName)
	_lcpterminalshelf.ID = field.NewString(tableName, "Id")
	_lcpterminalshelf.CreationTime = field.NewTime(tableName, "CreationTime")
	_lcpterminalshelf.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_lcpterminalshelf.LastModificationTime = field.NewTime(tableName, "LastModificationTime")
	_lcpterminalshelf.LastModifierUserID = field.NewInt64(tableName, "LastModifierUserId")
	_lcpterminalshelf.ShelfName = field.NewString(tableName, "ShelfName")
	_lcpterminalshelf.TerminalID = field.NewString(tableName, "TerminalId")
	_lcpterminalshelf.TerminalCode = field.NewString(tableName, "TerminalCode")
	_lcpterminalshelf.TerminalName = field.NewString(tableName, "TerminalName")
	_lcpterminalshelf.ItemID = field.NewString(tableName, "ItemId")
	_lcpterminalshelf.ItemTitle = field.NewString(tableName, "ItemTitle")
	_lcpterminalshelf.ItemAuthor = field.NewString(tableName, "ItemAuthor")
	_lcpterminalshelf.ItemISBN = field.NewString(tableName, "ItemISBN")
	_lcpterminalshelf.ItemBarcode = field.NewString(tableName, "ItemBarcode")
	_lcpterminalshelf.ItemCallNo = field.NewString(tableName, "ItemCallNo")
	_lcpterminalshelf.PlcCommand = field.NewString(tableName, "PlcCommand")
	_lcpterminalshelf.IsEmpty = field.NewField(tableName, "IsEmpty")
	_lcpterminalshelf.IsDisable = field.NewField(tableName, "IsDisable")
	_lcpterminalshelf.IsInterference = field.NewField(tableName, "IsInterference")
	_lcpterminalshelf.IsEnable = field.NewField(tableName, "IsEnable")
	_lcpterminalshelf.IsReserve = field.NewField(tableName, "IsReserve")
	_lcpterminalshelf.ReserveType = field.NewInt64(tableName, "ReserveType")
	_lcpterminalshelf.PatronBarcode = field.NewString(tableName, "PatronBarcode")
	_lcpterminalshelf.DisableReason = field.NewString(tableName, "DisableReason")
	_lcpterminalshelf.TenantID = field.NewInt64(tableName, "TenantId")
	_lcpterminalshelf.ReserveOverdueTime = field.NewTime(tableName, "ReserveOverdueTime")
	_lcpterminalshelf.DisablePatronBarcode = field.NewString(tableName, "DisablePatronBarcode")
	_lcpterminalshelf.DisplayName = field.NewString(tableName, "DisplayName")

	_lcpterminalshelf.fillFieldMap()

	return _lcpterminalshelf
}

type lcpterminalshelf struct {
	lcpterminalshelfDo lcpterminalshelfDo

	ALL                  field.Asterisk
	ID                   field.String
	CreationTime         field.Time
	CreatorUserID        field.Int64
	LastModificationTime field.Time
	LastModifierUserID   field.Int64
	ShelfName            field.String
	TerminalID           field.String
	TerminalCode         field.String
	TerminalName         field.String
	ItemID               field.String
	ItemTitle            field.String
	ItemAuthor           field.String
	ItemISBN             field.String
	ItemBarcode          field.String
	ItemCallNo           field.String
	PlcCommand           field.String
	IsEmpty              field.Field
	IsDisable            field.Field
	IsInterference       field.Field
	IsEnable             field.Field
	IsReserve            field.Field
	ReserveType          field.Int64
	PatronBarcode        field.String
	DisableReason        field.String
	TenantID             field.Int64
	ReserveOverdueTime   field.Time
	DisablePatronBarcode field.String
	DisplayName          field.String

	fieldMap map[string]field.Expr
}

func (l lcpterminalshelf) Table(newTableName string) *lcpterminalshelf {
	l.lcpterminalshelfDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l lcpterminalshelf) As(alias string) *lcpterminalshelf {
	l.lcpterminalshelfDo.DO = *(l.lcpterminalshelfDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *lcpterminalshelf) updateTableName(table string) *lcpterminalshelf {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewString(table, "Id")
	l.CreationTime = field.NewTime(table, "CreationTime")
	l.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	l.LastModificationTime = field.NewTime(table, "LastModificationTime")
	l.LastModifierUserID = field.NewInt64(table, "LastModifierUserId")
	l.ShelfName = field.NewString(table, "ShelfName")
	l.TerminalID = field.NewString(table, "TerminalId")
	l.TerminalCode = field.NewString(table, "TerminalCode")
	l.TerminalName = field.NewString(table, "TerminalName")
	l.ItemID = field.NewString(table, "ItemId")
	l.ItemTitle = field.NewString(table, "ItemTitle")
	l.ItemAuthor = field.NewString(table, "ItemAuthor")
	l.ItemISBN = field.NewString(table, "ItemISBN")
	l.ItemBarcode = field.NewString(table, "ItemBarcode")
	l.ItemCallNo = field.NewString(table, "ItemCallNo")
	l.PlcCommand = field.NewString(table, "PlcCommand")
	l.IsEmpty = field.NewField(table, "IsEmpty")
	l.IsDisable = field.NewField(table, "IsDisable")
	l.IsInterference = field.NewField(table, "IsInterference")
	l.IsEnable = field.NewField(table, "IsEnable")
	l.IsReserve = field.NewField(table, "IsReserve")
	l.ReserveType = field.NewInt64(table, "ReserveType")
	l.PatronBarcode = field.NewString(table, "PatronBarcode")
	l.DisableReason = field.NewString(table, "DisableReason")
	l.TenantID = field.NewInt64(table, "TenantId")
	l.ReserveOverdueTime = field.NewTime(table, "ReserveOverdueTime")
	l.DisablePatronBarcode = field.NewString(table, "DisablePatronBarcode")
	l.DisplayName = field.NewString(table, "DisplayName")

	l.fillFieldMap()

	return l
}

func (l *lcpterminalshelf) WithContext(ctx context.Context) ILcpterminalshelfDo {
	return l.lcpterminalshelfDo.WithContext(ctx)
}

func (l lcpterminalshelf) TableName() string { return l.lcpterminalshelfDo.TableName() }

func (l lcpterminalshelf) Alias() string { return l.lcpterminalshelfDo.Alias() }

func (l *lcpterminalshelf) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *lcpterminalshelf) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 28)
	l.fieldMap["Id"] = l.ID
	l.fieldMap["CreationTime"] = l.CreationTime
	l.fieldMap["CreatorUserId"] = l.CreatorUserID
	l.fieldMap["LastModificationTime"] = l.LastModificationTime
	l.fieldMap["LastModifierUserId"] = l.LastModifierUserID
	l.fieldMap["ShelfName"] = l.ShelfName
	l.fieldMap["TerminalId"] = l.TerminalID
	l.fieldMap["TerminalCode"] = l.TerminalCode
	l.fieldMap["TerminalName"] = l.TerminalName
	l.fieldMap["ItemId"] = l.ItemID
	l.fieldMap["ItemTitle"] = l.ItemTitle
	l.fieldMap["ItemAuthor"] = l.ItemAuthor
	l.fieldMap["ItemISBN"] = l.ItemISBN
	l.fieldMap["ItemBarcode"] = l.ItemBarcode
	l.fieldMap["ItemCallNo"] = l.ItemCallNo
	l.fieldMap["PlcCommand"] = l.PlcCommand
	l.fieldMap["IsEmpty"] = l.IsEmpty
	l.fieldMap["IsDisable"] = l.IsDisable
	l.fieldMap["IsInterference"] = l.IsInterference
	l.fieldMap["IsEnable"] = l.IsEnable
	l.fieldMap["IsReserve"] = l.IsReserve
	l.fieldMap["ReserveType"] = l.ReserveType
	l.fieldMap["PatronBarcode"] = l.PatronBarcode
	l.fieldMap["DisableReason"] = l.DisableReason
	l.fieldMap["TenantId"] = l.TenantID
	l.fieldMap["ReserveOverdueTime"] = l.ReserveOverdueTime
	l.fieldMap["DisablePatronBarcode"] = l.DisablePatronBarcode
	l.fieldMap["DisplayName"] = l.DisplayName
}

func (l lcpterminalshelf) clone(db *gorm.DB) lcpterminalshelf {
	l.lcpterminalshelfDo.ReplaceDB(db)
	return l
}

type lcpterminalshelfDo struct{ gen.DO }

type ILcpterminalshelfDo interface {
	gen.SubQuery
	Debug() ILcpterminalshelfDo
	WithContext(ctx context.Context) ILcpterminalshelfDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILcpterminalshelfDo
	Not(conds ...gen.Condition) ILcpterminalshelfDo
	Or(conds ...gen.Condition) ILcpterminalshelfDo
	Select(conds ...field.Expr) ILcpterminalshelfDo
	Where(conds ...gen.Condition) ILcpterminalshelfDo
	Order(conds ...field.Expr) ILcpterminalshelfDo
	Distinct(cols ...field.Expr) ILcpterminalshelfDo
	Omit(cols ...field.Expr) ILcpterminalshelfDo
	Join(table schema.Tabler, on ...field.Expr) ILcpterminalshelfDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILcpterminalshelfDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILcpterminalshelfDo
	Group(cols ...field.Expr) ILcpterminalshelfDo
	Having(conds ...gen.Condition) ILcpterminalshelfDo
	Limit(limit int) ILcpterminalshelfDo
	Offset(offset int) ILcpterminalshelfDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILcpterminalshelfDo
	Unscoped() ILcpterminalshelfDo
	Create(values ...*model.Lcpterminalshelf) error
	CreateInBatches(values []*model.Lcpterminalshelf, batchSize int) error
	Save(values ...*model.Lcpterminalshelf) error
	First() (*model.Lcpterminalshelf, error)
	Take() (*model.Lcpterminalshelf, error)
	Last() (*model.Lcpterminalshelf, error)
	Find() ([]*model.Lcpterminalshelf, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Lcpterminalshelf, err error)
	FindInBatches(result *[]*model.Lcpterminalshelf, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Lcpterminalshelf) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILcpterminalshelfDo
	Assign(attrs ...field.AssignExpr) ILcpterminalshelfDo
	Joins(fields ...field.RelationField) ILcpterminalshelfDo
	Preload(fields ...field.RelationField) ILcpterminalshelfDo
	FirstOrInit() (*model.Lcpterminalshelf, error)
	FirstOrCreate() (*model.Lcpterminalshelf, error)
	FindByPage(offset int, limit int) (result []*model.Lcpterminalshelf, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILcpterminalshelfDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l lcpterminalshelfDo) Debug() ILcpterminalshelfDo {
	return l.withDO(l.DO.Debug())
}

func (l lcpterminalshelfDo) WithContext(ctx context.Context) ILcpterminalshelfDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l lcpterminalshelfDo) ReadDB() ILcpterminalshelfDo {
	return l.Clauses(dbresolver.Read)
}

func (l lcpterminalshelfDo) WriteDB() ILcpterminalshelfDo {
	return l.Clauses(dbresolver.Write)
}

func (l lcpterminalshelfDo) Clauses(conds ...clause.Expression) ILcpterminalshelfDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l lcpterminalshelfDo) Returning(value interface{}, columns ...string) ILcpterminalshelfDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l lcpterminalshelfDo) Not(conds ...gen.Condition) ILcpterminalshelfDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l lcpterminalshelfDo) Or(conds ...gen.Condition) ILcpterminalshelfDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l lcpterminalshelfDo) Select(conds ...field.Expr) ILcpterminalshelfDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l lcpterminalshelfDo) Where(conds ...gen.Condition) ILcpterminalshelfDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l lcpterminalshelfDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ILcpterminalshelfDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l lcpterminalshelfDo) Order(conds ...field.Expr) ILcpterminalshelfDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l lcpterminalshelfDo) Distinct(cols ...field.Expr) ILcpterminalshelfDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l lcpterminalshelfDo) Omit(cols ...field.Expr) ILcpterminalshelfDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l lcpterminalshelfDo) Join(table schema.Tabler, on ...field.Expr) ILcpterminalshelfDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l lcpterminalshelfDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILcpterminalshelfDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l lcpterminalshelfDo) RightJoin(table schema.Tabler, on ...field.Expr) ILcpterminalshelfDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l lcpterminalshelfDo) Group(cols ...field.Expr) ILcpterminalshelfDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l lcpterminalshelfDo) Having(conds ...gen.Condition) ILcpterminalshelfDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l lcpterminalshelfDo) Limit(limit int) ILcpterminalshelfDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l lcpterminalshelfDo) Offset(offset int) ILcpterminalshelfDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l lcpterminalshelfDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILcpterminalshelfDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l lcpterminalshelfDo) Unscoped() ILcpterminalshelfDo {
	return l.withDO(l.DO.Unscoped())
}

func (l lcpterminalshelfDo) Create(values ...*model.Lcpterminalshelf) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l lcpterminalshelfDo) CreateInBatches(values []*model.Lcpterminalshelf, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l lcpterminalshelfDo) Save(values ...*model.Lcpterminalshelf) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l lcpterminalshelfDo) First() (*model.Lcpterminalshelf, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Lcpterminalshelf), nil
	}
}

func (l lcpterminalshelfDo) Take() (*model.Lcpterminalshelf, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Lcpterminalshelf), nil
	}
}

func (l lcpterminalshelfDo) Last() (*model.Lcpterminalshelf, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Lcpterminalshelf), nil
	}
}

func (l lcpterminalshelfDo) Find() ([]*model.Lcpterminalshelf, error) {
	result, err := l.DO.Find()
	return result.([]*model.Lcpterminalshelf), err
}

func (l lcpterminalshelfDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Lcpterminalshelf, err error) {
	buf := make([]*model.Lcpterminalshelf, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l lcpterminalshelfDo) FindInBatches(result *[]*model.Lcpterminalshelf, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l lcpterminalshelfDo) Attrs(attrs ...field.AssignExpr) ILcpterminalshelfDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l lcpterminalshelfDo) Assign(attrs ...field.AssignExpr) ILcpterminalshelfDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l lcpterminalshelfDo) Joins(fields ...field.RelationField) ILcpterminalshelfDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l lcpterminalshelfDo) Preload(fields ...field.RelationField) ILcpterminalshelfDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l lcpterminalshelfDo) FirstOrInit() (*model.Lcpterminalshelf, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Lcpterminalshelf), nil
	}
}

func (l lcpterminalshelfDo) FirstOrCreate() (*model.Lcpterminalshelf, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Lcpterminalshelf), nil
	}
}

func (l lcpterminalshelfDo) FindByPage(offset int, limit int) (result []*model.Lcpterminalshelf, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l lcpterminalshelfDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l lcpterminalshelfDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l lcpterminalshelfDo) Delete(models ...*model.Lcpterminalshelf) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *lcpterminalshelfDo) withDO(do gen.Dao) *lcpterminalshelfDo {
	l.DO = *do.(*gen.DO)
	return l
}
