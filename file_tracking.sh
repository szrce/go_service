#!/bin/sh
#
# PROVIDE: goprogram
# REQUIRE: networking
# KEYWORD:

. /etc/rc.subr

name="file_tracking"
rcvar="file_tracking_enable"
goprogram_user="file_tracking"
goprogram_command="/root/file_tracking"
pidfile="/var/run/${name}.pid"
command="/usr/sbin/daemon"
command_args="-P ${pidfile} -r -f ${goprogram_command}"

load_rc_config $name
: ${goprogram_enable:=no}

run_rc_command "$1"
