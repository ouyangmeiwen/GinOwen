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

func newLiblayerindexupdatelog(db *gorm.DB) liblayerindexupdatelog {
	_liblayerindexupdatelog := liblayerindexupdatelog{}

	_liblayerindexupdatelog.liblayerindexupdatelogDo.UseDB(db)
	_liblayerindexupdatelog.liblayerindexupdatelogDo.UseModel(&model.Liblayerindexupdatelog{})

	tableName := _liblayerindexupdatelog.liblayerindexupdatelogDo.TableName()
	_liblayerindexupdatelog.ALL = field.NewAsterisk(tableName)
	_liblayerindexupdatelog.ID = field.NewString(tableName, "Id")
	_liblayerindexupdatelog.CreationTime = field.NewTime(tableName, "CreationTime")
	_liblayerindexupdatelog.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_liblayerindexupdatelog.MinLayerID = field.NewString(tableName, "MinLayerId")
	_liblayerindexupdatelog.MinLayerName = field.NewString(tableName, "MinLayerName")
	_liblayerindexupdatelog.MaxLayerID = field.NewString(tableName, "MaxLayerId")
	_liblayerindexupdatelog.MaxLayerName = field.NewString(tableName, "MaxLayerName")
	_liblayerindexupdatelog.UpdatedLayerID = field.NewString(tableName, "UpdatedLayerId")
	_liblayerindexupdatelog.UpdatedLayerName = field.NewString(tableName, "UpdatedLayerName")
	_liblayerindexupdatelog.UpdateStartTime = field.NewTime(tableName, "UpdateStartTime")
	_liblayerindexupdatelog.UpdateEndTime = field.NewTime(tableName, "UpdateEndTime")
	_liblayerindexupdatelog.Remark = field.NewString(tableName, "Remark")
	_liblayerindexupdatelog.TenantID = field.NewInt64(tableName, "TenantId")
	_liblayerindexupdatelog.MaxLayerCode = field.NewString(tableName, "MaxLayerCode")
	_liblayerindexupdatelog.MinLayerCode = field.NewString(tableName, "MinLayerCode")
	_liblayerindexupdatelog.UpdatedLayerCode = field.NewString(tableName, "UpdatedLayerCode")

	_liblayerindexupdatelog.fillFieldMap()

	return _liblayerindexupdatelog
}

type liblayerindexupdatelog struct {
	liblayerindexupdatelogDo liblayerindexupdatelogDo

	ALL              field.Asterisk
	ID               field.String
	CreationTime     field.Time
	CreatorUserID    field.Int64
	MinLayerID       field.String
	MinLayerName     field.String
	MaxLayerID       field.String
	MaxLayerName     field.String
	UpdatedLayerID   field.String
	UpdatedLayerName field.String
	UpdateStartTime  field.Time
	UpdateEndTime    field.Time
	Remark           field.String
	TenantID         field.Int64
	MaxLayerCode     field.String
	MinLayerCode     field.String
	UpdatedLayerCode field.String

	fieldMap map[string]field.Expr
}

