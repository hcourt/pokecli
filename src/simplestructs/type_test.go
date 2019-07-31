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

func typeBuilder(name string, damage damageRelations) *SimpleType {
	t := SimpleType{Name: name}
	t.DamageRelations = damage
	return &t
}

func TestSimpleType_Effect(t *testing.T) {
	tests := []struct {
		name       string
		moveType   *SimpleType
		defendType *SimpleType
		expected   DamageEffect
	}{
		{
			name:     "no effect",
			moveType: typeBuilder("electric", damageRelations{}),
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
			moveType: typeBuilder("fire", damageRelations{}),
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
			moveType:   typeBuilder("normal", damageRelations{}),
			defendType: typeBuilder("dark", damageRelations{}),
			expected:   NormalDamage,
		},
		{
			name:     "double effect",
			moveType: typeBuilder("rock", damageRelations{}),
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
			result := test.moveType.Effect(test.defendType)
			assert.Equal(tt, test.expected, result)
		})
	}
}

func TestSimpleType_EffectMulti(t *testing.T) {
	tests := []struct {
		name        string
		moveType    *SimpleType
		defendTypes []*SimpleType
		expected    DamageEffect
	}{
		{
			name:     "quarter effect",
			moveType: typeBuilder("grass", damageRelations{}),
			defendTypes: []*SimpleType{
				typeBuilder("dragon", damageRelations{
					HalfDamageFrom: []interface{}{
						map[string]interface{}{
							"name": "grass",
						},
					},
				}),
				typeBuilder("steel", damageRelations{
					HalfDamageFrom: []interface{}{
						map[string]interface{}{
							"name": "grass",
						},
					},
				}),
			},
			expected: QuarterDamage,
		},
		{
			name:     "quadruple effect",
			moveType: typeBuilder("rock", damageRelations{}),
			defendTypes: []*SimpleType{
				typeBuilder("flying", damageRelations{
					DoubleDamageFrom: []struct {
						Name string `json:"name"`
						URL  string `json:"url"`
					}{{Name: "rock"}},
				}),
				typeBuilder("fire", damageRelations{
					DoubleDamageFrom: []struct {
						Name string `json:"name"`
						URL  string `json:"url"`
					}{{Name: "rock"}},
				}),
			},
			expected: QuadDamage,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			result := test.moveType.EffectMulti(test.defendTypes)
			assert.Equal(tt, test.expected, result)
		})
	}
}
