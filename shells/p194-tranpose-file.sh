# 194. Tranpose file
# https://leetcode.com/problems/transpose-file/

name=(name)
age=(age)
cat file.txt | while read -r line; do
  name+=(${line[1]})
  age+=(${line[2]})
done
echo "$name"
echo "$age"