#!/bin/sh

wget http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt
read n
tail -n $n hightemp.txt
rm hightemp.txt
