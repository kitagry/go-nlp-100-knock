#!/bin/sh

wget http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt
wc -l hightemp.txt
rm hightemp.txt
