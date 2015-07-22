QQ(){
	echo -n "Do you want to continue?(Yes/No):"
	read YN
	case "$YN" in 
		[Yy]|[Yy][Es][Ss])
		QQ;;
	*)
		exit 0;;
	esac
}
QQ
