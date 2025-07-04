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

func newLibiteminventorylog(db *gorm.DB) libiteminventorylog {
	_libiteminventorylog := libiteminventorylog{}

	_libiteminventorylog.libiteminventorylogDo.UseDB(db)
	_libiteminventorylog.libiteminventorylogDo.UseModel(&model.Libiteminventorylog{})

	tableName := _libiteminventorylog.libiteminventorylogDo.TableName()
	_libiteminventorylog.ALL = field.NewAsterisk(tableName)
	_libiteminventorylog.ID = field.NewString(tableName, "Id")
	_libiteminventorylog.CreationTime = field.NewTime(tableName, "CreationTime")
	_libiteminventorylog.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_libiteminventorylog.LayerID = field.NewString(tableName, "LayerId")
	_libiteminventorylog.LayerName = field.NewString(tableName, "LayerName")
	_libiteminventorylog.ItemBarcode = field.NewString(tableName, "ItemBarcode")
	_libiteminventorylog.InventoryWorkType = field.NewInt64(tableName, "InventoryWorkType")
	_libiteminventorylog.InventoryState = field.NewInt64(tableName, "InventoryState")
	_libiteminventorylog.OffShelfTime = field.NewTime(tableName, "OffShelfTime")
	_libiteminventorylog.Remark = field.NewString(tableName, "Remark")
	_libiteminventorylog.TenantID = field.NewInt64(tableName, "TenantId")
	_libiteminventorylog.InventoryWorkID = field.NewString(tableName, "InventoryWorkId")
	_libiteminventorylog.CreatorUserName = field.NewString(tableName, "CreatorUserName")
	_libiteminventorylog.ItemCallNo = field.NewString(tableName, "ItemCallNo")
	_libiteminventorylog.ItemTitle = field.NewString(tableName, "ItemTitle")
	_libiteminventorylog.LayerCode = field.NewString(tableName, "LayerCode")
	_libiteminventorylog.LocationID = field.NewString(tableName, "LocationId")
	_libiteminventorylog.LocationName = field.NewString(tableName, "LocationName")
	_libiteminventorylog.LocLayerCode = field.NewString(tableName, "LocLayerCode")
	_libiteminventorylog.LocLayerID = field.NewString(tableName, "LocLayerId")
	_libiteminventorylog.LocLayerName = field.NewString(tableName, "LocLayerName")
	_libiteminventorylog.ExceptionMsg = field.NewString(tableName, "ExceptionMsg")
	_libiteminventorylog.ItemAuthor = field.NewString(tableName, "ItemAuthor")
	_libiteminventorylog.ItemISBN = field.NewString(tableName, "ItemISBN")
	_libiteminventorylog.ItemPublisher = field.NewString(tableName, "ItemPublisher")
	_libiteminventorylog.OriginType = field.NewInt64(tableName, "OriginType")

	_libiteminventorylog.fillFieldMap()

	return _libiteminventorylog
}

type libiteminventorylog struct {
	libiteminventorylogDo libiteminventorylogDo

	ALL               field.Asterisk
	ID                field.String
	CreationTime      field.Time
	CreatorUserID     field.Int64
	LayerID           field.String
	LayerName         field.String
	ItemBarcode       field.String
	InventoryWorkType field.Int64
	InventoryState    field.Int64
	OffShelfTime      field.Time
	Remark            field.String
	TenantID          field.Int64
	InventoryWorkID   field.String
	CreatorUserName   field.String
	ItemCallNo        field.String
	ItemTitle         field.String
	LayerCode         field.String
	LocationID        field.String
	LocationName      field.String
	LocLayerCode      field.String
	LocLayerID        field.String
	LocLayerName      field.String
	ExceptionMsg      field.String
	ItemAuthor        field.String
	ItemISBN          field.String
	ItemPublisher     field.String
	OriginType        field.Int64

	fieldMap map[string]field.Expr
}

