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

func newSyslayertran(db *gorm.DB) syslayertran {
	_syslayertran := syslayertran{}

	_syslayertran.syslayertranDo.UseDB(db)
	_syslayertran.syslayertranDo.UseModel(&model.Syslayertran{})

	tableName := _syslayertran.syslayertranDo.TableName()
	_syslayertran.ALL = field.NewAsterisk(tableName)
	_syslayertran.ID = field.NewString(tableName, "Id")
	_syslayertran.Name = field.NewString(tableName, "Name")
	_syslayertran.IsShowBuilding = field.NewField(tableName, "IsShowBuilding")
	_syslayertran.IsShowFloor = field.NewField(tableName, "IsShowFloor")
	_syslayertran.IsShowRoom = field.NewField(tableName, "IsShowRoom")
	_syslayertran.Row = field.NewString(tableName, "Row")
	_syslayertran.Column = field.NewString(tableName, "Column")
	_syslayertran.Face = field.NewString(tableName, "Face")
	_syslayertran.Layer = field.NewString(tableName, "Layer")
	_syslayertran.ShowMode = field.NewString(tableName, "ShowMode")
	_syslayertran.LayerSort = field.NewString(tableName, "LayerSort")
	_syslayertran.ColumnSort = field.NewString(tableName, "ColumnSort")
	_syslayertran.Description = field.NewString(tableName, "Description")
	_syslayertran.IsSelected = field.NewField(tableName, "IsSelected")
	_syslayertran.TenantID = field.NewInt64(tableName, "TenantId")

	_syslayertran.fillFieldMap()

	return _syslayertran
}

type syslayertran struct {
	syslayertranDo syslayertranDo

	ALL            field.Asterisk
	ID             field.String
	Name           field.String
	IsShowBuilding field.Field
	IsShowFloor    field.Field
	IsShowRoom     field.Field
	Row            field.String
	Column         field.String
	Face           field.String
	Layer          field.String
	ShowMode       field.String
	LayerSort      field.String
	ColumnSort     field.String
	Description    field.String
	IsSelected     field.Field
	TenantID       field.Int64

	fieldMap map[string]field.Expr
}

func (s syslayertran) Table(newTableName string) *syslayertran {
	s.syslayertranDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s syslayertran) As(alias string) *syslayertran {
	s.syslayertranDo.DO = *(s.syslayertranDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *syslayertran) updateTableName(table string) *syslayertran {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewString(table, "Id")
	s.Name = field.NewString(table, "Name")
	s.IsShowBuilding = field.NewField(table, "IsShowBuilding")
	s.IsShowFloor = field.NewField(table, "IsShowFloor")
	s.IsShowRoom = field.NewField(table, "IsShowRoom")
	s.Row = field.NewString(table, "Row")
	s.Column = field.NewString(table, "Column")
	s.Face = field.NewString(table, "Face")
	s.Layer = field.NewString(table, "Layer")
	s.ShowMode = field.NewString(table, "ShowMode")
	s.LayerSort = field.NewString(table, "LayerSort")
	s.ColumnSort = field.NewString(table, "ColumnSort")
	s.Description = field.NewString(table, "Description")
	s.IsSelected = field.NewField(table, "IsSelected")
	s.TenantID = field.NewInt64(table, "TenantId")

	s.fillFieldMap()

	return s
}

func (s *syslayertran) WithContext(ctx context.Context) ISyslayertranDo {
	return s.syslayertranDo.WithContext(ctx)
}

func (s syslayertran) TableName() string { return s.syslayertranDo.TableName() }

func (s syslayertran) Alias() string { return s.syslayertranDo.Alias() }

func (s *syslayertran) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *syslayertran) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 15)
	s.fieldMap["Id"] = s.ID
	s.fieldMap["Name"] = s.Name
	s.fieldMap["IsShowBuilding"] = s.IsShowBuilding
	s.fieldMap["IsShowFloor"] = s.IsShowFloor
	s.fieldMap["IsShowRoom"] = s.IsShowRoom
	s.fieldMap["Row"] = s.Row
	s.fieldMap["Column"] = s.Column
	s.fieldMap["Face"] = s.Face
	s.fieldMap["Layer"] = s.Layer
	s.fieldMap["ShowMode"] = s.ShowMode
	s.fieldMap["LayerSort"] = s.LayerSort
	s.fieldMap["ColumnSort"] = s.ColumnSort
	s.fieldMap["Description"] = s.Description
	s.fieldMap["IsSelected"] = s.IsSelected
	s.fieldMap["TenantId"] = s.TenantID
}

func (s syslayertran) clone(db *gorm.DB) syslayertran {
	s.syslayertranDo.ReplaceDB(db)
	return s
}

type syslayertranDo struct{ gen.DO }

