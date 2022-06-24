package storage

import (
	"SimplyWebServer/library"
	"database/sql"
	"github.com/sirupsen/logrus"
	"time"
)

type transactionStorage struct {
	db      *sql.DB
	session map[int]User
	log     *logrus.Logger
}

func NewStorage(db *sql.DB, l *logrus.Logger) *transactionStorage {
	if db == nil {
		l.Fatal("db is nil")
	}
	session := make(map[int]User, 100)
	return &transactionStorage{
		db:      db,
		session: session,
		log:     l,
	}
}

func (s *transactionStorage) CreateUser() (int, error) {
	var userId int
	q := `INSERT INTO USERS (Balance) VALUES (0) returning id`
	err := s.db.QueryRow(q).Scan(&userId)
	if err != nil {
		s.log.Error("CreateUser:", err)
		return -1, err
	}
	s.session[userId] = User{userId, 0}
	return userId, nil
}

func (s *transactionStorage) UpdateBalanceByUserId(userId int, amount float64) (float64, error) {
	balance, err := s.getBalanceByUserId(userId)
	if err != nil {
		return 0, err
	}

	if amount < 0 && balance+amount < 0 {
		return 0, library.ErrBalanceBeZero

	}
	q := `UPDATE USERS SET BALANCE=$1 WHERE Id=$2`
	_,err = s.db.Exec(q, s.session[userId].Balance+amount, userId)
	if err != nil {
		s.log.Error("UpdateBalanceByUserId:", err)
		return 0, err
	}
	s.session[userId] = User{userId, s.session[userId].Balance + amount}
	return s.session[userId].Balance, nil
}

func (s *transactionStorage) AddTransaction(transactType library.TransactionType, userId int, amount float64) error {
	q := `INSERT INTO TRANSACTIONS (type, userId, amount, time) VALUES ($1,$2,$3,$4)`
	_, err := s.db.Exec(q, int(transactType), userId, amount, time.Now())
	if err != nil {
		s.log.Error("AddTransaction:", err)
		return err
	}
	return nil
}

func (s *transactionStorage) getBalanceByUserId(userId int) (float64, error) {
	var balance sql.NullFloat64
	if _, ok := s.session[userId]; !ok {
		q := `SELECT SUM(amount) FROM TRANSACTIONS WHERE userid=$1`
		err := s.db.QueryRow(q, userId).Scan(&balance)
		if err != nil {
			s.log.Error("getBalanceByUserId:", err)
			return -1, err
		}		
			s.session[userId] = User{userId, balance.Float64}
		
	}

	return s.session[userId].Balance, nil
}
