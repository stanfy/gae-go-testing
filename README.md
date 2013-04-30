appenginetesting
===============

Fork of [gae-go-testing](https://github.com/tenntenn/gae-go-testing) with a few minor changes:

- renamed for nicer import syntax (that IDEA's Go Plugin won't highlight as an error)
- added +build tags so that it compiles
- simplified install instructions. 

As of GAE 1.7.5, we now keep tags of the repository that are known to
be compatible with each GAE release. If you are not using the latest
GAE release, please use the associated tag.

Installation
------------

Set environment variables :

    $ export APPENGINE_SDK=/path/to/google_appengine
    $ export PATH=$PATH:$APPENGINE_SDK

Before installing this library, you have to install
[appengine SDK](https://developers.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go).
And copy appengine, appengine_internal and goprotobuf as followings :

    $ export APPENGINE_SDK=/path/to/google_appengine
    $ ln -s $APPENGINE_SDK/goroot/src/pkg/appengine
    $ ln -s $APPENGINE_SDK/goroot/src/pkg/appengine_internal


There is some incompatibility between 1.7.7 and go 1.0.3. You can fix
this by commenting out the following line in the file
*${APPENGINE\_SDK}/goroot/src/pkg/appengine\_internal/api_dev.go*:

    func init() { os.DisableWritesForAppEngine = true }
	
It should be:

	//func init() { os.DisableWritesForAppEngine = true }

You also need to change the top of func init():

	c := readConfig(os.Stdin)
	instanceConfig.AppID = string(c.AppId)
	instanceConfig.APIPort = int(*c.ApiPort)
	instanceConfig.VersionID = string(c.VersionId)
	instanceConfig.InstanceID = *c.InstanceId
	instanceConfig.Datacenter = *c.Datacenter

It should be something like, values don't really matter:
	instanceConfig.AppID = "testapp"
	instanceConfig.APIPort = 0
	instanceConfig.VersionID = "1.7.7"
	instanceConfig.InstanceID = "instanceid"
	instanceConfig.Datacenter = "instanceid"

This library can be installed as following :

    $ go get github.com/icub3d/appenginetesting


Usage
-----

The
[documentation](http://godoc.org/github.com/icub3d/appenginetesting)
has some basic examples.  You can also find complete test examples
within [gorca](https://github.com/icub3d/gorca)/(*_test.go). Finally,
[context_test.go](https://github.com/icub3d/appenginetesting/blob/master/context_test.go)
and
[recorder_test.go](https://github.com/icub3d/appenginetesting/blob/master/recorder_test.go)
show an example of usage.
