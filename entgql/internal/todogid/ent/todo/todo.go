// Copyright 2019-present Facebook
//
// Licensed x under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package todo

const (
	// Label holds the string label denoting the todo type in the database.
	Label = "todo"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the todo in the database.
	Table = "todos"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "todos"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "todo_owner"
)

// Columns holds all SQL columns for todo fields.
var Columns = []string{
	FieldID,
	FieldText,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "todos"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"todo_owner",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}