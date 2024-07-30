-- Create the database
CREATE DATABASE CoffeeLine;
USE CoffeeLine;

-- Create the Product table
CREATE TABLE Product (
  ProductID INT AUTO_INCREMENT PRIMARY KEY,  -- Primary key for the product table
  ProductName VARCHAR(255) NOT NULL,         -- Name of the product
  Stock INT UNSIGNED NOT NULL,               -- Stock quantity (must be non-negative)
  Price DECIMAL(10, 2) NOT NULL,             -- Price of the product
  CreatedAt DATE DEFAULT CURRENT_DATE,       -- Date when the product was created
  UpdatedAt DATE DEFAULT CURRENT_DATE        -- Date when the product was last updated
);

-- Create the Role table
CREATE TABLE Role (
  RoleID INT AUTO_INCREMENT PRIMARY KEY,     -- Primary key for the role table
  RoleName VARCHAR(255) NOT NULL             -- Name of the role
);

-- Create the Users table
CREATE TABLE Users (
  UserID INT AUTO_INCREMENT PRIMARY KEY,     -- Primary key for the users table
  Name VARCHAR(255) NOT NULL,                -- Name of the user
  Email VARCHAR(255) NOT NULL UNIQUE,        -- Email of the user (unique)
  Password CHAR(60) NOT NULL,                -- Password of the user
  Location VARCHAR(255) NOT NULL,            -- Location of the user
  RoleID INT NOT NULL,                       -- Role ID (foreign key)
  CreatedAt DATE DEFAULT CURRENT_DATE,       -- Date when the user was created
  UpdatedAt DATE DEFAULT CURRENT_DATE,       -- Date when the user was last updated
  FOREIGN KEY (RoleID) REFERENCES Role(RoleID) -- Foreign key constraint referencing Role table
);

-- Create the Payments table
CREATE TABLE Payments (
  PaymentID INT AUTO_INCREMENT PRIMARY KEY,  -- Primary key for the payments table
  PaymentAmount DECIMAL(10, 2) NOT NULL,     -- Amount of the payment
  PaymentDate DATE DEFAULT CURRENT_DATE      -- Date of the payment
);

-- Create the Shipping table
CREATE TABLE Shipping (
  ShippingID INT AUTO_INCREMENT PRIMARY KEY, -- Primary key for the shipping table
  ShippingStatus VARCHAR(255) NOT NULL       -- Status of the shipping
);

-- Create the Orders table
CREATE TABLE Orders (
  OrderID INT AUTO_INCREMENT PRIMARY KEY,    -- Primary key for the orders table
  UserID INT NOT NULL,                       -- User ID (foreign key)
  PaymentID INT NOT NULL,                    -- Payment ID (foreign key)
  ShippingID INT NOT NULL,                   -- Shipping ID (foreign key)
  OrderDate DATE DEFAULT CURRENT_DATE,       -- Date of the order
  FOREIGN KEY (UserID) REFERENCES Users(UserID),  -- Foreign key constraint referencing Users table
  FOREIGN KEY (PaymentID) REFERENCES Payments(PaymentID), -- Foreign key constraint referencing Payments table
  FOREIGN KEY (ShippingID) REFERENCES Shipping(ShippingID) -- Foreign key constraint referencing Shipping table
);

-- Create the OrderProduct table
CREATE TABLE OrderProduct (
  OrderProductID INT AUTO_INCREMENT PRIMARY KEY, -- Primary key for the order-product table
  OrderID INT NOT NULL,                          -- Order ID (foreign key)
  ProductID INT NOT NULL,                        -- Product ID (foreign key)
  Quantity INT UNSIGNED NOT NULL,                -- Quantity of the product in the order (must be non-negative)
  FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),       -- Foreign key constraint referencing Orders table
  FOREIGN KEY (ProductID) REFERENCES Product(ProductID)   -- Foreign key constraint referencing Product table
);
