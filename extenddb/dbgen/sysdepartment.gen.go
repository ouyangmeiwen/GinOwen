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

func newSysdepartment(db *gorm.DB) sysdepartment {
	_sysdepartment := sysdepartment{}

	_sysdepartment.sysdepartmentDo.UseDB(db)
	_sysdepartment.sysdepartmentDo.UseModel(&model.Sysdepartment{})

	tableName := _sysdepartment.sysdepartmentDo.TableName()
	_sysdepartment.ALL = field.NewAsterisk(tableName)
	_sysdepartment.ID = field.NewString(tableName, "Id")
	_sysdepartment.CreationTime = field.NewTime(tableName, "CreationTime")
	_sysdepartment.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_sysdepartment.LastModificationTime = field.NewTime(tableName, "LastModificationTime")
	_sysdepartment.LastModifierUserID = field.NewInt64(tableName, "LastModifierUserId")
	_sysdepartment.IsDeleted = field.NewField(tableName, "IsDeleted")
	_sysdepartment.DeleterUserID = field.NewInt64(tableName, "DeleterUserId")
	_sysdepartment.DeletionTime = field.NewTime(tableName, "DeletionTime")
	_sysdepartment.Code = field.NewString(tableName, "Code")
	_sysdepartment.Name = field.NewString(tableName, "Name")
	_sysdepartment.ParentID = field.NewString(tableName, "ParentId")
	_sysdepartment.Remark = field.NewString(tableName, "Remark")
	_sysdepartment.TenantID = field.NewInt64(tableName, "TenantId")

	_sysdepartment.fillFieldMap()

	return _sysdepartment
}

type sysdepartment struct {
	sysdepartmentDo sysdepartmentDo

	ALL                  field.Asterisk
	ID                   field.String
	CreationTime         field.Time
	CreatorUserID        field.Int64
	LastModificationTime field.Time
	LastModifierUserID   field.Int64
	IsDeleted            field.Field
	DeleterUserID        field.Int64
	DeletionTime         field.Time
	Code                 field.String
	Name                 field.String
	ParentID             field.String
	Remark               field.String
	TenantID             field.Int64

	fieldMap map[string]field.Expr
}

func (s sysdepartment) Table(newTableName string) *sysdepartment {
	s.sysdepartmentDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysdepartment) As(alias string) *sysdepartment {
	s.sysdepartmentDo.DO = *(s.sysdepartmentDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysdepartment) updateTableName(table string) *sysdepartment {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewString(table, "Id")
	s.CreationTime = field.NewTime(table, "CreationTime")
	s.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	s.LastModificationTime = field.NewTime(table, "LastModificationTime")
	s.LastModifierUserID = field.NewInt64(table, "LastModifierUserId")
	s.IsDeleted = field.NewField(table, "IsDeleted")
	s.DeleterUserID = field.NewInt64(table, "DeleterUserId")
	s.DeletionTime = field.NewTime(table, "DeletionTime")
	s.Code = field.NewString(table, "Code")
	s.Name = field.NewString(table, "Name")
	s.ParentID = field.NewString(table, "ParentId")
	s.Remark = field.NewString(table, "Remark")
	s.TenantID = field.NewInt64(table, "TenantId")

	s.fillFieldMap()

	return s
}

func (s *sysdepartment) WithContext(ctx context.Context) ISysdepartmentDo {
	return s.sysdepartmentDo.WithContext(ctx)
}

func (s sysdepartment) TableName() string { return s.sysdepartmentDo.TableName() }

func (s sysdepartment) Alias() string { return s.sysdepartmentDo.Alias() }

func (s *sysdepartment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysdepartment) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 13)
	s.fieldMap["Id"] = s.ID
	s.fieldMap["CreationTime"] = s.CreationTime
	s.fieldMap["CreatorUserId"] = s.CreatorUserID
	s.fieldMap["LastModificationTime"] = s.LastModificationTime
	s.fieldMap["LastModifierUserId"] = s.LastModifierUserID
	s.fieldMap["IsDeleted"] = s.IsDeleted
	s.fieldMap["DeleterUserId"] = s.DeleterUserID
	s.fieldMap["DeletionTime"] = s.DeletionTime
	s.fieldMap["Code"] = s.Code
	s.fieldMap["Name"] = s.Name
	s.fieldMap["ParentId"] = s.ParentID
	s.fieldMap["Remark"] = s.Remark
	s.fieldMap["TenantId"] = s.TenantID
}

func (s sysdepartment) clone(db *gorm.DB) sysdepartment {
	s.sysdepartmentDo.ReplaceDB(db)
	return s
}

type sysdepartmentDo struct{ gen.DO }

