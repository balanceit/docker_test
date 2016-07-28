
-- +migrate Up
create table bob (
	currenttime timestamp not null default current_timestamp
);

-- +migrate Down
DROP TABLE bob;
