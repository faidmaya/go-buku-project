CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(200) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  created_by VARCHAR(100),
  modified_at TIMESTAMP,
  modified_by VARCHAR(100)
);
