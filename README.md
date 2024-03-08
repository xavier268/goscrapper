# goscrapper
High level librairy for scrapping the web in go. 

Design goals :
* simplicity of use
* useable during discovery/design phase of a scrapper
* built-in concurrency
* easy to debug ( design phase, or later, when web site changes).
* os independant (linux or windows)
* default configurations should make sense, "convention over configuration" principle.
* compatible with sites using : http(s), js, ajax, authentication, ...

Architecture decisions taken :
* strictly follow semver versionning
* browser based
* rod based, but should not be visible directly
* as "declarative" as possible
* no implicit time outs
* extensive use fo context.Context **(to be implemented)**
* clear separation between capture of information (in scope) and processing of information (out of scope) via channels
* control resource utilization (timeouts, memory leaks, number of threads or of browser/tab instances, ...) via pagePool and WaitGroup
* build pipeline managed with "task"
* configuration driven (a yaml file)
* concurrency through multiple tabs in same browser (shared cookies) by default
* needed boiler plate or bindings should be autogenerated

Open choices :
* avoid init() as much as possible !!
* hijacking implemenation ?
* avoid loading images for performance (as option)
* verbose and detailled logging buit-in
* testable, included tests and examples

TODO / REWORK :

Redefine State, Job and Crapper Object
  * Scrapper is the top level object. It creates a job. 
  * The Job maintains a State value (from the configuration name, *meaning the state name should not be normalized ?*), a page, and a pointer to the scrapper.
    * The State in this case, becomes just a string, or an type int constant/enum for efficiency ?
  * When the Job runs, it performs Actions.
    * Its action could involve emitting data thru channels (Buses)
      * implementation is a forever loop, looking for and applying the States's actions. 
      * actions are implemented as func(*Job) ActionXX(p Parameter) and modify the Job internal state.
    * It may fork another Job, and continue its actions in parallele, 
    * It may move to another State, leaving current State and staying in same thread.
    * Before moving or forking to another State, the current Job evaluates the preconditions associated with the target State to ensure they are met.
      * There could be multiple fork/move targets associated with a single action : the first meeting its own preconditions will be used, the other will be ignored.
  * When the last Job has finished running, Scrapper ends its Run.
  * Page pooling seems attractive. Careful handling of self created pages by the web site ?
  * Action parameters as interface{}to allow for nesting (eg : action with conditions, hijack, foreach, oneof, ... => plus besoin de preconditions conditions de Stet ?)
  
Actions includes :
  * Noop (just use to wait for channel input and discard it)
  * Wait for duration
  * Select a reduced scope / reset to full page scope
  * Interactions : click, enter, ...
  * Hijack in or out data around an action ( implement as "startHiJack", action, "endHiJack" to avoid nested actions ?)
  * Emit ... to channel
  * Snapshots
  * Wait for input from channel - there should be a syntax for any string parameter (text, selector, content,urls,  ...) to indicate indirection from channel or from defined constants.
    * Ex : @@@NameOfSource@@@ is replaced with the bus content or the constant value ? (If Bus and Constants keys should never be identicals **to verify**, no need to specify if Bus or Define)
  * Load url (possibly new page ~ think about forking or closing previous page)
  * Fork to State(s) (with conditions ...)
  * Move to State(s) (with conditions ...)
  * Some kind of foreach loop to run a set of actions, defined in a State, on each Selection, while staying in same State ?
    * implement as Start for each , action1, action2, ..., end foreach to avoid nested actions. This will prevent nesting for loops in single State, but can be done using multiple States ...