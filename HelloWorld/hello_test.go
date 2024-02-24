package main

import "testing"

func TestHello(t *testing.T)  {
  
  t.Run("saying hello to people", func(t *testing.T){
    
  got := Hello("Luka")
  want := "Hello, Luka"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
  })
  
  t.Run("say 'Hello, world' when empty string is supplied", func(t *testing.T){
    got := Hello("")
    want := "Hello, world"

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })
}

