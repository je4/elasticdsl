package dsl

type QueryOperator string

const (
	QueryOperatorOR  QueryOperator = "OR"
	QueryOperatorAND QueryOperator = "AND"
)

type QueryZeroTermsQuery string

const (
	QeruyZeroTermsQueryNONE QueryZeroTermsQuery = "none"
	QueryZeroTermsQueryALL  QueryZeroTermsQuery = "all"
)

type QueryMultiMatchType string

const (
	QueryMultiMatchTypeBestFields   QueryMultiMatchType = "best_fields"
	QueryMultiMatchTypeMostFields   QueryMultiMatchType = "most_fields"
	QueryMultiMatchTypeCrossFields  QueryMultiMatchType = "cross_fields"
	QueryMultiMatchTypePhrase       QueryMultiMatchType = "phrase"
	QueryMultiMatchTypePhrasePrefix QueryMultiMatchType = "phrase_prefix"
	QueryMultiMatchTypeBoolPrefix   QueryMultiMatchType = "bool_prefix"
)

type QuerySimpleQueryStringFlag string

const (
	QuerySimpleQueryStringFlagALL        QuerySimpleQueryStringFlag = "ALL"
	QuerySimpleQueryStringFlagAND        QuerySimpleQueryStringFlag = "AND"
	QuerySimpleQueryStringFlagESCAPE     QuerySimpleQueryStringFlag = "ESCAPE"
	QuerySimpleQueryStringFlagFUZZY      QuerySimpleQueryStringFlag = "FUZZY"
	QuerySimpleQueryStringFlagNEAR       QuerySimpleQueryStringFlag = "NEAR"
	QuerySimpleQueryStringFlagNONE       QuerySimpleQueryStringFlag = "NONE"
	QuerySimpleQueryStringFlagNOT        QuerySimpleQueryStringFlag = "NOT"
	QuerySimpleQueryStringFlagOR         QuerySimpleQueryStringFlag = "OR"
	QuerySimpleQueryStringFlagPHRASE     QuerySimpleQueryStringFlag = "PHRASE"
	QuerySimpleQueryStringFlagPRECEDENCE QuerySimpleQueryStringFlag = "PRECEDENCE"
	QuerySimpleQueryStringFlagPREFIX     QuerySimpleQueryStringFlag = "PREFIX"
	QuerySimpleQueryStringFlagSLOP       QuerySimpleQueryStringFlag = "SLOP"
	QuerySimpleQueryStringFlagWHITESPACE QuerySimpleQueryStringFlag = "WHITESPACE"
)

type QueryRewrite string

const (
	QueryRewriteConstantScore         QueryRewrite = "constant_score"
	QueryRewriteScoringboolean        QueryRewrite = "scoring_boolean"
	QueryRewriteTopTermsBlendedFreqsN QueryRewrite = "top_terms_blended_freqs_N"
	QueryRewriteTopTermsBoostN        QueryRewrite = "top_terms_boost_N"
	QueryRewriteTopTermsN             QueryRewrite = "top_terms_N"
)

type QueryRelation string

const (
	QueryRelationINTERSECTS QueryRelation = "INTERSECTS"
	QueryRelationCONTAINS   QueryRelation = "CONTAINS"
	QueryRelationWITHIN     QueryRelation = "WITHIN"
)

type QueryExpandWildcards string

const (
	QueryExpandWildcardsAll    QueryExpandWildcards = "all"
	QueryExpandWildcardsOpen   QueryExpandWildcards = "open"
	QueryExpandWildcardsClosed QueryExpandWildcards = "closed"
	QueryExpandWildcardsHidden QueryExpandWildcards = "hidden"
	QueryExpandWildcardsNone   QueryExpandWildcards = "none"
)

type QuerySearchType string

const (
	QuerySearchTypeQueryThenFetch    QuerySearchType = "query_then_fetch"
	QuerySearchDFSTypeQueryThenFetch QuerySearchType = "dfs_query_then_fetch"
)

type SearchSuggestMode string

const (
	SearchSuggestModeAlways  SearchSuggestMode = "always"
	SearchSuggestModeMissing SearchSuggestMode = "missing"
	SearchSuggestModePopular SearchSuggestMode = "popular"
)

type ScoreMode string

const (
	ScoreModeAVG  ScoreMode = "avg"
	ScoreModeMAX  ScoreMode = "max"
	ScoreModeMIN  ScoreMode = "min"
	ScoreModeNONE ScoreMode = "none"
	ScoreModeSUM  ScoreMode = "sum"
)

type Order string

const (
	OrderASC  Order = "asc"
	OrderDESC Order = "desc"
)

type GapPolicy string

const (
	GapPolicySkip        GapPolicy = "skip"
	GapPolicyInsertZeros GapPolicy = "insert_zeros"
	GapPolicyKeepValues  GapPolicy = "keep_values"
)

type CollectMode string

const (
	CollectModeBreadthFirst CollectMode = "breadth_first"
	CollectModeDepthFirst   CollectMode = "depth_first"
)

type ExecutionHint string

const (
	ExecutionHintMap            ExecutionHint = "map"
	ExecutionHintGlobalOrdinals ExecutionHint = "global_ordinals"
)
