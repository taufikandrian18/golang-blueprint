package utility

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Language struct {
	EN string `json:"en"`
	ID string `json:"id"`
}

func GetNameOfDay(number int) (days interface{}) {
	switch number {
	case Sunday:
		days = Language{
			EN: "Sunday",
			ID: "Minggu",
		}
	case Monday:
		days = Language{
			EN: "Monday",
			ID: "Senin",
		}
	case Tuesday:
		days = Language{
			EN: "Tuesday",
			ID: "Selasa",
		}
	case Wednesday:
		days = Language{
			EN: "Wednesday",
			ID: "Rabu",
		}
	case Thursday:
		days = Language{
			EN: "Thursday",
			ID: "Kamis",
		}
	case Friday:
		days = Language{
			EN: "Friday",
			ID: "Jumat",
		}
	case Saturday:
		days = Language{
			EN: "Saturday",
			ID: "Sabtu",
		}
	}

	return
}
