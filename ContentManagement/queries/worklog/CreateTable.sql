CREATE TABLE IF NOT EXISTS
    worklogs (
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
    ) 