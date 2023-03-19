# joy-technologies-be

## Getting Started
- Using Go Version 1.20
- Run `go mod tidy`
- Run `go build`
- Run `joy-technologies-be`

## Assumptions
- There is only one book at library
- User only can borrow one book
- User can't borrow book again if user is borrowing book
- To reduce gatherings there is a maximum capacity for every user on pick up schedule (you can change maximum capacity at `internal/constant/constant.go`)
- If user choose pick up schedule on 15:40 so from system will set pick up schedule between 15:00-16:00
