# Test assignment description

You are given a number of inputs and corresponding expressions for substitutions between inputs and outputs, and the expected output.

The algorithm should be implemented using any approaches/ techniques that you find appropriate.

Assignment can be implemented in any language of your choice, and you can use any framework / libraries that you wish. It’s up to you to decide how to control quality of implementation and what design/ approach to pick/ invent.

#### Possible inputs
Following variables are acceptable as input:
- A: bool
- B: bool
- C: bool
- D: float
- E: int
- F: int

#### Expected outputs
The outputs are defined as:

**H:** one of these predefined values [M,P,T] (e.g. H could be equal to either of 3 values: M, P or T)

**K:** floating point number (e.g. float, decimal)

#### Substitution expressions
The assignment consists of base expressions set and two custom set of expressions that override / extend the base rules.

**Base**
```
A && B && !C => H = M A&&B&&C =>H=P !A && B && C => H = T [other] => [error]
H = M => K = D + (D * E / 10)
H = P => K = D + (D * (E - F) / 25.5) H = T => K = D - (D * F / 30)
```

**Custom 1**
```
H = P => K = 2 * D + (D * E / 100)
```

**Custom 2**
```
A && B && !C => H = T
A && !B && C => H = M
H = M => K = F + D + (D * E / 100)
```

#### Result
Assignment could be implemented in a way of backend application and be exposed in a simple REST API or as browser based application and run within the browser with or without communication to the backend. Shape of implementation is up to you.


## Running

### With docker

If you have installed docker, use make-helpers.

#### Build and run the API:
```
make build docker up
```

#### Down:
```
make down
```

#### To look at all available commands:
```
make help
```

### Without docker
In terminal from the root of this repo:
```
APP_PORT=8080 go run ./
```

## Dependencies
- github.com/go-chi/chi
- github.com/TV4/graceful (server graceful shutdown)

## Annotation
...