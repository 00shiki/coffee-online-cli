-- Insert sample data into Role table
INSERT INTO Role (RoleName) VALUES
('Customer'),
('Admin');

-- Insert sample data into Users table
INSERT INTO Users (Name, Email, Password, Location, RoleID) VALUES
('Alice Johnson', 'alice@example.com', 'password123', 'New York', 1),
('Bob Smith', 'bob@example.com', 'password456', 'Los Angeles', 1),
('Carol White', 'carol@example.com', 'password789', 'Chicago', 1),
('Dave Brown', 'dave@example.com', 'password101', 'Houston', 1),
('Eve Davis', 'eve@example.com', 'password202', 'Phoenix', 2);

-- Insert sample data into Product table
INSERT INTO Product (ProductName, Stock, Price) VALUES
('Espresso', 100, 3.50),
('Cappuccino', 50, 4.00),
('Latte', 75, 4.50),
('Americano', 60, 3.00),
('Mocha', 40, 5.00);

-- Insert sample data into Payments table
INSERT INTO Payments (PaymentAmount) VALUES
(17.50),
(8.00),
(9.00),
(10.50),
(7.50);

-- Insert sample data into Shipping table
INSERT INTO Shipping (ShippingStatus) VALUES
('Pending'),
('Shipped'),
('Delivered'),
('Pending'),
('Delivered');

-- Insert sample data into Orders table
INSERT INTO Orders (UserID, PaymentID, ShippingID) VALUES
(1, 1, 1),
(2, 2, 2),
(3, 3, 3),
(4, 4, 4),
(5, 5, 5);

-- Insert sample data into OrderProduct table
INSERT INTO OrderProduct (OrderID, ProductID, Quantity) VALUES
(1, 1, 2),
(1, 2, 1),
(2, 3, 2),
(3, 4, 1),
(3, 5, 1),
(4, 1, 1),
(4, 3, 1),
(5, 2, 2),
(5, 4, 1);
