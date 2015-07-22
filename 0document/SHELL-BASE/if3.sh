echo -n "Do you want to continue: Y or N"
read ANSWER
if [ $ANSWER=N -o $ANSWER=n ]
then
	exit
fi
