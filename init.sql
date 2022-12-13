CREATE TABLE user_profile (
	user_id SERIAL PRIMARY KEY,
	handle VARCHAR(15) NOT NULL,
	hash text NOT NULL,
	following integer NOT NULL DEFAULT 0,
	followers integer NOT NULL DEFAULT 0
);

CREATE TABLE token (
	token_id uuid PRIMARY KEY,
	value text NOT NULL,
	user_id integer REFERENCES user_profile(user_id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE post (
	post_id SERIAL PRIMARY KEY,
	user_id integer REFERENCES user_profile(user_id) ON DELETE CASCADE NOT NULL,
	content VARCHAR(280) NOT NULL,
	replies_to integer REFERENCES post(post_id) ON DELETE CASCADE,
	replies integer NOT NULL DEFAULT 0,
	likes integer NOT NULL DEFAULT 0,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE follower (
	user_id integer REFERENCES user_profile(user_id) ON DELETE CASCADE NOT NULL,
	follower_id integer REFERENCES user_profile(user_id) ON DELETE CASCADE CONSTRAINT different_user CHECK (follower_id != user_id) NOT NULL,
	followed_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (user_id, follower_id)
);

CREATE TABLE post_like (
	post_id integer REFERENCES post(post_id) ON DELETE CASCADE NOT NULL,
	user_id integer REFERENCES user_profile(user_id) ON DELETE CASCADE NOT NULL,
	liked_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (user_id, post_id)
);
