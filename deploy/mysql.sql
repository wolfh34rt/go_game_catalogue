CREATE USER 'user-gamecat'@'localhost' IDENTIFIED BY 'password';

CREATE database gamecatalogue;

GRANT usage on *.* to 'user-gamecat'@'localhost';

GRANT ALL PRIVILEGES ON gamecatalogue.* to 'user-gamecat'@'localhost';

USE gamecatalogue;

CREATE TABLE mediatypes (
	mediatypeid VARCHAR(30) NOT NULL PRIMARY KEY,
	mediatypename VARCHAR(50) NOT NULL
);

CREATE TABLE platforms (
	platformid VARCHAR(30) NOT NULL PRIMARY KEY,
	platformname VARCHAR(50) NOT NULL
);

CREATE TABLE items (
	itemid VARCHAR(30) NOT NULL PRIMARY KEY,
	itemname VARCHAR(30) NOT NULL,
	mediatype VARCHAR(30) NOT NULL,
	platform VARCHAR(30) NOT NULL
);

CREATE TABLE borrowed_items (
	borrowid VARCHAR(30) NOT NULL PRIMARY KEY,
	userid VARCHAR(30) NOT NULL,
	itemid VARCHAR(30) NOT NULL,
	borrowdate DATE NOT NULL
);


CREATE TABLE userprofile (
	profileid VARCHAR(30) NOT NULL PRIMARY KEY,
	userid VARCHAR(30) NOT NULL,
	email VARCHAR(100) NOT NULL
);

CREATE TABLE membership (
	userid VARCHAR(30) NOT NULL PRIMARY KEY,
	createdate DATE,
	confirmationtoken VARCHAR(128),
	isconfirmed BIT DEFAULT b'0',
	lastpasswordfailuredate DATE,
	passwordfailuressincelastsuccess INT,
	password VARCHAR(128) NOT NULL,
	passwordchangeddate DATE,
	passwordsalt VARCHAR(128),
	passwordverificationtoken VARCHAR(128),
	passwordverificationtokenexpirationdate DATE
);

CREATE TABLE oauthmembership (
	oauthmembershipid VARCHAR(30) NOT NULL PRIMARY KEY,
	provider VARCHAR(30) NOT NULL,
	provideruserid VARCHAR(30) NOT NULL,
	userid VARCHAR(30) NOT NULL
);

CREATE TABLE roles (
	roleid VARCHAR(30) NOT NULL PRIMARY KEY,
	rolename VARCHAR(30) NOT NULL
);

CREATE TABLE usersinroles (
	usersinroleid VARCHAR(30) NOT NULL PRIMARY KEY,
	userid VARCHAR(30) NOT NULL,
	roleid VARCHAR(30) NOT NULL
);

INSERT INTO `platforms` (`platformid`, `platformname`) VALUES
('20oaYV26Sy9Gxdkd4tbKpRThXXpyLN', 'PSX'),
('7Bp3peAG5YC6bTnq5SaZ92GGyFEdxN', 'PS3'),
('dKZGAxjN4mZnH8NS2BPk4q5MEdJWtP', 'XBOX'),
('eRm5AKUYOL2SDGuIaFcJyWi3453pcy', 'XBOX360'),
('EsUGL92pAo0asSdT9hyKGjTE65bTHD', 'DREAMCAST'),
('KV1IOZGIdGEioGBkLUfVoHftDqCc5N', 'PS2'),
('mrOWlqw5qQJgyNCOHNDyvqY19B74aH', 'PC');

INSERT INTO `mediatypes` (`mediatypeid`, `mediatypename`) VALUES
('jDZK9G2CFHdtAHk8KYpwmh1ThPk0Gf', 'STEAM'),
('Laf8UxMk96h0WVpN1bGjdYeK3alygG', 'Disc Media'),
('NYvtGleHNmAe3pPomi6esPGAb7YVwk', 'ORIGIN'),
('phngvXMSP0S7U2HKdOkAhcKhdqB1In', 'GOG.COM'),
('pZgRDy9uLFq2M21JDCi3X9dAA9mrxQ', 'GAMEFLY');

INSERT INTO `roles` (`roleid`, `rolename`) VALUES
('FIoiFidwqdjf234FgjifffeZfMdLM3', 'Borrower'),
('FIoiFidwqdjf376FgjidPieZfMdLM3', 'Administrator');