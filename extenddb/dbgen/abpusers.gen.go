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

func newAbpuser(db *gorm.DB) abpuser {
	_abpuser := abpuser{}

	_abpuser.abpuserDo.UseDB(db)
	_abpuser.abpuserDo.UseModel(&model.Abpuser{})

	tableName := _abpuser.abpuserDo.TableName()
	_abpuser.ALL = field.NewAsterisk(tableName)
	_abpuser.ID = field.NewInt64(tableName, "Id")
	_abpuser.CreationTime = field.NewTime(tableName, "CreationTime")
	_abpuser.CreatorUserID = field.NewInt64(tableName, "CreatorUserId")
	_abpuser.LastModificationTime = field.NewTime(tableName, "LastModificationTime")
	_abpuser.LastModifierUserID = field.NewInt64(tableName, "LastModifierUserId")
	_abpuser.IsDeleted = field.NewField(tableName, "IsDeleted")
	_abpuser.DeleterUserID = field.NewInt64(tableName, "DeleterUserId")
	_abpuser.DeletionTime = field.NewTime(tableName, "DeletionTime")
	_abpuser.AuthenticationSource = field.NewString(tableName, "AuthenticationSource")
	_abpuser.UserName = field.NewString(tableName, "UserName")
	_abpuser.TenantID = field.NewInt64(tableName, "TenantId")
	_abpuser.EmailAddress = field.NewString(tableName, "EmailAddress")
	_abpuser.Name = field.NewString(tableName, "Name")
	_abpuser.Surname = field.NewString(tableName, "Surname")
	_abpuser.Password = field.NewString(tableName, "Password")
	_abpuser.EmailConfirmationCode = field.NewString(tableName, "EmailConfirmationCode")
	_abpuser.PasswordResetCode = field.NewString(tableName, "PasswordResetCode")
	_abpuser.LockoutEndDateUtc = field.NewTime(tableName, "LockoutEndDateUtc")
	_abpuser.AccessFailedCount = field.NewInt64(tableName, "AccessFailedCount")
	_abpuser.IsLockoutEnabled = field.NewField(tableName, "IsLockoutEnabled")
	_abpuser.PhoneNumber = field.NewString(tableName, "PhoneNumber")
	_abpuser.IsPhoneNumberConfirmed = field.NewField(tableName, "IsPhoneNumberConfirmed")
	_abpuser.SecurityStamp = field.NewString(tableName, "SecurityStamp")
	_abpuser.IsTwoFactorEnabled = field.NewField(tableName, "IsTwoFactorEnabled")
	_abpuser.IsEmailConfirmed = field.NewField(tableName, "IsEmailConfirmed")
	_abpuser.IsActive = field.NewField(tableName, "IsActive")
	_abpuser.NormalizedUserName = field.NewString(tableName, "NormalizedUserName")
	_abpuser.NormalizedEmailAddress = field.NewString(tableName, "NormalizedEmailAddress")
	_abpuser.ConcurrencyStamp = field.NewString(tableName, "ConcurrencyStamp")
	_abpuser.ProfilePictureID = field.NewString(tableName, "ProfilePictureId")
	_abpuser.ShouldChangePasswordOnNextLogin = field.NewField(tableName, "ShouldChangePasswordOnNextLogin")
	_abpuser.SignInTokenExpireTimeUtc = field.NewTime(tableName, "SignInTokenExpireTimeUtc")
	_abpuser.SignInToken = field.NewString(tableName, "SignInToken")
	_abpuser.GoogleAuthenticatorKey = field.NewString(tableName, "GoogleAuthenticatorKey")

	_abpuser.fillFieldMap()

	return _abpuser
}

