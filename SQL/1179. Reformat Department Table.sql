/*Write an SQL query to reformat the table such that there is a department id column 
and a revenue column for each month.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reformat-department-table
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/*需要进一步了解MySQL语法*/
select id,
max(case when month='Jan' then revenue else null end) as Jan_Revenue,
max(case when month='Feb' then revenue else null end) as Feb_Revenue,
max(case when month='Mar' then revenue else null end) as Mar_Revenue,
max(case when month='Apr' then revenue else null end) as Apr_Revenue,
max(case when month='May' then revenue else null end) as May_Revenue,
max(case when month='Jun' then revenue else null end) as Jun_Revenue,
max(case when month='Jul' then revenue else null end) as Jul_Revenue,
max(case when month='Aug' then revenue else null end) as Aug_Revenue,
max(case when month='Sep' then revenue else null end) as Sep_Revenue,
max(case when month='Oct' then revenue else null end) as Oct_Revenue,
max(case when month='Nov' then revenue else null end) as Nov_Revenue,
max(case when month='Dec' then revenue else null end) as Dec_Revenue
from Department d group by id;