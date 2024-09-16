CREATE VIEW CustomerOrders AS
SELECT c.CustomerID, c.FirstName, c.LastName, o.OrderID, o.OrderDate, o.Amount
FROM Customers c
         JOIN Orders o ON c.CustomerID = o.CustomerID;

CREATE VIEW RecentCustomers AS
SELECT CustomerID, FirstName, LastName, SignupDate
FROM Customers
WHERE SignupDate > DATEADD(MONTH, -1, GETDATE());