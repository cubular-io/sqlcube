CREATE OR REPLACE PROCEDURE GetCustomerOrders(IN customer_id INT)
    LANGUAGE plpgsql
AS $$
BEGIN
    SELECT * FROM customer_orders WHERE customerid = customer_id;
END;
$$;

-- Create a procedure to add a new customer
CREATE OR REPLACE PROCEDURE AddCustomer(
    IN first_name VARCHAR,
    IN last_name VARCHAR,
    IN email VARCHAR,
    IN signup_date TIMESTAMP
)
    LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO customers (firstname, lastname, email, signupdate)
    VALUES (first_name, last_name, email, signup_date);
END;
$$;
-- End of 01_procedure.sql --

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
-- End of 02_procedure.sql --

