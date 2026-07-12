# Triggo 🔔

**Triggo** es un webhook de GitHub escrito en Go que escucha eventos de un repositorio (por ejemplo, `push`) y los reenvía como notificaciones formateadas a Discord mediante un *embed*, usando la URL de webhook de Discord registrada para ese repositorio.

Desplegada como función serverless en **Vercel**.

## ✨ Características

- Recibe y valida webhooks de GitHub usando la firma HMAC-SHA256 (`X-Hub-Signature-256`).
- Decodifica el payload del evento recibido (actualmente soporta `push`).
- Registro de la relación *repositorio ↔ webhook de Discord* en base de datos (`POST /api/setup`), lo que permite notificar distintos repositorios a distintos canales de Discord.
- Construye un *embed* de Discord con la información del evento (rama modificada y usuario que hizo el push) y lo envía al canal correspondiente a ese repositorio.
- Persistencia en PostgreSQL vía [GORM](https://gorm.io/), con auto-migración del modelo `RepositoryWebhook`.
- Middleware CORS configurable para permitir que un frontend externo consuma el endpoint de registro (`/api/setup`).
- Arquitectura desacoplada por capas (handler / services / ports / models), lo que facilita añadir soporte para nuevos eventos o nuevos destinos de notificación.

## 🏗️ Arquitectura

```
WebHook/
├── cmd/
│   └── main.go                  # Servidor local (multiplexor http.ServeMux)
├── api/
│   ├── webhook.go               # Entry point: recepción de eventos de GitHub
│   └── setup.go                 # Entry point: registro de repositorio ↔ webhook de Discord
├── pkg/
│   ├── config/
│   │   └── config.go             # Carga de variables de entorno
│   ├── ports/
│   │   ├── discord.go             # Interfaz de servicios de Discord
│   │   ├── github.go              # Interfaz de servicios de GitHub
│   │   └── repository.go          # Interfaz de servicios de persistencia
│   ├── github/
│   │   ├── handler/               # Manejo de la petición HTTP entrante del webhook
│   │   ├── services/               # Validación de firma y decodificación de eventos
│   │   └── model/                  # Estructuras del payload de GitHub (push, pusher, repository)
│   ├── discord/
│   │   ├── services/                # Construcción del embed, del payload y envío a Discord
│   │   └── model/                   # Estructuras del embed y payload de Discord
│   ├── repository/
│   │   ├── handler/                 # Manejo de la petición de registro (setup)
│   │   ├── services/                 # Conexión a la base de datos y operaciones CRUD
│   │   └── model/                    # Modelo persistido RepositoryWebhook
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

**Flujo de una petición de webhook (`/api/webhook`):**

1. GitHub envía un webhook a `api/webhook.go`.
2. `GithubServices.ValidatedHash` valida que la firma del payload coincide con el *secret* configurado.
3. `GithubServices.DecodeMessage` interpreta el evento (`push`, etc.) y arma un mensaje legible junto con la información de instalación/repositorio.
4. `DiscordServices.CreateEmbed` y `CreateDiscordPayload` transforman ese mensaje en un embed de Discord.
5. `DiscordServices.SendPayload` busca en base de datos la URL de Discord asociada a ese repositorio (`RepositoryServices.SearchRecord`) y envía el payload por POST.

## 🔧 Requisitos

- [Go](https://go.dev/) 1.26.4 o superior
- Una base de datos PostgreSQL accesible
- Un servidor de Discord con uno o más [webhooks configurados](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks) (uno por repositorio a notificar)
- Un repositorio de GitHub donde configurar el webhook

## ⚙️ Configuración

Triggo se configura mediante variables de entorno:

| Variable                 | Descripción                                                                                   |
| -------------------------- | ----------------------------------------------------------------------------------------------- |
| `GITHUB_WEBHOOK_SECRET`   | Secreto usado para validar la firma HMAC-SHA256 del webhook de GitHub                          |
| `POSTGRES_URL_NON_POOLING`| DSN de conexión a la base de datos PostgreSQL                                                  |
| `FRONT_URL`               | Origen permitido por el middleware CORS (URL del frontend que consume `/api/setup`)            |
| `PORT`                    | (Opcional) Puerto del servidor local. Por defecto `8080`                                       |

Puedes definirlas en un archivo `.env` local (ignorado por git, cargado con [godotenv](https://github.com/joho/godotenv)) o como variables de entorno en tu plataforma de despliegue (por ejemplo, Vercel).

## 🚀 Instalación y uso local

```bash
# Clonar el repositorio (rama v2)
git clone --branch v2 https://github.com/maikreyes/Triggo.git
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

## ☁️ Despliegue

El proyecto está preparado para desplegarse como funciones serverless en **Vercel** (de ahí la carpeta `api/`, con un archivo por endpoint). Basta con:

1. Importar el repositorio en Vercel.
2. Configurar las variables de entorno `GITHUB_WEBHOOK_SECRET`, `POSTGRES_URL_NON_POOLING` y `FRONT_URL` en el panel del proyecto.
3. Desplegar y usar las URLs generadas (`https://<tu-proyecto>.vercel.app/api/webhook`, `/api/setup`) como *Payload URL* al configurar el webhook en GitHub y como endpoint de registro desde el frontend.

### Configurar el webhook en GitHub

1. Ve a **Settings → Webhooks → Add webhook** en tu repositorio.
2. En **Payload URL**, coloca la URL de tu despliegue (ej. `https://<tu-proyecto>.vercel.app/api/webhook`).
3. En **Content type**, selecciona `application/json`.
4. En **Secret**, coloca el mismo valor que usaste en `GITHUB_WEBHOOK_SECRET`.
5. Selecciona los eventos que quieras enviar (por ejemplo, `push`).
6. Registra el par repositorio/Discord con una petición a `/api/setup` para que las notificaciones de ese repositorio lleguen al canal correcto.

## 📌 Eventos soportados

- ✅ `push`
- 🔜 Otros eventos pueden añadirse extendiendo `DecodeMessage` y `CreateEmbed`.

## 📄 Licencia

Este proyecto está bajo la licencia [MIT](./LICENSE).

## 👤 Autor

Desarrollado por **[Michael Estiven Reyes Escobar](https://github.com/maikreyes)**.