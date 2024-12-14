CREATE TABLE todo_items (
    id UUID PRIMARY KEY,
    description TEXT NOT NULL,
    due_date TIMESTAMP NOT NULL,
    file_id TEXT
);

