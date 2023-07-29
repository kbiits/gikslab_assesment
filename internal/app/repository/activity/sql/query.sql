-- name: AddActivity :exec
INSERT INTO activities (title, description, start_date, end_date, participants, skill_id) VALUES($1, $2, $3, $4, $5, $6);

-- name: UpdateActivity :exec
UPDATE activities SET (title, description, start_date, end_date, participants, skill_id) = ($1, $2, $3, $4, $5, $6) WHERE id = $7;

-- name: ListActivitiesBySkill :many
SELECT * FROM list_activities(sqlc.arg(skillId)::int8, sqlc.arg(sortBy)::varchar, sqlc.arg(orderCol)::varchar, sqlc.arg(limitVal)::int4, sqlc.arg(offsetVal)::int4);

-- name: FindActivityById :one
SELECT * FROM activities WHERE id = $1;

-- name: DeleteActivity :execrows
DELETE FROM activities WHERE id = $1;