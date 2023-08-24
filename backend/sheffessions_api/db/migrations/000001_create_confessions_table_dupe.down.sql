CREATE TABLE confessions (
    id SERIAL PRIMARY KEY,
    confession_text TEXT NOT NULL,
    date_of_confession DATE NOT NULL DEFAULT CURRENT_DATE,
    source_of_confession VARCHAR(255) NOT NULL
);
