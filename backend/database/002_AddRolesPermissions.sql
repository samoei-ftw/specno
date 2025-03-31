-- Add Roles to User
ALTER TABLE users ADD COLUMN role VARCHAR(50) NOT NULL DEFAULT 'user';

-- Permissions
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    type VARCHAR(255) UNIQUE NOT NULL
);

-- Map Permissions to Roles
CREATE TABLE IF NOT EXISTS rolepermissionlink (
    id SERIAL PRIMARY KEY,
    permission_id INT NOT NULL,
    role VARCHAR(255) NOT NULL
);