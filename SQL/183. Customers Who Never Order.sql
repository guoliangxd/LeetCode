/*Suppose that a website contains two tables, the Customers table and the Orders table. 
Write a SQL query to find all customers who never order anything.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/customers-who-never-order
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

select Name as Customers
from Customers left outer join Orders
on Customers.Id = Orders.CustomerId
where isnull(Orders.Id)