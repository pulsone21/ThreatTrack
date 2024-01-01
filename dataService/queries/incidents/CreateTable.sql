CREATE TABLE IF NOT EXISTS incidents (
		id VARCHAR(50) PRIMARY KEY,
		name VARCHAR(50),
		severity enum('Low','Medium','High', 'Critical') DEFAULT 'Low',
		status enum('Pending','Open','Active', 'Closed') DEFAULT 'Pending',
		type int DEFAULT 0,
		creationdate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (type) REFERENCES incidenttypes(id) ON DELETE SET DEFAULT
	);