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
				time_lte: %d
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

const addLiquidityQuery = `
	query {
		liquidityPositionSnapshots(
			first: %d,
			where: {
				liquidityTokenBalance_gt: 0,
				timestamp_gt: %d,
				timestamp_lte: %d
			}
		) {
			timestamp
			user {
				id
			}
			liquidityTokenBalance
		}
	}
`
