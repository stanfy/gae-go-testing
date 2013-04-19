#!/usr/bin/env python
# _*_ coding: utf-8 _*_

# Run this script to prepare everything:
#   curl https://raw.github.com/tenntenn/gae-go-testing/master/setup.py | python 

import os
import sys
import re

def main():

    # Check $APPENGINE_SDK
    if os.environ.get("APPENGINE_SDK") is None:
        print >>sys.stderr,"""Error: Please set path to \
                              GAE distribustion as APPENGINE_SDK such as \
                              'export APPENGINE_SDK=/usr/local/google_appengine/'."""
        return

    # Check $GOROOT
    if os.environ.get("GOROOT") is None:
        print >>sys.stderr, """Error: Please set path to \
                               Go distribustion as GOROOT such as \
                               'export GOROOT=/usr/local/go/'."""
        return

    # Check $PATH
    if os.environ.get("PATH") is None \
        or re.search(os.environ.get("APPENGINE_SDK"),
                 os.environ.get("PATH")) is None:

        print >>sys.stderr, """Error: Please add $APPENGINE_SDK to \
                               $PATH such as 'export PATH=$PATH:$APPENGINE_SDK'."""
        return

    # Check previous version
    packages = ["appengine", "appengine_internal", "code.google.com/p/goprotobuf"]
    for pkg in packages:
        dst = "{0}/src/pkg/{1}".format(os.environ.get("GOROOT"), pkg)
        if os.path.exists(dst):
            print >>sys.stderr, "Error: {0} already exists".format(dst)
            return 

    # Copy appengine to go distribustion
    for pkg in packages:
        src = "{0}/goroot/src/pkg/{1}/*".format(os.environ.get("APPENGINE_SDK"), pkg)
        dst = "{0}/src/pkg/{1}".format(os.environ.get("GOROOT"), pkg)
        os.makedirs(dst)
        cmd = "cp -r {0} {1}".format(src, dst)
        print cmd
        os.system(cmd)

    # Fix appengine internals
    internalsFile = "{0}/src/pkg/appengine_internal/api_dev.go".format(os.environ.get("GOROOT"))
    print "Fixing {0}...".format(internalsFile)
    fixedFile = open(internalsFile + ".tmp", "w")
    fixedFile.write(re.sub(r'(os\.DisableWritesForAppEngine.+?)\}', r'/* \1 */ }', open(internalsFile).read()))
    fixedFile.close()
    os.rename(internalsFile + ".tmp", internalsFile)

    # Install our packages
    for pkg in ["appenginetestinit", "appenginetesting"]:
        cmd = "{0}/bin/go get github.com/stanfy/gae-go-testing/{1}".format(os.environ.get("GOROOT"), pkg)
        print cmd
        os.system(cmd)

if __name__ == "__main__":
    main()
