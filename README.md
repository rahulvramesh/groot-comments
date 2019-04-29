# Groot Comments

1, Post Comment

` POST localhost:8005/orgs/my-company/comments/ `

```json
    { 
     "comment": "Looking to hire SE Asia's top dev talent!"
    }
```

2, Get Comments By Organization

` GET localhost:8005/orgs/my-company/comments/ `

3, Delete Comment / Soft 

`DELETE localhost:8005/orgs/my-company/comments/`

4, Get All Members In Comments

`GET localhost:8005/orgs/my-company/members/`


## Improvements Can Be Done
    -   Using Clean Code Architecture

## Deployments Plannings To Add
    -   K8
    -   Swarm / Compose 
    -   Native Docker

# Run Docker Compose

```sh
    make run
```

# For viewing documentation

```sh
    godoc -http=:6060
```

go and open 
`http://localhost:6060/pkg/github.com/rahulvramesh/groot-comments/`


# Deployment K8

`https://github.com/rahulvramesh/groot-comments/tree/master/docs/k8-deployment`