func (l libiteminventorylog) Table(newTableName string) *libiteminventorylog {
	l.libiteminventorylogDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l libiteminventorylog) As(alias string) *libiteminventorylog {
	l.libiteminventorylogDo.DO = *(l.libiteminventorylogDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *libiteminventorylog) updateTableName(table string) *libiteminventorylog {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewString(table, "Id")
	l.CreationTime = field.NewTime(table, "CreationTime")
	l.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	l.LayerID = field.NewString(table, "LayerId")
	l.LayerName = field.NewString(table, "LayerName")
	l.ItemBarcode = field.NewString(table, "ItemBarcode")
	l.InventoryWorkType = field.NewInt64(table, "InventoryWorkType")
	l.InventoryState = field.NewInt64(table, "InventoryState")
	l.OffShelfTime = field.NewTime(table, "OffShelfTime")
	l.Remark = field.NewString(table, "Remark")
	l.TenantID = field.NewInt64(table, "TenantId")
	l.InventoryWorkID = field.NewString(table, "InventoryWorkId")
	l.CreatorUserName = field.NewString(table, "CreatorUserName")
	l.ItemCallNo = field.NewString(table, "ItemCallNo")
	l.ItemTitle = field.NewString(table, "ItemTitle")
	l.LayerCode = field.NewString(table, "LayerCode")
	l.LocationID = field.NewString(table, "LocationId")
	l.LocationName = field.NewString(table, "LocationName")
	l.LocLayerCode = field.NewString(table, "LocLayerCode")
	l.LocLayerID = field.NewString(table, "LocLayerId")
	l.LocLayerName = field.NewString(table, "LocLayerName")
	l.ExceptionMsg = field.NewString(table, "ExceptionMsg")
	l.ItemAuthor = field.NewString(table, "ItemAuthor")
	l.ItemISBN = field.NewString(table, "ItemISBN")
	l.ItemPublisher = field.NewString(table, "ItemPublisher")
	l.OriginType = field.NewInt64(table, "OriginType")

	l.fillFieldMap()

	return l
}

func (l *libiteminventorylog) WithContext(ctx context.Context) ILibiteminventorylogDo {
	return l.libiteminventorylogDo.WithContext(ctx)
}

func (l libiteminventorylog) TableName() string { return l.libiteminventorylogDo.TableName() }

func (l libiteminventorylog) Alias() string { return l.libiteminventorylogDo.Alias() }

func (l *libiteminventorylog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *libiteminventorylog) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 26)
	l.fieldMap["Id"] = l.ID
	l.fieldMap["CreationTime"] = l.CreationTime
	l.fieldMap["CreatorUserId"] = l.CreatorUserID
	l.fieldMap["LayerId"] = l.LayerID
	l.fieldMap["LayerName"] = l.LayerName
	l.fieldMap["ItemBarcode"] = l.ItemBarcode
	l.fieldMap["InventoryWorkType"] = l.InventoryWorkType
	l.fieldMap["InventoryState"] = l.InventoryState
	l.fieldMap["OffShelfTime"] = l.OffShelfTime
	l.fieldMap["Remark"] = l.Remark
	l.fieldMap["TenantId"] = l.TenantID
	l.fieldMap["InventoryWorkId"] = l.InventoryWorkID
	l.fieldMap["CreatorUserName"] = l.CreatorUserName
	l.fieldMap["ItemCallNo"] = l.ItemCallNo
	l.fieldMap["ItemTitle"] = l.ItemTitle
	l.fieldMap["LayerCode"] = l.LayerCode
	l.fieldMap["LocationId"] = l.LocationID
	l.fieldMap["LocationName"] = l.LocationName
	l.fieldMap["LocLayerCode"] = l.LocLayerCode
	l.fieldMap["LocLayerId"] = l.LocLayerID
	l.fieldMap["LocLayerName"] = l.LocLayerName
	l.fieldMap["ExceptionMsg"] = l.ExceptionMsg
	l.fieldMap["ItemAuthor"] = l.ItemAuthor
	l.fieldMap["ItemISBN"] = l.ItemISBN
	l.fieldMap["ItemPublisher"] = l.ItemPublisher
	l.fieldMap["OriginType"] = l.OriginType
}

func (l libiteminventorylog) clone(db *gorm.DB) libiteminventorylog {
	l.libiteminventorylogDo.ReplaceDB(db)
	return l
}

type libiteminventorylogDo struct{ gen.DO }

