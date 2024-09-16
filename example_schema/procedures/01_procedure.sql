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