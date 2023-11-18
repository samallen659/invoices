package invoice

import "errors"

type Address struct {
	Street   string
	City     string
	PostCode string
	Country  string
}

func NewAddress(street string, city string, postCode string, country string) (*Address, error) {
	if street == "" {
		return nil, errors.New("Street cannot be emtpy")
	}
	if city == "" {
		return nil, errors.New("City cannot be emtpy")
	}
	if postCode == "" {
		return nil, errors.New("PostCode cannot be emtpy")
	}
	if country == "" {
		return nil, errors.New("Country cannot be emtpy")
	}

	return &Address{
		Street:   street,
		City:     city,
		PostCode: postCode,
		Country:  country,
	}, nil
}
