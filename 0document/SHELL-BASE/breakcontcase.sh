#!/bin/bash
#Interactive program to restore, backup, or unload
#a directory
echo "Welcome to the menu driven Archive program"
while true
do 
#Display a Menu
	echo 
	echo "Make a Choice from the Menu below"
	echo _
	echo "1 Restore Archive"
	echo "2 Backup directory"
	echo "3 Unload directory"
	echo "4 Quit"
	echo

# Read the user's selection
	echo -n "Enter Choice: "
read CHOICE
case $CHOICE in
[1-3]) echo 
#Read and validate the name of the directory
echo -n "What directory do you want?"
read WORKDIR
if [ ! -d "$WORKDIR" ]
then 
	echo "Sorry, $WORKDIR is not a directory"
	continue
fi
#Make the directory the current working directory
cd $WORKDIR;;
4) : ;;  # : is null,not excute everything
*) echo "Sorry, $CHOICE is not a valid choice"
continue
esac

case "$CHOICE" in
1) echo "Restoring..."
cpio -i </dev/myfile/0h;;
2) echo "Archiving..."
ls | cpio -o >/dev/myfile/0h;;
3) echo "Unloading..."
ls | cpio -o >/dev/myfile/0h;;
4) echo "Quitting"
break;;
esac

#Check for cpio errors
if [ $? -ne 0 ]
then 
	echo "A problem has occurred during the process"
	if [ $CHOICE = 3 ]
	then 
	echo "The directory will not be erased"
	fi
	echo "Please check the device and try again"
	continue
else
	if [ $CHOICE = 3 ]
	then 
		ls -l	#rm *
	fi
fi
done
