package xmrto

// *** RPC STRUCTS ***

// ResponseGetOrderParameters is for GetOrderParameters()
type ResponseGetOrderParameters struct {
	LowerLimit        float64 `json:"lower_limit"`
	Price             float64 `json:"price"`
	UpperLimit        float64 `json:"upper_limit"`
	ZeroConfEnabled   bool    `json:"zero_conf_enabled"`
	ZeroConfMaxAmount float64 `json:"zero_conf_max_amount"`
}

// RequestCreateOrder is for CreateOrder()
type RequestCreateOrder struct {
	BTCAmount      float64 `json:"btc_amount"`
	BTCDestAddress string  `json:"btc_dest_address"`
}

// ResponseCreateOrder is for CreateOrder()
type ResponseCreateOrder struct {
	State          string  `json:"state"`
	BTCAmount      float64 `json:"btc_amount"`
	BTCDestAddress string  `json:"btc_dest_address"`
	UUID           string  `json:"uuid"`
}

// RequestGetOrderStatus is for GetOrderStatus()
type RequestGetOrderStatus struct {
	UUID string `json:"uuid"`
}

// ResponseGetOrderStatus is for GetOrderStatus()
type ResponseGetOrderStatus struct {
	State                          string  `json:"state"`
	BTCAmount                      float64 `json:"btc_amount"`
	BTCDestAddress                 string  `json:"btc_dest_address"`
	UUID                           string  `json:"uuid"`
	BTCNumConfirmations            int64   `json:"btc_num_confirmations"`
	BTCNumConfirmationsBeforePurge int64   `json:"btc_num_confirmations_before_purge"`
	BTCTransactionID               string  `json:"btc_transaction_id"`
	CreatedAT                      string  `json:"created_at"`
	ExpiresAT                      string  `json:"expires_at"`
	SecondsTillTimeout             int64   `json:"seconds_till_timeout"`
	XMRAmountTotal                 float64 `json:"xmr_amount_total"`
	XMRAmountRemaining             float64 `json:"xmr_amount_remaining"`
	XMRNumConfirmationsRemaining   int64   `json:"xmr_num_confirmations_remaining"`
	XMRPriceBTC                    float64 `json:"xmr_price_btc"`
	XMRReceivingSubAddress         string  `json:"xmr_receiving_subaddress"`
	XMRReceivingAddress            string  `json:"xmr_receiving_address"`
	XMRRecommendedMixin            int64   `json:"xmr_recommended_mixin"`
	XMRRequiredAmount              float64 `json:"xmr_required_amount"`
}

// RequestGetOrderPrice is for GetOrderPrice()
type RequestGetOrderPrice struct {
	BTCAmount float64 `json:"btc_amount"`
}

// ResponseGetOrderPrice is for GetOrderPrice()
type ResponseGetOrderPrice struct {
	BTCAmount                    float64 `json:"btc_amount"`
	XMRAmountTotal               float64 `json:"xmr_amount_total"`
	XMRNumConfirmationsRemaining int64   `json:"xmr_num_confirmations_remaining"`
	XMRPriceBTC                  float64 `json:"xmr_price_btc"`
}
