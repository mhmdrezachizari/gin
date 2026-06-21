# Gin Shop API

A simple REST API built with **Go**, **Gin**, **GORM**, and **PostgreSQL** for managing products with multiple images.

## Features

- Create products
- Upload multiple images for each product
- One-to-many relationship between products and product images
- Store uploaded image paths in PostgreSQL
- Serve uploaded files statically
- Get all products with their images

## Tech Stack

- Go
- Gin Web Framework
- GORM
- PostgreSQL
- Multipart Form Data

## Project Structure

```text
.
├── cmd
│   └── main.go
├── database
│   └── postgres.go
├── handlers
│   └── product.go
├── models
│   └── product.go
├── uploads
│   └── products
├── go.mod
└── go.sum