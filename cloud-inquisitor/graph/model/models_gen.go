// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type HijackableResource struct {
	ID      string `json:"id"`
	Account string `json:"account"`
	Type    Type   `json:"type"`
	Value   *Value `json:"value"`
}

type HijackableResourceChain struct {
	ID        string                `json:"id"`
	Resource  *HijackableResource   `json:"resource"`
	Upstream  []*HijackableResource `json:"upstream"`
	Downsteam []*HijackableResource `json:"downsteam"`
}

type Type string

const (
	TypeAccount          Type = "ACCOUNT"
	TypeZone             Type = "ZONE"
	TypeRecord           Type = "RECORD"
	TypeDistribution     Type = "DISTRIBUTION"
	TypeOrigin           Type = "ORIGIN"
	TypeElasticbeanstalk Type = "ELASTICBEANSTALK"
)

var AllType = []Type{
	TypeAccount,
	TypeZone,
	TypeRecord,
	TypeDistribution,
	TypeOrigin,
	TypeElasticbeanstalk,
}

func (e Type) IsValid() bool {
	switch e {
	case TypeAccount, TypeZone, TypeRecord, TypeDistribution, TypeOrigin, TypeElasticbeanstalk:
		return true
	}
	return false
}

func (e Type) String() string {
	return string(e)
}

func (e *Type) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Type(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}

func (e Type) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
