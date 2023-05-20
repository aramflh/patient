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

	"patient/model"
)

func newTraitement(db *gorm.DB, opts ...gen.DOOption) traitement {
	_traitement := traitement{}

	_traitement.traitementDo.UseDB(db, opts...)
	_traitement.traitementDo.UseModel(&model.Traitement{})

	tableName := _traitement.traitementDo.TableName()
	_traitement.ALL = field.NewAsterisk(tableName)
	_traitement.DateDebut = field.NewTime(tableName, "date_debut")
	_traitement.DureeTraitement = field.NewString(tableName, "duree_traitement")
	_traitement.NNiss = field.NewString(tableName, "n_niss")
	_traitement.NomMedic = field.NewString(tableName, "nom_medic")
	_traitement.NInamiMed = field.NewString(tableName, "n_inami_med")
	_traitement.NInamiPha = field.NewString(tableName, "n_inami_pha")

	_traitement.fillFieldMap()

	return _traitement
}

type traitement struct {
	traitementDo traitementDo

	ALL             field.Asterisk
	DateDebut       field.Time
	DureeTraitement field.String
	NNiss           field.String
	NomMedic        field.String
	NInamiMed       field.String
	NInamiPha       field.String

	fieldMap map[string]field.Expr
}

func (t traitement) Table(newTableName string) *traitement {
	t.traitementDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t traitement) As(alias string) *traitement {
	t.traitementDo.DO = *(t.traitementDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *traitement) updateTableName(table string) *traitement {
	t.ALL = field.NewAsterisk(table)
	t.DateDebut = field.NewTime(table, "date_debut")
	t.DureeTraitement = field.NewString(table, "duree_traitement")
	t.NNiss = field.NewString(table, "n_niss")
	t.NomMedic = field.NewString(table, "nom_medic")
	t.NInamiMed = field.NewString(table, "n_inami_med")
	t.NInamiPha = field.NewString(table, "n_inami_pha")

	t.fillFieldMap()

	return t
}

func (t *traitement) WithContext(ctx context.Context) *traitementDo {
	return t.traitementDo.WithContext(ctx)
}

func (t traitement) TableName() string { return t.traitementDo.TableName() }

func (t traitement) Alias() string { return t.traitementDo.Alias() }

func (t *traitement) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *traitement) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 6)
	t.fieldMap["date_debut"] = t.DateDebut
	t.fieldMap["duree_traitement"] = t.DureeTraitement
	t.fieldMap["n_niss"] = t.NNiss
	t.fieldMap["nom_medic"] = t.NomMedic
	t.fieldMap["n_inami_med"] = t.NInamiMed
	t.fieldMap["n_inami_pha"] = t.NInamiPha
}

func (t traitement) clone(db *gorm.DB) traitement {
	t.traitementDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t traitement) replaceDB(db *gorm.DB) traitement {
	t.traitementDo.ReplaceDB(db)
	return t
}

type traitementDo struct{ gen.DO }

func (t traitementDo) Debug() *traitementDo {
	return t.withDO(t.DO.Debug())
}

func (t traitementDo) WithContext(ctx context.Context) *traitementDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t traitementDo) ReadDB() *traitementDo {
	return t.Clauses(dbresolver.Read)
}

func (t traitementDo) WriteDB() *traitementDo {
	return t.Clauses(dbresolver.Write)
}

func (t traitementDo) Session(config *gorm.Session) *traitementDo {
	return t.withDO(t.DO.Session(config))
}

func (t traitementDo) Clauses(conds ...clause.Expression) *traitementDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t traitementDo) Returning(value interface{}, columns ...string) *traitementDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t traitementDo) Not(conds ...gen.Condition) *traitementDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t traitementDo) Or(conds ...gen.Condition) *traitementDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t traitementDo) Select(conds ...field.Expr) *traitementDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t traitementDo) Where(conds ...gen.Condition) *traitementDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t traitementDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *traitementDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t traitementDo) Order(conds ...field.Expr) *traitementDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t traitementDo) Distinct(cols ...field.Expr) *traitementDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t traitementDo) Omit(cols ...field.Expr) *traitementDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t traitementDo) Join(table schema.Tabler, on ...field.Expr) *traitementDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t traitementDo) LeftJoin(table schema.Tabler, on ...field.Expr) *traitementDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t traitementDo) RightJoin(table schema.Tabler, on ...field.Expr) *traitementDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t traitementDo) Group(cols ...field.Expr) *traitementDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t traitementDo) Having(conds ...gen.Condition) *traitementDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t traitementDo) Limit(limit int) *traitementDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t traitementDo) Offset(offset int) *traitementDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t traitementDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *traitementDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t traitementDo) Unscoped() *traitementDo {
	return t.withDO(t.DO.Unscoped())
}

func (t traitementDo) Create(values ...*model.Traitement) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t traitementDo) CreateInBatches(values []*model.Traitement, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t traitementDo) Save(values ...*model.Traitement) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t traitementDo) First() (*model.Traitement, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Traitement), nil
	}
}

func (t traitementDo) Take() (*model.Traitement, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Traitement), nil
	}
}

func (t traitementDo) Last() (*model.Traitement, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Traitement), nil
	}
}

func (t traitementDo) Find() ([]*model.Traitement, error) {
	result, err := t.DO.Find()
	return result.([]*model.Traitement), err
}

func (t traitementDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Traitement, err error) {
	buf := make([]*model.Traitement, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t traitementDo) FindInBatches(result *[]*model.Traitement, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t traitementDo) Attrs(attrs ...field.AssignExpr) *traitementDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t traitementDo) Assign(attrs ...field.AssignExpr) *traitementDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t traitementDo) Joins(fields ...field.RelationField) *traitementDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t traitementDo) Preload(fields ...field.RelationField) *traitementDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t traitementDo) FirstOrInit() (*model.Traitement, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Traitement), nil
	}
}

func (t traitementDo) FirstOrCreate() (*model.Traitement, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Traitement), nil
	}
}

func (t traitementDo) FindByPage(offset int, limit int) (result []*model.Traitement, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t traitementDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t traitementDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t traitementDo) Delete(models ...*model.Traitement) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *traitementDo) withDO(do gen.Dao) *traitementDo {
	t.DO = *do.(*gen.DO)
	return t
}