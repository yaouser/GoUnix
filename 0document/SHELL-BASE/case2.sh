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
1|R) echo "Restore";;
2|B) echo "Backup";;
3|U) echo "Unload";;
*) echo "Sorry $CHOICE is not valid choice"
	exit 1;;
esac 

