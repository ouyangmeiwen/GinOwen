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

func newLibupdatefirstbooklog(db *gorm.DB) libupdatefirstbooklog {
	_libupdatefirstbooklog := libupdatefirstbooklog{}

	_libupdatefirstbooklog.libupdatefirstbooklogDo.UseDB(db)
	_libupdatefirstbooklog.libupdatefirstbooklogDo.UseModel(&model.Libupdatefirstbooklog{})

	tableName := _libupdatefirstbooklog.libupdatefirstbooklogDo.TableName()
	_libupdatefirstbooklog.ALL = field.NewAsterisk(tableName)
	_libupdatefirstbooklog.ID = field.NewString(tableName, "Id")
	_libupdatefirstbooklog.CreationTime = field.NewTime(tableName, "CreationTime")
	_libupdatefirstbooklog.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_libupdatefirstbooklog.LayerID = field.NewString(tableName, "LayerId")
	_libupdatefirstbooklog.LayerName = field.NewString(tableName, "LayerName")
	_libupdatefirstbooklog.ItemID = field.NewString(tableName, "ItemId")
	_libupdatefirstbooklog.ItemBarcode = field.NewString(tableName, "ItemBarcode")
	_libupdatefirstbooklog.ItemCallNo = field.NewString(tableName, "ItemCallNo")
	_libupdatefirstbooklog.PretendCallNo = field.NewString(tableName, "PretendCallNo")
	_libupdatefirstbooklog.PreItemID = field.NewString(tableName, "PreItemId")
	_libupdatefirstbooklog.PreItemBarcode = field.NewString(tableName, "PreItemBarcode")
	_libupdatefirstbooklog.PreItemCallNo = field.NewString(tableName, "PreItemCallNo")
	_libupdatefirstbooklog.PrePretendCallNo = field.NewString(tableName, "PrePretendCallNo")
	_libupdatefirstbooklog.Remark = field.NewString(tableName, "Remark")
	_libupdatefirstbooklog.TenantID = field.NewInt64(tableName, "TenantId")
	_libupdatefirstbooklog.CreatorUserName = field.NewString(tableName, "CreatorUserName")
	_libupdatefirstbooklog.ItemTitle = field.NewString(tableName, "ItemTitle")
	_libupdatefirstbooklog.LayerCode = field.NewString(tableName, "LayerCode")
	_libupdatefirstbooklog.PreItemTitle = field.NewString(tableName, "PreItemTitle")

	_libupdatefirstbooklog.fillFieldMap()

	return _libupdatefirstbooklog
}

type libupdatefirstbooklog struct {
	libupdatefirstbooklogDo libupdatefirstbooklogDo

	ALL              field.Asterisk
	ID               field.String
	CreationTime     field.Time
	CreatorUserID    field.Int64
	LayerID          field.String
	LayerName        field.String
	ItemID           field.String
	ItemBarcode      field.String
	ItemCallNo       field.String
	PretendCallNo    field.String
	PreItemID        field.String
	PreItemBarcode   field.String
	PreItemCallNo    field.String
	PrePretendCallNo field.String
	Remark           field.String
	TenantID         field.Int64
	CreatorUserName  field.String
	ItemTitle        field.String
	LayerCode        field.String
	PreItemTitle     field.String

	fieldMap map[string]field.Expr
}

