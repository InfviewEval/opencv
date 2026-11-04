package acceptance

import ("context";"database/sql")

func Lookup(ctx context.Context, db *sql.DB, name string) (string,error) {
 var email string
 err := db.QueryRowContext(ctx, "SELECT email FROM users WHERE name = ?",name).Scan(&email)
 return email,err
}
