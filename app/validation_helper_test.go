package app

import (
	"SimplyWebServer/library"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateParams(t *testing.T) {
	type args struct {
		amount float64
		ids    []int
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{"should return no err",
			args{10,[]int{2,1}},
			nil,
		},
		{"should return ErrEqualIDs if ids is equal",
			args{10,[]int{1,1}},
			library.ErrEqualIDs,
		},
		{"should return ErrSubZeroID if id<0 ",
			args{10,[]int{-1,2}},
			library.ErrSubZeroID,
		},
		{"should return ErrSubZeroAmount if amount<0",
			args{-10,[]int{1,2}},
			library.ErrSubZeroAmount,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			 err := ValidateParams(tt.args.amount, tt.args.ids...)
			 if err!=nil{
				assert.EqualError(t,err, tt.wantErr.Error())
			}
		})
	}
}