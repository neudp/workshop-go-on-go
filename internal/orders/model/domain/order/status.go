package order

type Status string

const (
	StatusCreated   Status = "created"
	StatusConfirmed Status = "confirmed"
	StatusPaid      Status = "paid"
	StatusShipped   Status = "shipped"
	StatusDelivered Status = "delivered"
	StatusCompleted Status = "completed"
	StatusReturned  Status = "returned"
	StatusCanceled  Status = "canceled"
)

func (status Status) CanTransitionTo(newStatus Status) bool {
	switch status {
	case StatusCreated:
		return newStatus == StatusConfirmed || newStatus == StatusCanceled
	case StatusConfirmed:
		return newStatus == StatusPaid || newStatus == StatusCanceled
	case StatusPaid:
		return newStatus == StatusShipped || newStatus == StatusCanceled
	case StatusShipped:
		return newStatus == StatusDelivered || newStatus == StatusReturned
	case StatusDelivered:
		return newStatus == StatusCompleted || newStatus == StatusReturned
	case StatusCompleted:
		return newStatus == StatusReturned
	case StatusReturned:
		return newStatus == StatusCanceled
	case StatusCanceled:
		return false
	default:
		return false
	}
}

func (status Status) IsFinal() bool {
	return status == StatusCompleted || status == StatusReturned || status == StatusCanceled
}

func (status Status) IsInitial() bool {
	return status == StatusCreated
}

func (status Status) IsImmutable() bool {
	return status.IsFinal() || status == StatusShipped || status == StatusDelivered
}

func (status Status) String() string {
	return string(status)
}
