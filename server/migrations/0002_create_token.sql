CREATE TABLE tokens (
	user_id INTEGER PRIMARY KEY,
	refresh VARCHAR(255),
	expires_at timestamp,
	FOREIGN KEY (user_id) REFERENCES Users (id)
)

