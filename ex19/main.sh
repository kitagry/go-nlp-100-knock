#!/bin/sh

wget http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt
cut -f 1 hightemp.txt | sort | uniq -c | sort -r -n
rm hightemp.txt
