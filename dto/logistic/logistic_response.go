package logisticdto

type LogisticResponse struct {
	ID              int    `json:"id"`
	LogisticName    string `json:"logistic_name" form:"logistic_name"`
	Amount          int    `json:"amount" form:"amount"`
	DestinationName string `json:"destination_name" form:"destination_name"`
	OriginName      string `json:"origin_name" form:"origin_name"`
	Duration        string `json:"duration" form:"duration"`
}
