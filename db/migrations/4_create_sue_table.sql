
-- +migrate Up
create table sue (
	currenttime timestamp not null default current_timestamp
);

-- +migrate Down
DROP TABLE sue;
