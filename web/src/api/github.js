import axios from 'axios'

const service = axios.create()

export function Commits(page) {
  return service({
    url: '',
    method: 'get'
  })
}

export function Members() {
  return service({
    url: '',
    method: 'get'
  })
}
