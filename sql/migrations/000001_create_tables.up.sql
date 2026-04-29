CREATE TABLE IF NOT EXISTS employees(
   id  SERIAL PRIMARY KEY NOT NULL,
   name VARCHAR (50) NOT NULL,
   email VARCHAR (50) UNIQUE NOT NULL,
   created_at timestamp DEFAULT now()
);

CREATE TYPE attendance_status AS ENUM ('LATE', 'ON_TIME');

CREATE TABLE IF NOT EXISTS attendances(
    id  SERIAL PRIMARY KEY NOT NULL,
    employee_id BIGINT NOT NULL REFERENCES employees(id),
    check_in timestamp,
    check_out timestamp,
    status attendance_status NOT NULL,
    created_at timestamp DEFAULT now()
);

CREATE INDEX attendances_emp_id ON attendances(employee_id);

INSERT INTO employees(name, email) VALUES
    ('test-1', 'test1@mail.com'),
    ('test-2', 'test2@mail.com'),
    ('test-3', 'test3@mail.com'),
    ('test-4', 'test4@mail.com'),
    ('test-5', 'test5@mail.com');