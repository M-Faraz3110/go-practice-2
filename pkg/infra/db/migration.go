package db

var migration = `
	CREATE TABLE IF NOT EXISTS public.users
	(
		id uuid PRIMARY KEY,
		email TEXT UNIQUE,
		deleted BOOL DEFAULT FALSE,
		created_at timestamp with time zone,
		updated_at timestamp with time zone
	);

	CREATE INDEX IF NOT EXISTS users_index ON users USING btree (email);

	CREATE TABLE IF NOT EXISTS public.books 
	(
		id uuid PRIMARY KEY,
		title TEXT,
		author TEXT,
		count INTEGER,
		deleted BOOL DEFAULT FALSE,
		created_at timestamp with time zone,
		updated_at timestamp with time zone
	);

	CREATE TABLE IF NOT EXISTS public.borrows
	(
		id uuid PRIMARY KEY,
		book_id uuid,
		user_id uuid,
		returned BOOL DEFAULT FALSE,
		created_at timestamp with time zone,
		updated_at timestamp with time zone
	);

`
