CREATE TABLE IF NOT EXISTS iocs (
  id varchar(50) NOT NULL,
  value varchar(50) NOT NULL,
  iocType int DEFAULT 0,
  verdict enum('Neutral','Benigne','Malicious') DEFAULT 'Neutral',
  PRIMARY KEY (id),
  KEY iocType (iocType),
 FOREIGN KEY (iocType) REFERENCES ioc_types (id) ON DELETE SET DEFAULT
);
