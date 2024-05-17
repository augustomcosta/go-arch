package vo

type Address struct {
	State string `json:"state"`
	City  string `json:"city"`
}

func NewAddress(state string, city string) (*Address, error) {
	return &Address{
		State: state,
		City:  city,
	}, nil
}
