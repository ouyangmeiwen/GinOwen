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

func newLibailibrarainsessionmetric(db *gorm.DB) libailibrarainsessionmetric {
	_libailibrarainsessionmetric := libailibrarainsessionmetric{}

	_libailibrarainsessionmetric.libailibrarainsessionmetricDo.UseDB(db)
	_libailibrarainsessionmetric.libailibrarainsessionmetricDo.UseModel(&model.Libailibrarainsessionmetric{})

	tableName := _libailibrarainsessionmetric.libailibrarainsessionmetricDo.TableName()
	_libailibrarainsessionmetric.ALL = field.NewAsterisk(tableName)
	_libailibrarainsessionmetric.ID = field.NewInt64(tableName, "Id")
	_libailibrarainsessionmetric.Date = field.NewTime(tableName, "Date")
	_libailibrarainsessionmetric.SecondsOfService = field.NewInt64(tableName, "SecondsOfService")
	_libailibrarainsessionmetric.SecondsOfServiceTotal = field.NewInt64(tableName, "SecondsOfServiceTotal")
	_libailibrarainsessionmetric.BreakCount = field.NewInt64(tableName, "BreakCount")
	_libailibrarainsessionmetric.SessionCount = field.NewInt64(tableName, "SessionCount")
	_libailibrarainsessionmetric.QuestionsCount = field.NewInt64(tableName, "QuestionsCount")
	_libailibrarainsessionmetric.TenantID = field.NewInt64(tableName, "TenantId")
	_libailibrarainsessionmetric.DeviceID = field.NewString(tableName, "DeviceId")

	_libailibrarainsessionmetric.fillFieldMap()

	return _libailibrarainsessionmetric
}

type libailibrarainsessionmetric struct {
	libailibrarainsessionmetricDo libailibrarainsessionmetricDo

	ALL                   field.Asterisk
	ID                    field.Int64
	Date                  field.Time
	SecondsOfService      field.Int64
	SecondsOfServiceTotal field.Int64
	BreakCount            field.Int64
	SessionCount          field.Int64
	QuestionsCount        field.Int64
	TenantID              field.Int64
	DeviceID              field.String

	fieldMap map[string]field.Expr
}

func (l libailibrarainsessionmetric) Table(newTableName string) *libailibrarainsessionmetric {
	l.libailibrarainsessionmetricDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l libailibrarainsessionmetric) As(alias string) *libailibrarainsessionmetric {
	l.libailibrarainsessionmetricDo.DO = *(l.libailibrarainsessionmetricDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *libailibrarainsessionmetric) updateTableName(table string) *libailibrarainsessionmetric {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt64(table, "Id")
	l.Date = field.NewTime(table, "Date")
	l.SecondsOfService = field.NewInt64(table, "SecondsOfService")
	l.SecondsOfServiceTotal = field.NewInt64(table, "SecondsOfServiceTotal")
	l.BreakCount = field.NewInt64(table, "BreakCount")
	l.SessionCount = field.NewInt64(table, "SessionCount")
	l.QuestionsCount = field.NewInt64(table, "QuestionsCount")
	l.TenantID = field.NewInt64(table, "TenantId")
	l.DeviceID = field.NewString(table, "DeviceId")

	l.fillFieldMap()

	return l
}

func (l *libailibrarainsessionmetric) WithContext(ctx context.Context) ILibailibrarainsessionmetricDo {
	return l.libailibrarainsessionmetricDo.WithContext(ctx)
}

func (l libailibrarainsessionmetric) TableName() string {
	return l.libailibrarainsessionmetricDo.TableName()
}

func (l libailibrarainsessionmetric) Alias() string { return l.libailibrarainsessionmetricDo.Alias() }

func (l *libailibrarainsessionmetric) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *libailibrarainsessionmetric) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 9)
	l.fieldMap["Id"] = l.ID
	l.fieldMap["Date"] = l.Date
	l.fieldMap["SecondsOfService"] = l.SecondsOfService
	l.fieldMap["SecondsOfServiceTotal"] = l.SecondsOfServiceTotal
	l.fieldMap["BreakCount"] = l.BreakCount
	l.fieldMap["SessionCount"] = l.SessionCount
	l.fieldMap["QuestionsCount"] = l.QuestionsCount
	l.fieldMap["TenantId"] = l.TenantID
	l.fieldMap["DeviceId"] = l.DeviceID
}

