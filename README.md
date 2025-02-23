# sat-verifier

---

## checks if a given boolean expression is satisfied by the given variables


## Build

    go build -o satverifier.exe main.go verifier.go


## Usage

    ./satverifier.exe 'x1 && (x2 && !x3 || x4)' '{"x1": "true", "x2": "false", "x3": "true", "x4": "true"}'

