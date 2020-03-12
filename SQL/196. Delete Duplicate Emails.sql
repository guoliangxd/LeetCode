/*Write a SQL query to delete all duplicate email entries in a table named Person, keeping only unique emails based on its smallest Id.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/delete-duplicate-emails
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

DELETE FROM Person
WHERE Id NOT IN (   -- 删除不在查询结果中的值
    SELECT id FROM
   (
       SELECT MIN(Id) AS Id -- 排除Email相同时中Id较大的行
       FROM Person
       GROUP BY Email
   ) AS temp    -- 此处需使用临时表，否则会发生报错
)