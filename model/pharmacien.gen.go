// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePharmacien = "Pharmacien"

// Pharmacien mapped from table <Pharmacien>
type Pharmacien struct {
	NInami     string  `gorm:"column:n_inami;type:character varying;not null" json:"n_inami"`
	Nom        string  `gorm:"column:nom;type:character varying;not null" json:"nom"`
	Prenom     string  `gorm:"column:prenom;type:character varying;not null" json:"prenom"`
	AMail      *string `gorm:"column:a_mail;type:character varying" json:"a_mail"`
	NTelephone *string `gorm:"column:n_telephone;type:character varying" json:"n_telephone"`
	NInamiPha  string  `gorm:"column:n_inami_pha;type:character varying;primaryKey" json:"n_inami_pha"`
}

// TableName Pharmacien's table name
func (*Pharmacien) TableName() string {
	return TableNamePharmacien
}