CREATE VIEW HighValueOrders AS
SELECT OrderID, CustomerID, OrderDate, Amount
FROM Orders
WHERE Amount > 1000;

CREATE VIEW CustomerStatistics AS
SELECT CustomerID, COUNT(OrderID) AS TotalOrders, SUM(Amount) AS TotalSpent
FROM Orders
GROUP BY CustomerID;
