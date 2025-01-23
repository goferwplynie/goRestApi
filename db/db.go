package db

import(
  "context"
  "time"
  "github.com/jackc/pgx/v5/pgxpool"
  "log"
)

//connection pool for db
//DB variable for no need to open new connections
var DB *pgxpool.Pool

func ConnectToDB(connString string){
  var err error
  //context ensures that db connection won't last forever
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  //if ConnectToDB function ends it cancels context
  defer cancel()

  //creating connection
  //context is required
  DB, err = pgxpool.New(ctx, connString)

  if err != nil{
    log.Fatalf("Can't connect to db: %v\n", err)
  }
}

func CloseDB(){
  if DB != nil{
    DB.Close()
  }
}
