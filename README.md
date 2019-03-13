# goipaddress
Golang module for IPv4 processing.
This module can parse IP addresses like 192.\*.1-21, and create IPv4Network instances.
Function ToInt converts an IP address to its integer form.
For example goipaddress.ToInt("192.168.0.1") will return 3232235521 as int64.
FromInt does the inverse transform.
Work in progress :)
