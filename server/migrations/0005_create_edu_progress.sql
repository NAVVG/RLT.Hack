CREATE TABLE edu_progress(
	user_id INTEGER PRIMARY KEY,
	edu_step INTEGER NOT NULL,
	is_completed BOOL,
	FOREIGN KEY (user_id) REFERENCES users (id)
)
