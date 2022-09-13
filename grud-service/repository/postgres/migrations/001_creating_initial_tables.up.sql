CREATE TABLE IF NOT EXISTS posts (
    id INT PRIMARY KEY,
    user_id INT,
    title TEXT,
    body TEXT
);