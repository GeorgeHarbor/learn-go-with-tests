package main

import "testing"

func TestHello(t *testing.T)  {
  
  got := Hello("Luka")
  want := "Hello, Luka"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}

