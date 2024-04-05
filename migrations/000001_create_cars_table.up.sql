-- migrations/000001_create_cars_table.up.sql
    CREATE TABLE IF NOT EXISTS cars (
		id SERIAL PRIMARY KEY,
		regNum VARCHAR(50),
		mark VARCHAR(50),
		model VARCHAR(50),
		year INT
	);
	CREATE TABLE IF NOT EXISTS people (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50),
		surname VARCHAR(50),
		patronymic VARCHAR(50),
		car_id INT,
        CONSTRAINT fk_car_id FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE
	);

--INSERT INTO cars (regNum, mark, model, year) VALUES ('Х001АС168', 'Lada', 'Vesta', 2019);
--INSERT INTO people (name, surname, patronymic, car_id) VALUES ('Иван', 'Иванов', 'Иванович', 1);
--
--INSERT INTO cars (regNum, mark, model, year) VALUES ('Х001АС169', 'Lada', 'Priora', 2019);
--INSERT INTO people (name, surname, patronymic, car_id) VALUES ('Петр', 'Иванов', 'Иванович', 2);
--
--INSERT INTO cars (regNum, mark, model, year) VALUES ('Х001АС170', 'Lada', 'Vesta', 2020);
--INSERT INTO people (name, surname, patronymic, car_id) VALUES ('Иван', 'Петров', 'Иванович', 3);
--
--INSERT INTO cars (regNum, mark, model, year) VALUES ('Х001АС171', 'Lada', 'Vesta', 2018);
--INSERT INTO people (name, surname, patronymic, car_id) VALUES ('Иван', 'Иванов', 'Петрович', 4);
--
--INSERT INTO cars (regNum, mark, model, year) VALUES ('Х001АС172', 'Lada', 'Priora', 2011);
--INSERT INTO people (name, surname, patronymic, car_id) VALUES ('Иван', 'Иванов', 'Петри', 5);
--
--INSERT INTO cars (regNum, mark, model, year) VALUES ('Х001АС173', 'Lada', 'Vesta', 2019);
--INSERT INTO people (name, surname, patronymic, car_id) VALUES ('Олег', 'Иванов', 'Иванович', 6);
--
--INSERT INTO cars (regNum, mark, model, year) VALUES ('Х001АС174', 'Lada', 'Xray', 2019);
--INSERT INTO people (name, surname, patronymic, car_id) VALUES ('Иван', 'Олегов', 'Иванович', 7);
--
--INSERT INTO cars (regNum, mark, model, year) VALUES ('Х001АС175', 'Kia', 'Spectra', 2011);
--INSERT INTO people (name, surname, patronymic, car_id) VALUES ('Владимир', 'Иванов', 'Иванович', 8);
--
--INSERT INTO cars (regNum, mark, model, year) VALUES ('Х001АС176', 'Lada', 'Vesta', 2019);
--INSERT INTO people (name, surname, patronymic, car_id) VALUES ('Иван', 'Владимир', 'Иванович', 9);
--
--INSERT INTO cars (regNum, mark, model, year) VALUES ('Х001АС177', 'Kia', 'Rio', 2019);
--INSERT INTO people (name, surname, patronymic, car_id) VALUES ('Дмитрий', 'Иванов', 'Иванович', 10);
--
--INSERT INTO cars (regNum, mark, model, year) VALUES ('Х001АС178', 'Kia', 'Rio', 2019);
--INSERT INTO people (name, surname, patronymic, car_id) VALUES ('Иван', 'Иванов', 'Иванович', 11);
--
--INSERT INTO cars (regNum, mark, model, year) VALUES ('Х001АС178', 'Volvo', 'XC90', 2019);
--INSERT INTO people (name, surname, patronymic, car_id) VALUES ('Олег', 'Петров', 'Иванович', 12);
