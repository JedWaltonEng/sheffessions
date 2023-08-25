CREATE TABLE IF NOT EXISTS articles(
   article_id serial PRIMARY KEY,
   title VARCHAR (255) NOT NULL,
   content TEXT NOT NULL
);

