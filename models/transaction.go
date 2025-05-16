package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID               string     `json:"id" bson:"id,omitempty"`
	EntryDate        time.Time  `json:"entryDate" bson:"entryDate"`
	CreatedAt        time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time  `json:"updatedAt" bson:"updatedAt"`
	TransactionDate  *time.Time `json:"transactionDate,omitempty" bson:"transactionDate,omitempty"`
	SerialNo         int        `json:"serialNo,omitempty" bson:"serialNo,omitempty"`
	Ledger           *string    `json:"ledger,omitempty" bson:"ledger,omitempty"`
	LedgerID         *string    `json:"ledgerId,omitempty" bson:"ledgerId,omitempty"`
	Journal          *Journal   `json:"journal,omitempty" bson:"journal,omitempty"`
	JournalID        *string    `json:"journalId,omitempty" bson:"journalId,omitempty"`
	Debit            float64    `json:"debit,omitempty" bson:"debit,omitempty"`
	Credit           float64    `json:"credit,omitempty" bson:"credit,omitempty"`
	Type             *string    `json:"type,omitempty" bson:"type,omitempty"`
	User             *string    `json:"user,omitempty" bson:"user,omitempty"`
	UserID           string     `json:"userId,omitempty" bson:"userId,omitempty"`
	Project          *string    `json:"project,omitempty" bson:"project,omitempty"`
	ProjectID        *string    `json:"projectId,omitempty" bson:"projectId,omitempty"`
	OpeningBalance   *string    `json:"openingBalance,omitempty" bson:"openingBalance,omitempty"`
	OpeningBalanceID *string    `json:"openingBalanceId,omitempty" bson:"openingBalanceId,omitempty"`
	From             *string    `json:"from,omitempty" bson:"from,omitempty"`
	FinancialYear    *string    `json:"financialYear,omitempty" bson:"financialYear,omitempty"`
	FinancialYearID  *string    `json:"financialYearId,omitempty" bson:"financialYearId,omitempty"`
	AgainstLedger    *string    `json:"againstLedger,omitempty" bson:"againstLedger,omitempty"`
	AgainstLedgerID  *string    `json:"againstLedgerId,omitempty" bson:"againstLedgerId,omitempty"`
}

type TransactionArchive struct {
	Id               primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	JournalArchiveId primitive.ObjectID `json:"journalArchiveId" bson:"journalArchiveId"`
	ID               string             `json:"id" bson:"id,omitempty"`
	EntryDate        time.Time          `json:"entryDate" bson:"entryDate"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt" bson:"updatedAt"`
	TransactionDate  *time.Time         `json:"transactionDate,omitempty" bson:"transactionDate,omitempty"`
	SerialNo         int                `json:"serialNo,omitempty" bson:"serialNo,omitempty"`
	Ledger           *string            `json:"ledger,omitempty" bson:"ledger,omitempty"`
	LedgerID         *string            `json:"ledgerId,omitempty" bson:"ledgerId,omitempty"`
	JournalID        *string            `json:"journalId,omitempty" bson:"journalId,omitempty"`
	Debit            float64            `json:"debit,omitempty" bson:"debit,omitempty"`
	Credit           float64            `json:"credit,omitempty" bson:"credit,omitempty"`
	Type             *string            `json:"type,omitempty" bson:"type,omitempty"`
	User             *string            `json:"user,omitempty" bson:"user,omitempty"`
	UserID           string             `json:"userId,omitempty" bson:"userId,omitempty"`
	Project          *string            `json:"project,omitempty" bson:"project,omitempty"`
	ProjectID        *string            `json:"projectId,omitempty" bson:"projectId,omitempty"`
	OpeningBalance   *string            `json:"openingBalance,omitempty" bson:"openingBalance,omitempty"`
	OpeningBalanceID *string            `json:"openingBalanceId,omitempty" bson:"openingBalanceId,omitempty"`
	From             *string            `json:"from,omitempty" bson:"from,omitempty"`
	FinancialYear    *string            `json:"financialYear,omitempty" bson:"financialYear,omitempty"`
	FinancialYearID  *string            `json:"financialYearId,omitempty" bson:"financialYearId,omitempty"`
	AgainstLedger    *string            `json:"againstLedger,omitempty" bson:"againstLedger,omitempty"`
	AgainstLedgerID  *string            `json:"againstLedgerId,omitempty" bson:"againstLedgerId,omitempty"`
}
