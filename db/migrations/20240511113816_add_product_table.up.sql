CREATE TABLE product
(
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(30) NOT NULL check ( char_length(name) >= 1 AND char_length(name) <= 30),
    sku VARCHAR(30) NOT NULL check ( char_length(sku) >= 1 AND char_length(sku) <= 30),
	category VARCHAR(20) NOT NULL,
	image_url TEXT NOT NULL,
	notes VARCHAR(200) NOT NULL check ( char_length(notes) >= 1 AND char_length(notes) <= 200),
	price INT NOT NULL check (price >= 1),
	stock INT NOT NULL check (stock >= 0 AND stock <= 100000),
	location VARCHAR(200) NOT NULL,
	is_available BOOLEAN NOT NULL,
    created_at timestamptz default now()
);
