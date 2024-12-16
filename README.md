# kringler

'tis the season for randomly generated lists

--

## Usage

### Initialising

```go

// Initialise empty kringler instance to add stuff to later
k := kringler.New()

// Initialise a kringler instance using a predefined list
k := kringler.NewFromList([]string{
  "Foo",
  "Bar",
})

// Initialise a kringler instance using a file, with each Kringler on a new line, e.g.
//   Foo
//   Bar
k := kringler.NewFromFile(path.Join("data", "kringlers.txt"))
```

### Shuffling

This step is optional if you already have a predefined list you want to work with.

```go
k.ShuffleKringlers()
log.Println(k.Kringlers)

// If you want a copy of the list, assign it to a variable
shuffledKringlers := k.ShuffleKringlers()
```

### Creating Assignments

This will create a unique slice of exclusive Giver and Receiver properties.

```go
k.CreateLinearAssignments()
log.Println(k.Assignments)

// If you want a copy of the list, assign it to a variable
assignments := k.CreateLinearAssignments()
```

### Using Assignments

```go
for _, assignment := range k.Assignments {
  fmt.Printf("%s -> %s", assignment.Giver, assignment.Receiver)
}
```
