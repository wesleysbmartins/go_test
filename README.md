# Golang Test
Aplicação desenvolvida para fins didáticos, onde é explorado o pacote de teste nativo de Golang e aplicados conceitos de Testes Unitários, Testes de Integração e Testes Automatizados.

## Testing Package
Utilizar o pacote de teste nativo do Go é bem simples, por exexplo:

Suponha que voce tenha um método ou função na qual deseja testar:
```go
// internal/math.go
package math

func Sum(a int, b int) int {
    return a + b
}
```

Para testar esta função voce deve:
 * Criar um arquivo com o a diretriz **_test.go** no final.
 * Importar o pacote **testing**.
 * Referenciar a sua função de teste passando **testing** como um parametro, por exemplo **func TestXXX(t *testing.T)**.
```go
// internal/math_test.go
package math

import "testing"

func TestSum(t *testing.T) {
    a := 1
    b := 1
    expected := 2

    result := Sum(a, b)

    if result != expected {
        msgErr := fmt.Sprintf(
            "ERROR\nOperation %s + %s\nExpected: %s\nResult: %s",
            strconv.Itoa(a),
            strconv.Itoa(b),
            strconv.Itoa(expected),
            strconv.Itoa(result)
        )

        t.Error(msgErr)
    } else {
        fmt.Println("SUCCESS")
    }
}
```

Para executar os testes você deve rodar em seu terminal:
```shell
user@abc:~$ go test
```
