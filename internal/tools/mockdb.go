package tools

import "time"

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails {
	"tejas" : {
		AuthToken: "123ABC",
		Username: "tejas",
		},
	"ujjwal" : {
		AuthToken: "456PQR",
		Username: "ujjwal",
	},
	"pawar" : {
		AuthToken: "098FGD",
		Username: "pawar",
	},
}

var mockCoinDetails = map[string]CoinDetails {
	"tejas" : {
		Coins: 200000,
		Username: "tejas",
	},
	"ujjwal" : {
		Coins: 250000,
		Username: "ujjwal",
	},
	"pawar" : {
		Coins: 10000000,
		Username: "pawar",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second + 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second + 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetUpDatabase() error {
	return nil
}