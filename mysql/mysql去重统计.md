# mysql count distinct 统计结果去重

mysql的sql语句中，count这个关键词能统计表中的数量,如

有一个tableA表，表中数据如下：

id	name	age
1	tony	18
2	jacky	19
3	jojo	18


SELECT COUNT(age) FROM tableA
以上这条语句能查出table表中有多少条数据。查询结果是3

而COUNT这个关键词与 DISTINCT一同使用时，可以将统计的数据中某字段不重复的数量。 如：

SELECT COUNT(DISTINCT age) from tableA

以上语句的查询结果是对age这字段的去重。结果是2