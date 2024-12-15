-- StoicMigration Up
CREATE TABLE User (
    id int AUTO_INCREMENT,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    email  varchar(255) NOT NULL,
    email_confirmed TINYINT(1),
    PRIMARY KEY(id)
);

-- StoicMigration Down
DROP TABLE User;