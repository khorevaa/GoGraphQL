package jwtType

import (
	"github.com/graphql-go/graphql"
	"github.com/NiciiA/GoGraphQL/domain/type/accountType"
	"github.com/NiciiA/GoGraphQL/domain/model/accountModel"
)

var Type *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
	Name:        "JWT",
	Description: "A JWT response",
	Fields: graphql.Fields{
		"jwt": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if account, ok := p.Source.(accountModel.Account); ok {
					return account.UserName, nil
				}
				return nil, nil
			},
		},
		"account": &graphql.Field{
			Type: accountType.Type,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if account, ok := p.Source.(accountModel.Account); ok {
					return account, nil
				}
				return nil, nil
			},
		},
	},
})