
-- +migrate Up
create table list (currenttime timestamp not null default current_timestamp);

-- +migrate Down
DROP TABLE list;
