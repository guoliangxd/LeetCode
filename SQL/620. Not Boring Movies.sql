/*X city opened a new cinema, many people would like to go to this cinema. The cinema also gives out a poster indicating the movies’ ratings and descriptions.
Please write a SQL query to output movies with an odd numbered ID and a description that is not 'boring'. Order the result by rating.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/not-boring-movies
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

select id, movie, description, rating 
from cinema 
where mod(id, 2) = 1 and description != 'boring'
order by rating desc