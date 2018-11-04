Script that reads the text file "logs.txt" (must be in the same directory as the script) and sends each text line to the Palo Alto Networks user ID agent (either firewall or standalone agent) for user to IP parsing using either UDP or TLS as transport layer.
https://www.paloaltonetworks.com/documentation/71/pan-os/pan-os/user-id/configure-user-id-to-receive-user-mappings-from-a-syslog-sender

Note: Syslog server certificate is not validated. Designed for testing ONLY!
___
Usage:

for udp
```
./sysPAN -tr=udp -ip=<Agent IP address>
```
for TLS
```
./sysPAN -tr=tls -ip=<Agent IP address>
```
if no arguments are provided it'll use UDP as a transport mechanism and send the syslog to ip address 192.168.1.1

___

Debugging in firewall:

FW> debug user-id set userid syslog

Debug level is info

FW> debug user-id on debug

debug level set to debug

FW> tail follow yes mp-log useridd.log
debug: pan_ssl_conn_accept_i(pan_ssl_utils.c:919): Accepting SSL connection from 10.16.1.5/49310
debug: pan_user_id_syslog_server_parse_msg(pan_user_id_syslog.c:989): msg <6> 2018-11-03T12:39:16-05:00 10.16.1.5 testtag[32851]: register user:username ip:1.1.1.1 Matched profile test with event type 0
___
