package vo

type DataTableVO struct {
	Draw            int         `json:"draw"`
	Data            interface{} `json:"data"`
	RecordsFiltered int         `json:"recordsFiltered"`
	RecordsTotal    int         `json:"recordsTotal"`
}

func DataTable(draw int, data interface{}, total int) DataTableVO {
	return DataTableVO{Draw: draw, Data: data, RecordsFiltered: total, RecordsTotal: total}
}
