CREATE SCHEMA bot;

CREATE TABLE bot.corporations (
	corp_id INTEGER PRIMARY KEY,
	created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE bot.corporation_infos (
	id SERIAL PRIMARY KEY,
	corp_id INTEGER NOT NULL REFERENCES bot.corporations(corp_id),
	name VARCHAR(255) NOT NULL,
	tag VARCHAR(255) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE bot.players (
	user_id INTEGER PRIMARY KEY,
	comment VARCHAR(255) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE bot.player_infos (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL REFERENCES bot.players(user_id),
	nickname VARCHAR(255) NOT NULL,
	corp_id INTEGER NULL REFERENCES bot.corporations(corp_id),
	info JSONB NOT NULL,
	created_at TIMESTAMPTZ NOT NULL
);
