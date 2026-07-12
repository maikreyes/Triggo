# Triggo 🔔

**Triggo** es una GitHub App escrita en Go que escucha eventos de repositorios de GitHub (por ejemplo, `push`) y los reenvía como notificaciones formateadas a Discord mediante un *embed*, usando un webhook de Discord configurado por cada repositorio.

A diferencia de un webhook manual de repositorio único, Triggo funciona como una **GitHub App multi-tenant**: se autentica ante GitHub mediante JWT + clave privada de la App, obtiene tokens de instalación por usuario/organización, y persiste en base de datos la asociación entre cada repositorio y su canal de Discord de destino.

Desplegada como función serverless en **Vercel**.

## ✨ Características

- **GitHub App real**: autenticación mediante JWT firmado con la clave privada de la App (`RS256`) y generación de tokens de instalación (`installation access tokens`) contra la API de GitHub.
- Listado de repositorios accesibles para una instalación (`GET /api/get_repositories`), con paginación automática contra la API de GitHub.
- Registro de la relación *repositorio ↔ webhook de Discord* en base de datos (`POST /api/setup`), permitiendo que cada repositorio notifique a un canal distinto.
- Recepción y validación de webhooks de GitHub mediante firma HMAC-SHA256 (`X-Hub-Signature-256`).
- Decodificación del payload del evento recibido (actualmente `push`) y envío del *embed* correspondiente al Discord asociado a ese repositorio.
- Persistencia en PostgreSQL vía [GORM](https://gorm.io/), con auto-migración del modelo `RepositoryWebhook`.
- Middleware CORS configurable para permitir que un frontend externo consuma la API (`/api/setup`, `/api/get_repositories`).
- Arquitectura desacoplada por capas (handler / services / ports / models) inspirada en un enfoque hexagonal, lo que facilita añadir soporte para nuevos eventos, nuevas fuentes de notificación o nuevos destinos.

## 🏗️ Arquitectura

```
WebHook/
├── cmd/
│   └── main.go                      # Servidor local (multiplexor http.ServeMux)
├── api/
│   ├── webhook.go                   # Entry point: recepción de eventos de GitHub
│   ├── setup.go                     # Entry point: registro de repositorio ↔ webhook de Discord
│   └── get_repositories.go          # Entry point: listado de repos accesibles por una instalación
├── pkg/
│   ├── config/
│   │   └── config.go                 # Carga y validación de variables de entorno
│   ├── ports/
│   │   ├── discord.go                 # Interfaz de servicios de Discord
│   │   ├── github.go                  # Interfaz de servicios de GitHub
│   │   ├── jwt.go                     # Interfaz de servicios JWT
│   │   └── repository.go              # Interfaz de servicios de persistencia
│   ├── jwt/
│   │   └── services/                  # Creación del JWT firmado para autenticar la App ante GitHub
│   ├── github/
│   │   ├── handler/                   # Manejo de peticiones HTTP (webhook y listado de repos)
│   │   ├── services/                  # Validación de firma, decodificación de eventos, tokens de instalación
│   │   └── model/                     # Estructuras de payloads de GitHub (push, pusher, installation, repository)
│   ├── discord/
│   │   ├── services/                  # Construcción del embed, del payload y envío a Discord
│   │   └── model/                     # Estructuras del embed y payload de Discord
│   ├── repository/
│   │   ├── handler/                   # Manejo de la petición de registro (setup)
│   │   ├── services/                  # Conexión a la base de datos y operaciones CRUD
│   │   └── model/                     # Modelo persistido RepositoryWebhook
│   └── middleware/
│       ├── middleware.go              # Estructura base del middleware
│       └── cors.go                    # Configuración de cabeceras CORS
├── go.mod
└── go.sum
```

**Flujo de registro de un repositorio (`/api/setup`):**

1. Un frontend (u otro cliente) envía el `installation_id`, el nombre del repositorio y la URL del webhook de Discord.
2. `RepositoryServices.DecodeRecord` decodifica el cuerpo de la petición.
3. `RepositoryServices.CreateRecord` persiste el registro en PostgreSQL.

**Flujo de un evento de webhook (`/api/webhook`):**

1. GitHub envía el evento a `api/webhook.go`.
2. `GithubServices.ValidatedHash` valida que la firma del payload coincide con el *secret* configurado.
3. `GithubServices.DecodeMessage` interpreta el evento (`push`, etc.) y arma un mensaje legible junto con la información de instalación/repositorio.
4. `DiscordServices.CreateEmbed` y `CreateDiscordPayload` transforman ese mensaje en un embed de Discord.
5. `DiscordServices.SendPayload` busca en base de datos la URL de Discord asociada a ese repositorio (`RepositoryServices.SearchRecord`) y envía el payload por POST.

**Flujo de listado de repositorios (`/api/get_repositories`):**

1. Se recibe un `installation_id` por query string.
2. `GithubServices.RequestAccessToken` genera un JWT (`JWTServices.CreateJWT`) y lo intercambia por un token de instalación en la API de GitHub.
3. `GithubServices.RequestInstallationRepositories` pagina sobre `/installation/repositories` usando ese token y devuelve la lista completa.

## 🔧 Requisitos

- [Go](https://go.dev/) 1.26.4 o superior
- Una [GitHub App](https://docs.github.com/en/apps/creating-github-apps) creada (App ID, clave privada y webhook secret)
- Una base de datos PostgreSQL accesible
- Un servidor de Discord con uno o más [webhooks configurados](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks) (uno por repositorio a notificar)

## ⚙️ Configuración

Triggo se configura mediante variables de entorno:

| Variable                     | Descripción                                                                                   |
| ----------------------------- | ----------------------------------------------------------------------------------------------- |
| `GITHUB_APP_ID`               | ID numérico de la GitHub App                                                                   |
| `GITHUB_PRIVATE_KEY_BASE64`   | Clave privada (PEM) de la GitHub App, codificada en Base64                                     |
| `GITHUB_WEBHOOK_SECRET`       | Secreto usado para validar la firma HMAC-SHA256 del webhook de GitHub                          |
| `POSTGRES_URL_NON_POOLING`    | DSN de conexión a la base de datos PostgreSQL                                                  |
| `FRONT_URL`                   | Origen permitido por el middleware CORS (URL del frontend que consume `/api/setup` y `/api/get_repositories`) |
| `PORT`                        | (Opcional) Puerto del servidor local. Por defecto `8080`                                       |

Puedes definirlas en un archivo `.env` local (ignorado por git, cargado con [godotenv](https://github.com/joho/godotenv)) o como variables de entorno en tu plataforma de despliegue (por ejemplo, Vercel).

## 🚀 Instalación y uso local

```bash
# Clonar el repositorio
git clone https://github.com/maikreyes/Triggo.git
cd Triggo/WebHook

# Instalar dependencias
go mod tidy

# Definir las variables de entorno necesarias en un archivo .env
# (ver tabla de Configuración)

# Ejecutar
go run cmd/main.go
```

El servidor local expone las siguientes rutas sobre `http://localhost:8080`:

- `POST /api/webhook` — recepción de eventos de GitHub
- `POST /api/setup` — registro de repositorio ↔ webhook de Discord
- `GET /api/get_repositories` — listado de repositorios accesibles por una instalación

## ☁️ Despliegue

El proyecto está preparado para desplegarse como funciones serverless en **Vercel** (de ahí la carpeta `api/`, con un archivo por endpoint). Basta con:

1. Importar el repositorio en Vercel.
2. Configurar las variables de entorno descritas arriba en el panel del proyecto.
3. Desplegar y usar las URLs generadas (`https://<tu-proyecto>.vercel.app/api/webhook`, `/api/setup`, `/api/get_repositories`).

### Configurar el webhook en GitHub

1. Crea o edita tu GitHub App en **Settings → Developer settings → GitHub Apps**.
2. En **Webhook → Webhook URL**, coloca la URL de tu despliegue (ej. `https://<tu-proyecto>.vercel.app/api/webhook`).
3. En **Webhook secret**, coloca el mismo valor que usaste en `GITHUB_WEBHOOK_SECRET`.
4. Selecciona los eventos (permisos) que quieras enviar, por ejemplo `push`.
5. Instala la App en el repositorio u organización deseada y registra el par repositorio/Discord con una petición a `/api/setup`.

## 📌 Eventos soportados

- ✅ `push`
- 🔜 Otros eventos pueden añadirse extendiendo `DecodeMessage` (paquete `github/services`) y `CreateEmbed` (paquete `discord/services`).

## 📄 Licencia

Este proyecto está bajo la licencia [MIT](./LICENSE).

## 👤 Autor

Desarrollado por **[Michael Estiven Reyes Escobar](https://github.com/maikreyes)**.