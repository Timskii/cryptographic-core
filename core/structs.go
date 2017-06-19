package core

import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type ConfidentialityLevel int32
type Transaction_Type int32

type Transaction struct{
	Type Transaction_Type
	ChaincodeID 					[]byte
	Payload 						[]byte
	Metadata						[]byte
	Txid							string
	Timestamp						*google_protobuf.Timestamp
	ConfidentialityLevel			ConfidentialityLevel       
	ConfidentialityProtocolVersion	string                     
	Nonce							[]byte                     
	ToValidators					[]byte                     
	Cert							[]byte                     
	Signature						[]byte       	
}


type ChaincodeID struct {
	Path string 
	Name string
}

