package utils

var GetAccountByNumberQuery = `
        SELECT 
            u.first_name,
            u.last_name,
            u.email,
            a.number,
            a.balance,
            a.password,
            a.admin
        FROM accounts a
        JOIN users u ON u.account_number = a.number
        WHERE number = ?
`

var GetAllAccountsQuery = `
        SELECT 
            u.first_name,
            u.last_name,
            u.email,
            a.number,
            a.balance,
            a.password,
            a.admin
        FROM accounts a
        JOIN users u ON u.account_number = a.number
`
