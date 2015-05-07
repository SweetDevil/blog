package tplfunc

import (
	"github.com/revel/revel"
)

func init() {
	revel.TemplateFuncs["timeStamp"] = TimeStamp
}
