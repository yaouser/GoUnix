#!/bin/bash
#NO SUCCUSS
#filename: newdate
if [ $# -lt 1 ]
then 
	date
else
	while getopt mdyDHMSTjJwahr OPTION
	do 
		case $OPTION
		in 
			m)date '+%m';;  #Month of Year
			d)date '+%d';;  #Day of Month
			y)date '+%y';;  #Year
			D)date '+%D';;  #MM/DD/YY
			H)date '+%H';;  #Hour
			M)date '+%M';;  #Minute
			S)date '+%S';;  #Second
			T)date '+%T';;  #HH:MM:SS
			j)date '+%j';;  #day of year
			J)date '+%y%j';; #digit Julian date
			w)date '+%w';;  #Day of the Week
			a)date '%a';;   #Day abbreviation
			h)date '+%h';;  #Month abbreviation
			r)date '+%r';;  #AM-PM time
			\?)echo "Invalid option $OPTION";;
		esac
	done
fi
