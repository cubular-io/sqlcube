-- Create a procedure to add a new order
CREATE OR REPLACE PROCEDURE AddOrder(
    IN customer_id INT,
    IN order_date TIMESTAMP,
    IN amount NUMERIC
)
    LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO orders (customerid, orderdate, amount)
    VALUES (customer_id, order_date, amount);
END;
$$;

-- Create a procedure to get high value orders
CREATE OR REPLACE PROCEDURE GetHighValueOrders()
    LANGUAGE plpgsql
AS $$
BEGIN
    SELECT * FROM high_value_orders;
END;
$$;