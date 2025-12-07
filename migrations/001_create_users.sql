CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(200) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  created_by VARCHAR(100),
  modified_at TIMESTAMP,
  modified_by VARCHAR(100)
);
