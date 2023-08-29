CREATE TABLE published_confessions (
    id SERIAL PRIMARY KEY,
    confession_id INT REFERENCES confessions(id),
    date_posted TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
