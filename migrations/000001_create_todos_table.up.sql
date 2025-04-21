CREATE TABLE IF NOT EXISTS todos (
                         id INTEGER PRIMARY KEY AUTOINCREMENT,
                         displayName TEXT NOT NULL,
                         dueAt TEXT,
                         updatedAt TEXT NOT NULL,
                         createdAt TEXT NOT NULL
);