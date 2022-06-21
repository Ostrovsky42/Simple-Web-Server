package storage

import (
	"SimplyWebServer/library"
	"database/sql"
	"log"
	"time"
)

type transactionStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *transactionStorage {
	return &transactionStorage{db: db}
}

func (s *transactionStorage) CreateUser() (int, error) {
	var userId int
	q := `INSERT INTO USERS (Balance) VALUES (0) returning id`
	err := s.db.QueryRow(q).Scan(&userId)
	if err != nil {
		log.Print("Query error", err)
		return -1, err
	}
	return userId, nil
}

func (s *transactionStorage) getBalanceByUserId(userId int) (float32, error) {
	var balance float32
	q := `SELECT balance FROM USERS  WHERE Id=$1`
	err := s.db.QueryRow(q, userId).Scan(&balance)
	if err != nil {
		log.Print("Query error", err)
		return -1, err
	}
	return balance, nil
}

func (s *transactionStorage) AddToBalance(userId int, amount float32) error {
	balance, err := s.getBalanceByUserId(userId)
	if err != nil {
		log.Print("GetBalanceByUserId Query error", err)
		return err
	}
	q := `UPDATE USERS SET BALANCE=$1 WHERE Id=$2`
	_, err = s.db.Exec(q, balance+amount, userId)
	if err != nil {
		log.Print("Query error", err)
		return err
	}
	return nil
}

func (s *transactionStorage) AddTransaction(transactType library.TransactionType, userId int, amount float32) error {
	q := `INSERT INTO TRANSACTIONS (type, userId, amount, time) VALUES ($1,$2,$3,$4)`
	_, err := s.db.Exec(q, int(transactType), userId, amount, time.Now())
	if err != nil {
		log.Print("Query error", err)
		return err
	}
	return nil
}
