Not sure what's going on here.

This repo:

  * Builds a static library (gofunction.a) containing an empty Go function
  * Links that into a C library which calls the Go function
  * Calls that from a main package

When you run the thing it crashes:

