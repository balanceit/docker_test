
-- +migrate Up
create table c_mingus (
	currenttime timestamp not null default current_timestamp
);

-- +migrate Down
DROP TABLE c_mingus;
