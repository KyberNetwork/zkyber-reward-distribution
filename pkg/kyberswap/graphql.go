package kyberswap

const (
	graphFirstLimit = 1000
)

const routerLogQuery = `
	query {
		routerExchanges(
			first: %d,
			where: {
				time_gt: %d,
				time_lt: %d
			},
			orderBy: time
			orderDirection: asc
		) {
			id
			token
			amount
			userAddress
			time
		}
	}
`
