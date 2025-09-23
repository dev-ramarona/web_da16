package mdlGlobal

import "github.com/golang-jwt/jwt/v5"

type MdlGlobalAllusrParams struct {
	Usrnme string `json:"usrnme,omitempty" bson:"usrnme,omitempty"`
	Psswrd string `json:"psswrd,omitempty" bson:"psswrd,omitempty"`
}

type MdlGlobalAllusrDtbase struct {
	Usrnme string   `json:"usrnme,omitempty" bson:"usrnme,omitempty"`
	Psswrd string   `json:"psswrd,omitempty" bson:"psswrd,omitempty"`
	Stfnme string   `json:"stfnme,omitempty" bson:"stfnme,omitempty"`
	Stfeml string   `json:"stfeml,omitempty" bson:"stfeml,omitempty"`
	Access []string `json:"access,omitempty" bson:"access,omitempty"`
	Keywrd []string `json:"keywrd,omitempty" bson:"keywrd,omitempty"`
}

type MdlGlobalAllusrInputs struct {
	Usrnme string `json:"usrnme,omitempty" bson:"usrnme,omitempty"`
	jwt.RegisteredClaims
}

type MdlGlobalAllusrTokens struct {
	Stfnme string   `json:"stfnme"`
	Usrnme string   `json:"usrnme"`
	Access []string `json:"access"`
	Keywrd []string `json:"keywrd"`
}

type MdlGlobalAllusrApplst struct {
	Prmkey string `json:"prmkey,omitempty" bson:"prmkey,omitempty"`
	Detail string `json:"detail,omitempty" bson:"detail,omitempty"`
}
