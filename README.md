appenginetesting
===============

Testing library for Go App Engine, giving you an appengine.Context fake that forwards to a dev_appserver.py child process.
This library is fixed for go1 based on http://code.google.com/p/gae-go-testing/ .
This library works on GAE/G 1.7.0 or higher and go1 and tested on:

    * GAE/G 1.7.0, go 1.0.3
    * GAE/G 1.7.1, go 1.0.3

Installation
------------

Set environment variables :

    $ export APPENGINE_SDK=/path/to/google_appengine
    $ export PATH=$PATH:$APPENGINE_SDK

Before installing this library, you have to install [appengine SDK](https://developers.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go).
And copy appengine, appengine_internal and goprotobuf as followings :

    $ export APPENGINE_SDK=/path/to/google_appengine
    $ ln -s $APPENGINE_SDK/goroot/src/pkg/appengine
    $ ln -s $APPENGINE_SDK/goroot/src/pkg/appengine_internal


This library can be installed as following :

    $ go get github.com/icub3d/appenginetesting


Usage
-----

context\_test.go and recorder\_test.go show an example of usage.