func (l liblayerindexupdatelog) Table(newTableName string) *liblayerindexupdatelog {
	l.liblayerindexupdatelogDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l liblayerindexupdatelog) As(alias string) *liblayerindexupdatelog {
	l.liblayerindexupdatelogDo.DO = *(l.liblayerindexupdatelogDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *liblayerindexupdatelog) updateTableName(table string) *liblayerindexupdatelog {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewString(table, "Id")
	l.CreationTime = field.NewTime(table, "CreationTime")
	l.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	l.MinLayerID = field.NewString(table, "MinLayerId")
	l.MinLayerName = field.NewString(table, "MinLayerName")
	l.MaxLayerID = field.NewString(table, "MaxLayerId")
	l.MaxLayerName = field.NewString(table, "MaxLayerName")
	l.UpdatedLayerID = field.NewString(table, "UpdatedLayerId")
	l.UpdatedLayerName = field.NewString(table, "UpdatedLayerName")
	l.UpdateStartTime = field.NewTime(table, "UpdateStartTime")
	l.UpdateEndTime = field.NewTime(table, "UpdateEndTime")
	l.Remark = field.NewString(table, "Remark")
	l.TenantID = field.NewInt64(table, "TenantId")
	l.MaxLayerCode = field.NewString(table, "MaxLayerCode")
	l.MinLayerCode = field.NewString(table, "MinLayerCode")
	l.UpdatedLayerCode = field.NewString(table, "UpdatedLayerCode")

	l.fillFieldMap()

	return l
}

func (l *liblayerindexupdatelog) WithContext(ctx context.Context) ILiblayerindexupdatelogDo {
	return l.liblayerindexupdatelogDo.WithContext(ctx)
}

func (l liblayerindexupdatelog) TableName() string { return l.liblayerindexupdatelogDo.TableName() }

func (l liblayerindexupdatelog) Alias() string { return l.liblayerindexupdatelogDo.Alias() }

func (l *liblayerindexupdatelog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *liblayerindexupdatelog) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 16)
	l.fieldMap["Id"] = l.ID
	l.fieldMap["CreationTime"] = l.CreationTime
	l.fieldMap["CreatorUserId"] = l.CreatorUserID
	l.fieldMap["MinLayerId"] = l.MinLayerID
	l.fieldMap["MinLayerName"] = l.MinLayerName
	l.fieldMap["MaxLayerId"] = l.MaxLayerID
	l.fieldMap["MaxLayerName"] = l.MaxLayerName
	l.fieldMap["UpdatedLayerId"] = l.UpdatedLayerID
	l.fieldMap["UpdatedLayerName"] = l.UpdatedLayerName
	l.fieldMap["UpdateStartTime"] = l.UpdateStartTime
	l.fieldMap["UpdateEndTime"] = l.UpdateEndTime
	l.fieldMap["Remark"] = l.Remark
	l.fieldMap["TenantId"] = l.TenantID
	l.fieldMap["MaxLayerCode"] = l.MaxLayerCode
	l.fieldMap["MinLayerCode"] = l.MinLayerCode
	l.fieldMap["UpdatedLayerCode"] = l.UpdatedLayerCode
}

func (l liblayerindexupdatelog) clone(db *gorm.DB) liblayerindexupdatelog {
	l.liblayerindexupdatelogDo.ReplaceDB(db)
	return l
}

type liblayerindexupdatelogDo struct{ gen.DO }

type ILiblayerindexupdatelogDo interface {
	gen.SubQuery
	Debug() ILiblayerindexupdatelogDo
	WithContext(ctx context.Context) ILiblayerindexupdatelogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILiblayerindexupdatelogDo
	Not(conds ...gen.Condition) ILiblayerindexupdatelogDo
	Or(conds ...gen.Condition) ILiblayerindexupdatelogDo
	Select(conds ...field.Expr) ILiblayerindexupdatelogDo
	Where(conds ...gen.Condition) ILiblayerindexupdatelogDo
	Order(conds ...field.Expr) ILiblayerindexupdatelogDo
	Distinct(cols ...field.Expr) ILiblayerindexupdatelogDo
	Omit(cols ...field.Expr) ILiblayerindexupdatelogDo
	Join(table schema.Tabler, on ...field.Expr) ILiblayerindexupdatelogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILiblayerindexupdatelogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILiblayerindexupdatelogDo
	Group(cols ...field.Expr) ILiblayerindexupdatelogDo
	Having(conds ...gen.Condition) ILiblayerindexupdatelogDo
	Limit(limit int) ILiblayerindexupdatelogDo
	Offset(offset int) ILiblayerindexupdatelogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILiblayerindexupdatelogDo
	Unscoped() ILiblayerindexupdatelogDo
	Create(values ...*model.Liblayerindexupdatelog) error
	CreateInBatches(values []*model.Liblayerindexupdatelog, batchSize int) error
	Save(values ...*model.Liblayerindexupdatelog) error
	First() (*model.Liblayerindexupdatelog, error)
	Take() (*model.Liblayerindexupdatelog, error)
	Last() (*model.Liblayerindexupdatelog, error)
	Find() ([]*model.Liblayerindexupdatelog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Liblayerindexupdatelog, err error)
	FindInBatches(result *[]*model.Liblayerindexupdatelog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Liblayerindexupdatelog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILiblayerindexupdatelogDo
	Assign(attrs ...field.AssignExpr) ILiblayerindexupdatelogDo
	Joins(fields ...field.RelationField) ILiblayerindexupdatelogDo
	Preload(fields ...field.RelationField) ILiblayerindexupdatelogDo
	FirstOrInit() (*model.Liblayerindexupdatelog, error)
	FirstOrCreate() (*model.Liblayerindexupdatelog, error)
	FindByPage(offset int, limit int) (result []*model.Liblayerindexupdatelog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILiblayerindexupdatelogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l liblayerindexupdatelogDo) Debug() ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Debug())
}

