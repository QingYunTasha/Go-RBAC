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
    const params = {
        headers: { 
            'Content-Type': 'application/json' 
        },
    }

    group('get resources', () => {
        const resp = http.get(url + 'resources')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('create resource', () => {
        const data = {
            name: 'res1',
            "permissions": [{
		        "operation": "read"
	        }]
        }
        
        const resp = http.post(url + 'resources', JSON.stringify(data), params)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('update resource', () => {
        const data = {
            name: 'res2'
        }

        const resp = http.put(url + 'resources' + '/' + 'res1', JSON.stringify(data), params)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('get resource', () => {
        const resp = http.get(url + 'resources' + '/' + 'res2')
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
            operation: 'write',
            resourcename: 'res2'
        }
        
        const resp = http.post(url + 'permissions', JSON.stringify(data), params)
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
        const resp = http.post(url + 'roles', JSON.stringify(data), params)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('update role', () => {
        const data = {
            permissions: [{
                operation: 'write',
                resourcename: 'res2'
            },
            {
                operation: 'read',
                resourcename: 'res2'
            }
        ]
        }

        const resp = http.put(url + 'roles' + '/' + 'role1', JSON.stringify(data), params)
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

    group('get users', () => {
        const resp = http.get(url + 'users')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('create user', () => {
        const data = {
            name: 'user1',
            email: 'user1@example.com'
        }
        const resp = http.post(url + 'users', JSON.stringify(data), params)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('update user', () => {
        const data = {
            rolename: 'role1'
        }

        const resp = http.put(url + 'users' + '/' + 'user1@example.com', JSON.stringify(data), params)
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('get user', () => {
        const resp = http.get(url + 'users' + '/' + 'user1@example.com')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('get permissions by resources', () => {
        const resp = http.get(url + 'permissions' + '?' + 'resource=res2')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('get permissions by role', () => {
        const resp = http.get(url + 'permissions' + '?' + 'role=role1')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('delete user', () => {
        const resp = http.del(url + 'users' + '/' + 'user1@example.com')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('delete role', () => {
        const resp = http.del(url + 'roles' + '/' + 'role1')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('delete permission', () => {
        const resp = http.del(url + 'permissions' + '/' + 'res2' + '/' + 'read')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('delete permission', () => {
        const resp = http.del(url + 'permissions' + '/' + 'res2' + '/' + 'write')
        check(resp, {
            'is success': (r) => r.status === 200,
        })
    })

    group('delete resource', () => {
        const resp = http.del(url + 'resources' + '/' + 'res2')
        check(resp, {
           'is success': (r) => r.status === 200,
       })
   })
}