package bitcoin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinP2PKH(t *testing.T) {
	addr := "1AGNa15ZQXAZUgFiqJ2i7Z2DPU2J6hW62i"
	err := ValidateP2PKH(addr)
	assert.Nil(t, err)

	addr = "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"
	err = ValidateP2PKH(addr)
	assert.Nil(t, err)

	addr = "1AGNa15ZQXAZUgFiqJ2i7Z2DPU2J6hW62j"
	err = ValidateP2PKH(addr)
	assert.NotNil(t, err)

	addr = "1AGNa15ZQXAZUgFiqJ2i7Z2DPU2J6hW62"
	err = ValidateP2PKH(addr)
	assert.NotNil(t, err)

	addr = "3BudsSxDY7JX9J8W8CBimHcDyLjTuvNALc"
	err = ValidateP2PKH(addr)
	assert.NotNil(t, err)

	addr = "bc1qaljg7k4pud76kjargxvvx54gz8tgnzx29cy9zn"
	err = ValidateP2PKH(addr)
	assert.NotNil(t, err)
}

func TestBitcoinP2SH(t *testing.T) {
	addr := "3BudsSxDY7JX9J8W8CBimHcDyLjTuvNALc"
	err := ValidateP2SH(addr)
	assert.Nil(t, err)

	addr = "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy"
	err = ValidateP2SH(addr)
	assert.Nil(t, err)

	addr = "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq"
	err = ValidateP2SH(addr)
	assert.NotNil(t, err)

	addr = "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"
	err = ValidateP2SH(addr)
	assert.NotNil(t, err)
}

func TestBitcoinBech32(t *testing.T) {
	addr := "bc1qaljg7k4pud76kjargxvvx54gz8tgnzx29cy9zn"
	err := ValidateSegwitBech32(addr)
	assert.Nil(t, err)

	addr = "bc1qar0srrr7xfkvy5l643lydnw9re59gtzzwf5mdq"
	err = ValidateSegwitBech32(addr)
	assert.Nil(t, err)

	addr = "bc12UEL5L"
	err = ValidateSegwitBech32(addr)
	assert.Nil(t, err)

	addr = "bc12uel5l"
	err = ValidateSegwitBech32(addr)
	assert.Nil(t, err)

	addr = "bc1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqc8247j"
	err = ValidateSegwitBech32(addr)
	assert.Nil(t, err)

	addr = "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy"
	err = ValidateSegwitBech32(addr)
	assert.NotNil(t, err)

	addr = "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"
	err = ValidateSegwitBech32(addr)
	assert.NotNil(t, err)

	addr = "BC13W508D6QEJXTDG4Y5R3ZARVARY0C5XW7KN40WF2"
	err = ValidateSegwitBech32(addr)
	assert.NotNil(t, err)
}
