-- List of customers located in Irvine city
SELECT * FROM customers WHERE customers.city = "Irvine";

-- List of customers whose order is handled by an employee named Adam Barr
SELECT * FROM customers
WHERE customers.id IN 
(SELECT customers.id FROM orders left join employees on orders.employee_id = employees.id 
left join customers on orders.customer_id = customers.id
WHERE employees.first_name = "Adam" and employees.last_name = "Barr");

-- List of products which are ordered by "Contonso, Ltd" Company
SELECT products.* FROM order_details left join products on order_details.product_id = products.id
WHERE order_details.order_id IN 
(SELECT orders.id FROM orders left join employees on orders.employee_id = employees.id 
left join customers on orders.customer_id = customers.id
WHERE customers.company_name = "Contoso, Ltd");

-- List of transactions (orders) which has "UPS Ground" as shipping method
SELECT orders.* FROM orders left join shipping_methods on orders.shipping_method_id = shipping_methods.id 
WHERE shipping_methods.shipping_method = "UPS Ground";
