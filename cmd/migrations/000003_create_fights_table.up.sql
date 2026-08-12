
CREATE TABLE IF NOT EXISTS fights (
    id SERIAL PRIMARY KEY,
    fighter1_id INT REFERENCES fighters(id) ON DELETE SET NULL,
    fighter2_id INT REFERENCES fighters(id) ON DELETE SET NULL,
    result INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS fight_rounds (
    id SERIAL PRIMARY KEY,
    fight_id INT REFERENCES fights(id) ON DELETE CASCADE, 
    round_number INT NOT NULL,
    fighter1_score INT NOT NULL,
    fighter2_score INT NOT NULL
);