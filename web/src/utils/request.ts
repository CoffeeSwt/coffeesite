import axios from 'axios' // 引入axios
import { BASE_URL } from 'global.config'

const service = axios.create({
  baseURL: BASE_URL,
  timeout: 99999,
})

// http request 拦截器
service.interceptors.request.use(
  (config) => {
    return config
  },
  (error) => {
    return error
  }
)

// http response 拦截器
service.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    switch (error.response.status) {
      case 500:
        break
      case 404:
        break
    }

    return error
  }
)

export default service
