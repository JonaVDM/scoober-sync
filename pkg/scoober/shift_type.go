package scoober

// Shift type is a object of a shift. I made this based of
// one of my shifts. This will most likely not going kept
// updated when stuff changes in the future. Unless it gets
// broken somehow, and I still work here
type Shift struct {
	ID               string  `json:"_id"`
	Region           string  `json:"region"`
	FromHour         int     `json:"fromHour"`
	FromMinute       int     `json:"fromMinute"`
	ToHour           int     `json:"toHour"`
	ToMinute         int     `json:"toMinute"`
	FromTimeExtended float32 `json:"fromTimeExtended"`
	ToTimeExtended   float32 `json:"toTimeExtended"`
	Week             int     `json:"week"`
	Date             string  `json:"date"`
	FromWithTimeZone string  `json:"fromWithTimeZone"`
	ToWithTimeZone   string  `json:"toWithTimeZone"`
	From             string  `json:"from"`
	To               string  `json:"to"`
	FromUnixOffset   int     `json:"fromUnixOffset"`
	ToUnixOffset     int     `json:"toUnixOffset"`
	Absence          bool    `json:"absence"`
	AbsenceReason    string  `json:"absenceReason"`
	Published        bool    `json:"published"`
	Type             string  `json:"string"`
	CreatedAt        string  `json:"createdAt"`
	CreatedBy        string  `json:"createdBy"`
	UpdatedAt        string  `json:"updatedAt"`
	SubType          string  `json:"subType"`
}
