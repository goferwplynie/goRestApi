package jsontools

import(
  "github.com/goferpwlynie/goRestApi/users"
  "os"
  "fmt"
  "io"
  "encoding/json"
)

func LoadFromJson(){
  file, err := os.Open("users.json")

  if err != nil{
    fmt.Println("error opening users.json: ", err)
    return
  }

  defer file.Close()

  byteValue, err := io.ReadAll(file)

  if err != nil{
    fmt.Println("error reading users.json: ", err)
    return
  }

  err = json.Unmarshal(byteValue, &users.Users)

  if err != nil{
    fmt.Println("error unmarshaling json: ", err)
    return
  }
}

func WriteJsonFile(){
  file ,err := os.Open("users.json")

  if err != nil{
    fmt.Println("error opening users.json: ", err)
    return
  }

  defer file.Close()

  updatedData, err := json.MarshalIndent(users.Users, "", " ")

  err = os.WriteFile("users.json", updatedData, 0644)

  if err != nil{
    fmt.Println("error writing to users.json: ", err)
    return
  }
}
