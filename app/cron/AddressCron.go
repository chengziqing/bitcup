package cron

import (
	"bitcup/app/models"
	"bitcup/app/tools"
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"github.com/revel/revel/modules/jobs/app/jobs"
)

var BtcAddress = AddressCron{}

type AddressCron struct {
	Address []*models.Address
}

func (a AddressCron) Run() {
	for _, adr := range a.Address {
		data, err := tools.Get(fmt.Sprintf("https://blockchain.info/address/%s?format=json&filter=5", adr.Address))
		if err != nil {
			return
		}
		json.Unmarshal(data, adr)
	}
}
func init() {
	//615-1
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "12kArNmKG3BozHk3d8hbaiThmsgU7qn2rB"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "16q3jpwpPJHehG48YTp3tB6s7bQ1gqd2zd"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1E1Xa9FyQRFUAgL1vYcfHukByMc2jY6dwD"})

	//615-2
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1GdtwcTZzJuBsGuc3zDupUhBBbJycFQwWR"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1JhrhwLHRm3zFBA1hAEp6vFGhb5971V6ZQ"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "17rYWUW22vmpGCaL4Vij6ThFrYiQJJLhuF"})

	//615-3
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1AXif9soHHGzfqvfAZBuHry3UVzrkZDdbL"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "14US9NBM7Km2J7Yk6v2cYiQsNvwYzXGGgz"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "13tEHizGkf3DwgsXiAHMeka5f3STDikbpt"})

	//615-4
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1LVNGVmn1vsASZcRFEyT33ue1gxvEakNhj"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1HTZNkFiMCyXASfBagRE7JyaPKtJ82JfA7"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "12Y7bLCY3sH8KQvMTNoC75H64id5tTNAGF"})

	//616-1
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1FcznbMy1j9odtof7gSmKCDRMwctYnZbbK"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "19X5y6F4dCqrwSaSGwVj1a4a2zgFLga4MS"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1BBCBGY1zP1andfB3CDqBKResAk51r4RLe"})

	//616-2
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "13ax3cHLMt93cdWb1QAdbgxSyyviAQNT9n"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1LvEGXdai87cDViqtVMwQnF8QbKN8moqLF"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "18X1YyACdJeTUu8tUhseM9YnudLJFznpHD"})

	//616-3
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "12GEXYvXwRqkM4a5oceWMKt3PLrbirCwy6"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "133A7dxRkVqcqcxWoBH2nQ1uTqhUQ3gzSA"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1GYWs17LQXm86652RSgJZrUVkLzHCC2FZj"})

	//617-1
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "17oyPDh3KXmH4nx76tYtQRPdK1M7HFvmw9"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1J4bdN4S1DduZaceCn42EhmaRiG4Q7Dapo"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "191Cyk7Eo15PwzsfgAzmguhVhzCLdcPZ9H"})

	//617-2
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1G1c5QZMCsNbGYwR8F8GwLJ9M3oaPqqD4t"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1KUYgRRmMTXbGn4LUPE4wn7yxMADWKycG8"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "16mQp1NQhXbh5nvJNSM32gRH4szvRdCs5B"})

	//617-3
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1PuUh8L9ciDg8z5barjQQg5svTdEQ12we1"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1NC5tjQiePN1deDVXa8v28347ZW2knRgTh"})
	//BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1PgGn2vSvfHKgwwvdyNGphrUW4b4cke8vu"})

	//618-1
	BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1Acf8iXtGTFHnoLx2afgua85qjsmtR58hL"})
	BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1EataB7GcshSnf25EnjMZ6tpEqivVJuHhb"})
	BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "12zSz5G8Xby5zTouSSnSFNjpRn2YkL19gY"})

	//618-2
	BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1HF32d4LYQ3kfxT4YxgZdmSJEi4RVdkQmR"})
	BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1ELo6nFneaRGxbJ18hoEU6sGfkTwa6styU"})
	BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1BP8oszrJfARH1RPHUixxEUZBCbCSzwwzi"})

	//618-3
	BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "18kynLrvJL2TWmY9Kg2uFGyPqynQ4XxNLB"})
	BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1GusspS5NJ4o63cT86JgeyHK9VMjruxERb"})
	BtcAddress.Address = append(BtcAddress.Address, &models.Address{Address: "1pCuehkwBsp99hA9jVU2CBxrHHHVAouj2"})

	BtcAddress.Run()
	revel.OnAppStart(func() {
		jobs.Schedule("@every 1m", BtcAddress)
	})
}
