
CREATE TABLE IF NOT EXISTS genders (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);


CREATE TABLE IF NOT EXISTS countries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);


CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    parent_id INT NULL REFERENCES categories(id) ON DELETE SET NULL,
    deleted_at TIMESTAMPTZ NULL
);


CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    fio VARCHAR(150) NOT NULL,
    birthday DATE NULL, 
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,
    gender_id INT NOT NULL REFERENCES genders(id)
);


CREATE TABLE IF NOT EXISTS films (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL, 
    duration INT NOT NULL DEFAULT 0, 
    year_of_issue INT NOT NULL CHECK (year_of_issue > 1800),
    age INT NOT NULL CHECK (age >= 0),
    link_img VARCHAR(255) NULL,
    link_kinopoisk VARCHAR(255) NULL,
    link_video VARCHAR(255) NOT NULL,
    country_id INT NULL REFERENCES countries(id), 
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE TABLE IF NOT EXISTS categories_films (
    id SERIAL PRIMARY KEY,
    film_id INT NOT NULL REFERENCES films(id) ON DELETE CASCADE,
    category_id INT NOT NULL REFERENCES categories(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS reviews (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    film_id INT NOT NULL REFERENCES films(id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    is_approved BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE TABLE IF NOT EXISTS ratings (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    film_id INT NOT NULL REFERENCES films(id) ON DELETE CASCADE,
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 10),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);


INSERT INTO genders (name) VALUES ('Male'), ('Female'), ('Other');

INSERT INTO countries (name) VALUES ('USA'), ('Russia'), ('France'), ('Germany'), ('Italy'), ('Japan');

INSERT INTO categories (name, parent_id) VALUES
('Action', NULL), ('Comedy', NULL), ('Drama', NULL), ('Sci-Fi', NULL),
('Horror', NULL), ('Romance', NULL), ('Thriller', NULL), ('Adventure', NULL),
('Animation', NULL), ('Fantasy', NULL);


INSERT INTO categories (name, parent_id) VALUES
('Superhero', (SELECT id FROM categories WHERE name = 'Action')),
('Cyberpunk', (SELECT id FROM categories WHERE name = 'Sci-Fi')),
('Romantic Comedy', (SELECT id FROM categories WHERE name = 'Romance'));

INSERT INTO users (fio, birthday, email, password, gender_id) VALUES
('John Doe', '1990-01-01', 'john.doe@example.com', 'password123', (SELECT id FROM genders WHERE name = 'Male')),
('Jane Smith', '1985-05-15', 'jane.smith@example.com', 'password456', (SELECT id FROM genders WHERE name = 'Female')),
('Иванов Иван', '1990-01-01', 'ivanov@example.com', 'hashed_pass', 1);

INSERT INTO films (name, country_id, duration, year_of_issue, age, link_img, link_kinopoisk, link_video, description) VALUES
('Inception', (SELECT id FROM countries WHERE name = 'USA'), 148, 2010, 13, 'https://example.com/inc.jpg', 'https://kinopoisk.ru/1', 'https://video.com/1', 'A thief who steals corporate secrets through the use of dream-sharing technology.'),
('The Matrix', (SELECT id FROM countries WHERE name = 'USA'), 136, 1999, 16, 'https://example.com/mat.jpg', 'https://kinopoisk.ru/2', 'https://video.com/2', 'A computer hacker learns about the true nature of his reality.'),
('Amélie', (SELECT id FROM countries WHERE name = 'France'), 122, 2001, 12, 'https://example.com/ame.jpg', 'https://kinopoisk.ru/3', 'https://video.com/3', 'A naive girl in Paris decides to help those around her.'),
('Spirited Away', (SELECT id FROM countries WHERE name = 'Japan'), 125, 2001, 10, 'https://example.com/spi.jpg', 'https://kinopoisk.ru/4', 'https://video.com/4', 'A young girl wanders into a world ruled by gods and spirits.');

INSERT INTO categories_films (film_id, category_id) VALUES
((SELECT id FROM films WHERE name = 'Inception'), (SELECT id FROM categories WHERE name = 'Sci-Fi')),
((SELECT id FROM films WHERE name = 'The Matrix'), (SELECT id FROM categories WHERE name = 'Sci-Fi')),
((SELECT id FROM films WHERE name = 'Amélie'), (SELECT id FROM categories WHERE name = 'Romantic Comedy')),
((SELECT id FROM films WHERE name = 'Spirited Away'), (SELECT id FROM categories WHERE name = 'Fantasy'));