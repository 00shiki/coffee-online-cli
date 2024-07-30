# Entity Relationship Diagram (ERD) : CoffeeLine

![image](ERD.png)

## 1. Database Creation:

    We start by creating a database named CoffeeLine.

## 2. Tables and Their Purpose:

- Product Table:

  - ProductID: Unique identifier for each product.
  - ProductName: Name of the product.
  - Stock: Quantity of the product in stock (must be non-negative).
  - Price: Price of the product.
  - CreatedAt: Timestamp of when the product was added.
  - UpdatedAt: Timestamp of when the product was last updated.

- Role Table:

  - RoleID: Unique identifier for each role.
  - RoleName: Name of the role.

- Users Table:

  - UserID: Unique identifier for each user.
  - Name: Name of the user.
  - Email: Email address of the user (must be unique).
  - Password: Hashed password of the user.
  - Location: Location of the user.
  - RoleID: Identifier for the role the user has (foreign key referencing Role table).
  - CreatedAt: Timestamp of when the user was created.
  - UpdatedAt: Timestamp of when the user was last updated.

- Orders Table:

  - OrderID: Unique identifier for each order.
  - UserID: Identifier of the user who made the order (foreign key referencing Users table).
  - PaymentID: Identifier for the payment method used (foreign key referencing Payments table).
  - ShippingID: Identifier for the shipping information (foreign key referencing Shipping table).
  - OrderDate: Date and time when the order was placed.

- Payments Table:

  - PaymentID: Unique identifier for each payment.
  - PaymentAmount: Amount of the payment.
  - PaymentDate: Date and time of the payment.

- Shipping Table:

  - ShippingID: Unique identifier for each shipping entry.
  - ShippingStatus: Status of the shipping (e.g., pending, shipped, delivered).

- OrderProduct Table:
  - OrderProductID: Unique identifier for each order-product entry.
  - OrderID: Identifier for the order (foreign key referencing Orders table).
  - ProductID: Identifier for the product (foreign key referencing Product table).
  - Quantity: Quantity of the product in the order (must be non-negative).

## 3. Relationships:

The Users table references the Role table through the RoleID foreign key.
The Orders table references the Users, Payments, and Shipping tables through the UserID, PaymentID, and ShippingID foreign keys, respectively.
The OrderProduct table references the Orders and Product tables through the OrderID and ProductID foreign keys, respectively.

## 4. Cardinality:

- Product to OrderProduct: One-to-Many
- OrderProduct to Orders: Many-to-One
- Users to Orders: One-to-Many
- Orders to Payments: One-to-One
- Orders to Shipping: One-to-One
- Role to Users: One-to-Many
