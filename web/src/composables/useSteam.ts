import { useFetch } from 'vue-composable'

export function useResolveVanityURL () {
  const { exec: requestExec, json, loading, status } = request('/ISteamUser/ResolveVanityURL/v1/')

  async function exec (value: string) {
    requestExec({ vanityurl: value })
  }

  return {
    exec, json, loading, status
  }
}

function request (path: string) {
  const { json, loading, exec: fetchExec, status } = useFetch()

  async function exec (params: any) {
    return fetchExec(buildURL(path, params), {
      method: 'GET'
    })
  }

  return {
    exec,
    json,
    loading,
    status
  }
}

function buildURL (path: string, params: any): string {
  const url = new URL(`${import.meta.env.VITE_STEAM_API_BASE_URL}${path}`)
  const query = { key: import.meta.env.VITE_STEAM_KEY, ...params }
  url.search = new URLSearchParams(query).toString()

  return url.toString()
}