func (l libailibrarainsessionmetric) clone(db *gorm.DB) libailibrarainsessionmetric {
	l.libailibrarainsessionmetricDo.ReplaceDB(db)
	return l
}

type libailibrarainsessionmetricDo struct{ gen.DO }

type ILibailibrarainsessionmetricDo interface {
	gen.SubQuery
	Debug() ILibailibrarainsessionmetricDo
	WithContext(ctx context.Context) ILibailibrarainsessionmetricDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILibailibrarainsessionmetricDo
	Not(conds ...gen.Condition) ILibailibrarainsessionmetricDo
	Or(conds ...gen.Condition) ILibailibrarainsessionmetricDo
	Select(conds ...field.Expr) ILibailibrarainsessionmetricDo
	Where(conds ...gen.Condition) ILibailibrarainsessionmetricDo
	Order(conds ...field.Expr) ILibailibrarainsessionmetricDo
	Distinct(cols ...field.Expr) ILibailibrarainsessionmetricDo
	Omit(cols ...field.Expr) ILibailibrarainsessionmetricDo
	Join(table schema.Tabler, on ...field.Expr) ILibailibrarainsessionmetricDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILibailibrarainsessionmetricDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILibailibrarainsessionmetricDo
	Group(cols ...field.Expr) ILibailibrarainsessionmetricDo
	Having(conds ...gen.Condition) ILibailibrarainsessionmetricDo
	Limit(limit int) ILibailibrarainsessionmetricDo
	Offset(offset int) ILibailibrarainsessionmetricDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILibailibrarainsessionmetricDo
	Unscoped() ILibailibrarainsessionmetricDo
	Create(values ...*model.Libailibrarainsessionmetric) error
	CreateInBatches(values []*model.Libailibrarainsessionmetric, batchSize int) error
	Save(values ...*model.Libailibrarainsessionmetric) error
	First() (*model.Libailibrarainsessionmetric, error)
	Take() (*model.Libailibrarainsessionmetric, error)
	Last() (*model.Libailibrarainsessionmetric, error)
	Find() ([]*model.Libailibrarainsessionmetric, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Libailibrarainsessionmetric, err error)
	FindInBatches(result *[]*model.Libailibrarainsessionmetric, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Libailibrarainsessionmetric) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILibailibrarainsessionmetricDo
	Assign(attrs ...field.AssignExpr) ILibailibrarainsessionmetricDo
	Joins(fields ...field.RelationField) ILibailibrarainsessionmetricDo
	Preload(fields ...field.RelationField) ILibailibrarainsessionmetricDo
	FirstOrInit() (*model.Libailibrarainsessionmetric, error)
	FirstOrCreate() (*model.Libailibrarainsessionmetric, error)
	FindByPage(offset int, limit int) (result []*model.Libailibrarainsessionmetric, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILibailibrarainsessionmetricDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l libailibrarainsessionmetricDo) Debug() ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Debug())
}

func (l libailibrarainsessionmetricDo) WithContext(ctx context.Context) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l libailibrarainsessionmetricDo) ReadDB() ILibailibrarainsessionmetricDo {
	return l.Clauses(dbresolver.Read)
}

func (l libailibrarainsessionmetricDo) WriteDB() ILibailibrarainsessionmetricDo {
	return l.Clauses(dbresolver.Write)
}

