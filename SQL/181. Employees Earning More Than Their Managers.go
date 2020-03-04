package leetcode

/*The Employee table holds all employees including their managers. Every employee has an Id,
and there is also a column for the manager Id.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/employees-earning-more-than-their-managers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

select e1.Name as Employee
from Employee as e1 left outer join Employee as e2 
on e1.ManagerId = e2.Id
where e1.Salary > e2.Salary