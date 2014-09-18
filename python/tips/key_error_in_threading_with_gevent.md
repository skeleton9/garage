http://stackoverflow.com/questions/8774958/keyerror-in-module-threading-after-a-successful-py-test-run
some error like: `Exception KeyError: KeyError(4427427920,) in <module 'threading' from '/System/Library/Frameworks/Python.framework/Versions/2.7/lib/python2.7/threading.pyc'> ignored`

It is indeed related to monkey-patching the threading module. In fact, I can easily trigger the exception by importing the threading module before monkey-patching threads. The following code will generate error:

	import threading
	import gevent
	import gevent.monkey
	gevent.monkey.patch_all()

	print "hello"


The code below will keep the error away
make sure threading is not loaded before monkey patch threading

	import sys
	# these two lines are needed...
	if 'threading' in sys.modules:
	    del sys.modules['threading']
	import gevent
	import gevent.monkey
	gevent.monkey.patch_all()

	print "hello again"