type ISyslayertranDo interface {
	gen.SubQuery
	Debug() ISyslayertranDo
	WithContext(ctx context.Context) ISyslayertranDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISyslayertranDo
	Not(conds ...gen.Condition) ISyslayertranDo
	Or(conds ...gen.Condition) ISyslayertranDo
	Select(conds ...field.Expr) ISyslayertranDo
	Where(conds ...gen.Condition) ISyslayertranDo
	Order(conds ...field.Expr) ISyslayertranDo
	Distinct(cols ...field.Expr) ISyslayertranDo
	Omit(cols ...field.Expr) ISyslayertranDo
	Join(table schema.Tabler, on ...field.Expr) ISyslayertranDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISyslayertranDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISyslayertranDo
	Group(cols ...field.Expr) ISyslayertranDo
	Having(conds ...gen.Condition) ISyslayertranDo
	Limit(limit int) ISyslayertranDo
	Offset(offset int) ISyslayertranDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISyslayertranDo
	Unscoped() ISyslayertranDo
	Create(values ...*model.Syslayertran) error
	CreateInBatches(values []*model.Syslayertran, batchSize int) error
	Save(values ...*model.Syslayertran) error
	First() (*model.Syslayertran, error)
	Take() (*model.Syslayertran, error)
	Last() (*model.Syslayertran, error)
	Find() ([]*model.Syslayertran, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Syslayertran, err error)
	FindInBatches(result *[]*model.Syslayertran, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Syslayertran) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISyslayertranDo
	Assign(attrs ...field.AssignExpr) ISyslayertranDo
	Joins(fields ...field.RelationField) ISyslayertranDo
	Preload(fields ...field.RelationField) ISyslayertranDo
	FirstOrInit() (*model.Syslayertran, error)
	FirstOrCreate() (*model.Syslayertran, error)
	FindByPage(offset int, limit int) (result []*model.Syslayertran, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISyslayertranDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s syslayertranDo) Debug() ISyslayertranDo {
	return s.withDO(s.DO.Debug())
}

func (s syslayertranDo) WithContext(ctx context.Context) ISyslayertranDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s syslayertranDo) ReadDB() ISyslayertranDo {
	return s.Clauses(dbresolver.Read)
}

func (s syslayertranDo) WriteDB() ISyslayertranDo {
	return s.Clauses(dbresolver.Write)
}

func (s syslayertranDo) Clauses(conds ...clause.Expression) ISyslayertranDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s syslayertranDo) Returning(value interface{}, columns ...string) ISyslayertranDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s syslayertranDo) Not(conds ...gen.Condition) ISyslayertranDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s syslayertranDo) Or(conds ...gen.Condition) ISyslayertranDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s syslayertranDo) Select(conds ...field.Expr) ISyslayertranDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s syslayertranDo) Where(conds ...gen.Condition) ISyslayertranDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s syslayertranDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISyslayertranDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s syslayertranDo) Order(conds ...field.Expr) ISyslayertranDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s syslayertranDo) Distinct(cols ...field.Expr) ISyslayertranDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s syslayertranDo) Omit(cols ...field.Expr) ISyslayertranDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s syslayertranDo) Join(table schema.Tabler, on ...field.Expr) ISyslayertranDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s syslayertranDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISyslayertranDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s syslayertranDo) RightJoin(table schema.Tabler, on ...field.Expr) ISyslayertranDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s syslayertranDo) Group(cols ...field.Expr) ISyslayertranDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s syslayertranDo) Having(conds ...gen.Condition) ISyslayertranDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s syslayertranDo) Limit(limit int) ISyslayertranDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s syslayertranDo) Offset(offset int) ISyslayertranDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s syslayertranDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISyslayertranDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s syslayertranDo) Unscoped() ISyslayertranDo {
	return s.withDO(s.DO.Unscoped())
}

func (s syslayertranDo) Create(values ...*model.Syslayertran) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s syslayertranDo) CreateInBatches(values []*model.Syslayertran, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s syslayertranDo) Save(values ...*model.Syslayertran) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s syslayertranDo) First() (*model.Syslayertran, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Syslayertran), nil
	}
}

func (s syslayertranDo) Take() (*model.Syslayertran, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Syslayertran), nil
	}
}

func (s syslayertranDo) Last() (*model.Syslayertran, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Syslayertran), nil
	}
}

func (s syslayertranDo) Find() ([]*model.Syslayertran, error) {
	result, err := s.DO.Find()
	return result.([]*model.Syslayertran), err
}

func (s syslayertranDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Syslayertran, err error) {
	buf := make([]*model.Syslayertran, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s syslayertranDo) FindInBatches(result *[]*model.Syslayertran, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s syslayertranDo) Attrs(attrs ...field.AssignExpr) ISyslayertranDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s syslayertranDo) Assign(attrs ...field.AssignExpr) ISyslayertranDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s syslayertranDo) Joins(fields ...field.RelationField) ISyslayertranDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s syslayertranDo) Preload(fields ...field.RelationField) ISyslayertranDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s syslayertranDo) FirstOrInit() (*model.Syslayertran, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Syslayertran), nil
	}
}

func (s syslayertranDo) FirstOrCreate() (*model.Syslayertran, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Syslayertran), nil
	}
}

func (s syslayertranDo) FindByPage(offset int, limit int) (result []*model.Syslayertran, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s syslayertranDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s syslayertranDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s syslayertranDo) Delete(models ...*model.Syslayertran) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *syslayertranDo) withDO(do gen.Dao) *syslayertranDo {
	s.DO = *do.(*gen.DO)
	return s
}
