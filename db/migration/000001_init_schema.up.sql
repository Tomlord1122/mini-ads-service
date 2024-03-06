CREATE TYPE gender_enum AS ENUM ('M', 'F');
CREATE TYPE country_enum AS ENUM ('TW', 'JP', 'US'); -- 假設這些是可能的值
CREATE TYPE platform_enum AS ENUM ('android', 'ios', 'web');

CREATE TABLE IF NOT EXISTS ads (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    start_at TIMESTAMP WITH TIME ZONE NOT NULL,
    end_at TIMESTAMP WITH TIME ZONE NOT NULL,
    age INT,
    gender gender_enum[],
    country country_enum[],
    platform platform_enum[]
);

CREATE INDEX idx_ads_start_at ON ads(start_at);
CREATE INDEX idx_ads_end_at ON ads(end_at);
CREATE INDEX idx_ads_age ON ads(age);
CREATE INDEX idx_ads_country ON ads USING GIN (country);
CREATE INDEX idx_ads_platform ON ads USING GIN (platform);
CREATE INDEX idx_ads_gender ON ads USING GIN (gender);

