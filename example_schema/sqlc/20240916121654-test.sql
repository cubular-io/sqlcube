
-- +migrate Up
CREATE TABLE Customers (
                           CustomerID INT PRIMARY KEY,
                           FirstName VARCHAR(50),
                           LastName VARCHAR(50),
                           Email VARCHAR(100),
                           SignupDate DATETIME
);
-- +migrate Down
