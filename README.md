# SetLoggerLevelfromsignals
simplest way to change the logger level in the application without break or restarting the pod are application. Here we are able to send by the signals (SIGUSR1).

In terminal 1

go run main.go

In terminal 2

kill -s USR1 <process id>

Can get below output. 

INFO[2023-01-26T23:31:08.672+05:30] Printing Info 0                              
INFO[2023-01-26T23:31:13.674+05:30] Printing Info 1                              
INFO[2023-01-26T23:31:18.675+05:30] Printing Info 2                              
INFO[2023-01-26T23:31:21.015+05:30] user defined signal 1                        
INFO[2023-01-26T23:31:23.676+05:30] Printing Info 3                              
DEBU[2023-01-26T23:31:23.676+05:30] Printing Debug3                              
INFO[2023-01-26T23:31:28.677+05:30] Printing Info 4                              
DEBU[2023-01-26T23:31:28.677+05:30] Printing Debug4                              
INFO[2023-01-26T23:31:33.677+05:30] Printing Info 5                              
DEBU[2023-01-26T23:31:33.678+05:30] Printing Debug5                              
INFO[2023-01-26T23:31:38.678+05:30] Printing Info 6                              
DEBU[2023-01-26T23:31:38.678+05:30] Printing Debug6                              
INFO[2023-01-26T23:31:43.679+05:30] Printing Info 7                              
DEBU[2023-01-26T23:31:43.679+05:30] Printing Debug7                              
INFO[2023-01-26T23:31:48.680+05:30] Printing Info 8                              
DEBU[2023-01-26T23:31:48.680+05:30] Printing Debug8                              
INFO[2023-01-26T23:31:53.682+05:30] Printing Info 9 


Basically sending signals after no:2 log and able to get debug level after that.
