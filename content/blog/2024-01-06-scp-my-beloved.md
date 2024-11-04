+++
title = "SCP, my beloved"
date = 2024-01-06
toc = true
[taxonomies]
tags = [ "linux", "scp", "ssh", "tech" ]
+++

# Introduction

Have you ever struggled to find the perfect tool for transferring files, which is easy to use and broadly available? If so, you are not alone, and I think I might have found the solution called "scp."

# What is SCP?

SCP is a program shipped with OpenSSH, allowing you to copy files between two hosts over a network. Since it's shipped with OpenSSH, it should be available on most of your computers, making it convenient to use. For SCP to work, you need to point it to a running SSH server. Now let's look at some examples of how to use SCP.

# How to use SCP:

Copying a local file to a remote host:  
```bash
scp localFile user@remotehost:path/to/target
``` 
And in reverse:  
```bash
scp user@remotehost:path/to/remotefile /path/to/target
```

To transfer whole folders, just point to a directory and use the `-r` flag to let scp run recursive.

In short, SCP simplifies file transfers with its straightforward approach. Next time you're stuck, remember SCP—it's a lifesaver.
