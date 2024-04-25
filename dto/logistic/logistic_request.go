package logisticdto

type CreateLogisticRequest struct {
	LogisticName    string `json:"logistic_name" form:"logistic_name"`
	Amount          int    `json:"amount" form:"amount"`
	DestinationName string `json:"destination_name" form:"destination_name"`
	OriginName      string `json:"origin_name" form:"origin_name"`
	Duration        string `json:"duration" form:"duration"`
}

type GetLogisticRequest struct {
	OriginName      string `json:"origin_name" form:"origin_name" validate:"required"`
	DestinationName string `json:"destination_name" form:"destination_name" validate:"required"`
}

type UpdateLogisticRequest struct {
	LogisticName    string `json:"logistic_name" form:"logistic_name"`
	Amount          int    `json:"amount" form:"amount"`
	DestinationName string `json:"destination_name" form:"destination_name"`
	OriginName      string `json:"origin_name" form:"origin_name"`
	Duration        string `json:"duration" form:"duration"`
}
