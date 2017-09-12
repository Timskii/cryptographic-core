package core

type Jsonobjects struct {
	objects	[]Jsonobject
}

type Jsonobject struct {
    Jsonrpc	string		`json:"jsonrpc"`
	Method	string		`json:"method"`
	Params	ParamsType	`json:"params"`
	Id	int			`json:"id"`
}

type ParamsType struct{
	Type int						`json:"type"`
	ChaincodeID 	ChaincodeIDType	`json:"chaincodeID"`
	CtorMsg			CtorMsgType		`json:"ctorMsg"`
}

type ChaincodeIDType struct{
	Name string			`json:"name"`
	}

type CtorMsgType struct {
	Function string		`json:"function"`
	Args []string		`json:"args"`
}

type Payload struct {
	ChaincodeSpec ChaincodeSpec `json:"chaincodeSpec"`
}

type ChaincodeSpec struct {
	ChaincodeID ChaincodeID		`json:"chaincodeID"`
	CtorMsg CtorMsgType			`json:"ctorMsg"`
}

type ChaincodeID struct {
	Path string 
	Name string
}


