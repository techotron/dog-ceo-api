CREATE TABLE IF NOT EXISTS breeds(
    breed_name VARCHAR(100) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS sub_breeds(
    breed_name VARCHAR(100) NOT NULL,
    sub_breed_name VARCHAR(100) NOT NULL,
    PRIMARY KEY(breed_name, sub_breed_name)
);

CREATE TABLE IF NOT EXISTS images(
    image_id SERIAL PRIMARY KEY,
    breed_name VARCHAR(100) NOT NULL,
    sub_breed_name VARCHAR(100),
    image_url TEXT NOT NULL
);

