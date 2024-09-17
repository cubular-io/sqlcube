CREATE VIEW CustomerOrders AS
SELECT c.CustomerID, c.FirstName, c.LastName, o.OrderID, o.OrderDate, o.Amount
FROM Customers c
         JOIN Orders o ON c.CustomerID = o.CustomerID;

CREATE VIEW RecentCustomers AS
SELECT CustomerID, FirstName, LastName, SignupDate
FROM Customers
WHERE SignupDate > DATEADD(MONTH, -1, GETDATE());
-- End of 01_views.sql --

CREATE VIEW HighValueOrders AS
SELECT OrderID, CustomerID, OrderDate, Amount
FROM Orders
WHERE Amount > 1000;

CREATE VIEW CustomerStatistics AS
SELECT CustomerID, COUNT(OrderID) AS TotalOrders, SUM(Amount) AS TotalSpent
FROM Orders
GROUP BY CustomerID;

-- End of 02_view2.sql --

