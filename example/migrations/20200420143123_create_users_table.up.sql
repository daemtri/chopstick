CREATE TABLE IF NOT EXISTS users(
   id INT,
   user_name VARCHAR(100) NOT NULL,
   password VARCHAR(40) NOT NULL,
   PRIMARY KEY ( id )
);

DROP TABLE IF exists users;
