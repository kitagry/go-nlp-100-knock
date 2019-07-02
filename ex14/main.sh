#!/bin/sh

wget http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt
head -n $1 hightemp.txt
rm hightemp.txt
