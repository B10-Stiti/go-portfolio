# Go Portfolio Website

A personal portfolio project built with **Go**.
This project comes in two versions:
- **Tailwind CSS** → [`tailwind` branch](https://github.com/B10-Stiti/go-portfolio/tree/migrate/tailwindcss) _(default)_  
- **Classic CSS** → [`classic` branch](https://github.com/B10-Stiti/go-portfolio/tree/master)  
## Features
- Server-side rendering with Go
- HTML template rendering
- Two styling approaches:
  - Tailwind CSS (modern)
  - Classic CSS (traditional)
## Local Development
Run Tailwind in watch mode and the Go server with [Air](https://github.com/air-verse/air)

```bash
make tailwind
make run
```
## 📦 Production
Build and run the project
```bash
make build
make start
```
## Preview
https://github.com/user-attachments/assets/4012d4f8-c6e8-4ee0-8c27-531fdccfff1f
## Built With
- **Backend & SSR** → Go
- **Styling (Tailwind branch)** → [Tailwind CSS](https://tailwindcss.com/)
- **Styling (Classic branch)** → Custom CSS
- **Hot reload** → [Air](https://github.com/air-verse/air)
