-- name: AddSkill :exec
INSERT INTO skills (name) VALUES ($1);

-- name: FindSkillByName :one
SELECT * FROM skills s WHERE s.name = $1; 

-- name: GetAllSkills :many
SELECT * FROM skills;

-- name: FindSkillByMultipleIds :many
SELECT * FROM skills WHERE id = ANY($1::int8[]);