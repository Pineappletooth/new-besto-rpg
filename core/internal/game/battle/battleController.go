package battle

import (
	"fmt"
	"github.com/google/uuid"
	"pineappletooth/bestoRpg/internal/model"
)

var BattlesStorage = make(map[string]*Battle)

type SkillPersistence interface {
	GetSkill(skill string) (model.Skill, error)
}
type Persistence interface {
	SaveBattle(battle *Battle) error
	GetBattle(battleId string) (*Battle, error)
}
type Controller struct {
	SkillPersistence  SkillPersistence
	BattlePersistence Persistence
}

func NewController(skillPersistence SkillPersistence, battlePersistence Persistence) Controller {
	return Controller{
		SkillPersistence:  skillPersistence,
		BattlePersistence: battlePersistence,
	}
}
func (controller Controller) processRound(battle *Battle) {
	for i := range battle.entities {
		entity := &battle.entities[i]
		for _, skill := range entity.ChosenSkills {
			battle.UseSkill(skill, entity)
		}
		entity.ChosenSkills = make([]string, 0)
	}
}

func NewBattle(controller Controller, entities []BattleEntity) (*Battle, error) {
	skills := make(map[string]*Skill)
	for i := range entities {
		for _, skill := range entities[i].Base.Skills {
			loadSkill, err := controller.loadSkill(skill)
			if err != nil {
				return nil, fmt.Errorf("battle could not load skills: %w", err)
			}
			skills[skill] = loadSkill
		}
	}

	battle := &Battle{
		Id:       uuid.NewString(),
		entities: entities,
		skills:   skills,
	}
	err := controller.BattlePersistence.SaveBattle(battle)
	if err != nil {
		return nil, err
	}

	return battle, nil
}

func (controller Controller) loadSkill(skill string) (*Skill, error) {
	skillModel, err := controller.SkillPersistence.GetSkill(skill)
	if err != nil {
		return nil, fmt.Errorf("skill %s not found: %w", skill, err)
	}
	return NewSkillFromModel(skillModel), nil
}

func (controller Controller) SelectSkill(battle *Battle, entityId string, skills []string) error {
	entity, ok := battle.getEntityById(entityId)
	if !ok {
		return fmt.Errorf("battle entity not found")
	}
	for _, skill := range skills {
		if _, ok := battle.GetSkill(skill); !ok {
			return fmt.Errorf("skill not found")
		}
	}
	entity.ChosenSkills = skills
	err := controller.onSelectSkill(battle)
	if err != nil {
		return err
	}
	return nil
}

func (controller Controller) onSelectSkill(battle *Battle) error {
	for i := range battle.entities {
		entity := &battle.entities[i]
		if !entity.isDead() && len(entity.ChosenSkills) == 0 {
			return nil
		}
	}
	err := controller.BattlePersistence.SaveBattle(battle)
	if err != nil {
		return err
	}
	controller.onRoundStart(battle)
	return nil
}

func (controller Controller) onRoundStart(battle *Battle) {
	controller.processRound(battle)
	if controller.checkEndBattle(battle) {
		controller.end(battle)
	}
}

func (controller Controller) checkEndBattle(battle *Battle) bool {
	teams := make(map[int]bool)
	deadTeams := make(map[int]bool)
	for i := range battle.entities {
		teams[battle.entities[i].Team] = true
		if battle.entities[i].isDead() {
			deadTeams[battle.entities[i].Team] = true
		}
	}

	return len(teams)-len(deadTeams) <= 1
}

func (controller Controller) end(battle *Battle) {
	//
}
