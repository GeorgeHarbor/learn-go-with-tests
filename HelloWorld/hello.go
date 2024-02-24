package main

import (
	"fmt"
	"strings"
)

const (
  serbianHelloPrefix = "Zdravo, "
  englishHelloPrefix = "Hello, "
  spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
  if name == "" {
    name = "world"
  }
  if language == "" {
    language = "english"
  }
  
  return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string){
  language = strings.ToLower(language)
  
  switch language{
  case "spanish":
      prefix = spanishHelloPrefix
  case "serbian":
      prefix = serbianHelloPrefix
  default:
      prefix = englishHelloPrefix
}
  return
}

func main()  {
  fmt.Println(Hello("Luka", "sPaNish")) 
}

