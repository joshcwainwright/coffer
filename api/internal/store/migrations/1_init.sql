-- +goose Up
CREATE TABLE connections (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  access_token TEXT NOT NULL
);

INSERT INTO
  connections (access_token)
VALUES
  ('testtoken');

-- +goose Down
DROP TABLE connections;
