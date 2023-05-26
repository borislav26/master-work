package error

var (
	DatabaseError     = AuthenticationError{Message: "Failed to execute database operation."}
	ServiceLayerError = AuthenticationError{Message: "Failed to execute operation on service layer."}
	APILayerError     = AuthenticationError{Message: "Failed to execute operation on API layer."}
)
