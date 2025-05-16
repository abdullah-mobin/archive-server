package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Journal struct {
	ID              string        `json:"id" bson:"id,omitempty"`
	EntryDate       time.Time     `json:"entryDate" bson:"entryDate"`
	CreatedAt       time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time     `json:"updatedAt" bson:"updatedAt"`
	SerialNo        *string       `json:"serialNo,omitempty" bson:"serialNo,omitempty"`
	References      *string       `json:"references,omitempty" bson:"references,omitempty"`
	Name            *string       `json:"name,omitempty" bson:"name,omitempty"`
	Debit           float64       `json:"debit" bson:"debit"`
	Credit          float64       `json:"credit" bson:"credit"`
	Transactions    []Transaction `json:"transactions,omitempty" bson:"transactions,omitempty"`
	Project         *string       `json:"project,omitempty" bson:"project,omitempty"`
	ProjectID       *string       `json:"projectId,omitempty" bson:"projectId,omitempty"`
	FinancialYear   *string       `json:"financialYear,omitempty" bson:"financialYear,omitempty"`
	FinancialYearID *string       `json:"financialYearId,omitempty" bson:"financialYearId,omitempty"`
	PayeeName       *string       `json:"payeeName,omitempty" bson:"payeeName,omitempty"`
	PaymentFrom     *string       `json:"paymentFrom,omitempty" bson:"paymentFrom,omitempty"`
	PaymentFromID   *string       `json:"paymentFromId,omitempty" bson:"paymentFromId,omitempty"`
	PaymentType     *string       `json:"paymentType,omitempty" bson:"paymentType,omitempty"`
	Type            *string       `json:"type,omitempty" bson:"type,omitempty"`
	CreatedBy       *string       `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	CreatedByID     *string       `json:"createdById,omitempty" bson:"createdById,omitempty"`
}

type JournalArchive struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	ID              string             `json:"id" bson:"id,omitempty"`
	EntryDate       time.Time          `json:"entryDate" bson:"entryDate"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt" bson:"updatedAt"`
	SerialNo        *string            `json:"serialNo,omitempty" bson:"serialNo,omitempty"`
	References      *string            `json:"references,omitempty" bson:"references,omitempty"`
	Name            *string            `json:"name,omitempty" bson:"name,omitempty"`
	Debit           float64            `json:"debit" bson:"debit"`
	Credit          float64            `json:"credit" bson:"credit"`
	Project         *string            `json:"project,omitempty" bson:"project,omitempty"`
	ProjectID       *string            `json:"projectId,omitempty" bson:"projectId,omitempty"`
	FinancialYear   *string            `json:"financialYear,omitempty" bson:"financialYear,omitempty"`
	FinancialYearID *string            `json:"financialYearId,omitempty" bson:"financialYearId,omitempty"`
	PayeeName       *string            `json:"payeeName,omitempty" bson:"payeeName,omitempty"`
	PaymentFrom     *string            `json:"paymentFrom,omitempty" bson:"paymentFrom,omitempty"`
	PaymentFromID   *string            `json:"paymentFromId,omitempty" bson:"paymentFromId,omitempty"`
	PaymentType     *string            `json:"paymentType,omitempty" bson:"paymentType,omitempty"`
	Type            *string            `json:"type,omitempty" bson:"type,omitempty"`
	CreatedBy       *string            `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	CreatedByID     *string            `json:"createdById,omitempty" bson:"createdById,omitempty"`
}
