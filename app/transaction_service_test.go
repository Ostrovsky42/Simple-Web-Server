package app

import (
	"SimplyWebServer/library"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)


func Test_transactionService_AddToBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	stubErr := errors.New("test error")
	type args struct {
		userId int
		amount float64
	}
	//arrangeq
	tests := []struct {
		name    string
		Storage  *MockStorage
		logger 	 *logrus.Logger
		args    args
		want    float64
		wantErr error
	}{
		{
			name: "should return ok",
			Storage: func(mock *MockStorage)*MockStorage {
				mock.EXPECT().UpdateBalanceByUserId(1,100.0).Return(100.0,nil)
				mock.EXPECT().AddTransaction(library.Adding,1,100.0).Return(nil)
				return mock
			}(NewMockStorage(ctrl)),
			args: args{userId: 1,amount: 100.0},
			want: 100.0,
			wantErr: nil,
		},
		{
			name: "should return ErrSubZeroID when id<0",
			Storage: func(mock *MockStorage)*MockStorage {
				return mock
			}(NewMockStorage(ctrl)),
			args: args{userId:0,amount: 100.0},
			want: 0.0,
			wantErr: library.ErrSubZeroID,
		},
		{
			name: "should return ErrSubZeroAmount when amount<0",
			Storage: func(mock *MockStorage)*MockStorage {

				return mock
			}(NewMockStorage(ctrl)),
			args: args{userId:0,amount: 100.0},
			want: 0.0,
			wantErr: library.ErrSubZeroAmount,
		},
		{
			name: "should return err when UpdateBalanceByUserId return err",
			Storage: func(mock *MockStorage)*MockStorage {
				mock.EXPECT().UpdateBalanceByUserId(1,100.0).Return(0.0,stubErr)
				return mock
			}(NewMockStorage(ctrl)),
			args: args{userId:1,amount: 100.0},
			want: 0.0,
			wantErr: stubErr,
		},
		{
			name: "should return err when AddTransaction return err",
			Storage: func(mock *MockStorage)*MockStorage {
				mock.EXPECT().UpdateBalanceByUserId(1,100.0).Return(100.0,nil)
				mock.EXPECT().AddTransaction(library.Adding,1,100.0).Return(stubErr)
				return mock
			}(NewMockStorage(ctrl)),
			args: args{userId:1,amount: 100.0},
			want: 0.0,
			wantErr: stubErr,
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &transactionService{
				Storage: tt.Storage,
				log: logrus.New(),
			}
			balance, err := t.AddToBalance(tt.args.userId, tt.args.amount)
			if err != nil{
				assert.Error(t1,err,tt.wantErr.Error())
				return
			}else {
				assert.Equal(t1,balance, tt.want)
			}
		})
	}
}

func Test_transactionService_AddTransfer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	stubErr := errors.New("test error")
	type args struct {
		fromId int
		toId int
		amount float64
	}
	type want struct {
		fromBalance float64
		toBalance 	float64
	}
	//arrangeq
	tests := []struct {
		name    string
		Storage  *MockStorage
		logger 	 *logrus.Logger
		args    args
		want    want
		wantErr error
	}{
		{
			name: "should return ok",
			Storage: func(mock *MockStorage)*MockStorage {
				mock.EXPECT().UpdateBalanceByUserId(1,-100.0).Return(0.0,nil)
				mock.EXPECT().UpdateBalanceByUserId(2,100.0).Return(100.0,nil)
				mock.EXPECT().AddTransaction(library.Transaction,1,-100.0).Return(nil)
				mock.EXPECT().AddTransaction(library.Transaction,2,100.0).Return(nil)
				return mock
			}(NewMockStorage(ctrl)),
			args: args{fromId:1, toId:2, amount: 100.0},
			want: want{0.0,100.0},
			wantErr: nil,
		},
		{
			name: "should return err if UpdateBalanceByUserId return err",
			Storage: func(mock *MockStorage)*MockStorage {
				mock.EXPECT().UpdateBalanceByUserId(1,-100.0).Return(0.0,stubErr)
				return mock
			}(NewMockStorage(ctrl)),
			args: args{fromId:1, toId:2, amount: 100.0},
			wantErr: stubErr,
		},
		{
			name: "should return err if AddTransaction return err",
			Storage: func(mock *MockStorage)*MockStorage {
				mock.EXPECT().UpdateBalanceByUserId(1,-100.0).Return(0.0,nil)
				mock.EXPECT().UpdateBalanceByUserId(2,100.0).Return(1000.0,nil)
				mock.EXPECT().AddTransaction(library.Transaction,1,-100.0).Return(stubErr)
				return mock
			}(NewMockStorage(ctrl)),
			args: args{fromId:1, toId:2, amount: 100.0},
			wantErr: stubErr,
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &transactionService{
				Storage: tt.Storage,
				log:     logrus.New(),
			}
			fromBalance,toBalance, err := t.AddTransfer(tt.args.fromId,tt.args.toId, tt.args.amount)
			if err != nil{
				assert.Error(t1,err,tt.wantErr.Error())
				return
			}else {
				assert.Equal(t1,fromBalance, tt.want.fromBalance)
				assert.Equal(t1,toBalance, tt.want.toBalance)
			}
		})
	}
}