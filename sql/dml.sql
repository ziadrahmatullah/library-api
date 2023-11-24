INSERT INTO authors (name, created_at, updated_at)
VALUES 
('Ziad', NOW(), NOW()),
('Ziyad', NOW(), NOW()),
('Jiad', NOW(), NOW());

INSERT INTO books (title, description, quantity, cover, author_id, created_at, updated_at)
VALUES 
('Buku 1', 'Tentang orang 1', 10, 'Book_One',1, NOW(), NOW()),
('Buku 2', 'Tentang orang 2', 10, 'Book_Two',1, NOW(), NOW()),
('Buku 3', 'Tentang orang 3', 10, 'Book_Three', 1,NOW(), NOW()),
('Buku 4', 'Tentang orang 4', 10, 'Book_Four', 1,NOW(), NOW()),
('Buku 5', 'Tentang orang 5', 10, 'Book_Five',1, NOW(), NOW()),
('Buku 6', 'Tentang orang 6', 10, 'Book_Six',1, NOW(), NOW()),
('Buku 7', 'Tentang orang 7', 10, 'Book_Seven',1, NOW(), NOW()),
('Buku 8', 'Tentang orang 8', 10, 'Book_Eight', 1,NOW(), NOW()),
('Buku 9', 'Tentang orang 9', 10, 'Book_Nine', 1,NOW(), NOW());

INSERT INTO users (name, email, phone, created_at, updated_at)
VALUES
('Alice', 'alice@gmail.com', '0877237373', NOW(), NOW()),
('Bob', 'alice@gmail.com', '0877237373', NOW(), NOW()),
('Celine', 'alice@gmail.com', '0877237373', NOW(), NOW()),
('Ferina', 'alice@gmail.com', '0877237373', NOW(), NOW());


INSERT INTO borrowing_books (book_id, user_id, created_at, updated_at)
VALUES
(1, 1, NOW(), NOW()),
(2, 1, NOW(), NOW()),
(3, 1, NOW(), NOW()),
(3, 1, NOW(), NOW()),
(4, 1, NOW(), NOW());
