package simplestructs

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// SimpleMove is a wrapper around a Move which simplifies functionality
type SimpleMove structs.Move

func (s *SimpleMove) String() string {
	return fmt.Sprintf("%s (class: %s, type: %v, power: %d, accuracy: %d)",
		s.Name, s.DamageClass.Name, s.Type.Name, s.Power, s.Accuracy,
	)
}
