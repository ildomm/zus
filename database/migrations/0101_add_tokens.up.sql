CREATE TABLE tokens (
    id int NOT NULL AUTO_INCREMENT,
    token varchar(255) NOT NULL,
	hash varchar(255) NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);