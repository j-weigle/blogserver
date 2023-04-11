-- UUID extension to store IDs as UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
SET TIMEZONE="America/New_York";

-- Blog Post table
CREATE TABLE IF NOT EXISTS blog_posts (
	id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
	created TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
	updated TIMESTAMP NULL,
	read_time INT NULL,
	author CHARACTER VARYING(100) NOT NULL,
	html TEXT NOT NULL
);

-- Blog Post Source table
CREATE TABLE IF NOT EXISTS blog_post_source (
	blog_post UUID REFERENCES blog_posts (id) ON DELETE CASCADE,
	source TEXT NOT NULL,
	PRIMARY KEY (blog_post)
);

-- Image Metadata table
CREATE TABLE IF NOT EXISTS blog_images (
	id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
	name CHARACTER VARYING(100) NOT NULL,
	url TEXT NOT NULL,
	size bigint NOT NULL
);
