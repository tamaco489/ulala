INSERT INTO users (name, email, auth_uuid, created_at, updated_at)
VALUES
	('John Doe 1', 'john1@example.com', 'abcd1234', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('Jane Smith 2', 'jane2@example.com', 'efgh5678', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('Alice Johnson 3', 'alice3@example.com', 'ijkl9012', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('Bob Brown 4', 'bob4@example.com', 'mnop3456', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('Sarah Davis 5', 'sarah5@example.com', 'qrst7890', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('Michael Wilson 6', 'michael6@example.com', 'uvwxy1234', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('Emily Thompson 7', 'emily7@example.com', 'zabcd5678', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('David Anderson 8', 'david8@example.com', 'efghi9012', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('Olivia Martinez 9', 'olivia9@example.com', 'ijklm3456', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('James Robinson 10', 'james10@example.com', 'nopqr7890', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO movie_types (type_name, created_at, updated_at)
VALUES
	('action', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('anime', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('science_fiction', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('horror', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('comedy', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('romance', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('fantasy', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('sports', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO movie_formats (format_name, created_at, updated_at)
VALUES
	('mp4', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('avi', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('mov', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('wmv', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
	('flv', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);