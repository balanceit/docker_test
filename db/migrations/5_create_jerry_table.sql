
-- +migrate Up
create table jerry (
	currenttime timestamp not null default current_timestamp
);

-- +migrate Down
DROP TABLE jerry;
