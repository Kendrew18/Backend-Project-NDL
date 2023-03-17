package struct_all_ndl

type Read_rekap struct {
	Ws_no           string    `json:"ws_no"`
	Date_rekap      string    `json:"date_rekap"`
	Customer_name   string    `json:"customer_name"`
	Item_name       string    `json:"item_name"`
	Order_rekap     int       `json:"order_rekap"`
	Ws_meter        float64   `json:"ws_meter"`
	Plant           string    `json:"plant"`
	Delivery_period string    `json:"delivery_period"`
	Status_rekap    string    `json:"status_rekap"`
	Comment_note    string    `json:"comment_note"`
	Meter           []float64 `json:"meter"`
	KG              []float64 `json:"kg"`
	Diff_price      []int64   `json:"diff_price"`
}

type Rekap_layer struct {
	Meter      float64 `json:"meter"`
	KG         float64 `json:"kg"`
	Diff_price int64   `json:"diff_price"`
}
