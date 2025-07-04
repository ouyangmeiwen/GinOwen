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

func newLiblayer(db *gorm.DB) liblayer {
	_liblayer := liblayer{}

	_liblayer.liblayerDo.UseDB(db)
	_liblayer.liblayerDo.UseModel(&model.Liblayer{})

	tableName := _liblayer.liblayerDo.TableName()
	_liblayer.ALL = field.NewAsterisk(tableName)
	_liblayer.ID = field.NewString(tableName, "Id")
	_liblayer.CreationTime = field.NewTime(tableName, "CreationTime")
	_liblayer.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_liblayer.LastModificationTime = field.NewTime(tableName, "LastModificationTime")
	_liblayer.LastModifierUserID = field.NewInt64(tableName, "LastModifierUserId")
	_liblayer.IsDeleted = field.NewField(tableName, "IsDeleted")
	_liblayer.DeleterUserID = field.NewInt64(tableName, "DeleterUserId")
	_liblayer.DeletionTime = field.NewTime(tableName, "DeletionTime")
	_liblayer.ShelfID = field.NewString(tableName, "ShelfId")
	_liblayer.Code = field.NewString(tableName, "Code")
	_liblayer.Name = field.NewString(tableName, "Name")
	_liblayer.Tid = field.NewString(tableName, "Tid")
	_liblayer.Side = field.NewString(tableName, "Side")
	_liblayer.LayerNo = field.NewInt64(tableName, "LayerNo")
	_liblayer.IsEnable = field.NewField(tableName, "IsEnable")
	_liblayer.Remark = field.NewString(tableName, "Remark")
	_liblayer.TenantID = field.NewInt64(tableName, "TenantId")
	_liblayer.ItemBarcode = field.NewString(tableName, "ItemBarcode")
	_liblayer.ItemCallNo = field.NewString(tableName, "ItemCallNo")
	_liblayer.PreCallNo = field.NewString(tableName, "PreCallNo")
	_liblayer.Barcode = field.NewString(tableName, "Barcode")
	_liblayer.OriginType = field.NewInt64(tableName, "OriginType")

	_liblayer.fillFieldMap()

	return _liblayer
}

type liblayer struct {
	liblayerDo liblayerDo

	ALL                  field.Asterisk
	ID                   field.String
	CreationTime         field.Time
	CreatorUserID        field.Int64
	LastModificationTime field.Time
	LastModifierUserID   field.Int64
	IsDeleted            field.Field
	DeleterUserID        field.Int64
	DeletionTime         field.Time
	ShelfID              field.String
	Code                 field.String
	Name                 field.String
	Tid                  field.String
	Side                 field.String
	LayerNo              field.Int64
	IsEnable             field.Field
	Remark               field.String
	TenantID             field.Int64
	ItemBarcode          field.String
	ItemCallNo           field.String
	PreCallNo            field.String
	Barcode              field.String
	OriginType           field.Int64

	fieldMap map[string]field.Expr
}

