package battle

import (
	"github.com/google/uuid"
	"pineappletooth/bestoRpg/internal/model"
)

type skillPersistence interface {
	GetSkill(skill string) (model.Skill, error)
}
type battleController struct {
	skillPersistence skillPersistence
}

func (controller battleController) processRound(battle *Battle) {
	for i := range battle.entities {
		entity := &battle.entities[i]
		for _, skill := range entity.ChosenSkills {
			battle.UseSkill(skill, entity)
		}
		entity.ChosenSkills = make([]string, 0)
	}
}

func NewBattle(controller battleController, entities []BattleEntity) *Battle {
	skills := make(map[string]*Skill)
	for i := range entities {
		for _, skill := range entities[i].Base.Skills {
			skills[skill] = controller.loadSkill(skill)
		}
	}
	return &Battle{
		Id:       uuid.NewString(),
		entities: entities,
		skills:   skills,
	}
}

func (controller battleController) loadSkill(skill string) *Skill {
	skillModel, err := controller.skillPersistence.GetSkill(skill)
	if err != nil {
		panic(err.Error())
	}
	return NewSkillFromModel(skillModel)
}

func (controller battleController) selectSkill(battle *Battle, battleId string, skills []string) {
	entity, ok := battle.getEntityById(battleId)
	if !ok {
		return
	}
	for _, skill := range skills {
		if _, ok := battle.GetSkill(skill); !ok {
			return
		}
	}
	entity.ChosenSkills = skills
	controller.onSelectSkill(battle)
}

func (controller battleController) onSelectSkill(battle *Battle) {
	for i := range battle.entities {
		entity := &battle.entities[i]
		if !entity.isDead() && len(entity.ChosenSkills) == 0 {
			return
		}
	}
	controller.onRoundStart(battle)
}

func (controller battleController) onRoundStart(battle *Battle) {
	controller.processRound(battle)
	if controller.checkEndBattle(battle) {
		controller.end(battle)
	}
}

func (controller battleController) checkEndBattle(battle *Battle) bool {
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

func (controller battleController) end(battle *Battle) {
	//
}
