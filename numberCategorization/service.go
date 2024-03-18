package numberCategorization

import "errors"

type Categorization string

type Service struct {
	numbers map[int]Categorization
}

func NewService() Service {
	return Service{
		numbers: map[int]Categorization{},
	}
}

func (s Service) AddNumber(value int) {
	category := ""
	if value%3 == 0 {
		category = "Type1"
	}
	if value%5 == 0 {
		category = "Type2"
	}
	if value%3 == 0 && value%5 == 0 {
		category = "Type3"
	}
	s.numbers[value] = Categorization(category)
}

func (s Service) GetNumber(v int) (Categorization, error) {
	if value, exists := s.numbers[v]; exists {
		return value, nil
	}
	return "", errors.New("number not found")
}

func (s Service) GetAllNumbers() map[int]Categorization {
	return s.numbers
}
