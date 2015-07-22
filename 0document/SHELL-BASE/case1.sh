#Display a menu
echo _
echo "1 Restore"
echo "2 Backup"
echo "3 Unload"
echo 

#Read and excute the user's selection
echo -n "Enter Choice:"
read CHOICE

case "$CHOICE" in
1) echo "Restore";;
2) echo "Backup";;
3) echo "Unload";;
*) echo "Sorry $CHOICE is not valid choice"
	exit 1;;
esac 

