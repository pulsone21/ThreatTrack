-- Active: 1702209715544@@localhost@3306@contentdb
CREATE TABLE IF NOT EXISTS iocs_incidents (
  id int NOT NULL AUTO_INCREMENT,
  iocId varchar(50) DEFAULT NULL,
  incidentId varchar(50) DEFAULT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (iocId) REFERENCES iocs (id),
  FOREIGN KEY (incidentId) REFERENCES incidents (id)
)