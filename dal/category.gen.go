// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/go-sonic/sonic/model/entity"
)

func newCategory(db *gorm.DB, opts ...gen.DOOption) category {
	_category := category{}

	_category.categoryDo.UseDB(db, opts...)
	_category.categoryDo.UseModel(&entity.Category{})

	tableName := _category.categoryDo.TableName()
	_category.ALL = field.NewAsterisk(tableName)
	_category.ID = field.NewInt32(tableName, "id")
	_category.CreateTime = field.NewTime(tableName, "create_time")
	_category.UpdateTime = field.NewTime(tableName, "update_time")
	_category.Description = field.NewString(tableName, "description")
	_category.Type = field.NewField(tableName, "type")
	_category.Name = field.NewString(tableName, "name")
	_category.ParentID = field.NewInt32(tableName, "parent_id")
	_category.Password = field.NewString(tableName, "password")
	_category.Slug = field.NewString(tableName, "slug")
	_category.Thumbnail = field.NewString(tableName, "thumbnail")
	_category.Priority = field.NewInt32(tableName, "priority")

	_category.fillFieldMap()

	return _category
}

type category struct {
	categoryDo categoryDo

	ALL         field.Asterisk
	ID          field.Int32
	CreateTime  field.Time
	UpdateTime  field.Time
	Description field.String
	Type        field.Field
	Name        field.String
	ParentID    field.Int32
	Password    field.String
	Slug        field.String
	Thumbnail   field.String
	Priority    field.Int32

	fieldMap map[string]field.Expr
}

func (c category) Table(newTableName string) *category {
	c.categoryDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c category) As(alias string) *category {
	c.categoryDo.DO = *(c.categoryDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *category) updateTableName(table string) *category {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.CreateTime = field.NewTime(table, "create_time")
	c.UpdateTime = field.NewTime(table, "update_time")
	c.Description = field.NewString(table, "description")
	c.Type = field.NewField(table, "type")
	c.Name = field.NewString(table, "name")
	c.ParentID = field.NewInt32(table, "parent_id")
	c.Password = field.NewString(table, "password")
	c.Slug = field.NewString(table, "slug")
	c.Thumbnail = field.NewString(table, "thumbnail")
	c.Priority = field.NewInt32(table, "priority")

	c.fillFieldMap()

	return c
}

func (c *category) WithContext(ctx context.Context) *categoryDo { return c.categoryDo.WithContext(ctx) }

func (c category) TableName() string { return c.categoryDo.TableName() }

func (c category) Alias() string { return c.categoryDo.Alias() }

func (c category) Columns(cols ...field.Expr) gen.Columns { return c.categoryDo.Columns(cols...) }

func (c *category) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *category) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 11)
	c.fieldMap["id"] = c.ID
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["update_time"] = c.UpdateTime
	c.fieldMap["description"] = c.Description
	c.fieldMap["type"] = c.Type
	c.fieldMap["name"] = c.Name
	c.fieldMap["parent_id"] = c.ParentID
	c.fieldMap["password"] = c.Password
	c.fieldMap["slug"] = c.Slug
	c.fieldMap["thumbnail"] = c.Thumbnail
	c.fieldMap["priority"] = c.Priority
}

func (c category) clone(db *gorm.DB) category {
	c.categoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c category) replaceDB(db *gorm.DB) category {
	c.categoryDo.ReplaceDB(db)
	return c
}

type categoryDo struct{ gen.DO }

func (c categoryDo) Debug() *categoryDo {
	return c.withDO(c.DO.Debug())
}

func (c categoryDo) WithContext(ctx context.Context) *categoryDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c categoryDo) ReadDB() *categoryDo {
	return c.Clauses(dbresolver.Read)
}

func (c categoryDo) WriteDB() *categoryDo {
	return c.Clauses(dbresolver.Write)
}

func (c categoryDo) Session(config *gorm.Session) *categoryDo {
	return c.withDO(c.DO.Session(config))
}

func (c categoryDo) Clauses(conds ...clause.Expression) *categoryDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c categoryDo) Returning(value interface{}, columns ...string) *categoryDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c categoryDo) Not(conds ...gen.Condition) *categoryDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c categoryDo) Or(conds ...gen.Condition) *categoryDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c categoryDo) Select(conds ...field.Expr) *categoryDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c categoryDo) Where(conds ...gen.Condition) *categoryDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c categoryDo) Order(conds ...field.Expr) *categoryDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c categoryDo) Distinct(cols ...field.Expr) *categoryDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c categoryDo) Omit(cols ...field.Expr) *categoryDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c categoryDo) Join(table schema.Tabler, on ...field.Expr) *categoryDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c categoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *categoryDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c categoryDo) RightJoin(table schema.Tabler, on ...field.Expr) *categoryDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c categoryDo) Group(cols ...field.Expr) *categoryDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c categoryDo) Having(conds ...gen.Condition) *categoryDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c categoryDo) Limit(limit int) *categoryDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c categoryDo) Offset(offset int) *categoryDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c categoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *categoryDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c categoryDo) Unscoped() *categoryDo {
	return c.withDO(c.DO.Unscoped())
}

func (c categoryDo) Create(values ...*entity.Category) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c categoryDo) CreateInBatches(values []*entity.Category, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c categoryDo) Save(values ...*entity.Category) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c categoryDo) First() (*entity.Category, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Category), nil
	}
}

func (c categoryDo) Take() (*entity.Category, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Category), nil
	}
}

func (c categoryDo) Last() (*entity.Category, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Category), nil
	}
}

func (c categoryDo) Find() ([]*entity.Category, error) {
	result, err := c.DO.Find()
	return result.([]*entity.Category), err
}

func (c categoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Category, err error) {
	buf := make([]*entity.Category, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c categoryDo) FindInBatches(result *[]*entity.Category, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c categoryDo) Attrs(attrs ...field.AssignExpr) *categoryDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c categoryDo) Assign(attrs ...field.AssignExpr) *categoryDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c categoryDo) Joins(fields ...field.RelationField) *categoryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c categoryDo) Preload(fields ...field.RelationField) *categoryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c categoryDo) FirstOrInit() (*entity.Category, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Category), nil
	}
}

func (c categoryDo) FirstOrCreate() (*entity.Category, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Category), nil
	}
}

func (c categoryDo) FindByPage(offset int, limit int) (result []*entity.Category, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c categoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c categoryDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c categoryDo) Delete(models ...*entity.Category) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *categoryDo) withDO(do gen.Dao) *categoryDo {
	c.DO = *do.(*gen.DO)
	return c
}
