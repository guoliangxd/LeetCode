/*Write a SQL query to find all duplicate emails in a table named Person.*/
select Email 
from Person
group by Email
having count(*) >= 2