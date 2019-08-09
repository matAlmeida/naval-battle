# shipbattle

Trabalho da turma de `Inteligência Artificial` ministrada pelo professor `Álvaro Degas`.

Feito por `Áurelio Chausse`, `Gabriel Figueiredo` e `Matheus Almeida`.

## Rodando

O projeto está sendo desenvolvido na linguagem `Go 1.x`

```sh
$ go version
go version go1.12.7 linux/amd64
$ go get -u github.com/matalmeida/shipbattle
$ cd $GOPATH/src/github.com/matalmeida/shipbattle
$ go run main.go
```

## Especificações do Projeto

O projeto consistirá em montar um algoritmo que jogue batalha naval contra um humano.

A avaliação será feita colocando-se um programa para jogar contra outro, com o auxílio de tabuleiros que serão montado a lápis no quadro em sala de aulas.

Um sorteio definirá as chaves, e as equipes se enfrentarão em partidas mata-mata. O campeão receberá nota 10 no projeto.

Os demais receberão notas proporcionais às demonstrações de ineficiência de seu algoritmo.

### O jogo.

Cada programa deverá dispor a sua própria esquadra num tabuleiro. A esquadra é composta conforme a figura abaixo. São 5 naves diferentes, sendo 1 porta-aviões, 2 cruzadores, 3 destroyers, 4 submarinos e 4 hidroaviões

![Tipos de naves](https://sites.google.com/site/alvarodegas/_/rsrc/1468757393193/home/academico/disciplinas/old-but-gold/2012-1/ia/projeto-2012-1/esquadra.png?height=175&width=200)

O tabuleiro onde cada nave deve ser alocada mede 10x10, conforme a figura abaixo.

![Disposição do campo](https://sites.google.com/site/alvarodegas/_/rsrc/1468757393254/home/academico/disciplinas/old-but-gold/2012-1/ia/projeto-2012-1/tabuleiro.png?height=197&width=200)

Cada nave pode ser disposta em qualquer posição possível.

### Restrições.

1. As naves não devem ser adjacentes umas às outras, seja na horizontal, seja na vertical, seja na diagonal.

2. Todas as naves devem ser posicionadas.

Um exemplo de disposição possível das peças pode ser visto na figura abaixo.

![Exemplo de disposição das naves](https://sites.google.com/site/alvarodegas/_/rsrc/1468757393482/home/academico/disciplinas/old-but-gold/2012-1/ia/projeto-2012-1/jogo_pronto.png?height=196&width=200)

## Como jogar.

Inicialmente cada programa vai dispor sua própria esquadra no tabuleiro, seguindo as restrições acima.

Cada equipe, a primeira delas deve ser definida em sorteio, vai tentar afundar a esquadra adversária. Alternadamente, cada programa deverá escolher três posições para atirar, e deverá receber o resultado de seu ataque. Cada posição atacada deve ser identificada pelas suas coordenadas: uma letra no intervalo de A a J e um número no intervalo de 1 a 10.

Um ataque pode atingir parte de uma nave, ou simplesmente a "água", que significa que o ataque não logrou êxito. Ao ser informado das coordenadas de um ataque, o programa deverá produzir uma informação a respeito de qual tipo de nave (porta-aviões, cruzador, destroyer, submarino e hidroavião) foi atingida. Caso uma nave tenha sido destruída inteiramente (sua última parte), isto deverá ser igualmente relatado.

Após receber a informação do resultado de seu ataque, que pode ser "água" ou parte de alguma nave inimiga (porta-aviões, cruzador, destroyer, submarino e hidroavião), o programa deve decidir sua estratégia a respeito.

Vence o jogo aquele que destruir inteiramente a esquadra adversária. O programa deverá se auto-declarar vencedor ou perdedor, caso isso aconteça.
