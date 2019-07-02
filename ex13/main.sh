#!/bin/sh

wget http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt
cut -f 1 -d '	' hightemp.txt > col1.txt
cut -f 2 -d '	' hightemp.txt > col2.txt
paste -d '	' col1.txt col2.txt
rm hightemp.txt col1.txt col2.txt
