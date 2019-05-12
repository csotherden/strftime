# strftime
[![License](https://img.shields.io/github/license/csotherden/strftime.svg)](https://github.com/csotherden/strftime/blob/master/LICENSE)
[![Build Status](https://travis-ci.com/csotherden/strftime.svg?branch=master)](https://travis-ci.com/csotherden/strftime)
[![GoDoc](http://godoc.org/github.com/csotherden/strftime?status.svg)](http://godoc.org/github.com/csotherden/strftime)
[![Coverage Status](https://codecov.io/gh/csotherden/strftime/branch/master/graphs/badge.svg?branch=master)](https://codecov.io/gh/csotherden/strftime/)
[![Go Report Card](https://goreportcard.com/badge/github.com/csotherden/strftime)](https://goreportcard.com/report/github.com/csotherden/strftime)

##A Go package for strftime formatting

Package strftime provides functionality for formatting time based on the C strftime(3) function. http://man7.org/linux/man-pages/man3/strftime.3.html

The Go standard library time package uses layout strings based on the reference time Mon Jan 2 15:04:05 MST 2006
(AKA 01/02 03:04:05PM ‘06 -0700). The purpose of the strftime package is not to replace the time package layout
based formatting or encourage Go developers to punt on learning to use the reference time. Instead this package
exists for those cases where it's necessary to use strftime conversion specifications, such as when receiving
timestamps from non Go systems which provide a format description using strftime conventions.

### Conversion Specifications

This package attempts to comply with the C strftime(3) function closely as reasonably possible. The format
specification strings contain special character sequences called conversion specifications. Conversion specifications
are prefixed by the % character. At this time all conversion specifications are supported with the exception of
modifiers.

### Localization

Support for localization is currently mixed. Conversion specifications for individual time fields should be fully
localized. However, conversion specifications that map to different timestamps formats depending on locale (for
example %c, %x, and %X) only support the en_US localization.