CREATE TABLE IF NOT EXISTS todos (
                         id INTEGER PRIMARY KEY AUTOINCREMENT,
                         displayName TEXT NOT NULL,
                         updatedAt INTEGER NOT NULL,
                         createdAt INTEGER NOT NULL
);