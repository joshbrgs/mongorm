package mongorm

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestModel_Create(t *testing.T) {
	type fields struct {
		ID        primitive.ObjectID
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		ctx            context.Context
		db             *mongo.Database
		collectionName string
		model          interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				ID:        tt.fields.ID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := m.Create(tt.args.ctx, tt.args.db, tt.args.collectionName, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("Model.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModel_Read(t *testing.T) {
	type fields struct {
		ID        primitive.ObjectID
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		ctx            context.Context
		db             *mongo.Database
		collectionName string
		filter         interface{}
		result         interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				ID:        tt.fields.ID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := m.Read(tt.args.ctx, tt.args.db, tt.args.collectionName, tt.args.filter, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("Model.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModel_Update(t *testing.T) {
	type fields struct {
		ID        primitive.ObjectID
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		ctx            context.Context
		db             *mongo.Database
		collectionName string
		filter         interface{}
		update         interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				ID:        tt.fields.ID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := m.Update(tt.args.ctx, tt.args.db, tt.args.collectionName, tt.args.filter, tt.args.update); (err != nil) != tt.wantErr {
				t.Errorf("Model.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModel_Delete(t *testing.T) {
	type fields struct {
		ID        primitive.ObjectID
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		ctx            context.Context
		db             *mongo.Database
		collectionName string
		filter         interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				ID:        tt.fields.ID,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := m.Delete(tt.args.ctx, tt.args.db, tt.args.collectionName, tt.args.filter); (err != nil) != tt.wantErr {
				t.Errorf("Model.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
