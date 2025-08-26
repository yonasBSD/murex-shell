# If Conditional: `if`

> Conditional statement to execute different blocks of code depending on the result of the condition

## Description

Conditional control flow

`if` can be utilized both as a method as well as a standalone function. As a
method, the conditional state is derived from the calling function (eg if the
previous function succeeds then the condition is `true`).

## Usage

### Function `if`:

```
if { code-block } then {
    # true
} else {
    # false
}
```

### Method `if`

```
command -> if {
    # true
} else {
    # false
}
```

### Negative Function `if`:

```
!if { code-block } then {
    # false
}
```

### Negative Method `if`:

```
command -> !if {
    # false
}
```

### Please Note:
the `then` and `else` statements are optional. So the first usage could
also be written as:

```
if { code-block } {
    # true
} {
    # false
}
```

However the practice of omitting those statements isn't recommended beyond
writing short one liners in the interactive command prompt.

## Examples

Check if a file exists:

```
if { g somefile.txt } then {
    out "File exists"
}
```

...or does not exist (both ways are valid):

```
!if { g somefile.txt } then {
    out "File does not exist"
}
```

```
if { g somefile.txt } else {
    out "File does not exist"
}
```

## Detail

### Pipelines and Output

The conditional block can contain entire pipelines - even multiple lines of code
let alone a single pipeline - as well as solitary commands as demonstrated in
the examples above. However the conditional block does not output stdout nor
stderr to the rest of the pipeline so you don't have to worry about redirecting
the output streams to `null`.

If you require output from the conditional blocks stdout then you will need to
use either a Murex named pipe to redirect the output, or test or debug flags
(depending on your use case) if you only need to occasionally inspect the
conditionals output.

### Exit Numbers

When evaluating a command or code block, `if` will treat an exit number less
than 0 as true, and one greater than 0 as false. When the exit number is 0, `if`
will examine the stdout of the command or code block. If there is no output, or
the output is one of the following strings, `if` will evaluate the command or
code block as false. Otherwise, it will be considered true.

* `0`
* `disabled`
* `fail`
* `failed`
* `false`
* `no`
* `null`
* `off`

## Synonyms

* `if`
* `!if`


## See Also

* [Caught Error Block: `catch`](../commands/catch.md):
  Handles the exception code raised by `try` or `trypipe`
* [Debug Mode: `debug`](../commands/debug.md):
  Debugging information
* [False: `false`](../commands/false.md):
  Returns a `false` value
* [Logic And Statements: `and`](../commands/and.md):
  Returns `true` or `false` depending on whether multiple conditions are met
* [Logic Or Statements: `or`](../commands/or.md):
  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.
* [Not: `!`](../commands/not-func.md):
  Reads the stdin and exit number from previous process and not's it's condition
* [Pipe Fail: `trypipe`](../commands/trypipe.md):
  Checks for non-zero exits of each function in a pipeline
* [Shell Script Tests: `test`](../commands/test.md):
  Murex's test framework - define tests, run tests and debug shell scripts
* [Switch Conditional: `switch`](../commands/switch.md):
  Blocks of cascading conditionals
* [True: `true`](../commands/true.md):
  Returns a `true` value
* [Try Block: `try`](../commands/try.md):
  Handles non-zero exits inside a block of code

<hr/>

This document was generated from [builtins/core/structs/if_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/structs/if_doc.yaml).