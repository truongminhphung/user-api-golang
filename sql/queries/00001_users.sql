-- sql/queries/users.sql

-- name: GetUsers :many
SELECT 
    id, first_name, last_name, email, age, phone, status, created_at, updated_at, deleted_at
FROM users
WHERE  deleted_at IS NULL
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: CreateUser :execresult
INSERT INTO users (
    first_name, last_name, email, password, age, phone, status
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
);

-- name: GetUserByID :one
SELECT 
    id, first_name, last_name, email, age, phone, status, created_at, updated_at, deleted_at
FROM users
WHERE id = ? AND deleted_at IS NULL; 

-- name: GetUserByEmail :one
SELECT 
    id, first_name, last_name, email, age, phone, status, created_at, updated_at, deleted_at
FROM users
WHERE email = ? AND deleted_at IS NULL;

-- name: GetUserByStatus :many
SELECT 
    id, first_name, last_name, email, age, phone, status, created_at, updated_at, deleted_at
FROM users
WHERE status = ? AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: CountUsers :one
SELECT COUNT(*) FROM users WHERE deleted_at IS NULL;

-- name: UpdateUser :execresult
UPDATE users
SET 
    first_name = ?, last_name = ?, email = ?, age = ?, phone = ?, status = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND deleted_at IS NULL;

-- name: UpdateUserPartial :exec
UPDATE users
SET
    first_name = COALESCE(?, first_name),
    last_name = COALESCE(?, last_name),
    email = COALESCE(?, email),
    age = COALESCE(?, age),
    phone = COALESCE(?, phone),
    status = COALESCE(?, status),
    updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND deleted_at IS NULL;

-- name: SoftDeleteUser :execresult
UPDATE users
SET
    deleted_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND deleted_at IS NULL;

-- name: RestoreUser :execresult
UPDATE users
SET
    deleted_at = NULL,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteUser :execresult
DELETE FROM users
WHERE id = ? AND deleted_at IS NULL;

-- name: CheckEmailExists :one
SELECT  COUNT(*) FROM users
WHERE email = ? AND deleted_at IS NULL;

-- name: CheckEmailExistsExceptID :one
SELECT COUNT(*) FROM users
WHERE email = ? AND id != ? AND deleted_at IS NULL;

-- name: GetUserWithPasswordByEmail :one
SELECT 
    id, first_name, last_name, email, password, age, phone, status,
    created_at, updated_at, deleted_at
FROM users
WHERE email = ? AND deleted_at IS NULL;