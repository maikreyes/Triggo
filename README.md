# Triggo

**Triggo** es una GitHub App que reenvía notificaciones de eventos `push` de tus repositorios hacia canales de Discord, mostrando la información como embeds formateados. La versión `v4` evoluciona el proyecto hacia una arquitectura completa de dos componentes: un backend en Go y un frontend en Next.js que permite configurar visualmente la integración por repositorio.

## Índice

- [Arquitectura](#arquitectura)
- [Características principales](#características-principales)
- [Estructura del proyecto](#estructura-del-proyecto)
- [Flujo de funcionamiento](#flujo-de-funcionamiento)
- [Endpoints de la API](#endpoints-de-la-api)
- [Variables de entorno](#variables-de-entorno)
- [Puesta en marcha](#puesta-en-marcha)
- [Stack tecnológico](#stack-tecnológico)
- [Licencia](#licencia)

## Arquitectura

El proyecto está dividido en dos módulos independientes dentro del mismo repositorio:

- **`WebHook/`** — API en Go que recibe los webhooks de GitHub, valida su firma, los transforma en un mensaje y los reenvía a Discord. También expone endpoints para que el frontend liste repositorios y registre la configuración de cada uno.
- **`frontend/`** — Aplicación en Next.js que sirve como panel de configuración: permite al usuario, luego de instalar la GitHub App, elegir un repositorio y asociarlo a una URL de webhook de Discord.

## Características principales

- **Autenticación como GitHub App**: generación de JWT firmado con RS256 (`golang-jwt/jwt`) usando el `App ID` y una clave privada RSA, para autenticarse ante la API de GitHub y obtener tokens de instalación.
- **Validación de firma de webhooks**: verificación del header `X-Hub-Signature-256` antes de procesar cualquier evento entrante.
- **Multi-tenant por repositorio**: cada combinación `installation_id` + `repository` se persiste en PostgreSQL (vía GORM) junto a su propia URL de Discord, permitiendo que múltiples repositorios e instalaciones envíen a distintos canales.
- **Reenvío a Discord**: construcción de un embed a partir del evento `push` (rama, usuario que hizo el push, repositorio) y envío como payload a la URL de webhook configurada.
- **Panel de configuración (frontend)**: formulario que consulta los repositorios disponibles de la instalación (`/api/repositories`) y registra la URL de Discord para el repositorio elegido (`/api/setup`).
- **CORS configurado** mediante middleware propio para permitir la comunicación entre el frontend y la API.

## Estructura del proyecto

```
Triggo/
├── WebHook/                     # Backend en Go
│   ├── cmd/
│   │   └── main.go              # Punto de entrada, registro de rutas HTTP
│   ├── api/                     # Handlers HTTP expuestos (webhook, setup, repositories)
│   └── pkg/
│       ├── config/              # Carga de configuración desde variables de entorno
│       ├── github/               # Modelos, servicios y handler relacionados a GitHub
│       │   ├── handler/         # WebhookHandler, GetRepositoriesHandler
│       │   ├── model/            # push, installation, repository, pusher, etc.
│       │   └── services/        # DecodeMessage, ValidatedHash, RequestAccessToken...
│       ├── jwt/services/        # Generación de JWT (RS256) para la GitHub App
│       ├── discord/              # Construcción y envío de embeds/payloads a Discord
│       ├── repository/           # Persistencia (GORM/PostgreSQL) de la config por repo
│       └── middleware/          # CORS y middlewares comunes
│
└── frontend/                    # Panel de configuración en Next.js
    └── app/
        ├── page.tsx
        └── setup/
            ├── page.tsx
            ├── component/       # AppHeader, LabeledInput, SelectInput, WebhookForm
            └── customHooks/     # useFetchRepositories, useSubmitWebhook
```

## Flujo de funcionamiento

1. El usuario instala la GitHub App en uno o varios repositorios y es redirigido al frontend (`/setup?installation_id=...`).
2. El frontend llama a `GET /api/repositories` para listar los repositorios disponibles de esa instalación (el backend genera un JWT, solicita un *installation access token* a GitHub y consulta los repositorios).
3. El usuario selecciona un repositorio e ingresa la URL del webhook de Discord destino; el frontend envía esta configuración a `POST /api/setup`, que la persiste en PostgreSQL.
4. Cuando ocurre un `push` en el repositorio, GitHub envía el evento a `POST /api/webhook`. El backend valida la firma, decodifica el mensaje, busca la configuración correspondiente (`installation_id` + `repository`) y reenvía un embed formateado a la URL de Discord asociada.

## Endpoints de la API

| Método | Ruta | Descripción |
|---|---|---|
| `POST` | `/api/webhook` | Recibe y valida los eventos de GitHub, y los reenvía a Discord. |
| `GET` | `/api/repositories` | Lista los repositorios de una instalación (`?installation_id=`). |
| `POST` | `/api/setup` | Registra la asociación repositorio → URL de Discord. |

## Variables de entorno

El backend (`WebHook/`) requiere las siguientes variables:

| Variable | Descripción |
|---|---|
| `GITHUB_APP_ID` | ID numérico de la GitHub App. |
| `GITHUB_PRIVATE_KEY_BASE64` | Clave privada RSA de la App, codificada en Base64. |
| `GITHUB_WEBHOOK_SECRET` | Secreto usado para validar la firma HMAC de los webhooks. |
| `POSTGRES_URL_NON_POOLING` | Cadena de conexión a PostgreSQL. |
| `FRONT_URL` | URL del frontend, usada para la configuración de CORS. |
| `PORT` | Puerto del servidor (por defecto `8080`). |

El frontend (`frontend/`) requiere:

| Variable | Descripción |
|---|---|
| `NEXT_PUBLIC_WEBHOOK_URL` | URL base de la API del backend. |

## Puesta en marcha

### Backend

```bash
cd WebHook
cp .env.example .env   # completar con tus credenciales
go run cmd/main.go
```

### Frontend

```bash
cd frontend
pnpm install
pnpm dev
```

## Stack tecnológico

- **Backend**: Go 1.26, `net/http`, `golang-jwt/jwt`, GORM + driver PostgreSQL, `godotenv`, `google/uuid`.
- **Frontend**: Next.js 16, React 19, TypeScript, Tailwind CSS 4.
- **Infraestructura**: Vercel (despliegue), PostgreSQL.

## Licencia

Este proyecto está bajo la licencia MIT. Ver el archivo [LICENSE](./LICENSE) para más detalles.