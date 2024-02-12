package models

type Clients struct {
	ID      int64 `json: "id"`
	Limit   int64 `json: "value"`
	Balance int64 `json: "value"`
}

type Transactions struct {
	ID          int64  `json: "id"`
	Value       int64  `json: "value"`
	Type        string `json: "type"`
	Description string `json: "description"`
	CreatedAt   int64  `json: "createdat"`
}
