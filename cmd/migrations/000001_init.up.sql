CREATE TABLE IF NOT EXISTS fighters (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    birth_date DATE,
    weight FLOAT,
    category VARCHAR(50),
    club_id INT
);

CREATE TABLE IF NOT EXISTS clubs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    address TEXT
);

CREATE TABLE IF NOT EXISTS tournaments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    date DATE NOT NULL,
    location VARCHAR(200),
    prize_fund FLOAT
);

CREATE TABLE IF NOT EXISTS tournament_participants (
    tournament_id INT REFERENCES tournaments(id) ON DELETE CASCADE,
    fighter_id INT REFERENCES fighters(id) ON DELETE CASCADE,
    place INT,
    PRIMARY KEY (tournament_id, fighter_id)
);