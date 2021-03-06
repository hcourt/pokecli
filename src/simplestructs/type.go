package simplestructs

import (
	"reflect"

	"github.com/mtslzr/pokeapi-go/structs"
)

const (
	NoDamage      = DamageEffect(0.00)
	QuarterDamage = DamageEffect(0.25)
	HalfDamage    = DamageEffect(0.50)
	NormalDamage  = DamageEffect(1.00)
	DoubleDamage  = DamageEffect(2.00)
	QuadDamage    = DamageEffect(4.00)
)

// DamageEffect is a multiplier for one type's damage against another type
type DamageEffect float32

// SimpleType is a wrapper around a Type which simplifies functionality
type SimpleType structs.Type

func (d DamageEffect) String() string {
	switch d {
	case NoDamage:
		return "not effective"
	case QuarterDamage:
		return "not very effective (25%)"
	case HalfDamage:
		return "not very effective (50%)"
	case NormalDamage:
		return "effective"
	case DoubleDamage:
		return "super effective"
	case QuadDamage:
		return "double super effective"
	default:
		return "unknown!"
	}
}

func (s *SimpleType) String() string {
	return s.Name
}

// Effect calculates the DamageEffect multiplier that this type produces against
// the defending type.
func (s *SimpleType) Effect(defendType *SimpleType) DamageEffect {
	defender := defendType.DamageRelations
	for _, t := range defender.NoDamageFrom {
		if t.Name == s.Name {
			return NoDamage
		}
	}
	for _, t := range defender.HalfDamageFrom {
		// this particular api interface is not a struct for some reason
		concreteT := reflect.ValueOf(t).Interface().(map[string]interface{})
		if concreteT["name"] == s.Name {
			return HalfDamage
		}
	}
	for _, t := range defender.DoubleDamageFrom {
		if t.Name == s.Name {
			return DoubleDamage
		}
	}
	return NormalDamage
}

// EffectMulti calculates the DamageEffect multiplier that this type produces
// against a set of defending types.
func (s *SimpleType) EffectMulti(defendTypes []*SimpleType) DamageEffect {
	result := NormalDamage
	for _, defendType := range defendTypes {
		result *= s.Effect(defendType)
	}
	return result
}
