CREATE TABLE deals(
	id SERIAL PRIMARY KEY,
	user_id INTEGER,
	header VARCHAR(100) NOT NULL,
	description TEXT,
	price FLOAT NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users (id)
)


CREATE TABLE sales(
	id SERIAL PRIMARY KEY,
	user_id INTEGER,
	deal_id INTEGER,
	volume 	FLOAT NOT NULL,
	date timestamp,
	FOREIGN KEY (deal_id) REFERENCES deals (id),
	FOREIGN KEY (user_id) REFERENCES users (id)
)
