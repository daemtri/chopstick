package fn

import (
	"github.com/src-d/go-mysql-server/sql"
)

type CurrentUser struct {
}

func (c *CurrentUser) Resolved() bool {
	return true
}

func (c *CurrentUser) IsNullable() bool {
	return false
}

func (c *CurrentUser) Children() []sql.Expression {
	return nil
}

func (c *CurrentUser) String() string {
	return "current_user()"
}

func (c *CurrentUser) Type() sql.Type {
	return sql.VarChar(50)
}

func (c *CurrentUser) Eval(ctx *sql.Context, row sql.Row) (interface{}, error) {
	return "user", nil
}

func (c *CurrentUser) WithChildren(e ...sql.Expression) (sql.Expression, error) {
	if len(e) != 0 {
		return nil, sql.ErrInvalidChildrenNumber.New(c, len(e), 0)
	}
	return c, nil
}

// NewIsBinary creates a new IsBinary expression.
func NewCurrentUser() sql.Expression {
	return &CurrentUser{}
}
