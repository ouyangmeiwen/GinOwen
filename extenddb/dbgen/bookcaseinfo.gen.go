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

func newBookcaseinfo(db *gorm.DB) bookcaseinfo {
	_bookcaseinfo := bookcaseinfo{}

	_bookcaseinfo.bookcaseinfoDo.UseDB(db)
	_bookcaseinfo.bookcaseinfoDo.UseModel(&model.Bookcaseinfo{})

	tableName := _bookcaseinfo.bookcaseinfoDo.TableName()
	_bookcaseinfo.ALL = field.NewAsterisk(tableName)
	_bookcaseinfo.NBookCaseInfoID = field.NewString(tableName, "nBookCaseInfoID")
	_bookcaseinfo.NEPCOrder = field.NewInt64(tableName, "nEPCOrder")
	_bookcaseinfo.SzBookCaseCode = field.NewString(tableName, "szBookCaseCode")
	_bookcaseinfo.NBookCaseLayers = field.NewInt64(tableName, "nBookCaseLayers")
	_bookcaseinfo.NBuildingNo = field.NewInt64(tableName, "nBuildingNo")
	_bookcaseinfo.NFloorNo = field.NewInt64(tableName, "nFloorNo")
	_bookcaseinfo.NRoomNo = field.NewInt64(tableName, "nRoomNo")
	_bookcaseinfo.SzMemo = field.NewString(tableName, "szMemo")
	_bookcaseinfo.NX1 = field.NewFloat64(tableName, "nX1")
	_bookcaseinfo.NY1 = field.NewFloat64(tableName, "nY1")
	_bookcaseinfo.NX2 = field.NewFloat64(tableName, "nX2")
	_bookcaseinfo.NY2 = field.NewFloat64(tableName, "nY2")
	_bookcaseinfo.NBookCount = field.NewInt64(tableName, "nBookCount")
	_bookcaseinfo.NID = field.NewInt64(tableName, "nID")
	_bookcaseinfo.NAngel = field.NewFloat64(tableName, "nAngel")
	_bookcaseinfo.SzCaseNoTrans = field.NewString(tableName, "szCaseNoTrans")
	_bookcaseinfo.BBosseyed = field.NewInt64(tableName, "bBosseyed")
	_bookcaseinfo.BNewShelves = field.NewInt64(tableName, "bNewShelves")

	_bookcaseinfo.fillFieldMap()

	return _bookcaseinfo
}

type bookcaseinfo struct {
	bookcaseinfoDo bookcaseinfoDo

	ALL             field.Asterisk
	NBookCaseInfoID field.String
	NEPCOrder       field.Int64
	SzBookCaseCode  field.String
	NBookCaseLayers field.Int64
	NBuildingNo     field.Int64
	NFloorNo        field.Int64
	NRoomNo         field.Int64
	SzMemo          field.String
	NX1             field.Float64
	NY1             field.Float64
	NX2             field.Float64
	NY2             field.Float64
	NBookCount      field.Int64
	NID             field.Int64
	NAngel          field.Float64
	SzCaseNoTrans   field.String
	BBosseyed       field.Int64
	BNewShelves     field.Int64

	fieldMap map[string]field.Expr
}

