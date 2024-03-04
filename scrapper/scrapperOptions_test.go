package scrapper

// Compiler check !
var _ ScrapperOption = SOHeadless(true)
var _ ScrapperOption = SOIgnore("un", "deux")
