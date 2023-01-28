import http from 'k6/http';
import { check, group } from 'k6'

export const options = {
    thresholds: {http_req_duration: ['p(95)<200']},

}

export function setup(){
    
}

export default function(){
    const url = 'http://localhost:8080/'
    const params = {
        headers: {
          'Content-Type': 'application/json',
        },
    }
    group('get resources', () => {
        const resp = http.get(url + 'resources')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('create resource', () =>{
        const data = {
            name: 'test1'
        }
    
        const resp = http.post(url + 'resources', data, params)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
}