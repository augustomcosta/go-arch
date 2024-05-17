-- name: CreateDeliveryman :exec
INSERT INTO deliverymen (
                      id,
                      name,
                      age,
                      address_state,
                      address_city
)
VALUES ($1, $2, $3, $4,$5);

-- name: GetDeliverymen :many
SELECT * FROM deliverymen;

-- name: GetDeliveryman :one
SELECT * FROM deliverymen WHERE id = $1;

-- name: UpdateDeliveryman :exec
UPDATE deliverymen
SET
    name = $1,
    age = $2,
    address_state = $3,
    address_city = $4
WHERE id = $5;

-- name: DeleteDeliveryman :exec
DELETE FROM deliverymen WHERE id = $1;