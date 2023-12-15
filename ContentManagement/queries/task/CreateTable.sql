CREATE TABLEIF NOT EXISTS task (
        id varchar(255) NOT NULL COMMENT 'Primary Key',
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
        PRIMARY KEY (id),
        KEY Assignee (Assignee),
        KEY Incident (Incident),
        FOREIGN KEY (Assignee) REFERENCES users (id),
        FOREIGN KEY (Incident) REFERENCES incidents (id)
    ) 