func (l libupdatefirstbooklog) Table(newTableName string) *libupdatefirstbooklog {
	l.libupdatefirstbooklogDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l libupdatefirstbooklog) As(alias string) *libupdatefirstbooklog {
	l.libupdatefirstbooklogDo.DO = *(l.libupdatefirstbooklogDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *libupdatefirstbooklog) updateTableName(table string) *libupdatefirstbooklog {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewString(table, "Id")
	l.CreationTime = field.NewTime(table, "CreationTime")
	l.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	l.LayerID = field.NewString(table, "LayerId")
	l.LayerName = field.NewString(table, "LayerName")
	l.ItemID = field.NewString(table, "ItemId")
	l.ItemBarcode = field.NewString(table, "ItemBarcode")
	l.ItemCallNo = field.NewString(table, "ItemCallNo")
	l.PretendCallNo = field.NewString(table, "PretendCallNo")
	l.PreItemID = field.NewString(table, "PreItemId")
	l.PreItemBarcode = field.NewString(table, "PreItemBarcode")
	l.PreItemCallNo = field.NewString(table, "PreItemCallNo")
	l.PrePretendCallNo = field.NewString(table, "PrePretendCallNo")
	l.Remark = field.NewString(table, "Remark")
	l.TenantID = field.NewInt64(table, "TenantId")
	l.CreatorUserName = field.NewString(table, "CreatorUserName")
	l.ItemTitle = field.NewString(table, "ItemTitle")
	l.LayerCode = field.NewString(table, "LayerCode")
	l.PreItemTitle = field.NewString(table, "PreItemTitle")

	l.fillFieldMap()

	return l
}

func (l *libupdatefirstbooklog) WithContext(ctx context.Context) ILibupdatefirstbooklogDo {
	return l.libupdatefirstbooklogDo.WithContext(ctx)
}

func (l libupdatefirstbooklog) TableName() string { return l.libupdatefirstbooklogDo.TableName() }

func (l libupdatefirstbooklog) Alias() string { return l.libupdatefirstbooklogDo.Alias() }

func (l *libupdatefirstbooklog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *libupdatefirstbooklog) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 19)
	l.fieldMap["Id"] = l.ID
	l.fieldMap["CreationTime"] = l.CreationTime
	l.fieldMap["CreatorUserId"] = l.CreatorUserID
	l.fieldMap["LayerId"] = l.LayerID
	l.fieldMap["LayerName"] = l.LayerName
	l.fieldMap["ItemId"] = l.ItemID
	l.fieldMap["ItemBarcode"] = l.ItemBarcode
	l.fieldMap["ItemCallNo"] = l.ItemCallNo
	l.fieldMap["PretendCallNo"] = l.PretendCallNo
	l.fieldMap["PreItemId"] = l.PreItemID
	l.fieldMap["PreItemBarcode"] = l.PreItemBarcode
	l.fieldMap["PreItemCallNo"] = l.PreItemCallNo
	l.fieldMap["PrePretendCallNo"] = l.PrePretendCallNo
	l.fieldMap["Remark"] = l.Remark
	l.fieldMap["TenantId"] = l.TenantID
	l.fieldMap["CreatorUserName"] = l.CreatorUserName
	l.fieldMap["ItemTitle"] = l.ItemTitle
	l.fieldMap["LayerCode"] = l.LayerCode
	l.fieldMap["PreItemTitle"] = l.PreItemTitle
}

func (l libupdatefirstbooklog) clone(db *gorm.DB) libupdatefirstbooklog {
	l.libupdatefirstbooklogDo.ReplaceDB(db)
	return l
}

type libupdatefirstbooklogDo struct{ gen.DO }

type ILibupdatefirstbooklogDo interface {
	gen.SubQuery
	Debug() ILibupdatefirstbooklogDo
	WithContext(ctx context.Context) ILibupdatefirstbooklogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILibupdatefirstbooklogDo
	Not(conds ...gen.Condition) ILibupdatefirstbooklogDo
	Or(conds ...gen.Condition) ILibupdatefirstbooklogDo
	Select(conds ...field.Expr) ILibupdatefirstbooklogDo
	Where(conds ...gen.Condition) ILibupdatefirstbooklogDo
	Order(conds ...field.Expr) ILibupdatefirstbooklogDo
	Distinct(cols ...field.Expr) ILibupdatefirstbooklogDo
	Omit(cols ...field.Expr) ILibupdatefirstbooklogDo
	Join(table schema.Tabler, on ...field.Expr) ILibupdatefirstbooklogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILibupdatefirstbooklogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILibupdatefirstbooklogDo
	Group(cols ...field.Expr) ILibupdatefirstbooklogDo
	Having(conds ...gen.Condition) ILibupdatefirstbooklogDo
	Limit(limit int) ILibupdatefirstbooklogDo
	Offset(offset int) ILibupdatefirstbooklogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILibupdatefirstbooklogDo
	Unscoped() ILibupdatefirstbooklogDo
	Create(values ...*model.Libupdatefirstbooklog) error
	CreateInBatches(values []*model.Libupdatefirstbooklog, batchSize int) error
	Save(values ...*model.Libupdatefirstbooklog) error
	First() (*model.Libupdatefirstbooklog, error)
	Take() (*model.Libupdatefirstbooklog, error)
	Last() (*model.Libupdatefirstbooklog, error)
	Find() ([]*model.Libupdatefirstbooklog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Libupdatefirstbooklog, err error)
	FindInBatches(result *[]*model.Libupdatefirstbooklog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Libupdatefirstbooklog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILibupdatefirstbooklogDo
	Assign(attrs ...field.AssignExpr) ILibupdatefirstbooklogDo
	Joins(fields ...field.RelationField) ILibupdatefirstbooklogDo
	Preload(fields ...field.RelationField) ILibupdatefirstbooklogDo
	FirstOrInit() (*model.Libupdatefirstbooklog, error)
	FirstOrCreate() (*model.Libupdatefirstbooklog, error)
	FindByPage(offset int, limit int) (result []*model.Libupdatefirstbooklog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILibupdatefirstbooklogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l libupdatefirstbooklogDo) Debug() ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Debug())
}

