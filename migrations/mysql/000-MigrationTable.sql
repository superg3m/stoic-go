-- StoicMigration Up
CREATE TABLE IF NOT EXISTS `Migration` (
    `ID` INT AUTO_INCREMENT NOT NULL,
    `MigrationFile` VARCHAR(512) NOT NULL,

    UNIQUE (`MigrationFile`),
    PRIMARY KEY (`ID`)
);

-- StoicMigration Down
DROP TABLE IF EXISTS `Migration`;