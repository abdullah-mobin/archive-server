package database

import (
	"archive-server/config"
	"archive-server/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ArchiveTransaction(transaction models.Transaction) error {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("Failed to get config: %v", err)
		return err
	}
	dbName := cfg.MongoDB.Database
	collectionName := cfg.MongoDB.TransactionCollection
	collection := DB.Database(dbName).Collection(collectionName)

	var journalArchive models.JournalArchive
	err = DB.Database(dbName).Collection(cfg.MongoDB.JournalCollection).
		FindOne(context.Background(), primitive.M{"id": transaction.JournalID}).Decode(&journalArchive)
	if err != nil {
		log.Printf("Failed to find journal archive: %v", err)
		return err
	}

	archive := models.TransactionArchive{
		Id:               primitive.NewObjectID(),
		JournalArchiveId: journalArchive.Id,
		ID:               transaction.ID,
		EntryDate:        transaction.EntryDate,
		CreatedAt:        transaction.CreatedAt,
		UpdatedAt:        transaction.UpdatedAt,
		TransactionDate:  transaction.TransactionDate,
		SerialNo:         transaction.SerialNo,
		Ledger:           transaction.Ledger,
		LedgerID:         transaction.LedgerID,
		JournalID:        transaction.JournalID,
		Debit:            transaction.Debit,
		Credit:           transaction.Credit,
		Type:             transaction.Type,
		User:             transaction.User,
		UserID:           transaction.UserID,
		Project:          transaction.Project,
		ProjectID:        transaction.ProjectID,
		OpeningBalance:   transaction.OpeningBalance,
		OpeningBalanceID: transaction.OpeningBalanceID,
		From:             transaction.From,
		FinancialYear:    transaction.FinancialYear,
		FinancialYearID:  transaction.FinancialYearID,
		AgainstLedger:    transaction.AgainstLedger,
		AgainstLedgerID:  transaction.AgainstLedgerID,
	}
	_, err = collection.InsertOne(context.Background(), archive)
	if err != nil {
		log.Printf("Failed to insert transaction archive: %v", err)
		return err
	}

	return nil
}

func ArchiveTransactions(journalArchiveID primitive.ObjectID, transactions []models.Transaction) error {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("Failed to get config: %v", err)
		return err
	}
	dbName := cfg.MongoDB.Database
	collectionName := cfg.MongoDB.TransactionCollection
	collection := DB.Database(dbName).Collection(collectionName)

	docs := make([]interface{}, len(transactions))
	for i, t := range transactions {
		docs[i] = models.TransactionArchive{
			Id:               primitive.NewObjectID(),
			JournalArchiveId: journalArchiveID,
			ID:               t.ID,
			EntryDate:        t.EntryDate,
			CreatedAt:        t.CreatedAt,
			UpdatedAt:        t.UpdatedAt,
			TransactionDate:  t.TransactionDate,
			SerialNo:         t.SerialNo,
			Ledger:           t.Ledger,
			LedgerID:         t.LedgerID,
			JournalID:        t.JournalID,
			Debit:            t.Debit,
			Credit:           t.Credit,
			Type:             t.Type,
			User:             t.User,
			UserID:           t.UserID,
			Project:          t.Project,
			ProjectID:        t.ProjectID,
			OpeningBalance:   t.OpeningBalance,
			OpeningBalanceID: t.OpeningBalanceID,
			From:             t.From,
			FinancialYear:    t.FinancialYear,
			FinancialYearID:  t.FinancialYearID,
			AgainstLedger:    t.AgainstLedger,
			AgainstLedgerID:  t.AgainstLedgerID,
		}
	}

	_, err = collection.InsertMany(context.Background(), docs)
	if err != nil {
		return err
	}

	return nil
}
