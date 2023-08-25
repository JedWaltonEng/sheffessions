CREATE TABLE IF NOT EXISTS comments(
   comment_id serial PRIMARY KEY,
   article_id INT REFERENCES articles(article_id),
   comment TEXT NOT NULL
);

