# UnDone Server

This provides a self-hostable server for the [UnDone][undone] App.
With this server you can sync the Tasks across devices and Platforms.

## Setup

The first thing you will need to do is to download and install
this server. There are two options:

- cloning the repository and compiling all of the code yourself
- downloading a fitting version for your OS from the [release page][release]

The next thing you need to do is create a password for the server.
This password can be changed at any time, but can not be restored,
meaning when you forget the password the only option is to set a new
one.

You can set the password with the provided `manager`. If you cloned
the repository you can find the `manager` in `manager/bin/manager.dart`.
You'll need to either run or compile it with dart.

In the release version you will find the `manager` executable in the
same directory as the `server` executable. The usage of the cli tool
is self explanatory, if you need help use `manager --help`.

**If you do not set a password the server will not run.**

[undone]: ""
[release]: ""