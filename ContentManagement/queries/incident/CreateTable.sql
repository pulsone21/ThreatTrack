CREATE TABLE IF NOT EXISTS incidents (
		id VARCHAR(50) PRIMARY KEY,
		name VARCHAR(50),
		severity enum('Low','Medium','High', 'Critical') DEFAULT 'Low',
		status enum('Pending','Open','Active', 'Closed') DEFAULT 'Pending',
		type int DEFAULT 0,
		FOREIGN KEY (type) REFERENCES incident_types(id) ON DELETE SET DEFAULT
	);