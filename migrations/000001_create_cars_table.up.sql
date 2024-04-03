-- migrations/000001_create_cars_table.up.sql
CREATE TABLE IF NOT EXISTS cars (
		id SERIAL PRIMARY KEY,
		reg_num VARCHAR(50),
		mark VARCHAR(50),
		model VARCHAR(50),
		year INT,
		owner_id INT
	);
	CREATE TABLE IF NOT EXISTS people (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50),
		surname VARCHAR(50),
		patronymic VARCHAR(50)
	);
	ALTER TABLE cars ADD CONSTRAINT owner_fk FOREIGN KEY (owner_id) REFERENCES people (id);
