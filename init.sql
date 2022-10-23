CREATE TABLE user_profile (
	user_id SERIAL PRIMARY KEY,
	handle VARCHAR(15) NOT NULL,
	hash text NOT NULL
);

CREATE TABLE token (
	token_id uuid PRIMARY KEY,
	value text NOT NULL,
	user_id integer REFERENCES user_profile(user_id) ON DELETE CASCADE NOT NULL
);
