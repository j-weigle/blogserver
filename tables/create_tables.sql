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
	created TIMESTAMP DEFAULT (NOW() AT TIME ZONE 'UTC') NOT NULL
);

-- Blog Post Full table
CREATE TABLE IF NOT EXISTS blog_posts_full (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	author UUID REFERENCES blog_users(id) NOT NULL,
	read_time INT NULL,
	content TEXT NOT NULL,
	source TEXT NOT NULL,
	created TIMESTAMP DEFAULT (NOW() AT TIME ZONE 'UTC') NOT NULL,
	updated TIMESTAMP NULL
);

-- Blog Post view
CREATE VIEW blog_posts AS SELECT (
	id,
	author,
	read_time,
	content,
	created,
	updated
) FROM blog_posts_full;

-- Image Metadata table
CREATE TABLE IF NOT EXISTS blog_images (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	url TEXT NOT NULL,
	name CHARACTER VARYING(100) NOT NULL,
	size bigint NOT NULL
);