func (l liblayer) Table(newTableName string) *liblayer {
	l.liblayerDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l liblayer) As(alias string) *liblayer {
	l.liblayerDo.DO = *(l.liblayerDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *liblayer) updateTableName(table string) *liblayer {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewString(table, "Id")
	l.CreationTime = field.NewTime(table, "CreationTime")
	l.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	l.LastModificationTime = field.NewTime(table, "LastModificationTime")
	l.LastModifierUserID = field.NewInt64(table, "LastModifierUserId")
	l.IsDeleted = field.NewField(table, "IsDeleted")
	l.DeleterUserID = field.NewInt64(table, "DeleterUserId")
	l.DeletionTime = field.NewTime(table, "DeletionTime")
	l.ShelfID = field.NewString(table, "ShelfId")
	l.Code = field.NewString(table, "Code")
	l.Name = field.NewString(table, "Name")
	l.Tid = field.NewString(table, "Tid")
	l.Side = field.NewString(table, "Side")
	l.LayerNo = field.NewInt64(table, "LayerNo")
	l.IsEnable = field.NewField(table, "IsEnable")
	l.Remark = field.NewString(table, "Remark")
	l.TenantID = field.NewInt64(table, "TenantId")
	l.ItemBarcode = field.NewString(table, "ItemBarcode")
	l.ItemCallNo = field.NewString(table, "ItemCallNo")
	l.PreCallNo = field.NewString(table, "PreCallNo")
	l.Barcode = field.NewString(table, "Barcode")
	l.OriginType = field.NewInt64(table, "OriginType")

	l.fillFieldMap()

	return l
}

func (l *liblayer) WithContext(ctx context.Context) ILiblayerDo { return l.liblayerDo.WithContext(ctx) }

func (l liblayer) TableName() string { return l.liblayerDo.TableName() }

func (l liblayer) Alias() string { return l.liblayerDo.Alias() }

func (l *liblayer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *liblayer) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 22)
	l.fieldMap["Id"] = l.ID
	l.fieldMap["CreationTime"] = l.CreationTime
	l.fieldMap["CreatorUserId"] = l.CreatorUserID
	l.fieldMap["LastModificationTime"] = l.LastModificationTime
	l.fieldMap["LastModifierUserId"] = l.LastModifierUserID
	l.fieldMap["IsDeleted"] = l.IsDeleted
	l.fieldMap["DeleterUserId"] = l.DeleterUserID
	l.fieldMap["DeletionTime"] = l.DeletionTime
	l.fieldMap["ShelfId"] = l.ShelfID
	l.fieldMap["Code"] = l.Code
	l.fieldMap["Name"] = l.Name
	l.fieldMap["Tid"] = l.Tid
	l.fieldMap["Side"] = l.Side
	l.fieldMap["LayerNo"] = l.LayerNo
	l.fieldMap["IsEnable"] = l.IsEnable
	l.fieldMap["Remark"] = l.Remark
	l.fieldMap["TenantId"] = l.TenantID
	l.fieldMap["ItemBarcode"] = l.ItemBarcode
	l.fieldMap["ItemCallNo"] = l.ItemCallNo
	l.fieldMap["PreCallNo"] = l.PreCallNo
	l.fieldMap["Barcode"] = l.Barcode
	l.fieldMap["OriginType"] = l.OriginType
}

func (l liblayer) clone(db *gorm.DB) liblayer {
	l.liblayerDo.ReplaceDB(db)
	return l
}

type liblayerDo struct{ gen.DO }

