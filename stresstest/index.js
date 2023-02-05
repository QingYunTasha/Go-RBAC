import http from 'k6/http';
import { check, group } from 'k6'
import { URL } from 'https://jslib.k6.io/url/1.0.0/index.js'

export const options = {
    thresholds: {http_req_duration: ['p(95)<200']},

}

export function setup(){
    
}

export default function(){
    const url = 'http://localhost:8080/'
    group('get resources', () => {
        const resp = http.get(url + 'resources')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('create resource', () => {
        const data = {
            name: 'test1'
        }
        
        const resp = http.post(url + 'resources', data)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('get resource', () => {
        const resp = http.get(url + 'resources' + '/' + 'test1')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('update resource', () => {
        const data = {
            name: 'test2'
        }

        const resp = http.put(url + 'resources' + '/' + 'test1', data)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('delete resource', () => {
         const resp = http.del(url + 'resources' + '/' + 'test2')
         check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('get permissions', () => {
        const resp = http.get(url + 'permissions')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
    group('create permission', () => {
        const data = {
            operation: 'read',
            resourcename: 'res1'
        }
        
        const resp = http.post(url + 'permissions', data)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
    group('get permission', () => {
        const resp = http.get(url + 'permissions' + '/' + 'res1')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
    group('delete permission', () => {
        const resp = http.del(url + 'permissions' + '/' + 'res1' + '/' + 'read')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('get roles', () => {
        const resp = http.get(url + 'roles')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
    group('create role', () => {
        const data = {
            name: 'role1'
        }
        const resp = http.post(url + 'roles', data)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
    group('get role', () => {
        const resp = http.get(url + 'roles' + '/' + 'role1')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
    group('update role', () => {
        const data = {
            name: 'role2'
        }

        const resp = http.put(url + 'roles' + '/' + 'role1', data)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
    group('delete role', () => {
        const resp = http.del(url + 'roles' + '/' + 'role2')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('get users', () => {
        const resp = http.get(url + 'users')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
    group('create user', () => {
        const data = {
            name: 'user1'
        }
        const resp = http.post(url + 'users', data)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
    group('get user', () => {
        const resp = http.get(url + 'users' + '/' + 'user1')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
    group('update user', () => {
        const data = {
            name: 'user2',
            rolename: 'role1'
        }

        const resp = http.put(url + 'users' + '/' + 'user1', data)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
    group('delete user', () => {
        const resp = http.del(url + 'users' + '/' + 'user2')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })
}