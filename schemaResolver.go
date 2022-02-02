package gorm

import "context"

type SchemaResolver interface {
	ResolveTable(ctx context.Context, tableName string) string
}

type SimpleSchemaResolver struct{}

func (resolver *SimpleSchemaResolver) ResolveTable(ctx context.Context, tableName string) string {
	return tableName
}
