package database

import (
	"archive-server/config"
	"archive-server/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ArchiveJournal(journal models.Journal) error {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("Failed to get config: %v", err)
		return err
	}
	dbName := cfg.MongoDB.Database
	collectionName := cfg.MongoDB.JournalCollection
	collection := DB.Database(dbName).Collection(collectionName)

	archive := models.JournalArchive{
		Id:              primitive.NewObjectID(),
		ID:              journal.ID,
		EntryDate:       journal.EntryDate,
		CreatedAt:       journal.CreatedAt,
		UpdatedAt:       journal.UpdatedAt,
		SerialNo:        journal.SerialNo,
		References:      journal.References,
		Name:            journal.Name,
		Debit:           journal.Debit,
		Credit:          journal.Credit,
		Project:         journal.Project,
		ProjectID:       journal.ProjectID,
		FinancialYear:   journal.FinancialYear,
		FinancialYearID: journal.FinancialYearID,
		PayeeName:       journal.PayeeName,
		PaymentFrom:     journal.PaymentFrom,
		PaymentFromID:   journal.PaymentFromID,
		PaymentType:     journal.PaymentType,
		Type:            journal.Type,
		CreatedBy:       journal.CreatedBy,
		CreatedByID:     journal.CreatedByID,
	}

	_, err = collection.InsertOne(context.Background(), archive)
	if err != nil {
		return err
	}

	if err := ArchiveTransactions(archive.Id, journal.Transactions); err != nil {
		log.Printf("Failed to archive transactions: %v", err)
		return err
	}

	return nil
}
