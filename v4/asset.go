package pongo2

import (
	"github.com/flosch/pongo2/v4"
	"github.com/go-webpack/webpack"
)

type tagAssetNode struct {
	Name string
}

func (self *tagAssetNode) Execute(ctx *pongo2.ExecutionContext, buffer pongo2.TemplateWriter) *pongo2.Error {
	tag, err := webpack.AssetHelper(self.Name)
	if err != nil {
		buffer.WriteString("asset error: " + self.Name + ":" + err.Error())
	} else {
		buffer.WriteString(string(tag))
	}

	return nil
}

func tagAssetParser(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	node := &tagAssetNode{}

	if filenameToken := arguments.MatchType(pongo2.TokenString); filenameToken != nil {
		node.Name = filenameToken.Val
	}

	return node, nil
}

func init() {
	pongo2.RegisterTag("asset", pongo2.TagParser(tagAssetParser))
}
