CREATE TABLE client (
    id SERIAL PRIMARY KEY,
	client_name TEXT NOT NULL,
	client_email TEXT NOT NULL
);

CREATE TABLE item (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	price MONEY NOT NULL
);

CREATE TABLE address (
	id SERIAL PRIMARY KEY,
	street TEXT NOT NULL,
	city TEXT NOT NULL,
	post_code TEXT NOT NULL,
	country TEXT NOT NULL
);

CREATE TABLE invoice (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    payment_due TIMESTAMP,
    description TEXT,
    client_id SERIAL REFERENCES client(id),
    payment_terms INT,
    status TEXT,
    total INT,
    sender_address_id SERIAL REFERENCES address(id),
    client_address_id SERIAL REFERENCES address(id)
);

CREATE TABLE invoice_item (
	invoice_id UUID REFERENCES invoice(id),
	item_id SERIAL REFERENCES item(id),
	quantity INT,
	total MONEY,
	PRIMARY KEY (invoice_id, item_id)
);

CREATE TABLE users (
	id UUID PRIMARY KEY,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	email TEXT NOT NULL
);