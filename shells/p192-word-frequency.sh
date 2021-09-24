# 194. word frequency
# https://leetcode.com/problems/word-frequency/

cat file.txt | tr -s ' ' '\n' | sort | uniq -c | sort -r | awk '{ print $2, $1 }'