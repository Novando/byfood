function setQueryParam(url: string, query: Record<string, string>) {
  const queryParam = (new URLSearchParams(query)).toString()
  const urls = [url]
  if (queryParam !== "") urls.push(queryParam)
  return urls
}

export async function postData(url: string, payload: object, query = {}, customHeaders?: object) {
  const urls = setQueryParam(url, query)
  try {
    const res = await fetch(urls.join('?'), {
      method: 'POST',
      body: JSON.stringify(payload),
      cache: 'no-cache',
      headers: {
        'Content-Type': 'application/json',
        ...customHeaders
      }
    })
    const resJson = await res.json()
    if (res.status >= 200 && res.status < 300) {
      return resJson
    }
    throw {...resJson, status: res.status}
  } catch (err: any) {
    throw {
      data: err.data || 'Contact support team',
      message: err.message || 'UNKNOWN_ERROR',
      status: err.status,
    }
  }
}

export async function getData(url: string, query = {}, customHeaders?: object) {
  const urls = setQueryParam(url, query)
  try {
    const res = await fetch(urls.join('?'), {
      method: 'GET',
      cache: 'no-cache',
      headers: {
        'Content-Type': 'application/json',
        ...customHeaders
      }
    })
    const resJson = await res.json()
    if (res.status >= 200 && res.status < 300) {
      return resJson
    }
    throw {...resJson, status: res.status}
  } catch (err: any) {
    throw {
      data: err.data || 'Contact support team',
      message: err.message || 'UNKNOWN_ERROR',
      status: err.status,
    }
  }
}

export async function delData(url: string, query = {}, customHeaders?: object) {
  const urls = setQueryParam(url, query)
  try {
    const res = await fetch(urls.join('?'), {
      method: 'DELETE',
      cache: 'no-cache',
      headers: {
        'Content-Type': 'application/json',
        ...customHeaders
      }
    })
    const resJson = await res.json()
    if (res.status >= 200 && res.status < 300) {
      return resJson
    }
    throw {...resJson, status: res.status}
  } catch (err: any) {
    throw {
      data: err.data || 'Contact support team',
      message: err.message || 'UNKNOWN_ERROR',
      status: err.status,
    }
  }
}

export async function putData(url: string, payload: object, query = {}, customHeaders?: object) {
  const urls = setQueryParam(url, query)
  try {
    const res = await fetch(urls.join('?'), {
      method: 'PUT',
      body: JSON.stringify(payload),
      cache: 'no-cache',
      headers: {
        'Content-Type': 'application/json',
        ...customHeaders
      }
    })
    const resJson = await res.json()
    if (res.status >= 200 && res.status < 300) {
      return resJson
    }
    throw {...resJson, status: res.status}
  } catch (err: any) {
    console.log(err)
    throw {
      data: err.data || 'Contact support team',
      message: err.message || 'UNKNOWN_ERROR',
      status: err.status,
    }
  }
}

export async function postFormData(url: string, form: FormData, query = {}, customHeaders?: object) {
  const urls = setQueryParam(url, query)
  try {
    const res = await fetch(urls.join('?'), {
      method: 'POST',
      body: form,
      cache: 'no-cache',
      headers: {
        ...customHeaders
      }
    })
    const resJson = await res.json()
    if (res.status >= 200 && res.status < 300) {
      return resJson
    }
    throw {...resJson, status: res.status}
  } catch (err: any) {
    throw {
      data: err.data || 'Contact support team',
      message: err.message || 'UNKNOWN_ERROR',
      status: err.status,
    }
  }
}