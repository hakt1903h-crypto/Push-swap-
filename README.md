# Push-swap

> Projet d'algorithmie en Go consistant à trier une pile

## 📋 Structure
piles.go : structure des piles, opérations autorisées, lecture des arguments et algorithme de tri.

push_swap.go : programme principal affichant la suite d'opérations pour trier la pile a.

checker.go : programme vérifiant si une suite d'opérations trie correctement la pile a avec b vide.

## ⚙️ Prérequis
Avoir le langage "GO" installé 

## 🚀 Utilisation et Verification
Pour lancer le programme, ouvrir le terminal et taper cette commande :

```bash
./push-swap "4 67 3 87 23"

Vérifier le résultat avec checker :

ARG="4 67 3 87 23"
./push-swap "$ARG" | ./checker "$ARG"

