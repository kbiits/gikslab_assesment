package dtoSkill

type ListSkillDtoResponse []SkillDtoResponse

type SkillDtoResponse struct {
	Id        int64  `json:"id"`
	SkillName string `json:"skill_name"`
}
