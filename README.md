gae-go-testing
==============

Testing library for Go App Engine, giving you an appengine.Context fake that forwards to a dev_appserver.py child process.
This library is based on https://github.com/tenntenn/gae-go-testing.
This library works on GAE/G 1.7.7.


Installation
-----

Make sure you have [appengine SDK](https://developers.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go) installed.

Run this script and set corresponding environment variables it asks for:

    curl https://raw.github.com/tenntenn/gae-go-testing/master/setup.py | python
This script will copy appengine, appengine_internal, and goprotobuf packages from GAE SDK to Go root and that install this library with commands

    go get github.com/stanfy/gae-go-testing/appenginetestinit
    go get github.com/stanfy/gae-go-testing/appenginetesting


Usage
-----

 * Import `github.com/stanfy/gae-go-testing/appenginetestinit` (making it to be the last in inports list) to your 
test file and call `appenginetestinit.Use()` from `init()` function. 

 * Create AppEngine context using `appenginetesting.NewContext`.
