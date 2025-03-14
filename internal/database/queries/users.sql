-- name: GetUser :one
SELECT id, name FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = $1
WHERE id = $2
RETURNING id, name;

-- name: CreateUser :one
INSERT INTO users
(id, name)
VALUES (gen_random_uuid(), $1)
RETURNING id, name;
