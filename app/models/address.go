package models

type Address struct {
	Hash160        string  `json:"hash160"`
	Address        string  `json:"address"`
	N_tx           int64   `json:"n_tx"`
	Total_received float64 `json:"total_received"`
	Total_sent     int64   `json:"total_sent"`
	Final_balance  float64 `json:"final_balance"`
}
