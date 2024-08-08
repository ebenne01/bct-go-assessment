CREATE SEQUENCE user_seq start with 10000;

CREATE TABLE user_status (
  user_status varchar(1) NOT NULL,
  description varchar(15) NOT NULL,
  PRIMARY KEY (user_status)
);

CREATE TABLE users (
  user_id bigint NOT NULL DEFAULT nextval('user_seq'),
  user_name varchar(50) UNIQUE NOT NULL,
  first_name varchar(255) NOT NULL,
  last_name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  user_status varchar(1) NOT NULL,
  department varchar(255),
  PRIMARY KEY (user_id),
  FOREIGN KEY (user_status) REFERENCES user_status(user_status)
  );

ALTER SEQUENCE user_seq OWNED BY users.user_id;

INSERT INTO user_status (user_status, description)
VALUES
('I', 'Inactive'),
('A', 'Active'),
('T', 'Terminated');

INSERT INTO users (user_name, first_name, last_name, email, user_status, department)
VALUES
('jdoe', 'John', 'Doe', 'jdoe@example.com', 'A', 'Sales'),
('ssmith', 'Steve', 'Smith', 'ssmith@example.com', 'T', 'Finance');
