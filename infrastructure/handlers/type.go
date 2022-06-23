package handlers

type AddToBalanceRequest struct {
	UserId int     `json:"userId"`
	Amount float64 `json:"amount"`
}

type AddTransferRequest struct {
	FromId int     `json:"fromId"`
	ToId   int     `json:"toId"`
	Amount float64 `json:"amount"`
}
