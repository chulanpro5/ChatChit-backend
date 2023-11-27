package util

import (
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

func NullStringToString(val sql.NullString) string {
	if val.Valid {
		return val.String
	} else {
		return ""
	}
}

func StrToInt64(val string) (int64, error) {
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, errors.New("can not parse int")
	}

	return i, nil
}

func StrToInt32(val string) (int32, error) {
	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return 0, errors.New("can not parse int")
	}

	return int32(i), nil
}

func StrToUInt(val string) (uint, error) {
	i, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, errors.New("can not parse int")
	}

	return uint(i), nil
}

func UIntToStr(val uint) (string, error) {
	return strconv.FormatUint(uint64(val), 10), nil
}

func ParseBigFloat(value string) (*big.Float, error) {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	_, err := fmt.Sscan(value, f)
	return f, err
}
