Webshare
========

Simple web sharing app.

Just tell it the directory to share and the port (80 by default) and that directory, it files and its subdirectories will be accessible via HTTP.

(NO password nor upload support yet...)

Usage samples:
```
webshare 
```
Will share the current directory through the URL :80/webshare/

```
webshare -port 1080
```
Will share the current directory through the URL :80/webshare/

```
webshare -dir /home/jdoe/sharedfiles -port 1080
```
Will share /home/jdoe/sharedfiles through the URL :1080/webshare/

```
webshare -prefix / -dir /home/jdoe/sharedfiles -port 1080
```
Will share /home/jdoe/sharedfiles through the URL :1080/
