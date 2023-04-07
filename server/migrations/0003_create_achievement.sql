
CREATE TABLE achievements(
	id SERIAL PRIMARY KEY, 
	header VARCHAR(100) NOT NULL, 	
	description TEXT,
	icon TEXT 
)

CREATE TABLE achievement_progress(
	achievement_id INTEGER,
	user_id INTEGER,
	current INTEGER,
	total INTEGER,
	PRIMARY KEY (achievement_id, user_id),
	FOREIGN KEY (user_id) REFERENCES users (id),
	FOREIGN KEY (achievement_id) REFERENCES achievements (id)
)