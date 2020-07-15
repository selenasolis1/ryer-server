package shared

import "fmt"

type Params []string

// Valid for this struct is arbitrary to demo an example error
func (p Params) Valid() error {
	for _, v := range p {
		if v == "" {
			return fmt.Errorf("param must be a non-empty string")
		}
	}
	return nil
}

type Str string
