package parser

import (
	"github.com/PUGE/logkit/conf"
	"github.com/PUGE/logkit/parser"
	. "github.com/PUGE/logkit/utils/models"
)

func init() {
	parser.RegisterConstructor(parser.TypeEmpty, NewParser)
}

type Parser struct {
	name string
}

func NewParser(c conf.MapConf) (parser.Parser, error) {
	name, _ := c.GetStringOr(parser.KeyParserName, "")
	return &Parser{
		name: name,
	}, nil
}

func (p *Parser) Name() string {
	return p.name
}

func (p *Parser) Parse(lines []string) (datas []Data, err error) {
	return
}
