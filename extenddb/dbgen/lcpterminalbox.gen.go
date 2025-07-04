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

func newLcpterminalbox(db *gorm.DB) lcpterminalbox {
	_lcpterminalbox := lcpterminalbox{}

	_lcpterminalbox.lcpterminalboxDo.UseDB(db)
	_lcpterminalbox.lcpterminalboxDo.UseModel(&model.Lcpterminalbox{})

	tableName := _lcpterminalbox.lcpterminalboxDo.TableName()
	_lcpterminalbox.ALL = field.NewAsterisk(tableName)
	_lcpterminalbox.ID = field.NewString(tableName, "Id")
	_lcpterminalbox.CreationTime = field.NewTime(tableName, "CreationTime")
	_lcpterminalbox.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_lcpterminalbox.LastModificationTime = field.NewTime(tableName, "LastModificationTime")
	_lcpterminalbox.LastModifierUserID = field.NewInt64(tableName, "LastModifierUserId")
	_lcpterminalbox.Name = field.NewString(tableName, "Name")
	_lcpterminalbox.TerminalID = field.NewString(tableName, "TerminalId")
	_lcpterminalbox.TerminalCode = field.NewString(tableName, "TerminalCode")
	_lcpterminalbox.TerminalName = field.NewString(tableName, "TerminalName")
	_lcpterminalbox.IsFull = field.NewField(tableName, "IsFull")
	_lcpterminalbox.IsDisable = field.NewField(tableName, "IsDisable")
	_lcpterminalbox.IsEnable = field.NewField(tableName, "IsEnable")
	_lcpterminalbox.DisableReason = field.NewString(tableName, "DisableReason")
	_lcpterminalbox.TenantID = field.NewInt64(tableName, "TenantId")

	_lcpterminalbox.fillFieldMap()

	return _lcpterminalbox
}

type lcpterminalbox struct {
	lcpterminalboxDo lcpterminalboxDo

	ALL                  field.Asterisk
	ID                   field.String
	CreationTime         field.Time
	CreatorUserID        field.Int64
	LastModificationTime field.Time
	LastModifierUserID   field.Int64
	Name                 field.String
	TerminalID           field.String
	TerminalCode         field.String
	TerminalName         field.String
	IsFull               field.Field
	IsDisable            field.Field
	IsEnable             field.Field
	DisableReason        field.String
	TenantID             field.Int64

	fieldMap map[string]field.Expr
}

func (l lcpterminalbox) Table(newTableName string) *lcpterminalbox {
	l.lcpterminalboxDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l lcpterminalbox) As(alias string) *lcpterminalbox {
	l.lcpterminalboxDo.DO = *(l.lcpterminalboxDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *lcpterminalbox) updateTableName(table string) *lcpterminalbox {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewString(table, "Id")
	l.CreationTime = field.NewTime(table, "CreationTime")
	l.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	l.LastModificationTime = field.NewTime(table, "LastModificationTime")
	l.LastModifierUserID = field.NewInt64(table, "LastModifierUserId")
	l.Name = field.NewString(table, "Name")
	l.TerminalID = field.NewString(table, "TerminalId")
	l.TerminalCode = field.NewString(table, "TerminalCode")
	l.TerminalName = field.NewString(table, "TerminalName")
	l.IsFull = field.NewField(table, "IsFull")
	l.IsDisable = field.NewField(table, "IsDisable")
	l.IsEnable = field.NewField(table, "IsEnable")
	l.DisableReason = field.NewString(table, "DisableReason")
	l.TenantID = field.NewInt64(table, "TenantId")

	l.fillFieldMap()

	return l
}

func (l *lcpterminalbox) WithContext(ctx context.Context) ILcpterminalboxDo {
	return l.lcpterminalboxDo.WithContext(ctx)
}

func (l lcpterminalbox) TableName() string { return l.lcpterminalboxDo.TableName() }

func (l lcpterminalbox) Alias() string { return l.lcpterminalboxDo.Alias() }

func (l *lcpterminalbox) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *lcpterminalbox) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 14)
	l.fieldMap["Id"] = l.ID
	l.fieldMap["CreationTime"] = l.CreationTime
	l.fieldMap["CreatorUserId"] = l.CreatorUserID
	l.fieldMap["LastModificationTime"] = l.LastModificationTime
	l.fieldMap["LastModifierUserId"] = l.LastModifierUserID
	l.fieldMap["Name"] = l.Name
	l.fieldMap["TerminalId"] = l.TerminalID
	l.fieldMap["TerminalCode"] = l.TerminalCode
	l.fieldMap["TerminalName"] = l.TerminalName
	l.fieldMap["IsFull"] = l.IsFull
	l.fieldMap["IsDisable"] = l.IsDisable
	l.fieldMap["IsEnable"] = l.IsEnable
	l.fieldMap["DisableReason"] = l.DisableReason
	l.fieldMap["TenantId"] = l.TenantID
}

