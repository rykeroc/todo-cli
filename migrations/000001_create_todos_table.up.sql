CREATE TABLE IF NOT EXISTS todos (
                         id INTEGER PRIMARY KEY AUTOINCREMENT,
                         displayName TEXT NOT NULL,
                         dueAt INTEGER,
                         updatedAt INTEGER NOT NULL,
                         createdAt INTEGER NOT NULL
);