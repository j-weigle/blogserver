-- UUID extension to store IDs as UUIDs
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- User table
CREATE TABLE IF NOT EXISTS blog_users (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	name CHARACTER VARYING(100) NOT NULL,
	username CHARACTER VARYING(100) NOT NULL,
	password TEXT NOT NULL,
	admin BOOLEAN DEFAULT FALSE NOT NULL,
	active BOOLEAN DEFAULT TRUE NOT NULL,
	created TIMESTAMP DEFAULT (NOW() AT TIME ZONE 'UTC') NOT NULL,
);

-- Blog Post table
CREATE TABLE IF NOT EXISTS blog_posts (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	author UUID REFERENCES blog_users (id) NOT NULL,
	read_time INT NULL,
	html TEXT NOT NULL
	created TIMESTAMP DEFAULT (NOW() AT TIME ZONE 'UTC') NOT NULL,
	updated TIMESTAMP NULL,
);

-- Blog Post Source table
CREATE TABLE IF NOT EXISTS blog_post_source (
	blog_post UUID REFERENCES blog_posts (id) ON DELETE CASCADE,
	source TEXT NOT NULL,
	PRIMARY KEY (blog_post)
);

-- Image Metadata table
CREATE TABLE IF NOT EXISTS blog_images (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	url TEXT NOT NULL,
	name CHARACTER VARYING(100) NOT NULL,
	size bigint NOT NULL
);
