-- Drop the year check constraint
ALTER TABLE games
    DROP CONSTRAINT IF EXISTS game_year_check;

-- Drop the genres length check constraint
ALTER TABLE games
    DROP CONSTRAINT IF EXISTS genres_length_check;

-- Drop the platform length check constraint
ALTER TABLE games
    DROP CONSTRAINT IF EXISTS platform_length_check;

-- Drop the price check constraint
ALTER TABLE games
    DROP CONSTRAINT IF EXISTS price_check;

-- Drop the rating count check constraint
ALTER TABLE games
    DROP CONSTRAINT IF EXISTS rating_count_check;
