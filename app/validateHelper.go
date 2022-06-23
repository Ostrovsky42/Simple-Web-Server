package app

import (
	"SimplyWebServer/library"
)

func ValidateParams(amount float64, ids ...int) error {
	if amount <= 0 {
		return library.ErrSubZeroAmount
	}
	if len(ids)>1&&ids[0]==ids[1]{
		return library.ErrEqualIDs
	}
	for _, id := range ids {
		if id <= 0 {
			return  library.ErrSubZeroID
		}
	}
	return nil
}
