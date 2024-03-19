# TODO list


## Redesign architecture to be more modular and manageable

* Design queries / program in a DomainSpecific Language
* Compile the queries into a go package
* Compile generated package, and link to it

## Langage definition

Inspired by the FERRET language : https://www.montferret.dev/docs/fql/ 

Contray to FERRET, the queries will be compiled into binary code by the golang compiler.

## Articulation between the queries and the main application program

* option 1: the main app will call the functions that have been given the names of the queries with the binding parameters ...
* option 2 : the queries are run asynchroneously, and results are retrieved from a channel ?