func (l lcpterminalbox) clone(db *gorm.DB) lcpterminalbox {
	l.lcpterminalboxDo.ReplaceDB(db)
	return l
}

type lcpterminalboxDo struct{ gen.DO }

type ILcpterminalboxDo interface {
	gen.SubQuery
	Debug() ILcpterminalboxDo
	WithContext(ctx context.Context) ILcpterminalboxDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILcpterminalboxDo
	Not(conds ...gen.Condition) ILcpterminalboxDo
	Or(conds ...gen.Condition) ILcpterminalboxDo
	Select(conds ...field.Expr) ILcpterminalboxDo
	Where(conds ...gen.Condition) ILcpterminalboxDo
	Order(conds ...field.Expr) ILcpterminalboxDo
	Distinct(cols ...field.Expr) ILcpterminalboxDo
	Omit(cols ...field.Expr) ILcpterminalboxDo
	Join(table schema.Tabler, on ...field.Expr) ILcpterminalboxDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILcpterminalboxDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILcpterminalboxDo
	Group(cols ...field.Expr) ILcpterminalboxDo
	Having(conds ...gen.Condition) ILcpterminalboxDo
	Limit(limit int) ILcpterminalboxDo
	Offset(offset int) ILcpterminalboxDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILcpterminalboxDo
	Unscoped() ILcpterminalboxDo
	Create(values ...*model.Lcpterminalbox) error
	CreateInBatches(values []*model.Lcpterminalbox, batchSize int) error
	Save(values ...*model.Lcpterminalbox) error
	First() (*model.Lcpterminalbox, error)
	Take() (*model.Lcpterminalbox, error)
	Last() (*model.Lcpterminalbox, error)
	Find() ([]*model.Lcpterminalbox, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Lcpterminalbox, err error)
	FindInBatches(result *[]*model.Lcpterminalbox, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Lcpterminalbox) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILcpterminalboxDo
	Assign(attrs ...field.AssignExpr) ILcpterminalboxDo
	Joins(fields ...field.RelationField) ILcpterminalboxDo
	Preload(fields ...field.RelationField) ILcpterminalboxDo
	FirstOrInit() (*model.Lcpterminalbox, error)
	FirstOrCreate() (*model.Lcpterminalbox, error)
	FindByPage(offset int, limit int) (result []*model.Lcpterminalbox, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILcpterminalboxDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l lcpterminalboxDo) Debug() ILcpterminalboxDo {
	return l.withDO(l.DO.Debug())
}

func (l lcpterminalboxDo) WithContext(ctx context.Context) ILcpterminalboxDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l lcpterminalboxDo) ReadDB() ILcpterminalboxDo {
	return l.Clauses(dbresolver.Read)
}

func (l lcpterminalboxDo) WriteDB() ILcpterminalboxDo {
	return l.Clauses(dbresolver.Write)
}

