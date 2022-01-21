# go-mysql

## go-mysql查询


```
type Users struct {
    ID            int64  `field:"id"`                      
    Username      string `field:"username"`           
    Password      string `field:"password"`           
    Tel           string `field:"tel"`                   
}
```

```
rows, err := db.Query(sql)  // select * from users
if err != nil {
    fmt.Println(err)
}
defer rows.Close()
for rows.Next() {
    user := new(Users)

    // works but I don't think it is good code for too many columns
    err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Tel)

    // TODO: How to scan in a simple way 


    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("user: ", user)
    list = append(list, *user)
}
if err := rows.Err(); err != nil {
    fmt.Println(err)
}
```
