-- Create database
DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'pglogrepl') THEN
        CREATE DATABASE pglogrepl;
    END IF;
END $$;
-- Create user
CREATE USER pglogrepl WITH REPLICATION PASSWORD 'secret';

-- Grant access to public schema
GRANT ALL ON SCHEMA public TO pglogrepl;