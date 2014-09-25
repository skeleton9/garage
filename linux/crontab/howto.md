## Use Crontab

See more at https://help.ubuntu.com/community/CronHowto

`crontab -l` to display all the crontab jobs

Start the crontab editor fromo command line:

`crontab -e`

Add one command that you need to run:

`* * * * * echo 'hello $(date)'`

The `*`s are minute(0-59), hour(0-23), day(1-31), month(1-12), weekday(0-6, 0 = Sunday)

You can add `MAILTO=user@example.com` to let cron to send mail `user@example.com`. If you set `MAILTO=""`, no mail will be sent.

## Common Issues

### Job not running?

* Newline at the end for each job

When adding a new entry to a blank crontab, forgetting to add a newline at the end is a common source for the job not running. If the last line in the crontab does not end with a newline, no errors will be reported at edit or runtime, but that line will never run. See man crontab for more information. This has already been suggested as a bug.

* User should in /etc/shadow

Another reason that your crontab jobs are not running is because the user is not in /etc/shadow

### Environment for the command you run

Add environment setting in crontab file is like:

`name = value`

The spaces around `name` are `value` are optional.

So you can add `PATH` like:

`PATH=$HOME/bin:$PATH`

See http://unixhelp.ed.ac.uk/CGI/man-cgi?crontab+5

```
Several environment  variables are set up automatically by the cron daemon. 
SHELL is set to /bin/sh, and LOGNAME and HOME are set from the /etc/passwd 
line  of the crontab?s owner.  HOME and SHELL may be overridden by settings 
in the crontab; LOGNAME may not.
```
