
-- +migrate Up
create table show (
	currenttime timestamp not null default current_timestamp
);

-- +migrate Down
DROP TABLE show;
