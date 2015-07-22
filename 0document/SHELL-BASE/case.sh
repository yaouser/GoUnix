case "$1" in
	start)
		start;;
	stop)
		stop;;
	status)
		rhstatus;;
	rstart|reload)
		restart;;
	condrestart)
		[ -f /var/lock/subsys/syslog ] && restart || : ;;
	*)
		echo "Usage:$0{start|stop|status|restart|condrestart}"
		exit 1
esac
