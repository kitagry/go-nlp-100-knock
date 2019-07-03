#!/bin/sh

wget http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt
n=$(cat hightemp.txt | wc -l)
count=$(($n / $1))
if [ $(($n % $1)) != 0 ]; then
  count=$count + 1
fi
split -l $count hightemp.txt
rm hightemp.txt
