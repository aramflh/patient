// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePathologie = "Pathologie"

// Pathologie mapped from table <Pathologie>
type Pathologie struct {
	NomPathologie string  `gorm:"column:nom_pathologie;type:character varying;primaryKey" json:"nom_pathologie"`
	NomSysAna     *string `gorm:"column:nom_sys_ana;type:character varying" json:"nom_sys_ana"`
}

// TableName Pathologie's table name
func (*Pathologie) TableName() string {
	return TableNamePathologie
}