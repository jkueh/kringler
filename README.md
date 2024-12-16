# kringler

'tis the season for randomly generated lists!

I tried to be clever and write something myself, and failed to catch an error I made - So here's a version of that code,
now with tests!

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

### Generating Messages

Manually writing your Kris Kringle messages is a recipe for RSI, so let the library do it for you with some handy-dandy
templating!

Here's an example template stored in `data/template.txt` on the local filesystem:

```
Season's Greetings, {{ .Giver }}!

You have been ~randomly~ selected to participate in an exclusive, top-secret, invite-only Kris Kringle event.

You _should_ buy a present for: *{{ .Receiver }}*

After the events that transpired last year, I am legally obligated to remind you:
*Do not tell this person you are buying them a gift!*

Please note: The budget for your gift is $50.00, or the following local equivalents:
- MYR 140
- USD 30
- JPY 4,800
- EUR 30
- IDR 500,000

Failure to adhere to the budget may result in a lump (or lumps, depending on the extent of the breach) of coal in your stocking.

_Yours sincerely,_
_The Harbinger of Christmas_
_(Jordan)_
```

Pro-tip: Run your message through a spell checker - Don't be like me, sending out messages with the same typo.

```go
k := kringler.NewFromList([]string{"Foo", "Bar", "Baz"})
k.ShuffleKringlers()
k.CreateMessageWithTemplate(assignment, path.Join("data", "template.txt"), "messages")
```
