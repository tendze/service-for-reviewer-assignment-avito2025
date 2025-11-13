CREATE TYPE pr_status AS ENUM ('OPEN', 'MERGED');

CREATE TABLE team (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE NOT NULL,
    team_id INT NOT NULL REFERENCES team(id),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE pull_request (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author_id INT NOT NULL REFERENCES "user"(id),
    status pr_status DEFAULT 'OPEN' NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE pr_reviewer (
    id SERIAL PRIMARY KEY,
    pr_id INT NOT NULL REFERENCES pull_request(id),
    reviewer_id INT NOT NULL REFERENCES "user"(id),
    created_at TIMESTAMP DEFAULT NOW()
);