func (b bookcaseinfo) Table(newTableName string) *bookcaseinfo {
	b.bookcaseinfoDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bookcaseinfo) As(alias string) *bookcaseinfo {
	b.bookcaseinfoDo.DO = *(b.bookcaseinfoDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bookcaseinfo) updateTableName(table string) *bookcaseinfo {
	b.ALL = field.NewAsterisk(table)
	b.NBookCaseInfoID = field.NewString(table, "nBookCaseInfoID")
	b.NEPCOrder = field.NewInt64(table, "nEPCOrder")
	b.SzBookCaseCode = field.NewString(table, "szBookCaseCode")
	b.NBookCaseLayers = field.NewInt64(table, "nBookCaseLayers")
	b.NBuildingNo = field.NewInt64(table, "nBuildingNo")
	b.NFloorNo = field.NewInt64(table, "nFloorNo")
	b.NRoomNo = field.NewInt64(table, "nRoomNo")
	b.SzMemo = field.NewString(table, "szMemo")
	b.NX1 = field.NewFloat64(table, "nX1")
	b.NY1 = field.NewFloat64(table, "nY1")
	b.NX2 = field.NewFloat64(table, "nX2")
	b.NY2 = field.NewFloat64(table, "nY2")
	b.NBookCount = field.NewInt64(table, "nBookCount")
	b.NID = field.NewInt64(table, "nID")
	b.NAngel = field.NewFloat64(table, "nAngel")
	b.SzCaseNoTrans = field.NewString(table, "szCaseNoTrans")
	b.BBosseyed = field.NewInt64(table, "bBosseyed")
	b.BNewShelves = field.NewInt64(table, "bNewShelves")

	b.fillFieldMap()

	return b
}

func (b *bookcaseinfo) WithContext(ctx context.Context) IBookcaseinfoDo {
	return b.bookcaseinfoDo.WithContext(ctx)
}

func (b bookcaseinfo) TableName() string { return b.bookcaseinfoDo.TableName() }

func (b bookcaseinfo) Alias() string { return b.bookcaseinfoDo.Alias() }

func (b *bookcaseinfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bookcaseinfo) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 18)
	b.fieldMap["nBookCaseInfoID"] = b.NBookCaseInfoID
	b.fieldMap["nEPCOrder"] = b.NEPCOrder
	b.fieldMap["szBookCaseCode"] = b.SzBookCaseCode
	b.fieldMap["nBookCaseLayers"] = b.NBookCaseLayers
	b.fieldMap["nBuildingNo"] = b.NBuildingNo
	b.fieldMap["nFloorNo"] = b.NFloorNo
	b.fieldMap["nRoomNo"] = b.NRoomNo
	b.fieldMap["szMemo"] = b.SzMemo
	b.fieldMap["nX1"] = b.NX1
	b.fieldMap["nY1"] = b.NY1
	b.fieldMap["nX2"] = b.NX2
	b.fieldMap["nY2"] = b.NY2
	b.fieldMap["nBookCount"] = b.NBookCount
	b.fieldMap["nID"] = b.NID
	b.fieldMap["nAngel"] = b.NAngel
	b.fieldMap["szCaseNoTrans"] = b.SzCaseNoTrans
	b.fieldMap["bBosseyed"] = b.BBosseyed
	b.fieldMap["bNewShelves"] = b.BNewShelves
}

func (b bookcaseinfo) clone(db *gorm.DB) bookcaseinfo {
	b.bookcaseinfoDo.ReplaceDB(db)
	return b
}

type bookcaseinfoDo struct{ gen.DO }

