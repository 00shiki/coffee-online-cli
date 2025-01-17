-- Insert sample data into Role table
INSERT INTO Role (RoleName)
VALUES ('Customer'),
       ('Admin');

-- Insert sample data into Users table
INSERT INTO Users (Name, Email, Password, Location, RoleID)
VALUES ('Alice Johnson', 'alice@example.com', '$2a$12$lNr07cHNenW7BDFNqMGl9OkHnC.9lPNRE13ojEmILMHvFfbdADE0G',
        'Jakarta Selatan', 1),
       ('Bob Smith', 'bob@example.com', '$2a$12$XryfxdUl0u9zWOOC3RmAlOP8gauDbkFCuwXhzccNr0cTExxyIIrDm
', 'Jakarta Pusat', 1),
       ('Carol White', 'carol@example.com', '$2a$12$B.Wx6DDTpMTW6x9M3LCnqOeEZntzX9ARcqQZXtj1.wzoTqum3wBTW
', 'Jakarta Barat', 1),
       ('Dave Brown', 'dave@example.com', '$2a$12$gKLM/msfA86KfPXiAYQ.Sus/x7fs0evjIxISiw4huHd.6pSXBVTVq
', 'Jakarta Timur', 1),
       ('Eve Davis', 'eve@example.com', '$2a$12$q5/W9r5y7RGwymcas0ty2eky.e865HZKQpxEGlF3Sq3PF0RVY40Na
', 'Jakarta Selatan', 2);

-- Insert sample data into Product table
INSERT INTO Product (ProductName, Stock, Price) 
VALUES ('Espresso', 100, 15000),
       ('Cappuccino', 50, 20000),
       ('Latte', 75, 22000),
       ('Americano', 60, 18000),
       ('Mocha', 40, 17500);

-- Insert sample data into Payments table
INSERT INTO Payments (PaymentAmount)
VALUES (59000),
       (53000),
       (44500),
       (46000),
       (67000);
       
-- Insert sample data into Shipping table
INSERT INTO Shipping (ShippingStatus)
VALUES ('Pending'),
       ('Shipped'),
       ('Delivered');

-- Insert sample data into Orders table
INSERT INTO Orders (UserID, PaymentID, ShippingID)
VALUES (1, 1, 1),
       (2, 2, 2),
       (3, 3, 3),
       (4, 4, 1),
       (5, 5, 2);

-- Insert sample data into OrderProduct table
INSERT INTO OrderProduct (OrderID, ProductID, Quantity)
VALUES (1, 1, 2),
       (1, 2, 1),
       (2, 3, 2),
       (3, 4, 1),
       (3, 5, 1),
       (4, 1, 1),
       (4, 3, 1),
       (5, 2, 2),
       (5, 4, 1);
