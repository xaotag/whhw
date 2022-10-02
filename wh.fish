#!/usr/bin/fish
function conn
 ping -c 1 baidu.com > /dev/null 2>&1
	if test $status = 0
		notify-send "已连接 $ip"
	else
		notify-send "已断开"
	end
end
function wh 
	set -g ip (ip a | grep -Eo "([0-9]{1,3}.){3}\.[0-9]{1,3}" | sed -n "1p") 
	set oldip  (cat  ~/.ip)
	if test "$ip" != "$oldip"
			echo $ip >~/.ip
		ping -c 1 baidu.com > /dev/null 2>&1
		if test $status != 0
			~/.script/whhw&  
		end
		conn	
	end
end
	
while true
	sleep 1
	wh
end
