### Variables
#### Declaring Variables
There are two ways to declare variables in Go.

1. The `:=` shorthand.

    The `:=` operator is used to declare and initialize a variable in one command. For example: `message := fmt.Springf("Hi, %v.", name)` creates a new variable, `message`, which is then set immediately to a string.
2. The default method.

    By default, variables are declared using the following structure: `var <name> <type>`. For example, `var message string` creates a new variable called `message` which has a type of `string`. 