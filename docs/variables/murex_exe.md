# `MUREX_EXE` (path)

> Absolute path to running shell

## Description

`MUREX_EXE` is very similar to the `$SHELL` environmental variable in that it
holds the full path to the running shell. The reason for defining a reserved
variable is so that the shell path cannot be overridden.

This is a [reserved variable](/docs/user-guide/reserved-vars.md) so it cannot be changed.

## See Also

* [Define Variable: `set`](../commands/set.md):
  Define a variable (typically local) and set it's value
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [`SHELL` (str)](../variables/shell.md):
  Path of current shell
* [`path`](../types/path.md):
  Structured object for working with file and directory paths

<hr/>

This document was generated from [gen/variables/MUREX_EXE_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/variables/MUREX_EXE_doc.yaml).