// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "number", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "balance", Type: field.TypeFloat64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "admin", Type: field.TypeBool, Default: false},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "from_account_number", Type: field.TypeString, Nullable: true},
		{Name: "to_account_number", Type: field.TypeString, Nullable: true},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_accounts_from_account",
				Columns:    []*schema.Column{TransactionsColumns[1]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "transactions_accounts_to_account",
				Columns:    []*schema.Column{TransactionsColumns[2]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TransactionDetailsColumns holds the columns for the "transaction_details" table.
	TransactionDetailsColumns = []*schema.Column{
		{Name: "transaction_id", Type: field.TypeInt, Increment: true},
		{Name: "amount", Type: field.TypeFloat64},
		{Name: "type", Type: field.TypeString, Size: 20},
		{Name: "transacted_at", Type: field.TypeTime},
	}
	// TransactionDetailsTable holds the schema information for the "transaction_details" table.
	TransactionDetailsTable = &schema.Table{
		Name:       "transaction_details",
		Columns:    TransactionDetailsColumns,
		PrimaryKey: []*schema.Column{TransactionDetailsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transaction_details_transactions_detail",
				Columns:    []*schema.Column{TransactionDetailsColumns[0]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "user_account", Type: field.TypeString},
		{Name: "first_name", Type: field.TypeString, Size: 50},
		{Name: "last_name", Type: field.TypeString, Size: 50},
		{Name: "email", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_accounts_account",
				Columns:    []*schema.Column{UsersColumns[0]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		TransactionsTable,
		TransactionDetailsTable,
		UsersTable,
	}
)

func init() {
	TransactionsTable.ForeignKeys[0].RefTable = AccountsTable
	TransactionsTable.ForeignKeys[1].RefTable = AccountsTable
	TransactionDetailsTable.ForeignKeys[0].RefTable = TransactionsTable
	UsersTable.ForeignKeys[0].RefTable = AccountsTable
}