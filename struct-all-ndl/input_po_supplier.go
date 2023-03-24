package struct_all_ndl

type Input_PO_Supplier struct {
	Ws_no            string `json:"ws_no"`
	Layer            string `json:"layer"`
	Nama_PO_Supplier string `json:"nama_po_supplier"`
	Tanggal_PO       string `json:"tanggal_po"`
	Meter            string `json:"meter"`
	Kg               string `json:"kg"`
	Diff_Price       string `json:"diff_price"`
	Total            string `json:"total"`
	Outstanding      string `json:"outstanding"`
}

type Outstanding struct {
	Meter    float64 `json:"meter"`
	Kg       string  `json:"kg"`
	Price_kg string  `json:"price_kg"`
	Layer    string  `json:"lyr"`
}
