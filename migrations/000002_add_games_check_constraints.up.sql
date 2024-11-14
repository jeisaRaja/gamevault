-- Check that year is between 1900 and the current year
ALTER TABLE games 
    ADD CONSTRAINT game_year_check CHECK (year BETWEEN 1900 AND EXTRACT(YEAR FROM CURRENT_DATE));

-- Check that genres array length is between 1 and 10
ALTER TABLE games 
    ADD CONSTRAINT genres_length_check CHECK (array_length(genres, 1) BETWEEN 1 AND 10);

-- Check that platforms array length is between 1 and 10
ALTER TABLE games 
    ADD CONSTRAINT platforms_length_check CHECK (array_length(platforms, 1) BETWEEN 1 AND 10);

-- Check that price is non-negative
ALTER TABLE games 
    ADD CONSTRAINT price_check CHECK (price >= 0);

-- Check that rating_count is non-negative
ALTER TABLE games 
    ADD CONSTRAINT rating_count_check CHECK (rating_count >= 0);
