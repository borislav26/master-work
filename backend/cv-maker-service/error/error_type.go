package error

var (
	DatabaseError     = CVMakerError{Message: "Failed to execute database operation."}
	ServiceLayerError = CVMakerError{Message: "Failed to execute operation on service layer."}
	APILayerError     = CVMakerError{Message: "Failed to execute operation on API layer."}
	RabbitMQError     = CVMakerError{Message: "Failed to send instruction to RabbitMQ."}
)
