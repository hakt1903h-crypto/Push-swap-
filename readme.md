# push-swap

Projet d'algorithmie en Go consistant à trier une pile d'entiers en utilisant deux piles (`a` et `b`) et un ensemble limité d'instructions.

## Structure du projet

* `piles.go` : structure des piles, opérations autorisées, lecture des arguments et algorithme de tri.

* `push_swap.go` : programme principal affichant la suite d'opérations pour trier la pile `a`.

* `checker.go` : programme vérifiant si une suite d'opérations trie correctement la pile `a` avec `b` vide.

## Compilation

```
go build -o push-swap piles.go push_swap.go
go build -o checker piles.go checker.go

```

## Utilisation

Générer les instructions de tri :

```
./push-swap "4 67 3 87 23"

```

Vérifier le résultat avec `checker` :

```
ARG="4 67 3 87 23"
./push-swap "$ARG" | ./checker "$ARG"

```

## Opérations autorisées

* `sa`, `sb`, `ss` : échange les deux premiers éléments de `a`, `b` ou les deux.

* `pa`, `pb` : déplace le sommet d'une pile vers l'autre.

* `ra`, `rb`, `rr` : décale la pile vers le haut (le premier passe en dernier).

* `rra`, `rrb`, `rrr` : décale la pile vers le bas (le dernier passe en premier).