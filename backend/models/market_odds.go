package models

type Market struct {
	Id   string `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	FxId int    `bun:"fx_id,notnull" json:"fx_id"`
	Bet 
}
