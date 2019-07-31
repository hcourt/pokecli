package simplestructs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type damageRelations struct {
	DoubleDamageFrom []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"double_damage_from"`
	DoubleDamageTo []interface{} `json:"double_damage_to"`
	HalfDamageFrom []interface{} `json:"half_damage_from"`
	HalfDamageTo   []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"half_damage_to"`
	NoDamageFrom []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"no_damage_from"`
	NoDamageTo []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"no_damage_to"`
}

func typeBuilder(name string, damage damageRelations) SimpleType {
	t := TypeFromName(name)
	t.DamageRelations = damage
	return t
}

func TestSimpleType_Effect(t *testing.T) {
	tests := []struct {
		name       string
		moveType   SimpleType
		defendType SimpleType
		expected   DamageEffect
	}{
		{
			name:     "no effect",
			moveType: TypeFromName("electric"),
			defendType: typeBuilder("ground", damageRelations{
				NoDamageFrom: []struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{{Name: "electric"}},
			}),
			expected: NoDamage,
		},
		{
			name:     "half effect",
			moveType: TypeFromName("fire"),
			defendType: typeBuilder("water", damageRelations{
				HalfDamageFrom: []interface{}{
					map[string]interface{}{
						"name": "fire",
					},
				},
			}),
			expected: HalfDamage,
		},
		{
			name:       "normal effect",
			moveType:   TypeFromName("normal"),
			defendType: typeBuilder("dark", damageRelations{}),
			expected:   NormalDamage,
		},
		{
			name:     "double effect",
			moveType: TypeFromName("rock"),
			defendType: typeBuilder("flying", damageRelations{
				DoubleDamageFrom: []struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{{Name: "rock"}}}),
			expected: DoubleDamage,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			result := test.moveType.Effect(&test.defendType)
			assert.Equal(tt, test.expected, result)
		})
	}
}
