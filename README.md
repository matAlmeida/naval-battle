# Ship Battle

## Goals

- O projeto consistirá em montar um algoritmo que jogue batalha naval contra um humano.
- A avaliação será feita colocando-se um programa para jogar contra outro, com o auxílio de tabuleiros que serão montado a lápis no quadro em sala de aulas.
- Um sorteio definirá as chaves, e as equipes se enfrentarão em partidas mata-mata. O campeão receberá nota 10 no projeto.
- Os demais receberão notas proporcionais às demonstrações de ineficiência de seu algoritmo.

## The Game

Cada programa deverá dispor a sua própria `esquadra` num tabuleiro. A esquadra é composta conforme a figura abaixo.

- 1 porta-aviões
- 2 cruzadores
- 3 destroyers
- 4 submarinos
- 4 hidroaviões

![Desenho das Embarcações](https://sites.google.com/site/alvarodegas/_/rsrc/1468757393193/home/academico/disciplinas/old-but-gold/2012-1/ia/projeto-2012-1/esquadra.png?height=175&width=200)

O tabuleiro onde cada nave deve ser alocada mede 10x10, conforme a figura abaixo.

![Tabuleiro](https://sites.google.com/site/alvarodegas/_/rsrc/1468757393254/home/academico/disciplinas/old-but-gold/2012-1/ia/projeto-2012-1/tabuleiro.png?height=197&width=200)

Cada nave pode estar disposta em qualquer posição possível.

## Restrictions

- As embarcações não podem ser adjacentes umas às outras, seja na horizonta, vertical ou na diagonal.
- Todas as embarcações devem ser posicionadas.

![Embarcações posicionadas](https://sites.google.com/site/alvarodegas/_/rsrc/1468757393482/home/academico/disciplinas/old-but-gold/2012-1/ia/projeto-2012-1/jogo_pronto.png?height=196&width=200)

## How to Play

- Cada programa deve dispor sua própria esquadra no tabuleiro, respeitando as restrições.

### Atacando

- Cada equipe vai tentar afundar a esquadra adversária.
- Os ataques serão compostos de `três` tiros, que devem ser indentificados por um par ordenado de `(A-J, 1-10)`.
- Após o ataque o programa deve receber o resultado de seu ataque.

### Recebendo o Ataque

- O ataque pode atingir a parte de uma nave ou a água.
- Ao receber as cordenadas que sofreu o ataque o programa devera retornar informações a respeito de qual tipo de nave foi atingida por cada ataque.
- Caso uma nave tenha sido totalmente destruida isso também deverá ser informado.

### Após finalizar o ataque

- O programa recebe as informações sobre o seu ataque e decide a sua próxima ação.

## Winning

- Ganha o jogo quem destruir toda a esquadra adversária.
- O próprio programa deverá se declarar vencedor ou perdedor, caso um dos casos ocorra.
