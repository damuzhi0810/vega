-- +goose Up

ALTER TYPE proposal_error ADD VALUE IF NOT EXISTS 'PROPOSAL_ERROR_INVALID_POSITION_DECIMAL_PLACES';

-- +goose Down