type abpuser struct {
	abpuserDo abpuserDo

	ALL                             field.Asterisk
	ID                              field.Int64
	CreationTime                    field.Time
	CreatorUserID                   field.Int64
	LastModificationTime            field.Time
	LastModifierUserID              field.Int64
	IsDeleted                       field.Field
	DeleterUserID                   field.Int64
	DeletionTime                    field.Time
	AuthenticationSource            field.String
	UserName                        field.String
	TenantID                        field.Int64
	EmailAddress                    field.String
	Name                            field.String
	Surname                         field.String
	Password                        field.String
	EmailConfirmationCode           field.String
	PasswordResetCode               field.String
	LockoutEndDateUtc               field.Time
	AccessFailedCount               field.Int64
	IsLockoutEnabled                field.Field
	PhoneNumber                     field.String
	IsPhoneNumberConfirmed          field.Field
	SecurityStamp                   field.String
	IsTwoFactorEnabled              field.Field
	IsEmailConfirmed                field.Field
	IsActive                        field.Field
	NormalizedUserName              field.String
	NormalizedEmailAddress          field.String
	ConcurrencyStamp                field.String
	ProfilePictureID                field.String
	ShouldChangePasswordOnNextLogin field.Field
	SignInTokenExpireTimeUtc        field.Time
	SignInToken                     field.String
	GoogleAuthenticatorKey          field.String

	fieldMap map[string]field.Expr
}

func (a abpuser) Table(newTableName string) *abpuser {
	a.abpuserDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a abpuser) As(alias string) *abpuser {
	a.abpuserDo.DO = *(a.abpuserDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *abpuser) updateTableName(table string) *abpuser {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "Id")
	a.CreationTime = field.NewTime(table, "CreationTime")
	a.CreatorUserID = field.NewInt64(table, "CreatorUserId")
	a.LastModificationTime = field.NewTime(table, "LastModificationTime")
	a.LastModifierUserID = field.NewInt64(table, "LastModifierUserId")
	a.IsDeleted = field.NewField(table, "IsDeleted")
	a.DeleterUserID = field.NewInt64(table, "DeleterUserId")
	a.DeletionTime = field.NewTime(table, "DeletionTime")
	a.AuthenticationSource = field.NewString(table, "AuthenticationSource")
	a.UserName = field.NewString(table, "UserName")
	a.TenantID = field.NewInt64(table, "TenantId")
	a.EmailAddress = field.NewString(table, "EmailAddress")
	a.Name = field.NewString(table, "Name")
	a.Surname = field.NewString(table, "Surname")
	a.Password = field.NewString(table, "Password")
	a.EmailConfirmationCode = field.NewString(table, "EmailConfirmationCode")
	a.PasswordResetCode = field.NewString(table, "PasswordResetCode")
	a.LockoutEndDateUtc = field.NewTime(table, "LockoutEndDateUtc")
	a.AccessFailedCount = field.NewInt64(table, "AccessFailedCount")
	a.IsLockoutEnabled = field.NewField(table, "IsLockoutEnabled")
	a.PhoneNumber = field.NewString(table, "PhoneNumber")
	a.IsPhoneNumberConfirmed = field.NewField(table, "IsPhoneNumberConfirmed")
	a.SecurityStamp = field.NewString(table, "SecurityStamp")
	a.IsTwoFactorEnabled = field.NewField(table, "IsTwoFactorEnabled")
	a.IsEmailConfirmed = field.NewField(table, "IsEmailConfirmed")
	a.IsActive = field.NewField(table, "IsActive")
	a.NormalizedUserName = field.NewString(table, "NormalizedUserName")
	a.NormalizedEmailAddress = field.NewString(table, "NormalizedEmailAddress")
	a.ConcurrencyStamp = field.NewString(table, "ConcurrencyStamp")
	a.ProfilePictureID = field.NewString(table, "ProfilePictureId")
	a.ShouldChangePasswordOnNextLogin = field.NewField(table, "ShouldChangePasswordOnNextLogin")
	a.SignInTokenExpireTimeUtc = field.NewTime(table, "SignInTokenExpireTimeUtc")
	a.SignInToken = field.NewString(table, "SignInToken")
	a.GoogleAuthenticatorKey = field.NewString(table, "GoogleAuthenticatorKey")

	a.fillFieldMap()

	return a
}

func (a *abpuser) WithContext(ctx context.Context) IAbpuserDo { return a.abpuserDo.WithContext(ctx) }

func (a abpuser) TableName() string { return a.abpuserDo.TableName() }

