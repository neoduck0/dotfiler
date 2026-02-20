if uwsm check may-start 1> /dev/null && uwsm select; then
	exec uwsm start default
fi
