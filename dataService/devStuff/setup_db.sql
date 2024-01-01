CREATE DATABASE threattrack;
USE threattrack;
CREATE TABLE IF NOT EXISTS incidenttypes (
		id int PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(50)
	);

CREATE TABLE IF NOT EXISTS incidents (
		id VARCHAR(50) PRIMARY KEY,
		name VARCHAR(50),
		severity enum('Low','Medium','High', 'Critical') DEFAULT 'Low',
		status enum('Pending','Open','Active', 'Closed') DEFAULT 'Pending',
		type int DEFAULT 0,
		creationdate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (type) REFERENCES incidenttypes(id) ON DELETE SET DEFAULT
	);

CREATE TABLE IF NOT EXISTS ioctypes (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(50) DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS iocs (
  id varchar(50) NOT NULL,
  value varchar(50) NOT NULL,
  iocType int DEFAULT 0,
  verdict enum('Neutral','Benigne','Malicious') DEFAULT 'Neutral',
  PRIMARY KEY (id),
  KEY iocType (iocType),
 FOREIGN KEY (iocType) REFERENCES ioctypes (id) ON DELETE SET DEFAULT
);




CREATE TABLE IF NOT EXISTS iocsincidents (
  id int NOT NULL AUTO_INCREMENT,
  iocId varchar(50) DEFAULT NULL,
  incidentId varchar(50) DEFAULT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (iocId) REFERENCES iocs (id),
  FOREIGN KEY (incidentId) REFERENCES incidents (id)
);

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(50) PRIMARY KEY NOT NULL,
    firstName VARCHAR(255) NOT NULL,
    lastName VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_At VARCHAR(50),
    fullname VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS task (
    id varchar(255) NOT NULL PRIMARY KEY,
    Title varchar(255) NOT NULL,
    Description text NOT NULL,
    Assignee varchar(255) DEFAULT NULL,
    Incident varchar(255) DEFAULT NULL,
    Status enum('Open', 'In Progress', 'Done')  NOT NULL DEFAULT 'Open',
    Priority enum(
        'Low',
        'Medium',
        'High',
        'Critical'
    ) DEFAULT 'Low',
    FOREIGN KEY (Assignee) REFERENCES users (id),
    FOREIGN KEY (Incident) REFERENCES incidents (id)
);

CREATE TABLE IF NOT EXISTS taskcomments (
    id varchar(255) NOT NULL COMMENT 'UUID',
    create_time timestamp NULL DEFAULT NULL COMMENT 'Create Time',
    content text NOT NULL,
    writer varchar(255) NOT NULL,
    task varchar(255) NOT NULL,
    PRIMARY KEY (id),
    KEY writer (writer),
    KEY task (task),
    FOREIGN KEY (writer) REFERENCES users (id),
    FOREIGN KEY (task) REFERENCES task (id)
); 

CREATE TABLE IF NOT EXISTS worklogs (
    id varchar(50) NOT NULL,
    writerId varchar(50) NOT NULL,
    incidentId varchar(50) NOT NULL,
    content text  NOT NULL,
    created_at varchar(50) DEFAULT NULL,
    PRIMARY KEY (id),
    KEY writerId (writerId),
    KEY incidentId (incidentId),
    FOREIGN KEY (writerId) REFERENCES users (id),
    FOREIGN KEY (incidentId) REFERENCES incidents (id)
);