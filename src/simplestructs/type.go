package simplestructs

import "github.com/mtslzr/pokeapi-go/structs"


// SimpleType is a wrapper around a Type which simplifies functionality
type SimpleType structs.Type

func typeFromName(name string) SimpleType {
	return SimpleType{Name: name}
}

func (s *SimpleType) String() string {
	return s.Name
}
