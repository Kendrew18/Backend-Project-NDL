package struct_all_ndl

type Read_NDL struct {
	Ws_no                  string    `json:"ws_no"`
	Tambah_data_tanggal    string    `json:"tambah_data_tanggal"`
	Customer_delivery_date string    `json:"customer_delivery_date"`
	Job_done               string    `json:"job_done"`
	Durasi                 int       `json:"durasi"`
	Analyzer_version       string    `json:"analyzer_version"`
	Order_status           string    `json:"order_status"`
	Cylider_status         string    `json:"cylider_status"`
	Gol                    string    `json:"gol"`
	Cust                   string    `json:"cust"`
	Item_name              string    `json:"item_name"`
	Model                  string    `json:"model"`
	Up                     int       `json:"up"`
	Repeat_ndl             int       `json:"repeat_ndl"`
	Toleransi              int       `json:"toleransi"`
	Order_ndl              []int     `json:"order_ndl"`
	W_s_order              float64   `json:"w_s_order"`
	Width                  float64   `json:"width"`
	Lenght                 float64   `json:"lenght"`
	Gusset                 float64   `json:"gusset"`
	Prod_size              float64   `json:"prod_size"`
	W                      float64   `json:"w"`
	C                      float64   `json:"c"`
	Color                  int       `json:"color"`
	Nama_layer             []string  `json:"nama_layer"`
	Layer_datail_1         []string  `json:"layer_datail1"`
	Layer_datail_2         []string  `json:"layer_datail2"`
	Layer_datail_3         []string  `json:"layer_datail3"`
	Layer_datail_4         []string  `json:"layer_datai4"`
	Width_layer            []float64 `json:"width_layer"`
	Rm_layer               []int     `json:"rm_layer"`
	Diff_layer             []float64 `json:"diff_layer"`
	Lyr_layer              []float64 `json:"lyr_layer"`
	Ink_layer              float64   `json:"ink_layer"`
	Adh_layer              []float64 `json:"adh_layer"`
	Total_layer            float64   `json:"total_layer"`
}
