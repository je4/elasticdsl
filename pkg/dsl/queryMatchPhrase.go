package dsl

type tMatchPhraseQuery tMatchBoolPrefixQuery

func (mp *tMatchPhraseQuery) GetQueryName() string { return "match_phrase" }
