package usecase

import (
	"context"
	"fmt"
	"github.com/alserov/parser/internal/utils"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

type Parser struct {
}

func (p *Parser) Parse(ctx context.Context, url, tag string, include ...string) ([]string, error) {
	res, err := utils.MakeRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, utils.NewError(utils.ErrBadRequest, fmt.Sprintf("failed to make a request: %v", err))
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, utils.NewError(utils.ErrNotFound, err.Error())
	}

	node, err := html.Parse(res.Body)
	if err != nil {
		return nil, utils.NewError(utils.ErrInternal, err.Error())
	}

	var (
		els []string
		inc bool
	)

	if len(include) != 0 {
		inc = true
	}

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tag {
		attr:
			for _, el := range n.Attr {
				if inc {
					for i := 0; i < len(include); i++ {
						if !strings.Contains(el.Val, include[i]) {
							continue attr
						}
					}
				}
				els = append(els, el.Val)
			}
		}
		for ch := n.FirstChild; ch != nil; ch = ch.NextSibling {
			traverse(ch)
		}
	}

	traverse(node)

	return els, nil
}
