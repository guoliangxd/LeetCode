/*Given a table salary, such as the one below, 
that has m=male and f=female values. Swap all f and m values (i.e., change all f values to m and vice versa) 
with a single update statement and no intermediate temp table.

Note that you must write a single update statement, DO NOT write any select statement for this problem.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-salary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/*case...when...then...else是MySQL的语法吗*/
update salary 
set 
    sex = case sex
        when 'm' then 'f'
        else 'm'
    end;