func (l libailibrarainsessionmetricDo) Clauses(conds ...clause.Expression) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l libailibrarainsessionmetricDo) Returning(value interface{}, columns ...string) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l libailibrarainsessionmetricDo) Not(conds ...gen.Condition) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l libailibrarainsessionmetricDo) Or(conds ...gen.Condition) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l libailibrarainsessionmetricDo) Select(conds ...field.Expr) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l libailibrarainsessionmetricDo) Where(conds ...gen.Condition) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l libailibrarainsessionmetricDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ILibailibrarainsessionmetricDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l libailibrarainsessionmetricDo) Order(conds ...field.Expr) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l libailibrarainsessionmetricDo) Distinct(cols ...field.Expr) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l libailibrarainsessionmetricDo) Omit(cols ...field.Expr) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l libailibrarainsessionmetricDo) Join(table schema.Tabler, on ...field.Expr) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l libailibrarainsessionmetricDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l libailibrarainsessionmetricDo) RightJoin(table schema.Tabler, on ...field.Expr) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l libailibrarainsessionmetricDo) Group(cols ...field.Expr) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l libailibrarainsessionmetricDo) Having(conds ...gen.Condition) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l libailibrarainsessionmetricDo) Limit(limit int) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l libailibrarainsessionmetricDo) Offset(offset int) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l libailibrarainsessionmetricDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l libailibrarainsessionmetricDo) Unscoped() ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Unscoped())
}

func (l libailibrarainsessionmetricDo) Create(values ...*model.Libailibrarainsessionmetric) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l libailibrarainsessionmetricDo) CreateInBatches(values []*model.Libailibrarainsessionmetric, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l libailibrarainsessionmetricDo) Save(values ...*model.Libailibrarainsessionmetric) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l libailibrarainsessionmetricDo) First() (*model.Libailibrarainsessionmetric, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libailibrarainsessionmetric), nil
	}
}

func (l libailibrarainsessionmetricDo) Take() (*model.Libailibrarainsessionmetric, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libailibrarainsessionmetric), nil
	}
}

func (l libailibrarainsessionmetricDo) Last() (*model.Libailibrarainsessionmetric, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libailibrarainsessionmetric), nil
	}
}

func (l libailibrarainsessionmetricDo) Find() ([]*model.Libailibrarainsessionmetric, error) {
	result, err := l.DO.Find()
	return result.([]*model.Libailibrarainsessionmetric), err
}

func (l libailibrarainsessionmetricDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Libailibrarainsessionmetric, err error) {
	buf := make([]*model.Libailibrarainsessionmetric, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l libailibrarainsessionmetricDo) FindInBatches(result *[]*model.Libailibrarainsessionmetric, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l libailibrarainsessionmetricDo) Attrs(attrs ...field.AssignExpr) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l libailibrarainsessionmetricDo) Assign(attrs ...field.AssignExpr) ILibailibrarainsessionmetricDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l libailibrarainsessionmetricDo) Joins(fields ...field.RelationField) ILibailibrarainsessionmetricDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l libailibrarainsessionmetricDo) Preload(fields ...field.RelationField) ILibailibrarainsessionmetricDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l libailibrarainsessionmetricDo) FirstOrInit() (*model.Libailibrarainsessionmetric, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libailibrarainsessionmetric), nil
	}
}

func (l libailibrarainsessionmetricDo) FirstOrCreate() (*model.Libailibrarainsessionmetric, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Libailibrarainsessionmetric), nil
	}
}

func (l libailibrarainsessionmetricDo) FindByPage(offset int, limit int) (result []*model.Libailibrarainsessionmetric, count int64, err error) {
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

func (l libailibrarainsessionmetricDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l libailibrarainsessionmetricDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l libailibrarainsessionmetricDo) Delete(models ...*model.Libailibrarainsessionmetric) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *libailibrarainsessionmetricDo) withDO(do gen.Dao) *libailibrarainsessionmetricDo {
	l.DO = *do.(*gen.DO)
	return l
}
