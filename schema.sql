CREATE TABLE profile
(
    id varchar(10) UNIQUE PRIMARY KEY, first_name varchar(30) NOT NULL, last_name varchar(30) NULL,
    email varchar(50) UNIQUE NOT NULL, password varchar(255) NOT NULL, created_at TIMESTAMP WITH TIME ZONE, 
    updated_at TIMESTAMP WITH TIME ZONE
)