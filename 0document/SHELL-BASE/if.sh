QQ(){
	echo -n "Do you want to continue?(Yes/No):"
	read YN
	if echo "$YN" | grep -q '^[Yy]\([Ee][Ss]\)*$'
		QQ
	then
		exit 0
	fi
}
QQ
