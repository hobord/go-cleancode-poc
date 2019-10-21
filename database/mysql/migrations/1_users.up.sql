CREATE TABLE IF NOT EXISTS users (
  id CHAR(36),
  email VARCHAR(255),
  name VARCHAR(255),
  created_at DATETIME,
  updated_at DATETIME,
);