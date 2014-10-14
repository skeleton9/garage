[mod_auth_baisc](http://httpd.apache.org/docs/2.2/mod/mod_auth_basic.html) is
useful when you want to protect a url with username/password.

See how-tos https://wiki.apache.org/httpd/PasswordBasicAuth

Basic usage exmaple:

```
<Location /subdir>
  AuthType Basic
  AuthName "Authentication Required"
  AuthUserFile "/etc/htpasswd/.htpasswd"
  Require valid-user
</Location>
```

Use the tool [htpasswd](http://httpd.apache.org/docs/2.4/programs/htpasswd.html)
to help manage users and passwords file `/etc/htpsswd/.htpasswd`.

You can also limit the IPs that do not need to input a password like:

```
  Order deny,allow
  Deny from all
  Allow from 1.2.3
  AuthType Basic
  AuthName "Authentication Required"
  AuthUserFile "/etc/htpasswd/.htpasswd"
  Require valid-user
  Satisfy Any
```
`Satisfy Any` option enables that client will be granted access if they either
pass the host restriction or enter a valid username and password. This can be
used to password restrict an area, but to let clients from particular addresses
in without prompting for a password.
See http://httpd.apache.org/docs/2.2/mod/core.html#satisfy