func (a abpuser) Alias() string { return a.abpuserDo.Alias() }

func (a *abpuser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *abpuser) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 34)
	a.fieldMap["Id"] = a.ID
	a.fieldMap["CreationTime"] = a.CreationTime
	a.fieldMap["CreatorUserId"] = a.CreatorUserID
	a.fieldMap["LastModificationTime"] = a.LastModificationTime
	a.fieldMap["LastModifierUserId"] = a.LastModifierUserID
	a.fieldMap["IsDeleted"] = a.IsDeleted
	a.fieldMap["DeleterUserId"] = a.DeleterUserID
	a.fieldMap["DeletionTime"] = a.DeletionTime
	a.fieldMap["AuthenticationSource"] = a.AuthenticationSource
	a.fieldMap["UserName"] = a.UserName
	a.fieldMap["TenantId"] = a.TenantID
	a.fieldMap["EmailAddress"] = a.EmailAddress
	a.fieldMap["Name"] = a.Name
	a.fieldMap["Surname"] = a.Surname
	a.fieldMap["Password"] = a.Password
	a.fieldMap["EmailConfirmationCode"] = a.EmailConfirmationCode
	a.fieldMap["PasswordResetCode"] = a.PasswordResetCode
	a.fieldMap["LockoutEndDateUtc"] = a.LockoutEndDateUtc
	a.fieldMap["AccessFailedCount"] = a.AccessFailedCount
	a.fieldMap["IsLockoutEnabled"] = a.IsLockoutEnabled
	a.fieldMap["PhoneNumber"] = a.PhoneNumber
	a.fieldMap["IsPhoneNumberConfirmed"] = a.IsPhoneNumberConfirmed
	a.fieldMap["SecurityStamp"] = a.SecurityStamp
	a.fieldMap["IsTwoFactorEnabled"] = a.IsTwoFactorEnabled
	a.fieldMap["IsEmailConfirmed"] = a.IsEmailConfirmed
	a.fieldMap["IsActive"] = a.IsActive
	a.fieldMap["NormalizedUserName"] = a.NormalizedUserName
	a.fieldMap["NormalizedEmailAddress"] = a.NormalizedEmailAddress
	a.fieldMap["ConcurrencyStamp"] = a.ConcurrencyStamp
	a.fieldMap["ProfilePictureId"] = a.ProfilePictureID
	a.fieldMap["ShouldChangePasswordOnNextLogin"] = a.ShouldChangePasswordOnNextLogin
	a.fieldMap["SignInTokenExpireTimeUtc"] = a.SignInTokenExpireTimeUtc
	a.fieldMap["SignInToken"] = a.SignInToken
	a.fieldMap["GoogleAuthenticatorKey"] = a.GoogleAuthenticatorKey
}

func (a abpuser) clone(db *gorm.DB) abpuser {
	a.abpuserDo.ReplaceDB(db)
	return a
}

type abpuserDo struct{ gen.DO }

