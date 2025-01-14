# frontend-trigo

Template TRIGO WEB APP. 


## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Customize configuration

See [Vite Configuration Reference](https://vite.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Compile and Minify for Production

```sh
npm run build
```
# frontend-trigo

L'application possède deux API afin qu'il puisse fonctionner, une API pour vérifier l'authentification (SSO) et une autre pour tous les traitements avec la base de données. A savoir qu'il faut ajouter dans le dossier SSO, le fichier .env.dev afin de pouvoir accéder à l'API d'authentification.

### Commande de lancement entière
# Pour l'application VUEJS : 
```sh
npm run dev
```
# Pour l'API qui gère les traitements avec la base de données (en étant dans le dossier root du projet)
```sh
cd backend 
go run main.go
```
# Pour l'API qui gère l'authentification de l'application  (en étant dans le dossier root du projet)
```sh
cd SSO
npm start
```

