INSERT INTO albums (album)
SELECT '{"title": "Blue Train", "artist": "Gerry Mulligan", "price": 56.99 }'
WHERE NOT EXISTS
        (SELECT id
         FROM albums
         WHERE album ->> 'title' = 'Blue Train');

INSERT INTO albums (album)
SELECT '{"title": "Jeru", "artist": "Gerry Mulligan", "price": 17.99 }'
WHERE NOT EXISTS
        (SELECT id
         FROM albums
         WHERE album ->> 'title' = 'Jeru');

INSERT INTO albums (album)
SELECT '{"title": "Sarah Vaughan and Clifford Brown", "artist": "Sarah Vaughan", "price": 39.99 }'
WHERE NOT EXISTS
        (SELECT id
         FROM albums
         WHERE album ->> 'title' = 'Sarah Vaughan and Clifford Brown');