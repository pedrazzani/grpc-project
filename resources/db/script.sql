create table main.categories
(
    id          string,
    name        string,
    description string
);

create table main.products
(
    id          string,
    name        string,
    description string,
    category_id string
);

INSERT INTO categories (id, name, description)
VALUES ('5001a184-1242-42e2-a1da-200d3017fb4e', 'Category 1', 'Category 1');
INSERT INTO categories (id, name, description)
VALUES ('e70e27b3-7ced-47c4-acbc-875de4640131', 'Category 2', 'Category 2');

INSERT INTO products (id, name, description, category_id)
VALUES ('e377227e-79a2-4c3d-92e6-133a79baea1d', 'Product 1', 'Product 1', '5001a184-1242-42e2-a1da-200d3017fb4e');
INSERT INTO products (id, name, description, category_id)
VALUES ('e9b03e26-1858-4678-8d87-bf3280029159', 'Product 2', 'Product 2', '5001a184-1242-42e2-a1da-200d3017fb4e');

