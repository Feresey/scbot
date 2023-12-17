CREATE TABLE corporation (
	corp_id INTEGER PRIMARY KEY,
	created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE corporation_info (
	id SERIAL PRIMARY KEY,
	corp_id INTEGER NOT NULL REFERENCES corporation(corp_id),
	name VARCHAR(255) NOT NULL,
	tag VARCHAR(255) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE player (
	user_id INTEGER PRIMARY KEY,
	comment VARCHAR(255) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE player_info (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL REFERENCES player(user_id),
	nickname VARCHAR(255) NOT NULL,
	corp_id INTEGER NULL REFERENCES corporation(corp_id),
	info JSONB NOT NULL,
	created_at TIMESTAMPTZ NOT NULL
);
