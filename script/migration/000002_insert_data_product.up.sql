BEGIN;

INSERT INTO "products" ("name", "price", "stock", "created_at", "updated_at") 
VALUES 
    ('Sarung Sapphire Motif Lilin', 76000, 50, extract(epoch FROM NOW()), extract(epoch FROM NOW())),
    ('Sarung Atlas 550', 47500, 75, extract(epoch FROM NOW()), extract(epoch FROM NOW())),
    ('Sarung Wadimor Hujan Gerimis Gunungan', 58000, 65, extract(epoch FROM NOW()), extract(epoch FROM NOW())),
    ('Sarung Gajah Duduk Signature',  128000, 35, extract(epoch FROM NOW()), extract(epoch FROM NOW())),
    ('Sarung BHS', 485000, 2, extract(epoch FROM NOW()), extract(epoch FROM NOW()));

COMMIT;