type IAbpuserDo interface {
	gen.SubQuery
	Debug() IAbpuserDo
	WithContext(ctx context.Context) IAbpuserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAbpuserDo
	Not(conds ...gen.Condition) IAbpuserDo
	Or(conds ...gen.Condition) IAbpuserDo
	Select(conds ...field.Expr) IAbpuserDo
	Where(conds ...gen.Condition) IAbpuserDo
	Order(conds ...field.Expr) IAbpuserDo
	Distinct(cols ...field.Expr) IAbpuserDo
	Omit(cols ...field.Expr) IAbpuserDo
	Join(table schema.Tabler, on ...field.Expr) IAbpuserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAbpuserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAbpuserDo
	Group(cols ...field.Expr) IAbpuserDo
	Having(conds ...gen.Condition) IAbpuserDo
	Limit(limit int) IAbpuserDo
	Offset(offset int) IAbpuserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAbpuserDo
	Unscoped() IAbpuserDo
	Create(values ...*model.Abpuser) error
	CreateInBatches(values []*model.Abpuser, batchSize int) error
	Save(values ...*model.Abpuser) error
	First() (*model.Abpuser, error)
	Take() (*model.Abpuser, error)
	Last() (*model.Abpuser, error)
	Find() ([]*model.Abpuser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Abpuser, err error)
	FindInBatches(result *[]*model.Abpuser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Abpuser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAbpuserDo
	Assign(attrs ...field.AssignExpr) IAbpuserDo
	Joins(fields ...field.RelationField) IAbpuserDo
	Preload(fields ...field.RelationField) IAbpuserDo
	FirstOrInit() (*model.Abpuser, error)
	FirstOrCreate() (*model.Abpuser, error)
	FindByPage(offset int, limit int) (result []*model.Abpuser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAbpuserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a abpuserDo) Debug() IAbpuserDo {
	return a.withDO(a.DO.Debug())
}

func (a abpuserDo) WithContext(ctx context.Context) IAbpuserDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a abpuserDo) ReadDB() IAbpuserDo {
	return a.Clauses(dbresolver.Read)
}

func (a abpuserDo) WriteDB() IAbpuserDo {
	return a.Clauses(dbresolver.Write)
}

func (a abpuserDo) Clauses(conds ...clause.Expression) IAbpuserDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a abpuserDo) Returning(value interface{}, columns ...string) IAbpuserDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a abpuserDo) Not(conds ...gen.Condition) IAbpuserDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a abpuserDo) Or(conds ...gen.Condition) IAbpuserDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a abpuserDo) Select(conds ...field.Expr) IAbpuserDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a abpuserDo) Where(conds ...gen.Condition) IAbpuserDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a abpuserDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAbpuserDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a abpuserDo) Order(conds ...field.Expr) IAbpuserDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a abpuserDo) Distinct(cols ...field.Expr) IAbpuserDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a abpuserDo) Omit(cols ...field.Expr) IAbpuserDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a abpuserDo) Join(table schema.Tabler, on ...field.Expr) IAbpuserDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a abpuserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAbpuserDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a abpuserDo) RightJoin(table schema.Tabler, on ...field.Expr) IAbpuserDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a abpuserDo) Group(cols ...field.Expr) IAbpuserDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a abpuserDo) Having(conds ...gen.Condition) IAbpuserDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a abpuserDo) Limit(limit int) IAbpuserDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a abpuserDo) Offset(offset int) IAbpuserDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a abpuserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAbpuserDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a abpuserDo) Unscoped() IAbpuserDo {
	return a.withDO(a.DO.Unscoped())
}

func (a abpuserDo) Create(values ...*model.Abpuser) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a abpuserDo) CreateInBatches(values []*model.Abpuser, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a abpuserDo) Save(values ...*model.Abpuser) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a abpuserDo) First() (*model.Abpuser, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpuser), nil
	}
}

func (a abpuserDo) Take() (*model.Abpuser, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpuser), nil
	}
}

func (a abpuserDo) Last() (*model.Abpuser, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpuser), nil
	}
}

func (a abpuserDo) Find() ([]*model.Abpuser, error) {
	result, err := a.DO.Find()
	return result.([]*model.Abpuser), err
}

func (a abpuserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Abpuser, err error) {
	buf := make([]*model.Abpuser, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a abpuserDo) FindInBatches(result *[]*model.Abpuser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a abpuserDo) Attrs(attrs ...field.AssignExpr) IAbpuserDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a abpuserDo) Assign(attrs ...field.AssignExpr) IAbpuserDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a abpuserDo) Joins(fields ...field.RelationField) IAbpuserDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a abpuserDo) Preload(fields ...field.RelationField) IAbpuserDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a abpuserDo) FirstOrInit() (*model.Abpuser, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpuser), nil
	}
}

func (a abpuserDo) FirstOrCreate() (*model.Abpuser, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Abpuser), nil
	}
}

func (a abpuserDo) FindByPage(offset int, limit int) (result []*model.Abpuser, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a abpuserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a abpuserDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a abpuserDo) Delete(models ...*model.Abpuser) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *abpuserDo) withDO(do gen.Dao) *abpuserDo {
	a.DO = *do.(*gen.DO)
	return a
}
