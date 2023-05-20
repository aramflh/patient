// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTraitement = "Traitement"

// Traitement mapped from table <Traitement>
type Traitement struct {
	DateDebut       time.Time `gorm:"column:date_debut;type:date;primaryKey" json:"date_debut"`
	DureeTraitement *string   `gorm:"column:duree_traitement;type:character varying" json:"duree_traitement"`
	NNiss           string    `gorm:"column:n_niss;type:character varying;primaryKey" json:"n_niss"`
	NomMedic        string    `gorm:"column:nom_medic;type:character varying;primaryKey" json:"nom_medic"`
	NInamiMed       string    `gorm:"column:n_inami_med;type:character varying;primaryKey" json:"n_inami_med"`
	NInamiPha       string    `gorm:"column:n_inami_pha;type:character varying;primaryKey" json:"n_inami_pha"`
}

// TableName Traitement's table name
func (*Traitement) TableName() string {
	return TableNameTraitement
}