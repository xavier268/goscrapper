# goscrapper
High level librairy for scrapping the web in go. 

Design goals :
* simplicity of use
* useable during discovery/design phase of a scrapper
* built-in concurrency
* easy to debug ( design phase, or later, when web site changes).
* os independant (linux or windows)
* testable, included tests and examples
* default configurations should make sense, "convention over configuration" principle.
* compatible with sites using : http(s), js, ajax, authentication, ...

Architecture decisions taken :
* strictly follow semver versionning
* browser based
* rod based, but should not be visible directly
* as "declarative" as possible
* verbose and detailled logging buit-in
* no implicite time outs
* extensive use fo context.Context
* clear separation between capture of information (in scope) and processing of information (out of scope)
* control resource utilization (timeouts, memory leaks, number of threads or of brwser/tab instances, ...)
* build pipeline managed with "task"

Open choices :
* avoid init() as much as possible !!
* configuration driven (a go struct or a yaml file ?)
* needed boiler plate or bindings should be autogenerated
* concurrency through multople tabs in same browser (shared cookies) by default
* hijacking implemenation ?
* avoid loading images for performance (as option)