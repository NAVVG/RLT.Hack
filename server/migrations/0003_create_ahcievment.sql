
CREATE TABLE achievements(
	id SERIAL PRIMARY KEY, 
	header VARCHAR(100) NOT NULL, 	
	description TEXT,
	icon TEXT 
)

CREATE TABLE achievement_progress(
	ahcievment_id INTEGER,
	user_id INTEGER,
	PRIMARY KEY (ahcievment_id, user_id),
	FOREIGN KEY (user_id) REFERENCES users (id),
	FOREIGN KEY (ahcievment_id) REFERENCES ahcievments (id)
)