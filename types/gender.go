package types

import (
	"encoding/json"
	"fmt"
)

type GenderName string

const (
	GenderMale        = "Male"
	GenderFemale      = "Female"
	GenderNonBinary   = "Non-binary/Prefer not to specify"
	GenderUnspecified = "Unspecified"
)

type Gender struct {
	ID   int
	Name GenderName
}

var gendersByID = map[int]Gender{
	0: {ID: 0, Name: GenderUnspecified},
	1: {ID: 1, Name: GenderFemale},
	2: {ID: 2, Name: GenderMale},
	3: {ID: 3, Name: GenderNonBinary},
}

// Genders is a collection of all possible gender values.
var Genders = struct {
	Unspecified Gender
	Female      Gender
	Male        Gender
	NonBinary   Gender
}{
	gendersByID[0],
	gendersByID[1],
	gendersByID[2],
	gendersByID[3],
}

func (g *Gender) UnmarshalJSON(b []byte) error {
	var id int
	if err := json.Unmarshal(b, &id); err != nil {
		return err
	}

	gender, ok := gendersByID[id]
	if !ok {
		return fmt.Errorf("invalid gender ID: %d", id)
	}
	*g = gender
	return nil
}

func (g Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.ID)
}
