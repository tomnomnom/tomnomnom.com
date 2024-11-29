package main

import (
	"bytes"
	"embed"
	"errors"
	"html/template"
	"path"
	"strings"

	"github.com/gomarkdown/markdown"
	mdhtml "github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)

//go:embed blog-posts
var blogPostContent embed.FS

const blogPostDir = "blog-posts"

type blogPost struct {
	ID      string
	Title   string
	Content template.HTML
}

type blogs []*blogPost

func (b blogs) find(id string) *blogPost {
	for _, p := range b {
		if p.ID == id {
			return p
		}
	}
	return nil
}

func getBlogPosts() (blogs, error) {
	out := make(blogs, 0)

	entries, err := blogPostContent.ReadDir(blogPostDir)
	if err != nil {
		return out, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		post, err := parseBlogPost(entry.Name())
		if err != nil {
			return out, err
		}

		out = append(out, post)
	}

	return out, nil
}

func parseBlogPost(filename string) (*blogPost, error) {
	raw, err := blogPostContent.ReadFile(path.Join(blogPostDir, filename))
	if err != nil {
		return nil, err
	}

	parts := strings.Split(filename, ".")
	if len(parts) != 2 {
		return nil, errors.New("blog post file lacks extension")
	}

	id := parts[0]
	format := parts[1]

	if format == "mkd" {
		raw = parseMarkdown(raw)
	}

	title, err := extractTitle(raw)
	if err != nil {
		return nil, err
	}

	// TODO: parse metadata from... somewhere.
	return &blogPost{
		ID:      id,
		Title:   title,
		Content: template.HTML(raw),
	}, nil
}

func parseMarkdown(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := mdhtml.CommonFlags | mdhtml.HrefTargetBlank
	opts := mdhtml.RendererOptions{Flags: htmlFlags}
	renderer := mdhtml.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func extractTitle(src []byte) (string, error) {
	sel, err := css.Parse("h1:first-child")
	if err != nil {
		return "", err
	}

	doc, err := html.Parse(bytes.NewReader(src))
	if err != nil {
		return "", err
	}

	elements := sel.Select(doc)
	if len(elements) == 0 {
		return "", errors.New("no title found")
	}

	return elements[0].FirstChild.Data, nil
}
