# Simple Go API

<div align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Go Logo" width="200"/>

  <p><em>Ejemplo de API REST en Go con Clean Architecture</em></p>
</div>

---

## Descripcion

Este proyecto es un ejemplo de una **API REST en Go** que implementa las mejores practicas de desarrollo:

- **Arquitectura Hexagonal (Clean Architecture)**: Separacion clara entre dominio, aplicacion e infraestructura
- **Inyeccion de Dependencias (DI)**: Contenedor manual para gestion centralizada de dependencias
- **Patron Repository**: Abstraccion de la capa de datos
- **Use Cases**: Logica de negocio aislada
- **Gin Framework**: Router HTTP de alto rendimiento
- **Hot Reload**: Desarrollo agil con Air

## Estructura del Proyecto

```
simple-go-api/
├── cmd/
│   └── api/
│       └── main.go              # Entry point
├── internal/
│   ├── infrastructure/
│   │   ├── config/              # Configuracion
│   │   ├── container/           # DI Container
│   │   ├── middlewares/         # Middlewares HTTP
│   │   └── router/              # Rutas
│   └── users/                   # Dominio de usuarios
│       ├── domain/              # Entidades e interfaces
│       ├── application/         # Casos de uso
│       └── infrastructure/      # Handlers y repositorios
└── docs/
    └── http/                    # Ejemplos de requests
```

## Instalacion y Uso

### Prerrequisitos

- Go 1.25.3 o superior
- Air (opcional, para hot reload)

### Instalacion

```bash
# Clonar el repositorio
git clone https://github.com/tuusuario/simple-go-api.git
cd simple-go-api

# Instalar dependencias
go mod download
```

### Ejecutar

```bash
# Con Go directamente
go run cmd/api/main.go

# Con Air (hot reload)
air
```

El servidor estara disponible en `http://localhost:3000`

## Endpoints

### Crear Usuario

```http
POST /api/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com"
}
```

### Ejemplo de respuesta

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "John Doe",
  "email": "john@example.com"
}
```

## Tecnologias

- [Go](https://golang.org/) - Lenguaje de programacion
- [Gin](https://gin-gonic.com/) - Framework HTTP
- [Air](https://github.com/cosmtrek/air) - Hot reload
- [Zap](https://github.com/uber-go/zap) - Logger

## Arquitectura

Este proyecto sigue los principios de **Clean Architecture**:

1. **Domain Layer**: Entidades de negocio e interfaces (contratos)
2. **Application Layer**: Casos de uso con logica de negocio
3. **Infrastructure Layer**: Implementaciones concretas (HTTP handlers, repositorios, config)

### Flujo de una Request

```
HTTP Request → Router → Handler → Use Case → Repository → Domain
```

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un issue o PR.

## Licencia

MIT License
