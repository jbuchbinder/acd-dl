# ACD-DL

[![Build Status](https://secure.travis-ci.org/jbuchbinder/acd-dl.png)](http://travis-ci.org/jbuchbinder/acd-dl)

Recursive downloader for shared content from Amazon Cloud Drive.

## SYNTAX

```
acd-dl [-debug] ID [ID [ID ... ]]
```

IDs are the alphanumeric identifier present in the share URL, so ``https://www.amazon.com/clouddrive/share/XXXXXXXXXXXXXXXXX222222222222XXXXXXXXXX?ref_=cd_ph_share_link_copy`` would yield an id of ``XXXXXXXXXXXXXXXXX222222222222XXXXXXXXXX``.

