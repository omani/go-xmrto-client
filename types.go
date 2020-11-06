package xmrto

// *** RPC STRUCTS ***

// ResponseGetOrderParameters is for GetOrderParameters()
type ResponseGetOrderParameters struct {
	LowerLimit        string `json:"lower_limit"`
	Price             string `json:"price"`
	UpperLimit        string `json:"upper_limit"`
	ZeroConfEnabled   bool   `json:"zero_conf_enabled"`
	ZeroConfMaxAmount string `json:"zero_conf_max_amount"`
}

// RequestCreateOrder is for CreateOrder()
type RequestCreateOrder struct {
	Amount         float64 `json:"amount"`
	AmountCurrency string  `json:"amount_currency"`
	BTCDestAddress string  `json:"btc_dest_address"`
}

// ResponseCreateOrder is for CreateOrder()
type ResponseCreateOrder struct {
	State          string `json:"state"`
	BTCAmount      string `json:"btc_amount"`
	BTCDestAddress string `json:"btc_dest_address"`
	UUID           string `json:"uuid"`
}

// RequestGetOrderStatus is for GetOrderStatus()
type RequestGetOrderStatus struct {
	UUID string `json:"uuid"`
}

// ResponseGetOrderStatus is for GetOrderStatus()
type ResponseGetOrderStatus struct {
	State                        string `json:"state"`
	BTCAmount                    string `json:"btc_amount"`
	BTCAmountPartial             string `json:"btc_amount_partial"`
	BTCDestAddress               string `json:"btc_dest_address"`
	UUID                         string `json:"uuid"`
	BTCNumConfirmationsThreshold int64  `json:"btc_num_confirmations_threshold"`
	CreatedAT                    string `json:"created_at"`
	ExpiresAT                    string `json:"expires_at"`
	SecondsTillTimeout           int64  `json:"seconds_till_timeout"`
	BTCTransactionID             string `json:"btc_transaction_id"`
	XMRAmountTotal               string `json:"incoming_amount_total"`
	XMRAmountRemaining           string `json:"remaining_amount_incoming"`
	XMRNumConfirmationsRemaining int64  `json:"incoming_num_confirmations_remaining"`
	XMRPriceBTC                  string `json:"incoming_price_btc"`
	XMRReceivingSubAddress       string `json:"receiving_subaddress"`
	XMRRecommendedMixin          int64  `json:"recommended_mixin"`
}

// RequestGetOrderPrice is for GetOrderPrice()
type RequestGetOrderPrice struct {
	Amount         float64 `json:"amount"`
	AmountCurrency string  `json:"amount_currency"`
}

// ResponseGetOrderPrice is for GetOrderPrice()
type ResponseGetOrderPrice struct {
	BTCAmount                    string `json:"btc_amount"`
	XMRAmountTotal               string `json:"incoming_amount_total"`
	XMRNumConfirmationsRemaining int64  `json:"incoming_num_confirmations_remaining"`
	XMRPriceBTC                  string `json:"incoming_price_btc"`
}
