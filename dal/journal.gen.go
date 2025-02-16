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

func newJournal(db *gorm.DB, opts ...gen.DOOption) journal {
	_journal := journal{}

	_journal.journalDo.UseDB(db, opts...)
	_journal.journalDo.UseModel(&entity.Journal{})

	tableName := _journal.journalDo.TableName()
	_journal.ALL = field.NewAsterisk(tableName)
	_journal.ID = field.NewInt32(tableName, "id")
	_journal.CreateTime = field.NewTime(tableName, "create_time")
	_journal.UpdateTime = field.NewTime(tableName, "update_time")
	_journal.Content = field.NewString(tableName, "content")
	_journal.Likes = field.NewInt64(tableName, "likes")
	_journal.SourceContent = field.NewString(tableName, "source_content")
	_journal.Type = field.NewField(tableName, "type")

	_journal.fillFieldMap()

	return _journal
}

type journal struct {
	journalDo journalDo

	ALL           field.Asterisk
	ID            field.Int32
	CreateTime    field.Time
	UpdateTime    field.Time
	Content       field.String
	Likes         field.Int64
	SourceContent field.String
	Type          field.Field

	fieldMap map[string]field.Expr
}

func (j journal) Table(newTableName string) *journal {
	j.journalDo.UseTable(newTableName)
	return j.updateTableName(newTableName)
}

func (j journal) As(alias string) *journal {
	j.journalDo.DO = *(j.journalDo.As(alias).(*gen.DO))
	return j.updateTableName(alias)
}

func (j *journal) updateTableName(table string) *journal {
	j.ALL = field.NewAsterisk(table)
	j.ID = field.NewInt32(table, "id")
	j.CreateTime = field.NewTime(table, "create_time")
	j.UpdateTime = field.NewTime(table, "update_time")
	j.Content = field.NewString(table, "content")
	j.Likes = field.NewInt64(table, "likes")
	j.SourceContent = field.NewString(table, "source_content")
	j.Type = field.NewField(table, "type")

	j.fillFieldMap()

	return j
}

func (j *journal) WithContext(ctx context.Context) *journalDo { return j.journalDo.WithContext(ctx) }

func (j journal) TableName() string { return j.journalDo.TableName() }

func (j journal) Alias() string { return j.journalDo.Alias() }

func (j journal) Columns(cols ...field.Expr) gen.Columns { return j.journalDo.Columns(cols...) }

func (j *journal) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := j.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (j *journal) fillFieldMap() {
	j.fieldMap = make(map[string]field.Expr, 7)
	j.fieldMap["id"] = j.ID
	j.fieldMap["create_time"] = j.CreateTime
	j.fieldMap["update_time"] = j.UpdateTime
	j.fieldMap["content"] = j.Content
	j.fieldMap["likes"] = j.Likes
	j.fieldMap["source_content"] = j.SourceContent
	j.fieldMap["type"] = j.Type
}

func (j journal) clone(db *gorm.DB) journal {
	j.journalDo.ReplaceConnPool(db.Statement.ConnPool)
	return j
}

func (j journal) replaceDB(db *gorm.DB) journal {
	j.journalDo.ReplaceDB(db)
	return j
}

type journalDo struct{ gen.DO }

func (j journalDo) Debug() *journalDo {
	return j.withDO(j.DO.Debug())
}

func (j journalDo) WithContext(ctx context.Context) *journalDo {
	return j.withDO(j.DO.WithContext(ctx))
}

func (j journalDo) ReadDB() *journalDo {
	return j.Clauses(dbresolver.Read)
}

func (j journalDo) WriteDB() *journalDo {
	return j.Clauses(dbresolver.Write)
}

func (j journalDo) Session(config *gorm.Session) *journalDo {
	return j.withDO(j.DO.Session(config))
}

func (j journalDo) Clauses(conds ...clause.Expression) *journalDo {
	return j.withDO(j.DO.Clauses(conds...))
}

func (j journalDo) Returning(value interface{}, columns ...string) *journalDo {
	return j.withDO(j.DO.Returning(value, columns...))
}

func (j journalDo) Not(conds ...gen.Condition) *journalDo {
	return j.withDO(j.DO.Not(conds...))
}

func (j journalDo) Or(conds ...gen.Condition) *journalDo {
	return j.withDO(j.DO.Or(conds...))
}

func (j journalDo) Select(conds ...field.Expr) *journalDo {
	return j.withDO(j.DO.Select(conds...))
}