type ILibiteminventorylogDo interface {
	gen.SubQuery
	Debug() ILibiteminventorylogDo
	WithContext(ctx context.Context) ILibiteminventorylogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILibiteminventorylogDo
	Not(conds ...gen.Condition) ILibiteminventorylogDo
	Or(conds ...gen.Condition) ILibiteminventorylogDo
	Select(conds ...field.Expr) ILibiteminventorylogDo
	Where(conds ...gen.Condition) ILibiteminventorylogDo
	Order(conds ...field.Expr) ILibiteminventorylogDo
	Distinct(cols ...field.Expr) ILibiteminventorylogDo
	Omit(cols ...field.Expr) ILibiteminventorylogDo
	Join(table schema.Tabler, on ...field.Expr) ILibiteminventorylogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILibiteminventorylogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILibiteminventorylogDo
	Group(cols ...field.Expr) ILibiteminventorylogDo
	Having(conds ...gen.Condition) ILibiteminventorylogDo
	Limit(limit int) ILibiteminventorylogDo
	Offset(offset int) ILibiteminventorylogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILibiteminventorylogDo
	Unscoped() ILibiteminventorylogDo
	Create(values ...*model.Libiteminventorylog) error
	CreateInBatches(values []*model.Libiteminventorylog, batchSize int) error
	Save(values ...*model.Libiteminventorylog) error
	First() (*model.Libiteminventorylog, error)
	Take() (*model.Libiteminventorylog, error)
	Last() (*model.Libiteminventorylog, error)
	Find() ([]*model.Libiteminventorylog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Libiteminventorylog, err error)
	FindInBatches(result *[]*model.Libiteminventorylog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Libiteminventorylog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILibiteminventorylogDo
	Assign(attrs ...field.AssignExpr) ILibiteminventorylogDo
	Joins(fields ...field.RelationField) ILibiteminventorylogDo
	Preload(fields ...field.RelationField) ILibiteminventorylogDo
	FirstOrInit() (*model.Libiteminventorylog, error)
	FirstOrCreate() (*model.Libiteminventorylog, error)
	FindByPage(offset int, limit int) (result []*model.Libiteminventorylog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILibiteminventorylogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l libiteminventorylogDo) Debug() ILibiteminventorylogDo {
	return l.withDO(l.DO.Debug())
}

func (l libiteminventorylogDo) WithContext(ctx context.Context) ILibiteminventorylogDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l libiteminventorylogDo) ReadDB() ILibiteminventorylogDo {
	return l.Clauses(dbresolver.Read)
}

func (l libiteminventorylogDo) WriteDB() ILibiteminventorylogDo {
	return l.Clauses(dbresolver.Write)
}

func (l libiteminventorylogDo) Clauses(conds ...clause.Expression) ILibiteminventorylogDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l libiteminventorylogDo) Returning(value interface{}, columns ...string) ILibiteminventorylogDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l libiteminventorylogDo) Not(conds ...gen.Condition) ILibiteminventorylogDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l libiteminventorylogDo) Or(conds ...gen.Condition) ILibiteminventorylogDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l libiteminventorylogDo) Select(conds ...field.Expr) ILibiteminventorylogDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l libiteminventorylogDo) Where(conds ...gen.Condition) ILibiteminventorylogDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l libiteminventorylogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ILibiteminventorylogDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l libiteminventorylogDo) Order(conds ...field.Expr) ILibiteminventorylogDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l libiteminventorylogDo) Distinct(cols ...field.Expr) ILibiteminventorylogDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l libiteminventorylogDo) Omit(cols ...field.Expr) ILibiteminventorylogDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l libiteminventorylogDo) Join(table schema.Tabler, on ...field.Expr) ILibiteminventorylogDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l libiteminventorylogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILibiteminventorylogDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l libiteminventorylogDo) RightJoin(table schema.Tabler, on ...field.Expr) ILibiteminventorylogDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l libiteminventorylogDo) Group(cols ...field.Expr) ILibiteminventorylogDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l libiteminventorylogDo) Having(conds ...gen.Condition) ILibiteminventorylogDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l libiteminventorylogDo) Limit(limit int) ILibiteminventorylogDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l libiteminventorylogDo) Offset(offset int) ILibiteminventorylogDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l libiteminventorylogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILibiteminventorylogDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l libiteminventorylogDo) Unscoped() ILibiteminventorylogDo {
	return l.withDO(l.DO.Unscoped())
}

func (l libiteminventorylogDo) Create(values ...*model.Libiteminventorylog) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l libiteminventorylogDo) CreateInBatches(values []*model.Libiteminventorylog, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l libiteminventorylogDo) Save(values ...*model.Libiteminventorylog) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l libiteminventorylogDo) First() (*model.Libiteminventorylog, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libiteminventorylog), nil
	}
}

func (l libiteminventorylogDo) Take() (*model.Libiteminventorylog, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libiteminventorylog), nil
	}
}

func (l libiteminventorylogDo) Last() (*model.Libiteminventorylog, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libiteminventorylog), nil
	}
}

func (l libiteminventorylogDo) Find() ([]*model.Libiteminventorylog, error) {
	result, err := l.DO.Find()
	return result.([]*model.Libiteminventorylog), err
}

func (l libiteminventorylogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Libiteminventorylog, err error) {
	buf := make([]*model.Libiteminventorylog, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l libiteminventorylogDo) FindInBatches(result *[]*model.Libiteminventorylog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l libiteminventorylogDo) Attrs(attrs ...field.AssignExpr) ILibiteminventorylogDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l libiteminventorylogDo) Assign(attrs ...field.AssignExpr) ILibiteminventorylogDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l libiteminventorylogDo) Joins(fields ...field.RelationField) ILibiteminventorylogDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l libiteminventorylogDo) Preload(fields ...field.RelationField) ILibiteminventorylogDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l libiteminventorylogDo) FirstOrInit() (*model.Libiteminventorylog, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libiteminventorylog), nil
	}
}

func (l libiteminventorylogDo) FirstOrCreate() (*model.Libiteminventorylog, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libiteminventorylog), nil
	}
}

func (l libiteminventorylogDo) FindByPage(offset int, limit int) (result []*model.Libiteminventorylog, count int64, err error) {
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

func (l libiteminventorylogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l libiteminventorylogDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l libiteminventorylogDo) Delete(models ...*model.Libiteminventorylog) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *libiteminventorylogDo) withDO(do gen.Dao) *libiteminventorylogDo {
	l.DO = *do.(*gen.DO)
	return l
}
