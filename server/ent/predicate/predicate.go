// Code generated by ent, DO NOT EDIT.

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// Transaction is the predicate function for transaction builders.
type Transaction func(*sql.Selector)

// TransactionOrErr calls the predicate only if the error is not nit.
func TransactionOrErr(p Transaction, err error) Transaction {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}