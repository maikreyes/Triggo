# Triggo 🔔

**Triggo** es un webhook GitHub App escrito en Go que escucha eventos de un repositorio de GitHub (por ejemplo, `push`) y los reenvía como notificaciones formateadas a un canal de Discord mediante un *embed*.

Está desplegado como función serverless en Vercel: **[triggo-webhook.vercel.app](https://triggo-webhook.vercel.app)**

## ✨ Características

- Recibe y valida webhooks de GitHub usando la firma HMAC-SHA256 (`X-Hub-Signature-256`).
- Decodifica el payload del evento recibido (actualmente soporta `push`).
- Construye un *embed* de Discord con la información del evento (rama modificada y usuario que hizo el push).
- Envía el mensaje formateado a un canal de Discord vía Discord Webhooks.
- Arquitectura desacoplada por capas (handler / services / ports / models), lo que facilita añadir soporte para nuevos eventos o nuevos destinos de notificación.

## 🏗️ Arquitectura

El proyecto sigue un enfoque inspirado en arquitectura hexagonal, separando responsabilidades en paquetes independientes:

```
WebHook/
├── api/
│   └── webhook.go              # Entry point (función serverless de Vercel)
├── pkg/
│   ├── config/
│   │   └── config.go            # Carga de variables de entorno
│   ├── ports/
│   │   ├── discord.go           # Interfaz de servicios de Discord
│   │   └── github.go            # Interfaz de servicios de GitHub
│   ├── github/
│   │   ├── handler/             # Manejo de la petición HTTP entrante
│   │   ├── services/             # Validación de firma y decodificación de eventos
│   │   └── model/                # Estructuras del payload de GitHub (push, pusher, repository)
│   └── discord/
│       ├── services/              # Construcción del embed y payload de Discord
│       └── model/                 # Estructuras del embed y payload de Discord
├── go.mod
└── go.sum
```

**Flujo de una petición:**

1. GitHub envía un webhook a `WebHook/api/webhook.go`.
2. `GithubServices.ValidatedHash` valida que la firma del payload coincide con el *secret* configurado.
3. `GithubServices.DecodeMessage` interpreta el evento (`push`, etc.) y arma un mensaje legible.
4. `DiscordServices.CreateEmbed` y `CreateDiscordPayload` transforman ese mensaje en un embed de Discord.
5. El payload resultante se envía por POST a la URL del webhook de Discord configurada.

## 🔧 Requisitos

- [Go](https://go.dev/) 1.26.4 o superior
- Un servidor de Discord con un [webhook configurado](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks)
- Un repositorio de GitHub donde configurar el webhook

## ⚙️ Configuración

Triggo se configura mediante variables de entorno:

| Variable                | Descripción                                                                 |
| ------------------------ | ---------------------------------------------------------------------------- |
| `GITHUB_WEBHOOK_SECRET`  | Secreto usado para validar la firma HMAC-SHA256 del webhook de GitHub       |
| `DISCORD_WEBHOOK_URL`    | URL del webhook de Discord al que se enviarán las notificaciones            |

Puedes definirlas en un archivo `.env` local (ignorado por git) o como variables de entorno en tu plataforma de despliegue (por ejemplo, Vercel).

## 🚀 Instalación y uso local

```bash
# Clonar el repositorio
git clone https://github.com/maikreyes/Triggo.git
cd Triggo/WebHook

# Instalar dependencias
go mod tidy

# Definir las variables de entorno necesarias
export GITHUB_WEBHOOK_SECRET="tu_secreto"
export DISCORD_WEBHOOK_URL="https://discord.com/api/webhooks/..."

# Ejecutar
go run api/webhook.go
```

## ☁️ Despliegue

El proyecto está preparado para desplegarse como función serverless en **Vercel** (de ahí la carpeta `api/`). Basta con:

1. Importar el repositorio en Vercel.
2. Configurar las variables de entorno `GITHUB_WEBHOOK_SECRET` y `DISCORD_WEBHOOK_URL` en el panel del proyecto.
3. Desplegar y usar la URL generada (`https://<tu-proyecto>.vercel.app/api/webhook`) como *Payload URL* al configurar el webhook en GitHub.

### Configurar el webhook en GitHub

1. Ve a **Settings → Webhooks → Add webhook** en tu repositorio.
2. En **Payload URL**, coloca la URL de tu despliegue (ej. `https://triggo-webhook.vercel.app/api/webhook`).
3. En **Content type**, selecciona `application/json`.
4. En **Secret**, coloca el mismo valor que usaste en `GITHUB_WEBHOOK_SECRET`.
5. Selecciona los eventos que quieras enviar (por ejemplo, `push`).

## 📌 Eventos soportados

- ✅ `push`
- 🔜 Otros eventos pueden añadirse extendiendo `DecodeMessage` y `CreateEmbed`.

## 📄 Licencia

Este proyecto está bajo la licencia [MIT](./LICENSE).

## 👤 Autor

Desarrollado por **[Michael Estiven Reyes Escobar](https://github.com/maikreyes)**.