func (l liblayerindexupdatelogDo) WithContext(ctx context.Context) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l liblayerindexupdatelogDo) ReadDB() ILiblayerindexupdatelogDo {
	return l.Clauses(dbresolver.Read)
}

func (l liblayerindexupdatelogDo) WriteDB() ILiblayerindexupdatelogDo {
	return l.Clauses(dbresolver.Write)
}

func (l liblayerindexupdatelogDo) Clauses(conds ...clause.Expression) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l liblayerindexupdatelogDo) Returning(value interface{}, columns ...string) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l liblayerindexupdatelogDo) Not(conds ...gen.Condition) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l liblayerindexupdatelogDo) Or(conds ...gen.Condition) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l liblayerindexupdatelogDo) Select(conds ...field.Expr) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l liblayerindexupdatelogDo) Where(conds ...gen.Condition) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l liblayerindexupdatelogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ILiblayerindexupdatelogDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l liblayerindexupdatelogDo) Order(conds ...field.Expr) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l liblayerindexupdatelogDo) Distinct(cols ...field.Expr) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l liblayerindexupdatelogDo) Omit(cols ...field.Expr) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l liblayerindexupdatelogDo) Join(table schema.Tabler, on ...field.Expr) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l liblayerindexupdatelogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l liblayerindexupdatelogDo) RightJoin(table schema.Tabler, on ...field.Expr) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l liblayerindexupdatelogDo) Group(cols ...field.Expr) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l liblayerindexupdatelogDo) Having(conds ...gen.Condition) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l liblayerindexupdatelogDo) Limit(limit int) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l liblayerindexupdatelogDo) Offset(offset int) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l liblayerindexupdatelogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l liblayerindexupdatelogDo) Unscoped() ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Unscoped())
}

func (l liblayerindexupdatelogDo) Create(values ...*model.Liblayerindexupdatelog) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l liblayerindexupdatelogDo) CreateInBatches(values []*model.Liblayerindexupdatelog, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l liblayerindexupdatelogDo) Save(values ...*model.Liblayerindexupdatelog) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l liblayerindexupdatelogDo) First() (*model.Liblayerindexupdatelog, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Liblayerindexupdatelog), nil
	}
}

func (l liblayerindexupdatelogDo) Take() (*model.Liblayerindexupdatelog, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Liblayerindexupdatelog), nil
	}
}

func (l liblayerindexupdatelogDo) Last() (*model.Liblayerindexupdatelog, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Liblayerindexupdatelog), nil
	}
}

func (l liblayerindexupdatelogDo) Find() ([]*model.Liblayerindexupdatelog, error) {
	result, err := l.DO.Find()
	return result.([]*model.Liblayerindexupdatelog), err
}

func (l liblayerindexupdatelogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Liblayerindexupdatelog, err error) {
	buf := make([]*model.Liblayerindexupdatelog, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l liblayerindexupdatelogDo) FindInBatches(result *[]*model.Liblayerindexupdatelog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l liblayerindexupdatelogDo) Attrs(attrs ...field.AssignExpr) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l liblayerindexupdatelogDo) Assign(attrs ...field.AssignExpr) ILiblayerindexupdatelogDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l liblayerindexupdatelogDo) Joins(fields ...field.RelationField) ILiblayerindexupdatelogDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l liblayerindexupdatelogDo) Preload(fields ...field.RelationField) ILiblayerindexupdatelogDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l liblayerindexupdatelogDo) FirstOrInit() (*model.Liblayerindexupdatelog, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Liblayerindexupdatelog), nil
	}
}

func (l liblayerindexupdatelogDo) FirstOrCreate() (*model.Liblayerindexupdatelog, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Liblayerindexupdatelog), nil
	}
}

func (l liblayerindexupdatelogDo) FindByPage(offset int, limit int) (result []*model.Liblayerindexupdatelog, count int64, err error) {
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

func (l liblayerindexupdatelogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l liblayerindexupdatelogDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l liblayerindexupdatelogDo) Delete(models ...*model.Liblayerindexupdatelog) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *liblayerindexupdatelogDo) withDO(do gen.Dao) *liblayerindexupdatelogDo {
	l.DO = *do.(*gen.DO)
	return l
}
