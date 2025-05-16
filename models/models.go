package models

// model Journal {
//   id              String                @id @default(cuid())
//   entryDate       DateTime              @default(now()) @db.Date
//   createdAt       DateTime              @default(now()) @db.Timestamp(0)
//   updatedAt       DateTime              @default(now()) @updatedAt @db.Timestamp(0)
//   serialNo        String?
//   references      String?
//   name            String?
//   debit           Float                 @default(0)
//   credit          Float                 @default(0)
//   Transaction     Transaction[]
//   Project         Project?              @relation(fields: [projectId], references: [id], onDelete: Cascade)
//   projectId       String?
//   FinancialYear   FinancialYear?        @relation(fields: [financialYearId], references: [id], onDelete: Cascade)
//   financialYearId String?
//   payeeName       String?
//   PaymentFrom     Ledger?               @relation(fields: [paymentFromId], references: [id], onDelete: Cascade)
//   paymentFromId   String?
//   paymentType     VOUCHER_PAYMENT_TYPE?
//   type            JOURNAL_TYPE?
//   CreatedBy       User?                 @relation(fields: [createdById], references: [id])
//   createdById     String?

//   @@map("Journals")
// }

// model Transaction {
//   id               String            @id @default(cuid())
//   entryDate        DateTime          @default(now()) @db.Date
//   createdAt        DateTime          @default(now()) @db.Timestamp(0)
//   updatedAt        DateTime          @default(now()) @updatedAt @db.Timestamp(0)
//   transactionDate  DateTime?         @db.Date
//   serialNo         Int
//   Ledger           Ledger?           @relation(name: "MainLedger", fields: [ledgerId], references: [id], onDelete: Cascade)
//   ledgerId         String?
//   Journal          Journal?          @relation(fields: [journalId], references: [id], onDelete: Cascade)
//   journalId        String?
//   debit            Float             @default(0)
//   credit           Float             @default(0)
//   type             TRANSACTION_TYPE?
//   User             User              @relation(fields: [userId], references: [id])
//   userId           String
//   Project          Project?          @relation(fields: [projectId], references: [id], onDelete: Cascade)
//   projectId        String?
//   OpeningBalance   OpeningBalance?   @relation(fields: [openingBalanceId], references: [id], onDelete: Cascade)
//   openingBalanceId String?           @unique
//   from             TRANSACTION_FROM?
//   FinancialYear    FinancialYear?    @relation(fields: [financialYearId], references: [id], onDelete: Cascade)
//   financialYearId  String?
//   AgainstLedger    Ledger?           @relation(name: "AgainstLedger", fields: [againstLedgerId], references: [id], onDelete: Cascade)
//   againstLedgerId  String?

//   @@map("Transactions")
// }
