-- Creating tables
CREATE TABLE IF NOT EXISTS candidates (
    id BIGSERIAL PRIMARY KEY, 
	name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    city_id INTEGER REFERENCES cities(id),
    experience INTEGER,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS cities(
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
    city_timezone VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);


CREATE TABLE IF NOT EXISTS hiring_managers (
	id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    company VARCHAR(50) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS vacancies (
	id BIGSERIAL PRIMARY KEY,
    hiring_manager_id INTEGER REFERENCES hiring_managers(id),
    position VARCHAR(50) NOT NULL,
    description TEXT,
    salary INT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);


CREATE TABLE IF NOT EXISTS applications (
    id BIGSERIAL PRIMARY KEY,
    candidate_id INTEGER REFERENCES candidates(id),
    vacancy_id INTEGER REFERENCES vacancies(id),
    status_id INTEGER REFERENCES statuses(id),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);


CREATE TABLE IF NOT EXISTS statuses(
    id BIGSERIAL PRIMARY KEY,
	name VARCHAR(100) UNIQUE NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- Сreating indexes

-- If the database will receive frequent queries like this:
-- "SELECT * FROM vacancies WHERE hiring_manager_id = <some id>;"
-- then creating this index will speed up the queries.
CREATE INDEX vacancies_hiring_manager_id_index ON vacancies (hiring_manager_id);