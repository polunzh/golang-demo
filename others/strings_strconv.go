package main

import (
  "fmt"
  "strings"
  "strconv"
)

func main() {
  var str string = "This This This This is an example of string张"
  prefix := "This"
  suffix := "ing"
  contains := "amp"
  not_contains := "zhang"
  character := "a"
  utf8Str := '张'
  messy := " dogggy "

  fmt.Printf("T/F? Does the string \"%s\" have prefix %s? %t\n", str, prefix, strings.HasPrefix(str, prefix))
  fmt.Printf("T/F? Does the string \"%s\" have suffix %s? %t\n", str, suffix, strings.HasSuffix(str, suffix))
  fmt.Printf("T/F? Does the string \"%s\" contains %s? %t\n", str, contains, strings.Contains(str, contains))
  fmt.Printf("Index of \"%s\" in \"%s\": %d\n", contains, str, strings.Index(str, contains))
  fmt.Printf("Index of \"%s\" in \"%s\": %d\n", not_contains, str, strings.Index(str, not_contains))
  fmt.Printf("First Index of \"%s\" in \"%s\": %d\n", character, str, strings.Index(str, character))
  fmt.Printf("Last Index of \"%s\" in \"%s\": %d\n", character, str, strings.LastIndex(str, character))
  fmt.Printf("Last Index of \"%s\" in \"%s\": %d\n", character, str, strings.LastIndex(str, character))
  fmt.Printf("Index of \"%s\" in \"%s\": %d\n", utf8Str, str, strings.IndexRune(str, rune(utf8Str)))
  fmt.Printf("Replace first 5 characters of \"%s\": %s\n", str, strings.Replace(str, "This", "GAI", 1))
  fmt.Printf("Count of \"%s\" in \"%s\": %d\n", prefix, str, strings.Count(str, prefix))
  fmt.Printf("Repeat \"%s\": \"%s\"\n", not_contains, strings.Repeat(not_contains, 5))
  fmt.Printf("\"%s\" to lower: \"%s\"\n", prefix, strings.ToLower(prefix))
  fmt.Printf("\"%s\" to upper: \"%s\"\n", prefix, strings.ToUpper(prefix))
  fmt.Printf("\"%s\" to upper: \"%s\"\n", prefix, strings.ToUpper(prefix))
  fmt.Printf("Trim spaces of \"%s\": \"%s\"\n", messy, strings.TrimSpace(messy))
  fmt.Printf("Trim left spaces of \"%s\": \"%s\"\n", messy, strings.TrimLeft(messy, " "))
  fmt.Printf("Trim right spaces of \"%s\": \"%s\"\n", messy, strings.TrimRight(messy, " "))
  fmt.Printf("Trim \"gy\" of \"%s\": \"%s\"\n", messy, strings.TrimRight(messy, "gy "))
  fmt.Printf("Split \"%s\" with space:\n", str);
  sl := strings.Fields(str)
  for _, val := range sl {
    fmt.Printf("%s - ", val)
  }
  fmt.Println()
  fmt.Printf("Joined: %s\n", strings.Join(sl, "|"))
  digitStr := "12"
  digit := 12
  val, _ := strconv.Atoi(digitStr)
  fmt.Printf("Atoi: \"%s\" => %d\n", digitStr, val)
  fmt.Printf("Itoa: %d => \"%s\"", digit, strconv.Itoa(digit))
}
