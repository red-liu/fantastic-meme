package firstmod

import "fmt"

func Hello(name string) string {
    
    return  fmt.Sprintf("Hello %s", name)

}

func HelloV2(name string) string{

   return fmt.Sprintf("<h1>Hello %s</h1>", name)
}
