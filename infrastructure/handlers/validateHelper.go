package handlers

func ValidateParams(amount float32, ids ...int) bool {
	if amount <= 0 {
		return false
	}
	for _, id := range ids {
		if id <= 0 {
			return false
		}
	}
	return true
}
