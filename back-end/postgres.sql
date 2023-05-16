CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    login VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    auth BOOLEAN NOT NULL DEFAULT FALSE,
    date_auth timestamp with time zone,
    date_out timestamp with time zone
);

CREATE TABLE IF NOT EXISTS variants (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS tests (
    id serial PRIMARY KEY,
    user_id INT REFERENCES users(id),
    variant_id INT REFERENCES variants(id),
    date timestamp with time zone
);

CREATE TABLE IF NOT EXISTS answers (
    id serial PRIMARY KEY,
    test_id INT REFERENCES tests(id),
    user_id INT REFERENCES users(id),
    answer VARCHAR(255) NOT NULL,
    correct_answer VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks (
    id serial PRIMARY KEY,
    count INT NOT NULL,
    variant_id INT REFERENCES variants(id),
    task VARCHAR(255) NOT NULL,
    correct_answer VARCHAR(255) NOT NULL,
    answer_1 VARCHAR(255) NOT NULL,
    answer_2 VARCHAR(255) NOT NULL,
    answer_3 VARCHAR(255) NOT NULL,
    answer_4 VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS results (
    id serial PRIMARY KEY,
    test_id INT REFERENCES tests(id),
    user_id INT REFERENCES users(id),
    percent INT NOT NULL
);


---user

INSERT INTO users VALUES (1, 'admin', 'admin', FALSE, NULL, NULL) ON CONFLICT DO NOTHING;


---variants

INSERT INTO variants (id, name) VALUES (1, 'Билет 1') ON CONFLICT DO NOTHING;
INSERT INTO variants (id, name) VALUES (2, 'Билет 2') ON CONFLICT DO NOTHING;


---answers

-- INSERT INTO answers (answer) 
-- VALUES (
--     '2'
-- ) ON CONFLICT DO NOTHING;

-- INSERT INTO answers (answer) 
-- VALUES (
--     '4'
-- ) ON CONFLICT DO NOTHING;

-- INSERT INTO answers (answer) 
-- VALUES (
--     '6'
-- ) ON CONFLICT DO NOTHING;

-- INSERT INTO answers (answer) 
-- VALUES (
--     '8'
-- ) ON CONFLICT DO NOTHING;

-- INSERT INTO answers (answer) 
-- VALUES (
--     '10'
-- ) ON CONFLICT DO NOTHING;

-- INSERT INTO answers (answer) 
-- VALUES (
--     '12'
-- ) ON CONFLICT DO NOTHING;


---tasks

INSERT INTO tasks (
    id, count, variant_id, task, correct_answer, answer_1, answer_2, answer_3, answer_4) 
VALUES (
    1, 
    1,
    (SELECT id from variants WHERE id=1), 
    'Сколько будет 1+1?',
    '2',
    '1',
    '2',
    '3',
    '4'
) ON CONFLICT DO NOTHING;

INSERT INTO tasks (
    id, count, variant_id, task, correct_answer, answer_1, answer_2, answer_3, answer_4) 
VALUES (
    2, 
    2,
    (SELECT id from variants WHERE id=1), 
    'Сколько будет 2+2?', 
    '4',
    '8',
    '4',
    '1',
    '6'
) ON CONFLICT DO NOTHING;

INSERT INTO tasks (
    id, count, variant_id, task, correct_answer, answer_1, answer_2, answer_3, answer_4) 
VALUES (
    3, 
    3,
    (SELECT id from variants WHERE id=1), 
    'Сколько будет 3+3?', 
    '6',
    '9',
    '6',
    '12',
    '8'
) ON CONFLICT DO NOTHING;

INSERT INTO tasks (
    id, count, variant_id, task, correct_answer, answer_1, answer_2, answer_3, answer_4) 
VALUES (
    4, 
    1,
    (SELECT id from variants WHERE id=2), 
    'Сколько будет 4+4?',
    '8',
    '16',
    '12',
    '10',
    '8'
) ON CONFLICT DO NOTHING;

INSERT INTO tasks (
    id, count, variant_id, task, correct_answer, answer_1, answer_2, answer_3, answer_4) 
VALUES (
    5, 
    2,
    (SELECT id from variants WHERE id=2), 
    'Сколько будет 5+5?', 
    '10',
    '25',
    '12',
    '10',
    '15'
) ON CONFLICT DO NOTHING;

INSERT INTO tasks (
    id, count, variant_id, task, correct_answer, answer_1, answer_2, answer_3, answer_4) 
VALUES (
    6, 
    3,
    (SELECT id from variants WHERE id=2), 
    'Сколько будет 6+6?', 
    '12',
    '10',
    '12',
    '14',
    '36'
) ON CONFLICT DO NOTHING;


---tests

-- INSERT INTO tests (user_id, variant_id, date) 
-- VALUES (
--     (SELECT id from users WHERE id=1), 
--     (SELECT id from variants WHERE id=1), 
--     current_timestamp
-- ) ON CONFLICT DO NOTHING;


---results

-- INSERT INTO results (test_id, percent) 
-- VALUES (
--     (SELECT id from tests WHERE id=1), 
--     99
-- ) ON CONFLICT DO NOTHING;