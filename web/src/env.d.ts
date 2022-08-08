/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_STEAM_KEY: string
  readonly VITE_STEAM_API_BASE_URL: string
  // more env variables...
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
