CREATE TRIGGER IncrementTotalOrders
    AFTER INSERT ON Orders
    FOR EACH ROW
BEGIN
UPDATE Customers
SET TotalOrders = TotalOrders + 1
WHERE CustomerID = NEW.CustomerID;
END;