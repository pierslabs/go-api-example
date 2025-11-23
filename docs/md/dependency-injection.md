# Dependency Injection Pattern

## Arquitectura

Este proyecto implementa un patrón de inyección de dependencias manual y explícito, siguiendo las mejores prácticas de Go.

## Estructura

```
internal/infrastructure/
├── config/       # Gestión de configuración
└── container/    # Contenedor de dependencias
```

## Componentes

### 1. Config ([config/config.go](../internal/infrastructure/config/config.go))

Gestiona toda la configuración de la aplicación usando variables de entorno con valores por defecto.

```go
cfg, err := config.Load()
```

**Variables de entorno:**
- `PORT`: Puerto del servidor (default: 3000)
- `GIN_MODE`: Modo de Gin (default: debug)

### 2. Container ([container/container.go](../internal/infrastructure/container/container.go))

Contenedor que gestiona el ciclo de vida y las dependencias de todos los componentes.

**Ventajas:**
- ✅ Inicialización centralizada y ordenada
- ✅ Dependencias explícitas y fáciles de rastrear
- ✅ Fácil de testear (puedes crear contenedores de prueba)
- ✅ No requiere librerías externas
- ✅ El orden de inicialización es claro y visible

**Flujo de inicialización:**
1. `initInfrastructure()` - Componentes base (DB, HTTP engine)
2. `initRepositories()` - Capa de persistencia
3. `initUseCases()` - Lógica de negocio
4. `initHandlers()` - Controladores HTTP
5. `initRouter()` - Configuración de rutas

## Uso

### En main.go

```go
// 1. Cargar configuración
cfg, err := config.Load()

// 2. Crear contenedor con todas las dependencias
app, err := container.New(cfg)

// 3. Usar los componentes
app.Router.SetupRoutes()
app.Engine.Run(port)
```

### En tests

```go
func TestSomething(t *testing.T) {
    cfg := &config.Config{
        Server: config.ServerConfig{Port: "8080", Mode: "test"},
    }

    container, _ := container.New(cfg)
    // Usa container.UserRepository, container.CreateUserUseCase, etc.
}
```

## Añadir nuevas dependencias

### 1. Agregar al Container struct:

```go
type Container struct {
    // ... existing fields
    NewRepository domain.NewRepository
    NewUseCase    *application.NewUseCase
    NewHandler    *handlers.NewHandler
}
```

### 2. Inicializar en el método correspondiente:

```go
func (c *Container) initRepositories() {
    // ... existing repos
    c.NewRepository = repositories.NewImplementation()
}

func (c *Container) initUseCases() {
    // ... existing use cases
    c.NewUseCase = application.NewUseCase(c.NewRepository)
}
```

### 3. Inyectar en handlers o router según sea necesario

## Principios seguidos

- **Single Responsibility**: Cada componente tiene una responsabilidad clara
- **Dependency Inversion**: Los módulos de alto nivel no dependen de los de bajo nivel
- **Explicit Dependencies**: Todas las dependencias son explícitas en los constructores
- **Composition over Inheritance**: Usamos composición para construir el grafo de dependencias
- **Separation of Concerns**: Config, inicialización, y lógica de negocio están separados

## Comparación: Antes vs Después

### ❌ Antes (main.go con DI manual):
```go
userRepo := repositories.NewInMemoryUserRepository()
createUserUC := application.NewCreateUserUseCase(userRepo)
userHandler := handlers.NewUserHttpHandler(createUserUC)
appRouter := router.NewRouter(engine, userHandler)
```

**Problemas:**
- Main crece con cada nueva feature
- Difícil de testear
- Sin configuración centralizada
- Orden de inicialización mezclado con lógica

### ✅ Después (Container Pattern):
```go
cfg, _ := config.Load()
app, _ := container.New(cfg)
app.Router.SetupRoutes()
```

**Beneficios:**
- Main minimalista y enfocado
- Fácil de testear cada capa
- Configuración centralizada
- Orden de inicialización claro y mantenible