type ISysdepartmentDo interface {
	gen.SubQuery
	Debug() ISysdepartmentDo
	WithContext(ctx context.Context) ISysdepartmentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysdepartmentDo
	Not(conds ...gen.Condition) ISysdepartmentDo
	Or(conds ...gen.Condition) ISysdepartmentDo
	Select(conds ...field.Expr) ISysdepartmentDo
	Where(conds ...gen.Condition) ISysdepartmentDo
	Order(conds ...field.Expr) ISysdepartmentDo
	Distinct(cols ...field.Expr) ISysdepartmentDo
	Omit(cols ...field.Expr) ISysdepartmentDo
	Join(table schema.Tabler, on ...field.Expr) ISysdepartmentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysdepartmentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysdepartmentDo
	Group(cols ...field.Expr) ISysdepartmentDo
	Having(conds ...gen.Condition) ISysdepartmentDo
	Limit(limit int) ISysdepartmentDo
	Offset(offset int) ISysdepartmentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysdepartmentDo
	Unscoped() ISysdepartmentDo
	Create(values ...*model.Sysdepartment) error
	CreateInBatches(values []*model.Sysdepartment, batchSize int) error
	Save(values ...*model.Sysdepartment) error
	First() (*model.Sysdepartment, error)
	Take() (*model.Sysdepartment, error)
	Last() (*model.Sysdepartment, error)
	Find() ([]*model.Sysdepartment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Sysdepartment, err error)
	FindInBatches(result *[]*model.Sysdepartment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Sysdepartment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysdepartmentDo
	Assign(attrs ...field.AssignExpr) ISysdepartmentDo
	Joins(fields ...field.RelationField) ISysdepartmentDo
	Preload(fields ...field.RelationField) ISysdepartmentDo
	FirstOrInit() (*model.Sysdepartment, error)
	FirstOrCreate() (*model.Sysdepartment, error)
	FindByPage(offset int, limit int) (result []*model.Sysdepartment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysdepartmentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysdepartmentDo) Debug() ISysdepartmentDo {
	return s.withDO(s.DO.Debug())
}

func (s sysdepartmentDo) WithContext(ctx context.Context) ISysdepartmentDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysdepartmentDo) ReadDB() ISysdepartmentDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysdepartmentDo) WriteDB() ISysdepartmentDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysdepartmentDo) Clauses(conds ...clause.Expression) ISysdepartmentDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysdepartmentDo) Returning(value interface{}, columns ...string) ISysdepartmentDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysdepartmentDo) Not(conds ...gen.Condition) ISysdepartmentDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysdepartmentDo) Or(conds ...gen.Condition) ISysdepartmentDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysdepartmentDo) Select(conds ...field.Expr) ISysdepartmentDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysdepartmentDo) Where(conds ...gen.Condition) ISysdepartmentDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysdepartmentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISysdepartmentDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysdepartmentDo) Order(conds ...field.Expr) ISysdepartmentDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysdepartmentDo) Distinct(cols ...field.Expr) ISysdepartmentDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysdepartmentDo) Omit(cols ...field.Expr) ISysdepartmentDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysdepartmentDo) Join(table schema.Tabler, on ...field.Expr) ISysdepartmentDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysdepartmentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysdepartmentDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysdepartmentDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysdepartmentDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysdepartmentDo) Group(cols ...field.Expr) ISysdepartmentDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysdepartmentDo) Having(conds ...gen.Condition) ISysdepartmentDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysdepartmentDo) Limit(limit int) ISysdepartmentDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysdepartmentDo) Offset(offset int) ISysdepartmentDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysdepartmentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysdepartmentDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysdepartmentDo) Unscoped() ISysdepartmentDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysdepartmentDo) Create(values ...*model.Sysdepartment) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysdepartmentDo) CreateInBatches(values []*model.Sysdepartment, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysdepartmentDo) Save(values ...*model.Sysdepartment) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysdepartmentDo) First() (*model.Sysdepartment, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysdepartment), nil
	}
}

func (s sysdepartmentDo) Take() (*model.Sysdepartment, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysdepartment), nil
	}
}

func (s sysdepartmentDo) Last() (*model.Sysdepartment, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysdepartment), nil
	}
}

func (s sysdepartmentDo) Find() ([]*model.Sysdepartment, error) {
	result, err := s.DO.Find()
	return result.([]*model.Sysdepartment), err
}

func (s sysdepartmentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Sysdepartment, err error) {
	buf := make([]*model.Sysdepartment, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysdepartmentDo) FindInBatches(result *[]*model.Sysdepartment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysdepartmentDo) Attrs(attrs ...field.AssignExpr) ISysdepartmentDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysdepartmentDo) Assign(attrs ...field.AssignExpr) ISysdepartmentDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysdepartmentDo) Joins(fields ...field.RelationField) ISysdepartmentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysdepartmentDo) Preload(fields ...field.RelationField) ISysdepartmentDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysdepartmentDo) FirstOrInit() (*model.Sysdepartment, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysdepartment), nil
	}
}

func (s sysdepartmentDo) FirstOrCreate() (*model.Sysdepartment, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Sysdepartment), nil
	}
}

func (s sysdepartmentDo) FindByPage(offset int, limit int) (result []*model.Sysdepartment, count int64, err error) {
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

func (s sysdepartmentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysdepartmentDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysdepartmentDo) Delete(models ...*model.Sysdepartment) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysdepartmentDo) withDO(do gen.Dao) *sysdepartmentDo {
	s.DO = *do.(*gen.DO)
	return s
}
