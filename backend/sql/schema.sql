-- sql/schema.sql
CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255)
    -- Other fields as needed
);

CREATE TABLE IF NOT EXISTS Questions (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Answers (
    id SERIAL PRIMARY KEY,
    question_id INTEGER REFERENCES Questions(id) ON DELETE CASCADE,
    text TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS DailyInformation (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES Users(id) ON DELETE CASCADE,
    timestamp TIMESTAMP NOT NULL,
    question_id INTEGER REFERENCES Questions(id) ON DELETE CASCADE,
    answer_id INTEGER REFERENCES Answers(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS YearStatus (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES Users(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    color VARCHAR(20) NOT NULL
);



