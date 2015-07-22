iscontinue(){
while true 
do 
	echo -n "Continue?(Y/N)"
	read ANSWER
	case $ANSWER in
		[Yy]) return 0;;
		[Nn]) return 1;;
		*) echo "Answer Y or N";;
	esac
done
}
