# Vue 3 + TypeScript + Vite

This application is designed to provide a robust and efficient development environment using Vue 3, TypeScript, Go lang and POstgresSQL and Vite. It leverages modern front-end tools and frameworks to deliver a seamless user experience.

![App Preview](https://i.imgur.com/DdK8vNv.png)

## Purpose of the Application

The purpose of this application is to host a site for our client to easily list items for sale.  The site can be used to reserve items to be bought online or in person. 

## Tools and Technologies Used

- **Vue 3**: A progressive JavaScript framework for building user interfaces.
- **TypeScript**: A strongly typed programming language that builds on JavaScript.
- **Vite**: A next-generation front-end tooling for faster and leaner development.
- **Preline**: Utility-first CSS framework.
- **Tailwind CSS**: A utility-first CSS framework for rapid UI development.
- **Vue Router**: The official router for Vue.js for building single-page applications.
- **Vuex**: State management pattern + library for Vue.js applications.

## Creating the Vite App:

## Step 1: Create the Vite App

Start by creating the Vite app with the Vue 3 and TypeScript template:

```bash
pnpm create vite@latest -- --template vue-ts

```

## Step 2: Install Tailwind CSS

Install Tailwind CSS along with PostCSS and Autoprefixer:

```bash
pnpm install -D tailwindcss postcss autoprefixer
npx tailwindcss init -p
```

## Step 3: Install Vue Router

Install Vue Router:

```bash
pnpm install vue-router@4
```

## URLs

- Development URL: [https://prep-app-dev.nuri-sw.com/](https://prep-app-dev.nuri-sw.com/)
- Production URL: [https://prep-app-prod.nuri-sw.com/](https://prep-app-prod.nuri-sw.com/)

## Learn More

To learn more about the tools and frameworks used, check out the following resources:

- [Vue 3 Documentation](https://vuejs.org/)
- [TypeScript Documentation](https://www.typescriptlang.org/)
- [Vite Documentation](https://vitejs.dev/)
- [Tailwind CSS Documentation](https://tailwindcss.com/)
- [Preline Documentation](https://preline.dev/)
- [Vue Router Documentation](https://router.vuejs.org/)
- [Vuex Documentation](https://vuex.vuejs.org/)

## Getting Started

To get started with development, clone the repository and install the dependencies:

```sh
pnpm install
vite