
-- +goose Up
CREATE TABLE other_fancier_post (
    id int NOT NULL,
    title text,
    body text,
    created_on timestamp without time zone,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE other_fancier_post;
