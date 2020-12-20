package godestone

import (
	"encoding/json"

	"github.com/karashiiro/godestone/pack"
)

type profileSelectors struct {
	Achievements map[string](interface{})
	Attributes   map[string](interface{})
	Character    map[string](interface{})
	ClassJob     map[string](interface{})
	Gearset      map[string](interface{})
	Minion       map[string](interface{})
	Mount        map[string](interface{})
}

func loadProfileSelectors() (*profileSelectors, error) {
	achievementsBytes, err := pack.Asset("profile/achievements.json")
	if err != nil {
		return nil, err
	}
	achievements := make(map[string](interface{}))
	json.Unmarshal(achievementsBytes, &achievements)

	attributesBytes, err := pack.Asset("profile/attributes.json")
	if err != nil {
		return nil, err
	}
	attributes := make(map[string](interface{}))
	json.Unmarshal(attributesBytes, &attributes)

	characterBytes, err := pack.Asset("profile/character.json")
	if err != nil {
		return nil, err
	}
	character := make(map[string](interface{}))
	json.Unmarshal(characterBytes, &character)

	classJobBytes, err := pack.Asset("profile/classjob.json")
	if err != nil {
		return nil, err
	}
	classJob := make(map[string](interface{}))
	json.Unmarshal(classJobBytes, &classJob)

	gearsetBytes, err := pack.Asset("profile/gearset.json")
	if err != nil {
		return nil, err
	}
	gearset := make(map[string](interface{}))
	json.Unmarshal(gearsetBytes, &gearset)

	minionBytes, err := pack.Asset("profile/minion.json")
	if err != nil {
		return nil, err
	}
	minion := make(map[string](interface{}))
	json.Unmarshal(minionBytes, &minion)

	mountBytes, err := pack.Asset("profile/mount.json")
	if err != nil {
		return nil, err
	}
	mount := make(map[string](interface{}))
	json.Unmarshal(mountBytes, &mount)

	return &profileSelectors{
		Achievements: achievements,
		Attributes:   attributes,
		Character:    character,
		ClassJob:     classJob,
		Gearset:      gearset,
		Minion:       minion,
		Mount:        mount,
	}, nil
}
