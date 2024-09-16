
-- +migrate Up
CREATE TABLE Orders (
                        OrderID INT PRIMARY KEY,
                        CustomerID INT,
                        OrderDate DATETIME,
                        Amount DECIMAL(10, 2),
                        FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID)
);

-- +migrate Down
