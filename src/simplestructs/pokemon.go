package simplestructs

import (
	"fmt"

	"github.com/mtslzr/pokeapi-go/structs"
)

// SimplePokemon is a wrapper around a Pokemon which simplifies functionality
type SimplePokemon structs.Pokemon

func (s *SimplePokemon) String() string {
	var sTypes []string
	for _, t := range s.Types {
		sTypes = append(sTypes, t.Type.Name)
	}
	return fmt.Sprintf("%s (#%d) %v", s.Name, s.ID, sTypes)
}