func (j journalDo) Where(conds ...gen.Condition) *journalDo {
	return j.withDO(j.DO.Where(conds...))
}

func (j journalDo) Order(conds ...field.Expr) *journalDo {
	return j.withDO(j.DO.Order(conds...))
}

func (j journalDo) Distinct(cols ...field.Expr) *journalDo {
	return j.withDO(j.DO.Distinct(cols...))
}

func (j journalDo) Omit(cols ...field.Expr) *journalDo {
	return j.withDO(j.DO.Omit(cols...))
}

func (j journalDo) Join(table schema.Tabler, on ...field.Expr) *journalDo {
	return j.withDO(j.DO.Join(table, on...))
}

func (j journalDo) LeftJoin(table schema.Tabler, on ...field.Expr) *journalDo {
	return j.withDO(j.DO.LeftJoin(table, on...))
}

func (j journalDo) RightJoin(table schema.Tabler, on ...field.Expr) *journalDo {
	return j.withDO(j.DO.RightJoin(table, on...))
}

func (j journalDo) Group(cols ...field.Expr) *journalDo {
	return j.withDO(j.DO.Group(cols...))
}

func (j journalDo) Having(conds ...gen.Condition) *journalDo {
	return j.withDO(j.DO.Having(conds...))
}

func (j journalDo) Limit(limit int) *journalDo {
	return j.withDO(j.DO.Limit(limit))
}

func (j journalDo) Offset(offset int) *journalDo {
	return j.withDO(j.DO.Offset(offset))
}

func (j journalDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *journalDo {
	return j.withDO(j.DO.Scopes(funcs...))
}

func (j journalDo) Unscoped() *journalDo {
	return j.withDO(j.DO.Unscoped())
}

func (j journalDo) Create(values ...*entity.Journal) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Create(values)
}

func (j journalDo) CreateInBatches(values []*entity.Journal, batchSize int) error {
	return j.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (j journalDo) Save(values ...*entity.Journal) error {
	if len(values) == 0 {
		return nil
	}
	return j.DO.Save(values)
}

func (j journalDo) First() (*entity.Journal, error) {
	if result, err := j.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Journal), nil
	}
}

func (j journalDo) Take() (*entity.Journal, error) {
	if result, err := j.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Journal), nil
	}
}

func (j journalDo) Last() (*entity.Journal, error) {
	if result, err := j.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Journal), nil
	}
}

func (j journalDo) Find() ([]*entity.Journal, error) {
	result, err := j.DO.Find()
	return result.([]*entity.Journal), err
}

func (j journalDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Journal, err error) {
	buf := make([]*entity.Journal, 0, batchSize)
	err = j.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (j journalDo) FindInBatches(result *[]*entity.Journal, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return j.DO.FindInBatches(result, batchSize, fc)
}

func (j journalDo) Attrs(attrs ...field.AssignExpr) *journalDo {
	return j.withDO(j.DO.Attrs(attrs...))
}

func (j journalDo) Assign(attrs ...field.AssignExpr) *journalDo {
	return j.withDO(j.DO.Assign(attrs...))
}

func (j journalDo) Joins(fields ...field.RelationField) *journalDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Joins(_f))
	}
	return &j
}

func (j journalDo) Preload(fields ...field.RelationField) *journalDo {
	for _, _f := range fields {
		j = *j.withDO(j.DO.Preload(_f))
	}
	return &j
}

func (j journalDo) FirstOrInit() (*entity.Journal, error) {
	if result, err := j.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Journal), nil
	}
}

func (j journalDo) FirstOrCreate() (*entity.Journal, error) {
	if result, err := j.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Journal), nil
	}
}

func (j journalDo) FindByPage(offset int, limit int) (result []*entity.Journal, count int64, err error) {
	result, err = j.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = j.Offset(-1).Limit(-1).Count()
	return
}

func (j journalDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = j.Count()
	if err != nil {
		return
	}

	err = j.Offset(offset).Limit(limit).Scan(result)
	return
}

func (j journalDo) Scan(result interface{}) (err error) {
	return j.DO.Scan(result)
}

func (j journalDo) Delete(models ...*entity.Journal) (result gen.ResultInfo, err error) {
	return j.DO.Delete(models)
}

func (j *journalDo) withDO(do gen.Dao) *journalDo {
	j.DO = *do.(*gen.DO)
	return j
}
