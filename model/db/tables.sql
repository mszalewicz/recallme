-- PRAGMA encoding;
PRAGMA foreign_keys = ON;

CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(18) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password CHAR(64) NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);

CREATE TABLE cards (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	userid INTEGER,
	question_text TEXT,
	reveal_text TEXT,
	FOREIGN KEY(userid) REFERENCES users(id)
);