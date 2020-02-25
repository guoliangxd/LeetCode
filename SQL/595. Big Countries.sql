/*A country is big if it has an area of bigger than 3 million square km or a population of more than 25 million.

Write a SQL solution to output big countries' name, population and area.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/big-countries
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

select name, population, area
from World
where area > 3000000 or population > 25000000