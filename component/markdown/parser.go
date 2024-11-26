package markdown

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	mdhtml "github.com/gomarkdown/markdown/html"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var splitter = regexp.MustCompile(`[\s/_]`)

func titleId(s string) string {
	return cases.Lower(language.Und).String(strings.Join(splitter.Split(s, -1), "-"))
}

var parser = mdhtml.NewRenderer(mdhtml.RendererOptions{
	Flags:          mdhtml.CommonFlags,
	RenderNodeHook: myRenderHook,
})

var htmlFormatter = html.New(html.WithClasses(false), html.TabWidth(2))
var highlightStyle = styles.Get("monokailight")

// based on https://github.com/alecthomas/chroma/blob/master/quick/quick.go
func htmlHighlight(w io.Writer, source, lang, defaultLang string) error {
	if lang == "" {
		lang = defaultLang
	}
	l := lexers.Get(lang)
	if l == nil {
		l = lexers.Analyse(source)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return htmlFormatter.Format(w, highlightStyle, it)
}

func renderCode(w io.Writer, codeBlock *ast.CodeBlock) {
	defaultLang := ""
	lang := string(codeBlock.Info)
	htmlHighlight(w, string(codeBlock.Literal), lang, defaultLang)
}

func renderTitle(w io.Writer, heading *ast.Heading, entering bool) {
	if !entering {
		return
	}
	content := []byte{}
	for _, child := range heading.Children {
		content = append(content, child.AsLeaf().Literal...)
	}

	_, _ = io.WriteString(w, fmt.Sprintf(
		`<h%d id="%s">
			%s
		</h%d>
	`, heading.Level, titleId(string(content)), content, heading.Level))

}

func myRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if code, ok := node.(*ast.CodeBlock); ok {
		renderCode(w, code)
		return ast.GoToNext, true
	}
	if heading, ok := node.(*ast.Heading); ok {
		renderTitle(w, heading, entering)
		return ast.SkipChildren, true
	}
	return ast.GoToNext, false
}

func Parse(content []byte) string {
	return string(markdown.ToHTML(content, nil, parser))
}

type Heading struct {
	Content string
	Level   uint8
}

func (h Heading) ID() string {
	return titleId(h.Content)
}

func GetHeadings(content []byte) []Heading {
	var headings []Heading
	var headingLevel uint8 = 0
	isCollectingHeading := false
	var heading []byte
	for i, char := range content {
		isNewLine := i == 0 || content[i-1] == '\n'
		if char == '\n' {
			if headingLevel > 0 && len(heading) > 0 {
				headings = append(headings, Heading{
					Content: string(heading),
					Level:   headingLevel,
				})
				headingLevel = 0
				isCollectingHeading = false
				heading = nil
			}
		} else if isCollectingHeading {
			heading = append(heading, char)
		} else if (isNewLine || headingLevel > 0) && char == '#' {
			headingLevel += 1
		} else if headingLevel > 0 && char == ' ' {
			isCollectingHeading = true
		}
	}

	return headings

}
