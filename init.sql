-- init.sql
CREATE TABLE IF NOT EXISTS attributes (
    id UUID PRIMARY KEY,
    value TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS questions (
    id UUID PRIMARY KEY,
    text TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS question_attributes (
    question_id UUID REFERENCES questions(id),
    attribute_id UUID REFERENCES attributes(id),
    PRIMARY KEY (question_id, attribute_id)
);

CREATE TABLE IF NOT EXISTS comments (
    id UUID PRIMARY KEY,
    author TEXT NOT NULL,
    email TEXT,
    telephone TEXT,
    text TEXT NOT NULL,
    visibility BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sums (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    amount INT NOT NULL
);
