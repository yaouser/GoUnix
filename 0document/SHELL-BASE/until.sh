num=1
until [ ! "$num" -le 10 ];
do 
	echo "num is $num"
	num=$(($num+1))
done
