#!/bin/sh

wget http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt
cut -f 3 hightemp.txt | sort -r
rm hightemp.txt