func (l libupdatefirstbooklogDo) WithContext(ctx context.Context) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l libupdatefirstbooklogDo) ReadDB() ILibupdatefirstbooklogDo {
	return l.Clauses(dbresolver.Read)
}

func (l libupdatefirstbooklogDo) WriteDB() ILibupdatefirstbooklogDo {
	return l.Clauses(dbresolver.Write)
}

func (l libupdatefirstbooklogDo) Clauses(conds ...clause.Expression) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l libupdatefirstbooklogDo) Returning(value interface{}, columns ...string) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l libupdatefirstbooklogDo) Not(conds ...gen.Condition) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l libupdatefirstbooklogDo) Or(conds ...gen.Condition) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l libupdatefirstbooklogDo) Select(conds ...field.Expr) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l libupdatefirstbooklogDo) Where(conds ...gen.Condition) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l libupdatefirstbooklogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ILibupdatefirstbooklogDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l libupdatefirstbooklogDo) Order(conds ...field.Expr) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l libupdatefirstbooklogDo) Distinct(cols ...field.Expr) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l libupdatefirstbooklogDo) Omit(cols ...field.Expr) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l libupdatefirstbooklogDo) Join(table schema.Tabler, on ...field.Expr) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l libupdatefirstbooklogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l libupdatefirstbooklogDo) RightJoin(table schema.Tabler, on ...field.Expr) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l libupdatefirstbooklogDo) Group(cols ...field.Expr) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l libupdatefirstbooklogDo) Having(conds ...gen.Condition) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l libupdatefirstbooklogDo) Limit(limit int) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l libupdatefirstbooklogDo) Offset(offset int) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l libupdatefirstbooklogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l libupdatefirstbooklogDo) Unscoped() ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Unscoped())
}

func (l libupdatefirstbooklogDo) Create(values ...*model.Libupdatefirstbooklog) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l libupdatefirstbooklogDo) CreateInBatches(values []*model.Libupdatefirstbooklog, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l libupdatefirstbooklogDo) Save(values ...*model.Libupdatefirstbooklog) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l libupdatefirstbooklogDo) First() (*model.Libupdatefirstbooklog, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libupdatefirstbooklog), nil
	}
}

func (l libupdatefirstbooklogDo) Take() (*model.Libupdatefirstbooklog, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libupdatefirstbooklog), nil
	}
}

func (l libupdatefirstbooklogDo) Last() (*model.Libupdatefirstbooklog, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libupdatefirstbooklog), nil
	}
}

func (l libupdatefirstbooklogDo) Find() ([]*model.Libupdatefirstbooklog, error) {
	result, err := l.DO.Find()
	return result.([]*model.Libupdatefirstbooklog), err
}

func (l libupdatefirstbooklogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Libupdatefirstbooklog, err error) {
	buf := make([]*model.Libupdatefirstbooklog, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l libupdatefirstbooklogDo) FindInBatches(result *[]*model.Libupdatefirstbooklog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l libupdatefirstbooklogDo) Attrs(attrs ...field.AssignExpr) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l libupdatefirstbooklogDo) Assign(attrs ...field.AssignExpr) ILibupdatefirstbooklogDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l libupdatefirstbooklogDo) Joins(fields ...field.RelationField) ILibupdatefirstbooklogDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l libupdatefirstbooklogDo) Preload(fields ...field.RelationField) ILibupdatefirstbooklogDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l libupdatefirstbooklogDo) FirstOrInit() (*model.Libupdatefirstbooklog, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libupdatefirstbooklog), nil
	}
}

func (l libupdatefirstbooklogDo) FirstOrCreate() (*model.Libupdatefirstbooklog, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libupdatefirstbooklog), nil
	}
}

func (l libupdatefirstbooklogDo) FindByPage(offset int, limit int) (result []*model.Libupdatefirstbooklog, count int64, err error) {
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

func (l libupdatefirstbooklogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l libupdatefirstbooklogDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l libupdatefirstbooklogDo) Delete(models ...*model.Libupdatefirstbooklog) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *libupdatefirstbooklogDo) withDO(do gen.Dao) *libupdatefirstbooklogDo {
	l.DO = *do.(*gen.DO)
	return l
}
