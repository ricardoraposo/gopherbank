package models

type NewDepositRequestParams struct {
	AccountId string  `json:"toAccount"`
	Amount    float64 `json:"amount"`
}

type DepositParam struct {
    Account string `json:"account"`
}
