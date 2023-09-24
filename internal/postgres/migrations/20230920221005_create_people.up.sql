CREATE TABLE people (
  id SERIAL NOT NULL,
  first_name VARCHAR ( 50 ) NOT NULL,
  surname VARCHAR ( 50 ) NOT NULL,
  patronymic VARCHAR ( 50 ),
  age INT,
  gender VARCHAR ( 50 ),
  country VARCHAR ( 50 )
);