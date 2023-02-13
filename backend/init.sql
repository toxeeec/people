CREATE EXTENSION pg_trgm;

CREATE TABLE user_profile (
	user_id SERIAL PRIMARY KEY,
	handle VARCHAR(15) NOT NULL,
	hash TEXT NOT NULL,
	following INTEGER NOT NULL DEFAULT 0,
	followers INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX user_profile_handle_idx ON user_profile USING GIN(handle gin_trgm_ops);

CREATE TABLE token (
	token_id uuid PRIMARY KEY,
	value TEXT NOT NULL,
	user_id INTEGER REFERENCES user_profile(user_id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE post (
	post_id SERIAL PRIMARY KEY,
	user_id INTEGER REFERENCES user_profile(user_id) ON DELETE CASCADE NOT NULL,
	content VARCHAR(280) NOT NULL,
	replies_to INTEGER REFERENCES post(post_id) ON DELETE CASCADE,
	replies INTEGER NOT NULL DEFAULT 0,
	likes INTEGER NOT NULL DEFAULT 0,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	ts TSVECTOR GENERATED ALWAYS AS (TO_TSVECTOR('english', content)) STORED
);

CREATE INDEX post_ts_idx ON post USING GIN(ts);

CREATE TABLE follower (
	user_id INTEGER REFERENCES user_profile(user_id) ON DELETE CASCADE NOT NULL,
	follower_id INTEGER REFERENCES user_profile(user_id) ON DELETE CASCADE CONSTRAINT different_user CHECK (follower_id != user_id) NOT NULL,
	followed_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (user_id, follower_id)
);

CREATE TABLE post_like (
	post_id INTEGER REFERENCES post(post_id) ON DELETE CASCADE NOT NULL,
	user_id INTEGER REFERENCES user_profile(user_id) ON DELETE CASCADE NOT NULL,
	liked_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (user_id, post_id)
);

CREATE TABLE image (
	image_id SERIAL PRIMARY KEY,
	-- TODO: remove / 
	name VARCHAR(41) NOT NULL, -- / + 36 characters for UUID + 4 for file extension
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	user_id INTEGER REFERENCES user_profile(user_id) ON DELETE CASCADE NOT NULL,
	in_use boolean NOT NULL DEFAULT false
);

CREATE TABLE post_image (
	post_id INTEGER REFERENCES post(post_id) ON DELETE CASCADE NOT NULL,
	image_id INTEGER REFERENCES image(image_id) ON DELETE CASCADE NOT NULL,
	PRIMARY KEY (post_id, image_id)
);

CREATE TABLE message (
	message_id SERIAL PRIMARY KEY,
	from_id INTEGER REFERENCES user_profile(user_id) ON DELETE CASCADE NOT NULL,
	to_id INTEGER REFERENCES user_profile(user_id) ON DELETE CASCADE NOT NULL,
	content TEXT NOT NULL,
	sent_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
