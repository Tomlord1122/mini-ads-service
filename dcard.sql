CREATE TYPE gender_enum AS ENUM ('M', 'F');
CREATE TYPE country_enum AS ENUM ('TW', 'JP', 'US'); -- 假設這些是可能的值
CREATE TYPE platform_enum AS ENUM ('android', 'ios', 'web');

CREATE TABLE IF NOT EXISTS ads (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    start_at TIMESTAMP WITH TIME ZONE NOT NULL,
    end_at TIMESTAMP WITH TIME ZONE NOT NULL,
    age INT,
    gender gender_enum,
    country country_enum[],
    platform platform_enum[]
);