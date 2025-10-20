# Power4 Web 🎮

Un jeu de Puissance 4 (Connect Four) développé en Go avec une interface web

### Démarrage du serveur

Lancez le serveur avec la commande :
```bash
go run .
```

Ou alternativement :
```bash
go run main.go
```

Le serveur démarre par défaut sur **http://localhost:8080**

### Accès au jeu

Ouvrez votre navigateur et accédez à :
```
http://localhost:8080
```

## 🎯 Fonctionnalités

- **3 modes de difficulté** :
  - Normal (6x7)
  - Difficile (6x9)
  - Expert (7x8)
- Noms personnalisables pour les joueurs

## 🎮 Comment jouer

1. Entrez les noms des joueurs
2. Choisissez une difficulté
3. Cliquez sur "Commencer la partie"
4. Les joueurs jouent à tour de rôle en cliquant sur les colonnes
5. Le premier à aligner 4 jetons gagne !

