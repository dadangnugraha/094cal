lue='\e[1;34m'
green='\e[1;32m'
purple='\e[1;35m'
cyan='\e[1;36m'
red='\e[1;31m'
white='\e[1;37m'
yellow='\e[1;33m' NOW=`date "+%d.%m.%Y"`
TIME=`date "+%H:%M"`
echo -e $green
echo -e $red
figlet -f slant "the94"
echo -e $green


 #!/bin/bash

echo "masukkan angka"
read p
echo "+"
read q
tambah=$(echo "$p + $q" | bc)
echo "hasil = $tambah"
