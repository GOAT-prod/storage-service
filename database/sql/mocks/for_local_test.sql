INSERT INTO material (name, country)
VALUES ('Leather', 'Italy'),
       ('Rubber', 'Germany'),
       ('Textile', 'China'),
       ('Synthetic', 'USA'),
       ('Mesh', 'Vietnam');

INSERT INTO brand (name)
VALUES ('Nike'),
       ('Adidas'),
       ('Puma'),
       ('Reebok'),
       ('New Balance');

INSERT INTO factory (name, country, city, address)
VALUES ('Nike Factory 1', 'Vietnam', 'Ho Chi Minh City', '123 Main St'),
       ('Adidas Factory', 'China', 'Guangzhou', '456 Second St'),
       ('Puma Factory', 'Germany', 'Berlin', '789 Third St'),
       ('Reebok Factory', 'USA', 'Los Angeles', '321 Sunset Blvd'),
       ('New Balance Factory', 'UK', 'London', '654 Thames Rd');

INSERT INTO product (brand_id, factory_id, name, description, price, is_approved, is_deleted)
VALUES (1, 1, 'Nike Air Max', 'High-quality running sneakers with Air Max cushioning.', 120.99, true, false),
       (2, 2, 'Adidas Ultraboost', 'Premium running shoes with Boost technology.', 180.50, true, false),
       (3, 3, 'Puma Suede Classic', 'Retro-style sneakers with suede upper.', 75.00, true, false),
       (4, 4, 'Reebok Nano X1', 'Cross-training shoes for all-around performance.', 130.00, true, false),
       (5, 5, 'New Balance 990v5', 'Classic running shoes made in the USA.', 175.99, true, false);

INSERT INTO product_item (product_id, stock_count, size, weight, color)
VALUES (1, 50, 42, 0.9, 'Black'),
       (1, 30, 44, 1.0, 'White'),
       (2, 40, 43, 0.85, 'Blue'),
       (2, 35, 45, 0.9, 'Gray'),
       (3, 25, 40, 0.8, 'Red'),
       (4, 60, 41, 0.95, 'Green'),
       (5, 20, 43, 0.88, 'Navy');

INSERT INTO product_material (product_id, material_id)
VALUES (1, 1), -- Nike Air Max: Leather
       (1, 2), -- Nike Air Max: Rubber
       (2, 3), -- Adidas Ultraboost: Textile
       (2, 2), -- Adidas Ultraboost: Rubber
       (3, 1), -- Puma Suede Classic: Leather
       (4, 4), -- Reebok Nano X1: Synthetic
       (5, 5); -- New Balance 990v5: Mesh

INSERT INTO images (url, product_id)
VALUES ('https://example.com/images/nike_air_max.jpg', 1),
       ('https://example.com/images/adidas_ultraboost.jpg', 2),
       ('https://example.com/images/puma_suede_classic.jpg', 3),
       ('https://example.com/images/reebok_nano_x1.jpg', 4),
       ('https://example.com/images/new_balance_990v5.jpg', 5);
