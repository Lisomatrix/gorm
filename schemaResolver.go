package gorm

type SchemaResolver interface {
	ResolveJoin(stmt *Statement, join string) string
}

type SimpleSchemaResolver struct{}

func (resolver *SimpleSchemaResolver) ResolveJoin(stmt *Statement, join string) string {
	return join
}
