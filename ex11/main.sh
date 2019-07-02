#!/bin/sh

wget http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt
sed 's/	/ /g' hightemp.txt
rm hightemp.txt
