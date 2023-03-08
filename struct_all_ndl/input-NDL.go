package struct_all_ndl

type Input_NDL struct {
	Ws_no                  string  `json:"ws_no"`
	Tambah_data_tanggal    string  `json:"tambah_data_tanggal"`
	Customer_delivery_date string  `json:"customer_delivery_date"`
	Job_done               string  `json:"job_done"`
	Durasi                 int     `json:"durasi"`
	Analyzer_version       string  `json:"analyzer_version"`
	Order_status           string  `json:"order_status"`
	Cylider_status         string  `json:"cylider_status"`
	Gol                    string  `json:"gol"`
	Cust                   string  `json:"cust"`
	Item_name              string  `json:"item_name"`
	Model                  string  `json:"model"`
	Up                     int     `json:"up"`
	Repeat_ndl             int     `json:"repeat_ndl"`
	Toleransi              int     `json:"toleransi"`
	Order_ndl              string  `json:"order_ndl"`
	W_s_order              float64 `json:"w_s_order"`
	Width                  float64 `json:"width"`
	Lenght                 float64 `json:"lenght"`
	Gusset                 float64 `json:"gusset"`
	Prod_size              float64 `json:"prod_size"`
	W                      float64 `json:"w"`
	C                      float64 `json:"c"`
	Color                  int     `json:"color"`
	Layer                  string  `json:"layer"`
}

type Ws_no struct {
	Ws_no string `json:"ws_no"`
}
