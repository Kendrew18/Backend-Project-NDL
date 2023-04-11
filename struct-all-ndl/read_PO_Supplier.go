package struct_all_ndl

type Read_PO_supplier_str struct {
	Ws_no            string `json:"ws_no"`
	Layer            string `json:"layer"`
	Nama_po_supplier string `json:"nama_po_supplier"`
	Tanggal_PO       string `json:"tanggal_po"`
	Meter            string `json:"meter"`
	Kg               string `json:"kg"`
	Diff_Price       string `json:"diff_price"`
	Total            string `json:"total"`
	Outstanding      string `json:"outstanding"`
}

type Read_PO_supplier_fix struct {
	Ws_no             string    `json:"ws_no"`
	Nama_po_supplier  []string  `json:"nama_po_supplier"`
	Tanggal_PO        []string  `json:"tanggal_po"`
	Meter             []float64 `json:"meter"`
	Kg                []float64 `json:"kg"`
	Diff_Price        []int64   `json:"diff_price"`
	Total_kg          float64   `json:"total_kg"`
	Total_meter       float64   `json:"total_meter"`
	Total_price       int64     `json:"total_price"`
	Outstanding_kg    float64   `json:"outstanding_kg"`
	Outstanding_meter float64   `json:"outstanding_meter"`
	Outstanding_price int64     `json:"outstanding_price"`
}

type Layer_PO struct {
	Lyr_PO []string `json:"lyr_po"`
}

type Layer_PO_Str struct {
	Lyr_PO string `json:"lyr_po"`
}
