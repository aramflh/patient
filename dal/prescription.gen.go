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

func newPrescription(db *gorm.DB, opts ...gen.DOOption) prescription {
	_prescription := prescription{}

	_prescription.prescriptionDo.UseDB(db, opts...)
	_prescription.prescriptionDo.UseModel(&model.Prescription{})

	tableName := _prescription.prescriptionDo.TableName()
	_prescription.ALL = field.NewAsterisk(tableName)
	_prescription.DateEmission = field.NewTime(tableName, "date_emission")
	_prescription.DureeTraitement = field.NewString(tableName, "duree_traitement")
	_prescription.NNiss = field.NewString(tableName, "n_niss")
	_prescription.NomMedic = field.NewString(tableName, "nom_medic")
	_prescription.NInamiMed = field.NewString(tableName, "n_inami_med")
	_prescription.NInamiPha = field.NewString(tableName, "n_inami_pha")

	_prescription.fillFieldMap()

	return _prescription
}

type prescription struct {
	prescriptionDo prescriptionDo

	ALL             field.Asterisk
	DateEmission    field.Time
	DureeTraitement field.String
	NNiss           field.String
	NomMedic        field.String
	NInamiMed       field.String
	NInamiPha       field.String

	fieldMap map[string]field.Expr
}

func (p prescription) Table(newTableName string) *prescription {
	p.prescriptionDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p prescription) As(alias string) *prescription {
	p.prescriptionDo.DO = *(p.prescriptionDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *prescription) updateTableName(table string) *prescription {
	p.ALL = field.NewAsterisk(table)
	p.DateEmission = field.NewTime(table, "date_emission")
	p.DureeTraitement = field.NewString(table, "duree_traitement")
	p.NNiss = field.NewString(table, "n_niss")
	p.NomMedic = field.NewString(table, "nom_medic")
	p.NInamiMed = field.NewString(table, "n_inami_med")
	p.NInamiPha = field.NewString(table, "n_inami_pha")

	p.fillFieldMap()

	return p
}

func (p *prescription) WithContext(ctx context.Context) *prescriptionDo {
	return p.prescriptionDo.WithContext(ctx)
}

func (p prescription) TableName() string { return p.prescriptionDo.TableName() }

func (p prescription) Alias() string { return p.prescriptionDo.Alias() }

func (p *prescription) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *prescription) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["date_emission"] = p.DateEmission
	p.fieldMap["duree_traitement"] = p.DureeTraitement
	p.fieldMap["n_niss"] = p.NNiss
	p.fieldMap["nom_medic"] = p.NomMedic
	p.fieldMap["n_inami_med"] = p.NInamiMed
	p.fieldMap["n_inami_pha"] = p.NInamiPha
}

func (p prescription) clone(db *gorm.DB) prescription {
	p.prescriptionDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p prescription) replaceDB(db *gorm.DB) prescription {
	p.prescriptionDo.ReplaceDB(db)
	return p
}

type prescriptionDo struct{ gen.DO }

func (p prescriptionDo) Debug() *prescriptionDo {
	return p.withDO(p.DO.Debug())
}

func (p prescriptionDo) WithContext(ctx context.Context) *prescriptionDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p prescriptionDo) ReadDB() *prescriptionDo {
	return p.Clauses(dbresolver.Read)
}

func (p prescriptionDo) WriteDB() *prescriptionDo {
	return p.Clauses(dbresolver.Write)
}

func (p prescriptionDo) Session(config *gorm.Session) *prescriptionDo {
	return p.withDO(p.DO.Session(config))
}

func (p prescriptionDo) Clauses(conds ...clause.Expression) *prescriptionDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p prescriptionDo) Returning(value interface{}, columns ...string) *prescriptionDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p prescriptionDo) Not(conds ...gen.Condition) *prescriptionDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p prescriptionDo) Or(conds ...gen.Condition) *prescriptionDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p prescriptionDo) Select(conds ...field.Expr) *prescriptionDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p prescriptionDo) Where(conds ...gen.Condition) *prescriptionDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p prescriptionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *prescriptionDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p prescriptionDo) Order(conds ...field.Expr) *prescriptionDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p prescriptionDo) Distinct(cols ...field.Expr) *prescriptionDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p prescriptionDo) Omit(cols ...field.Expr) *prescriptionDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p prescriptionDo) Join(table schema.Tabler, on ...field.Expr) *prescriptionDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p prescriptionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *prescriptionDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p prescriptionDo) RightJoin(table schema.Tabler, on ...field.Expr) *prescriptionDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p prescriptionDo) Group(cols ...field.Expr) *prescriptionDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p prescriptionDo) Having(conds ...gen.Condition) *prescriptionDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p prescriptionDo) Limit(limit int) *prescriptionDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p prescriptionDo) Offset(offset int) *prescriptionDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p prescriptionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *prescriptionDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p prescriptionDo) Unscoped() *prescriptionDo {
	return p.withDO(p.DO.Unscoped())
}

func (p prescriptionDo) Create(values ...*model.Prescription) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p prescriptionDo) CreateInBatches(values []*model.Prescription, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p prescriptionDo) Save(values ...*model.Prescription) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p prescriptionDo) First() (*model.Prescription, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Prescription), nil
	}
}

func (p prescriptionDo) Take() (*model.Prescription, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Prescription), nil
	}
}

func (p prescriptionDo) Last() (*model.Prescription, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Prescription), nil
	}
}

func (p prescriptionDo) Find() ([]*model.Prescription, error) {
	result, err := p.DO.Find()
	return result.([]*model.Prescription), err
}

func (p prescriptionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Prescription, err error) {
	buf := make([]*model.Prescription, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p prescriptionDo) FindInBatches(result *[]*model.Prescription, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p prescriptionDo) Attrs(attrs ...field.AssignExpr) *prescriptionDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p prescriptionDo) Assign(attrs ...field.AssignExpr) *prescriptionDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p prescriptionDo) Joins(fields ...field.RelationField) *prescriptionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p prescriptionDo) Preload(fields ...field.RelationField) *prescriptionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p prescriptionDo) FirstOrInit() (*model.Prescription, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Prescription), nil
	}
}

func (p prescriptionDo) FirstOrCreate() (*model.Prescription, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Prescription), nil
	}
}

func (p prescriptionDo) FindByPage(offset int, limit int) (result []*model.Prescription, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p prescriptionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p prescriptionDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p prescriptionDo) Delete(models ...*model.Prescription) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *prescriptionDo) withDO(do gen.Dao) *prescriptionDo {
	p.DO = *do.(*gen.DO)
	return p
}