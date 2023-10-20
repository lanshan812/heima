package patent

type GetCarScheduleListResp struct {
	Total     int        `json:"total"`
	Schedules []Schedule `json:"schedules"`
	CarList   []Car      `json:"carList"`
}
type Schedule struct {
	Route       string `json:"route"`
	Finish      string `json:"finish"`
	ErrorReason string `json:"errorReason"`
	CarList     []Car  `json:"carList"`
}

type Car struct {
	Route  string `json:"route"`
	Type   string `json:"type"`
	CarId  string `json:"carId"`
	Status string `json:"status"`
}
