SELECT  
    users.id as UserID,
    users.firstname as Firstname,
    users.lastname as Lastname,
    users.email as Email,
    users.created_at as CreatedAt,
    users.fullname as Fullname,
FROM users 
ORDER BY users.created_at
LIMIT ? OFFSET ?