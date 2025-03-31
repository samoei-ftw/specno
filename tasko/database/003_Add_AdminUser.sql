INSERT INTO users (email, password, role)
SELECT 'admin@example.com', '******', 'admin'
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'admin@example.com');