func (l lcpterminalboxDo) Clauses(conds ...clause.Expression) ILcpterminalboxDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l lcpterminalboxDo) Returning(value interface{}, columns ...string) ILcpterminalboxDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l lcpterminalboxDo) Not(conds ...gen.Condition) ILcpterminalboxDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l lcpterminalboxDo) Or(conds ...gen.Condition) ILcpterminalboxDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l lcpterminalboxDo) Select(conds ...field.Expr) ILcpterminalboxDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l lcpterminalboxDo) Where(conds ...gen.Condition) ILcpterminalboxDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l lcpterminalboxDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ILcpterminalboxDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l lcpterminalboxDo) Order(conds ...field.Expr) ILcpterminalboxDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l lcpterminalboxDo) Distinct(cols ...field.Expr) ILcpterminalboxDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l lcpterminalboxDo) Omit(cols ...field.Expr) ILcpterminalboxDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l lcpterminalboxDo) Join(table schema.Tabler, on ...field.Expr) ILcpterminalboxDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l lcpterminalboxDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILcpterminalboxDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l lcpterminalboxDo) RightJoin(table schema.Tabler, on ...field.Expr) ILcpterminalboxDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l lcpterminalboxDo) Group(cols ...field.Expr) ILcpterminalboxDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l lcpterminalboxDo) Having(conds ...gen.Condition) ILcpterminalboxDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l lcpterminalboxDo) Limit(limit int) ILcpterminalboxDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l lcpterminalboxDo) Offset(offset int) ILcpterminalboxDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l lcpterminalboxDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILcpterminalboxDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l lcpterminalboxDo) Unscoped() ILcpterminalboxDo {
	return l.withDO(l.DO.Unscoped())
}

func (l lcpterminalboxDo) Create(values ...*model.Lcpterminalbox) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l lcpterminalboxDo) CreateInBatches(values []*model.Lcpterminalbox, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l lcpterminalboxDo) Save(values ...*model.Lcpterminalbox) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l lcpterminalboxDo) First() (*model.Lcpterminalbox, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Lcpterminalbox), nil
	}
}

func (l lcpterminalboxDo) Take() (*model.Lcpterminalbox, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Lcpterminalbox), nil
	}
}

func (l lcpterminalboxDo) Last() (*model.Lcpterminalbox, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Lcpterminalbox), nil
	}
}

func (l lcpterminalboxDo) Find() ([]*model.Lcpterminalbox, error) {
	result, err := l.DO.Find()
	return result.([]*model.Lcpterminalbox), err
}

func (l lcpterminalboxDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Lcpterminalbox, err error) {
	buf := make([]*model.Lcpterminalbox, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l lcpterminalboxDo) FindInBatches(result *[]*model.Lcpterminalbox, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l lcpterminalboxDo) Attrs(attrs ...field.AssignExpr) ILcpterminalboxDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l lcpterminalboxDo) Assign(attrs ...field.AssignExpr) ILcpterminalboxDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l lcpterminalboxDo) Joins(fields ...field.RelationField) ILcpterminalboxDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l lcpterminalboxDo) Preload(fields ...field.RelationField) ILcpterminalboxDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l lcpterminalboxDo) FirstOrInit() (*model.Lcpterminalbox, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Lcpterminalbox), nil
	}
}

func (l lcpterminalboxDo) FirstOrCreate() (*model.Lcpterminalbox, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Lcpterminalbox), nil
	}
}

func (l lcpterminalboxDo) FindByPage(offset int, limit int) (result []*model.Lcpterminalbox, count int64, err error) {
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

func (l lcpterminalboxDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l lcpterminalboxDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l lcpterminalboxDo) Delete(models ...*model.Lcpterminalbox) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *lcpterminalboxDo) withDO(do gen.Dao) *lcpterminalboxDo {
	l.DO = *do.(*gen.DO)
	return l
}
