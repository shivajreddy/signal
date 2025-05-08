-- Set project name variable
\set project_name 'signal'

-- Create database if not exists
SELECT format('CREATE DATABASE %I', :'project_name')
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = :'project_name');

-- Connect to the database
\c :project_name;

-- Create extension for UUID if not exists
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create users table (this will be managed by GORM, but here's the SQL equivalent)
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    last_login TIMESTAMP WITH TIME ZONE
); 
