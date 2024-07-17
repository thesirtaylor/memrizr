package interfaces

import (
	"context"

	"github.com/google/uuid"

	"github.com/thesirtaylor/memrizr/model"
)

type UserService interface { //service contract for the user service
	//defines method the handler layer expects any service 
	//it interacts with to implement
	Get(ctx context.Context, uid uuid.UUID) (*model.User, error)
	// Post(ctx context.Context, user *model.User)  (*model.User,error)
}

type UserRepository interface { //repository contract for the user repository
	//defines methods the service layer expects any repository
	//it interacts with to implement
	FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error)
}