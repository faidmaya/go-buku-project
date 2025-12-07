CREATE TABLE IF NOT EXISTS books (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(1000),
  image_url VARCHAR(500),
  release_year INTEGER,
  price INTEGER,
  total_page INTEGER,
  thickness VARCHAR(50),
  category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  created_by VARCHAR(100),
  modified_at TIMESTAMP,
  modified_by VARCHAR(100)
);