type IBookcaseinfoDo interface {
	gen.SubQuery
	Debug() IBookcaseinfoDo
	WithContext(ctx context.Context) IBookcaseinfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBookcaseinfoDo
	Not(conds ...gen.Condition) IBookcaseinfoDo
	Or(conds ...gen.Condition) IBookcaseinfoDo
	Select(conds ...field.Expr) IBookcaseinfoDo
	Where(conds ...gen.Condition) IBookcaseinfoDo
	Order(conds ...field.Expr) IBookcaseinfoDo
	Distinct(cols ...field.Expr) IBookcaseinfoDo
	Omit(cols ...field.Expr) IBookcaseinfoDo
	Join(table schema.Tabler, on ...field.Expr) IBookcaseinfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBookcaseinfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBookcaseinfoDo
	Group(cols ...field.Expr) IBookcaseinfoDo
	Having(conds ...gen.Condition) IBookcaseinfoDo
	Limit(limit int) IBookcaseinfoDo
	Offset(offset int) IBookcaseinfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBookcaseinfoDo
	Unscoped() IBookcaseinfoDo
	Create(values ...*model.Bookcaseinfo) error
	CreateInBatches(values []*model.Bookcaseinfo, batchSize int) error
	Save(values ...*model.Bookcaseinfo) error
	First() (*model.Bookcaseinfo, error)
	Take() (*model.Bookcaseinfo, error)
	Last() (*model.Bookcaseinfo, error)
	Find() ([]*model.Bookcaseinfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Bookcaseinfo, err error)
	FindInBatches(result *[]*model.Bookcaseinfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Bookcaseinfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBookcaseinfoDo
	Assign(attrs ...field.AssignExpr) IBookcaseinfoDo
	Joins(fields ...field.RelationField) IBookcaseinfoDo
	Preload(fields ...field.RelationField) IBookcaseinfoDo
	FirstOrInit() (*model.Bookcaseinfo, error)
	FirstOrCreate() (*model.Bookcaseinfo, error)
	FindByPage(offset int, limit int) (result []*model.Bookcaseinfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBookcaseinfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b bookcaseinfoDo) Debug() IBookcaseinfoDo {
	return b.withDO(b.DO.Debug())
}

func (b bookcaseinfoDo) WithContext(ctx context.Context) IBookcaseinfoDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bookcaseinfoDo) ReadDB() IBookcaseinfoDo {
	return b.Clauses(dbresolver.Read)
}

func (b bookcaseinfoDo) WriteDB() IBookcaseinfoDo {
	return b.Clauses(dbresolver.Write)
}

func (b bookcaseinfoDo) Clauses(conds ...clause.Expression) IBookcaseinfoDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bookcaseinfoDo) Returning(value interface{}, columns ...string) IBookcaseinfoDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bookcaseinfoDo) Not(conds ...gen.Condition) IBookcaseinfoDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bookcaseinfoDo) Or(conds ...gen.Condition) IBookcaseinfoDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bookcaseinfoDo) Select(conds ...field.Expr) IBookcaseinfoDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bookcaseinfoDo) Where(conds ...gen.Condition) IBookcaseinfoDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bookcaseinfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IBookcaseinfoDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b bookcaseinfoDo) Order(conds ...field.Expr) IBookcaseinfoDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bookcaseinfoDo) Distinct(cols ...field.Expr) IBookcaseinfoDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bookcaseinfoDo) Omit(cols ...field.Expr) IBookcaseinfoDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bookcaseinfoDo) Join(table schema.Tabler, on ...field.Expr) IBookcaseinfoDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bookcaseinfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBookcaseinfoDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bookcaseinfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IBookcaseinfoDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bookcaseinfoDo) Group(cols ...field.Expr) IBookcaseinfoDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bookcaseinfoDo) Having(conds ...gen.Condition) IBookcaseinfoDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bookcaseinfoDo) Limit(limit int) IBookcaseinfoDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bookcaseinfoDo) Offset(offset int) IBookcaseinfoDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bookcaseinfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBookcaseinfoDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bookcaseinfoDo) Unscoped() IBookcaseinfoDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bookcaseinfoDo) Create(values ...*model.Bookcaseinfo) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bookcaseinfoDo) CreateInBatches(values []*model.Bookcaseinfo, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bookcaseinfoDo) Save(values ...*model.Bookcaseinfo) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bookcaseinfoDo) First() (*model.Bookcaseinfo, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bookcaseinfo), nil
	}
}

func (b bookcaseinfoDo) Take() (*model.Bookcaseinfo, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bookcaseinfo), nil
	}
}

func (b bookcaseinfoDo) Last() (*model.Bookcaseinfo, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bookcaseinfo), nil
	}
}

func (b bookcaseinfoDo) Find() ([]*model.Bookcaseinfo, error) {
	result, err := b.DO.Find()
	return result.([]*model.Bookcaseinfo), err
}

func (b bookcaseinfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Bookcaseinfo, err error) {
	buf := make([]*model.Bookcaseinfo, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bookcaseinfoDo) FindInBatches(result *[]*model.Bookcaseinfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bookcaseinfoDo) Attrs(attrs ...field.AssignExpr) IBookcaseinfoDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bookcaseinfoDo) Assign(attrs ...field.AssignExpr) IBookcaseinfoDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bookcaseinfoDo) Joins(fields ...field.RelationField) IBookcaseinfoDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bookcaseinfoDo) Preload(fields ...field.RelationField) IBookcaseinfoDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bookcaseinfoDo) FirstOrInit() (*model.Bookcaseinfo, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bookcaseinfo), nil
	}
}

func (b bookcaseinfoDo) FirstOrCreate() (*model.Bookcaseinfo, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Bookcaseinfo), nil
	}
}

func (b bookcaseinfoDo) FindByPage(offset int, limit int) (result []*model.Bookcaseinfo, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bookcaseinfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bookcaseinfoDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bookcaseinfoDo) Delete(models ...*model.Bookcaseinfo) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bookcaseinfoDo) withDO(do gen.Dao) *bookcaseinfoDo {
	b.DO = *do.(*gen.DO)
	return b
}
