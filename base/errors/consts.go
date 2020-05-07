package errors

var (
	BadRequest            = Error{400, "BadRequest", "Bad Request"}
	InvalidParameter      = Error{400, "InvalidParameter", "Parameters in a request is invalid, unsupported, or cannot be used."}
	InvalidQueryParameter = Error{400, "InvalidQueryParameter", "The query string invalid"}
	MalformedParameter    = Error{400, "MalformedParameter", "Parameters in a request is invalid, contains a syntax error, or cannot be decoded."}
	Unauthorized          = Error{401, "Unauthorized", "Unauthorized"}
	AccessDenied          = Error{403, "AccessDenied", "Access Denied"}
	AccessConflict        = Error{409, "AccessConflict", "Access Conflict"}
	InternalError         = Error{500, "InternalError", "An internal error has occurred. Retry your request, but if the problem persists, contact us with details by posting a message on the service forums."}
)
