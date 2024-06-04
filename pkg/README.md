# /pkg:

While /internal holds code that cannot be imported into other applications, /pkg stores libraries that are used in third-party applications.
This is necessary to import them into another project instead of duplicating code from project to project.
Contents:
Primarily, custom or shared libraries are stored here, which can be useful in various projects.
Usage:

You may choose not to use this directory if the project is very small and adding a new level of nesting does not make practical sense.
