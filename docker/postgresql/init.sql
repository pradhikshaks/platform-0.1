CREATE TABLE emp(
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL,
  role TEXT
);

INSERT INTO emp (name,role)
  VALUES ('vibav','developer'),
  ('madhan','developer'),
  ('shanker','manager');

CREATE TABLE projects(
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name TEXT NOT NULL,
  project TEXT
);

INSERT INTO projects (name, project)
  VALUES ('vibav','platform backend'),
  ('madhan','platform frontend'),
  ('shanker','platform research');