-- StoicMigration Up
CREATE TABLE IF NOT EXISTS User (
    id int AUTO_INCREMENT,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    email  varchar(255) NOT NULL,
    email_confirmed TINYINT(1),
    PRIMARY KEY(id)
);

-- StoicMigration Down
DROP TABLE IF EXISTS User;