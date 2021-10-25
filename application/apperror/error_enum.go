package apperror

const (
	ERR500                            ErrorType = "ER500 %s"                                      // custom err500 message
	ERR400                            ErrorType = "ER400 %s"                                      // custom err500 messag
	FailUnmarshalResponseBodyError    ErrorType = "ER400 fail to unmarshal response body"         // used by controller
	ObjectNotFound                    ErrorType = "ER404 object %s is not found"                  // used by injected repo in interactor
	UnrecognizedEnum                  ErrorType = "ER500 %s is not recognized %s enum"            // used by enum
	DatabaseNotFoundInContextError    ErrorType = "ER500 database is not found in context"        // used by repoimpl
	JWTSecretKeyMustNotEmpty          ErrorType = "ER500 JWT secret key must not empty"           // used by repoimpl
	CannotGenerateID                  ErrorType = "ER500 cannot generate error"                   // used by repoimpl
	DateMustNotEmpty                  ErrorType = "ER400 date must not empty"                     // used by repoimpl
	AuditIDMustNotEmpty               ErrorType = "ER400 audit id must not empty"                 //
	RoomIDMustNotEmpty                ErrorType = "ER400 room id must not empty"                  //
	ItemIDMustNotEmpty                ErrorType = "ER400 item id must not empty"                  //
	ItemIsAlreadyScannedInThisSession ErrorType = "ER400 item is already scanned in this session" //
)