type ILiblayerDo interface {
	gen.SubQuery
	Debug() ILiblayerDo
	WithContext(ctx context.Context) ILiblayerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILiblayerDo
	Not(conds ...gen.Condition) ILiblayerDo
	Or(conds ...gen.Condition) ILiblayerDo
	Select(conds ...field.Expr) ILiblayerDo
	Where(conds ...gen.Condition) ILiblayerDo
	Order(conds ...field.Expr) ILiblayerDo
	Distinct(cols ...field.Expr) ILiblayerDo
	Omit(cols ...field.Expr) ILiblayerDo
	Join(table schema.Tabler, on ...field.Expr) ILiblayerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILiblayerDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILiblayerDo
	Group(cols ...field.Expr) ILiblayerDo
	Having(conds ...gen.Condition) ILiblayerDo
	Limit(limit int) ILiblayerDo
	Offset(offset int) ILiblayerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILiblayerDo
	Unscoped() ILiblayerDo
	Create(values ...*model.Liblayer) error
	CreateInBatches(values []*model.Liblayer, batchSize int) error
	Save(values ...*model.Liblayer) error
	First() (*model.Liblayer, error)
	Take() (*model.Liblayer, error)
	Last() (*model.Liblayer, error)
	Find() ([]*model.Liblayer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Liblayer, err error)
	FindInBatches(result *[]*model.Liblayer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Liblayer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILiblayerDo
	Assign(attrs ...field.AssignExpr) ILiblayerDo
	Joins(fields ...field.RelationField) ILiblayerDo
	Preload(fields ...field.RelationField) ILiblayerDo
	FirstOrInit() (*model.Liblayer, error)
	FirstOrCreate() (*model.Liblayer, error)
	FindByPage(offset int, limit int) (result []*model.Liblayer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILiblayerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l liblayerDo) Debug() ILiblayerDo {
	return l.withDO(l.DO.Debug())
}

func (l liblayerDo) WithContext(ctx context.Context) ILiblayerDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l liblayerDo) ReadDB() ILiblayerDo {
	return l.Clauses(dbresolver.Read)
}

func (l liblayerDo) WriteDB() ILiblayerDo {
	return l.Clauses(dbresolver.Write)
}

func (l liblayerDo) Clauses(conds ...clause.Expression) ILiblayerDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l liblayerDo) Returning(value interface{}, columns ...string) ILiblayerDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l liblayerDo) Not(conds ...gen.Condition) ILiblayerDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l liblayerDo) Or(conds ...gen.Condition) ILiblayerDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l liblayerDo) Select(conds ...field.Expr) ILiblayerDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l liblayerDo) Where(conds ...gen.Condition) ILiblayerDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l liblayerDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ILiblayerDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l liblayerDo) Order(conds ...field.Expr) ILiblayerDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l liblayerDo) Distinct(cols ...field.Expr) ILiblayerDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l liblayerDo) Omit(cols ...field.Expr) ILiblayerDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l liblayerDo) Join(table schema.Tabler, on ...field.Expr) ILiblayerDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l liblayerDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILiblayerDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l liblayerDo) RightJoin(table schema.Tabler, on ...field.Expr) ILiblayerDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l liblayerDo) Group(cols ...field.Expr) ILiblayerDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l liblayerDo) Having(conds ...gen.Condition) ILiblayerDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l liblayerDo) Limit(limit int) ILiblayerDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l liblayerDo) Offset(offset int) ILiblayerDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l liblayerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILiblayerDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l liblayerDo) Unscoped() ILiblayerDo {
	return l.withDO(l.DO.Unscoped())
}

func (l liblayerDo) Create(values ...*model.Liblayer) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l liblayerDo) CreateInBatches(values []*model.Liblayer, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l liblayerDo) Save(values ...*model.Liblayer) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l liblayerDo) First() (*model.Liblayer, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Liblayer), nil
	}
}

func (l liblayerDo) Take() (*model.Liblayer, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Liblayer), nil
	}
}

func (l liblayerDo) Last() (*model.Liblayer, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Liblayer), nil
	}
}

func (l liblayerDo) Find() ([]*model.Liblayer, error) {
	result, err := l.DO.Find()
	return result.([]*model.Liblayer), err
}

func (l liblayerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Liblayer, err error) {
	buf := make([]*model.Liblayer, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l liblayerDo) FindInBatches(result *[]*model.Liblayer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l liblayerDo) Attrs(attrs ...field.AssignExpr) ILiblayerDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l liblayerDo) Assign(attrs ...field.AssignExpr) ILiblayerDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l liblayerDo) Joins(fields ...field.RelationField) ILiblayerDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l liblayerDo) Preload(fields ...field.RelationField) ILiblayerDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l liblayerDo) FirstOrInit() (*model.Liblayer, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Liblayer), nil
	}
}

func (l liblayerDo) FirstOrCreate() (*model.Liblayer, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Liblayer), nil
	}
}

func (l liblayerDo) FindByPage(offset int, limit int) (result []*model.Liblayer, count int64, err error) {
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

func (l liblayerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l liblayerDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l liblayerDo) Delete(models ...*model.Liblayer) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *liblayerDo) withDO(do gen.Dao) *liblayerDo {
	l.DO = *do.(*gen.DO)
	return l
}
